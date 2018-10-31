package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/nfnt/resize"
)

func main() {
	path, _ := os.Getwd()
	fmt.Println(path)

	err := filepath.Walk(".", func(filepath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a filepath %q: %v\n", filepath, err)
			return err
		}

		if info.IsDir() != true {
			thumbnail(path, info.Name())
			resample(path, info.Name())
		}

		return nil
	})
	if err != nil {
		fmt.Println("error walking the path", err)
		return
	}

}

func thumbnail(wpath, name string) {

	fp := path.Join(wpath, name)

	file, err := os.Open(fp)
	if err != nil {
		log.Println(err)
		return
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
		return
	}

	out := resize.Resize(1000, 0, img, resize.Bilinear)

	file, err = os.Create(fp)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	jpeg.Encode(file, out, nil)
}

func resample(wpath, name string) {
	fp := path.Join(wpath, name)

	file, err := os.Open(fp)
	if err != nil {
		log.Println(err)
		return
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
		return
	}

	out := resize.Resize(300, 0, img, resize.Bilinear)

	resampledDir := path.Join(wpath, "_vorschaubilder")
	os.MkdirAll(resampledDir, 0755)

	resampledFp := path.Join(resampledDir, name)
	file, err = os.Create(resampledFp)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	jpeg.Encode(file, out, nil)
}
