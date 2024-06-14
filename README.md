# Apple refurbished

Simple Golang app for pulling Apple's refurbished computers and dumping them in GCS.

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

# Query

```
| rename filters.dimensions.dimensionRelYear as filters.dimensions.year, dimensionScreensize as dimension, filters.dimensions.tsMemorySize as ram, filters.dimensions.dimensionCapacity as disk, price.currentPrice.amount as chf
| table _time, title, year, dimension, ram, disk, chf
```
