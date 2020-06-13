# Memegen

A simple web service that generates a meme image given text and an image URL. 

Run with `docker run -p 8080:8080 gcr.io/as-a-service-dev/memegen`

[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://deploy.cloud.run)

### URL parameters:

* `image`: URL of the image
* `top`:  text to add at the top of the image
* `bottom`:  text to add at the bottom of the image

## Running the server locally

* Build with `docker build . -t memegen`
* Start with `docker run -p 8080:8080 memegen`
* Open in your browser at `http://localhost:8080/`

## Deploy to your server

The following container image always reflects the latest version of the `master` branch of this repo: `gcr.io/as-a-service-dev/memegen`

![Cloud Build](https://badger-l7zawt5jsq-uw.a.run.app/build/status?project=as-a-service-dev&id=fdd11a9f-8f5d-46f2-974c-0ab041c02b30)
