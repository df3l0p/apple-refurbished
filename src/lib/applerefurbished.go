package applerefurbished

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/net/html"
)

const DefaultUrl = "https://www.apple.com/ch-fr/shop/refurbished/mac/14-pouces-macbook-pro"

func Dump(url string, bucketName string) (string, error) {
	filename := fmt.Sprintf("%s.json", time.Now().Format("2006-01-02"))
	return DumpWithFilename(url, bucketName, filename)
}

func DumpWithFilename(url string, bucketName string, filename string) (string, error) {
	data, err := fetchURL(url)
	if err != nil {
		return "", err
	}
	res, err := getJsonWithComputers(data)
	if err != nil {
		return "", err
	}
	return storeDataWithFilename(context.Background(), res, bucketName, filename)
}

func fetchURL(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error when creating request: %v", err)
	}

	// todo(dferreiralopes): create const
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error when querying '%s': %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return "", fmt.Errorf("HTTP response is not 2xx: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error when reading body: %v", err)
	}
	return string(body), nil
}

func getJsonWithComputers(htmlContent string) (map[string]interface{}, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	var traverse func(*html.Node)
	re := regexp.MustCompile(`\{.*\}`)
	js := ""

	traverse = func(n *html.Node) {
		if n.Type == html.TextNode && strings.Contains(n.Data, "REFURB_GRID_BOOTSTRAP") {
			js = re.FindString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	if js == "" {
		return nil, fmt.Errorf("unable to find valid JSON next to 'REFURB_GRID_BOOTSTRAP' string")
	}

	var res map[string]interface{}
	err = json.Unmarshal([]byte(js), &res)
	if err != nil {
		return nil, fmt.Errorf("unable to parse JSON '%s': %v", js[:10], err)
	}

	return res, nil
}

func storeDataWithFilename(ctx context.Context, data map[string]interface{}, bucketName string, filename string) (string, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to create storage object: %v", err)
	}
	defer client.Close()

	filepath := fmt.Sprintf("dump/%s", filename)
	bucket := client.Bucket(bucketName)
	object := bucket.Object(filepath)
	timeoutCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	writer := object.NewWriter(timeoutCtx)

	writer.ObjectAttrs.ContentType = "application/json"

	b, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("unable to serialize data")
	}

	if _, err := io.Copy(writer, bytes.NewReader(b)); err != nil {
		return "", fmt.Errorf("failed to copy file to bucket: %v", err)
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	return filepath, nil
}

func storeData(ctx context.Context, data map[string]interface{}, bucketName string) (string, error) {
	filename := fmt.Sprintf("%s.json", time.Now().Format("2006-01-02"))
	return storeDataWithFilename(ctx, data, bucketName, filename)
}
