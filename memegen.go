package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/fogleman/gg"
)

func createMeme(im image.Image, textTop string, textBottom string) image.Image {
	bounds := im.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	max := math.Round(float64(width) / 10)

	dc := gg.NewContextForImage(im)

	if err := dc.LoadFontFace("/usr/share/fonts/truetype/msttcorefonts/impact.ttf", max); err != nil {
		panic(err)
	}

	positionX := float64(width / 2)
	positionTopY := float64(height / 6)
	positionBottomY := float64(5 * height / 6)

	dc.SetRGB(0, 0, 0)
	n := 6 // "stroke" size
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := positionX + float64(dx)
			ytop := positionTopY + float64(dy)
			ybottom := positionBottomY + float64(dy)
			dc.DrawStringAnchored(strings.ToUpper(textTop), x, ytop, 0.5, 0)
			dc.DrawStringAnchored(strings.ToUpper(textBottom), x, ybottom, 0.5, 1)
		}
	}

	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(strings.ToUpper(textTop), positionX, positionTopY, 0.5, 0)
	dc.DrawStringAnchored(strings.ToUpper(textBottom), positionX, positionBottomY, 0.5, 1)

	return dc.Image()
}

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	log.Print("New meme ", q)

	textTop := q.Get("top")
	textBottom := q.Get("bottom")

	// Download image
	imgURL := q.Get("image")
	if imgURL == "" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "Generate meme by providing an image URL, top and bottom text using query parameters. See <a href=\"/?top=I'm in ur cloud&bottom=creating ur memes&image=https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Cat_on_laptop_-_Just_Browsing.jpg/800px-Cat_on_laptop_-_Just_Browsing.jpg\">example</a>")
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

	meme := createMeme(im, textTop, textBottom)

	w.Header().Set("Content-Type", "image/jpeg")
	jpeg.Encode(w, meme, nil)
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
