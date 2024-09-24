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
index=* title="*macbook*"
| sort _time 0
| rename filters.dimensions.dimensionRelYear as year, filters.dimensions.dimensionScreensize as dimension, filters.dimensions.tsMemorySize as ram, filters.dimensions.dimensionCapacity as disk, price.currentPrice.raw_amount as chf
| fillnull value="" title, year, dimension, ram, disk, chf
| stats min(_time) as first_seen, max(_time) as last_seen, values(title) as title, values(year) as year, values(dimension) as dimension, values(ram) as ram, values(disk) as disk, values(chf) as chf by partNumber
| eval is_available=if(last_seen>relative_time(now(), "-1d"), 1, 0)
| convert ctime(first_seen), ctime(last_seen)
| sort - chf
```
