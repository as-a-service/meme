package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/jpeg"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("New meme")

	const S = 1024

	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)

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
