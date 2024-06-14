package main

import (
	cmd "apple-refurbished/src/cmd"
	applerefurbished "apple-refurbished/src/lib"
	server "apple-refurbished/src/server"
	"flag"
	"fmt"
	"log"
)

var (
	url      = flag.String("url", applerefurbished.DefaultUrl, "the URL on Apple with the refurbished items")
	bucket   = flag.String("bucket", "", "the bucket where to store the dump")
	filename = flag.String("filename", "", "(optional) the filename for the dump")

	srv = flag.Bool("server", false, "Server mode")

	file = flag.String("file", "", "A file containing the JSON from Apple's refurbished website")
)

func xor(a, b bool) bool {
	return (a || b) && !(a && b)
}

func main() {
	flag.Parse()

	if !xor(*file != "", xor(*srv, (*bucket != "" || *filename != ""))) {
		log.Fatalf("set either filer or server or bucket/filename")
	}

	if *srv {
		server.Run()
	}

	if *file != "" {
		res, err := applerefurbished.ProcessJsonFile(*file)
		if err != nil {
			log.Fatalf("unable to process json file '%s': %v", *file, err)
		}
		fmt.Print(res)
		return
	}

	err := cmd.Run(*url, *bucket, *filename)
	if err != nil {
		log.Fatal(err)
	}
}
