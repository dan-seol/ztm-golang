package main

import (
	"bytes"
	"database/sql"
	"image"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/chai2010/webp"

	_ "github.com/mattn/go-sqlite3"

	"keiko/keikodb"
)

type WebpImage = bytes.Buffer

// Scans the given directory and returns all entries.
func getImageNames(path string, cache map[string][]os.DirEntry) []os.DirEntry {
	cachedEntries, ok := cache[path]
	if ok {
		return cachedEntries
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		cache[path] = entries
	}()
	return entries
}

// Read raw image data from disk.
func loadImg(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Parse the raw bytes of an image into an Image format.
func parseImg(rawImg []byte) image.Image {
	img, _, err := image.Decode(bytes.NewReader(rawImg))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

// Convert an image to WebP format.
func imgToWebp(img image.Image) WebpImage {
	var buf bytes.Buffer
	if err := webp.Encode(&buf, img, &webp.Options{Lossless: true}); err != nil {
		log.Fatal(err)
	}
	return buf
}

// Returns a random image from the `data` directory.
// The first return value is the name of the image,
// the second return value is the image data in WebP format.
func getImage(imgBase string, imageNameCache map[string][]os.DirEntry, imageCache map[string]WebpImage) (string, WebpImage) {
	// find all the names of the images
	imgNames := getImageNames(imgBase, imageNameCache)
	// pick an index at random
	index := rand.Intn(len(imgNames))

	// get the randomly picked image
	imgEntry := imgNames[index]
	imgName := imgEntry.Name()
	// construct a path to the image in the `data` directory
	imgPath := path.Join(imgBase, imgEntry.Name())
	cachedWebpImage, ok := imageCache[imgPath]
	if ok {
		return imgName, cachedWebpImage
	}
	// load the image into memory
	//rawImg :=

	// parse into an image.Image
	//parsedImg :=
	// convert to WebP
	webpImg := imgToWebp(parseImg(loadImg(imgPath)))
	defer func() {
		imageCache[imgPath] = webpImg
	}()
	return imgName, webpImg
}

// handler for a GET request on `/`
func handleRoot(db *sql.DB, imageNameCache map[string][]os.DirEntry, imageCache map[string]WebpImage, totalRequest *int, counterMap sync.Map) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			return
		}

		// get a random image
		name, img := getImage("data", imageNameCache, imageCache)
		// defer keikodb.IncrementHitCount(db, name)
		// serve it
		w.Write(img.Bytes())
		defer func() {
			if req.Method != "GET" {
				return
			}
			//log.Println("Served", name)
			valueToStore, _ := counterMap.LoadOrStore(name, 0)
			counterMap.Store(name, valueToStore.(int)+1)
		}()
		return
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())

	log.Println("Using database 'keiko.db'")
	db, err := sql.Open("sqlite3", "keiko.db")
	if err != nil {
		log.Fatal(err)
	}

	// try to create a new database
	keikodb.MakeNew(db)

	// set caches
	// imageNameCache map[string][]os.DirEntry, rawImageCache map[string][]byte
	imageNameCache := make(map[string][]os.DirEntry)
	imageCache := make(map[string]WebpImage)
	var counterMap sync.Map
	var totalRequest int
	// set channel
	// ch := make(chan string, 50)
	// serve on `/`
	http.Handle("/", handleRoot(db, imageNameCache, imageCache, &totalRequest, counterMap))
	// increment the hit counter for this image

	log.Println("Listening on localhost:8099")
	err = http.ListenAndServe("localhost:8099", nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {

		keikodb.IncrementInBatch(db, counterMap)
	}()
}
