# Apple refurbished

Simple Golang golagn for pulling Apple's refurbished website and dumping the Computers there in GCS.

# Docker

Instructions for running docker image
```bash
make docker-build
docker run --rm -it -e GOOGLE_APPLICATION_CREDENTIALS="/app/credentials.json" -v $HOME/.config/gcloud/application_default_credentials.json:/app/credentials.json applerefurbished -bucket <your_bucket> [-filename <test_filename>]
```
