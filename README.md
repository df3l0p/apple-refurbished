# Apple refurbished

Simple Golang golagn for pulling Apple's refurbished website and dumping the Computers there in GCS.

# Docker

Instructions for running docker image
```bash
make docker-build
docker run --rm -it -e GOOGLE_APPLICATION_CREDENTIALS="/app/credentials.json" -v $HOME/.config/gcloud/application_default_credentials.json:/app/credentials.json applerefurbished -bucket <your_bucket> [-filename <test_filename>]
```


# Usage

Run directly on CLI
```bash
go run src/main.go -bucket "<some-bucket>" -filename "testcli"
```

Run via web service
```bash
go run src/main.go -server

curl "http://localhost:8080/?bucket=<some-bucket>&filename=testhttp"
```
