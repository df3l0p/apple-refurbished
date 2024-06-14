package cmd

import (
	applerefurbished "apple-refurbished/src/lib"
	"fmt"
	"log"
)

func Run(url string, bucket string, filename string) error {
	if bucket == "" {
		return fmt.Errorf("missing bucket")
	}

	log.Printf("fetching data from: '%s'", url)

	var filepath string
	var err error
	if filename != "" {
		filepath, err = applerefurbished.DumpWithFilename(url, bucket, filename)
	} else {
		filepath, err = applerefurbished.Dump(url, bucket)
	}
	if err != nil {
		fmt.Errorf("unable do dump: %v", err)
	}

	log.Printf("Data available under: 'gs://%s/%s'", bucket, filepath)
	return nil
}
