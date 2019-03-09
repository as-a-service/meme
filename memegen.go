package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/fogleman/gg"
)

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	log.Print("New meme ", q)

	textTop := q.Get("text-top")
	textBottom := q.Get("text-bottom")

	// Download image
	imgURL := q.Get("image")
	if imgURL == "" {
		fmt.Fprintf(w, "Please provide an image with ?image=URL")
		return
	}
	resp, err := http.Get(imgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	im, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bounds := im.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	dc := gg.NewContextForImage(im)

	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(textTop, float64(width/2), float64(height/4), 0.5, 0.5)
	dc.DrawStringAnchored(textBottom, float64(width/2), float64(3*height/4), 0.5, 0.5)

	w.Header().Set("Content-Type", "image/jpeg")
	jpeg.Encode(w, dc.Image(), nil)
}

func main() {

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Print("Starting memegen.")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
