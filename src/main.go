package main

import (
	cmd "apple-refurbished/src/cmd"
	applerefurbished "apple-refurbished/src/lib"
	server "apple-refurbished/src/server"
	"flag"
	"log"
)

var (
	url      = flag.String("url", applerefurbished.DefaultUrl, "the URL on Apple with the refurbished items")
	bucket   = flag.String("bucket", "", "the bucket where to store the dump")
	filename = flag.String("filename", "", "(optional) the filename for the dump")

	srv = flag.Bool("server", false, "Server mode")
)

func xor(a, b bool) bool {
	return (a || b) && !(a && b)
}

func main() {
	flag.Parse()

	if !xor(*srv, (*bucket != "" || *filename != "")) {
		log.Fatalf("set either server or bucket/filename")
	}

	if *srv {
		server.Run()
	}

	cmd.Run(*url, *bucket, *filename)
}
