# Memegen

A simple web service that generates a meme image given text and an image URL. 

Run with `docker run -p 8080:8080 gcr.io/steren-demo/memegen`

### URL parameters:

* `image`: URL of the image
* `text-top`:  text to add at the top of the image
* `text-bottom`:  text to add at the bottom of the image

## Running the server locally

* Build with `docker build . -t memegen`
* Start with `docker run -p 8080:8080 memegen`
* Open in your browser at `http://localhost:8080/`

## Deploy to your server

The following container image always reflects the latest version of the `master` branch of this repo: `gcr.io/steren-demo/memegen`