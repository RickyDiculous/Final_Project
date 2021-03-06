package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)


// The resources struct will hold textures and potentially fonts and sounds
type Resources struct {
	images map[string]*sf.Texture
	textures map[string]*sf.Image
}

func NewResources() *Resources {
	r := new(Resources)

	r.images = make(map[string]*sf.Texture)
	r.textures = make (map[string]*sf.Image, 0)

	r.LoadAllImages("./assets/images")
	r.LoadTextures ("./assets/images")

	return r
}

// Recursively searches assets/images for .png files and appends them to images slice
func (r *Resources) LoadAllImages(dir string) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			r.LoadAllImages(dir + "/" + f.Name())
		} else if filepath.Ext(f.Name()) == ".png" {
			texture := sf.NewTexture(dir + "/" + f.Name())
			r.images[f.Name()] = texture
		}
	}
}

// Does the same as function above, only appends a type *sf.Texture to the textures slice
func (r *Resources) LoadTextures(dir string) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			r.LoadTextures(dir + "/" + f.Name())
		} else if filepath.Ext(f.Name()) == ".png" {
			filename := dir + "/" + f.Name()
			r.textures[f.Name ()] = sf.NewImage (filename)
		}
	}
}