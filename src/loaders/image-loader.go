package loaders

import rl "github.com/gen2brain/raylib-go/raylib"

type ImageLoader struct {
	images map[AppImage]*rl.Image
}

func NewImageLoader() *ImageLoader {
	loader := &ImageLoader{images: make(map[AppImage]*rl.Image)}
	loader._load()
	return loader
}

func (il *ImageLoader) GetImage(imgName AppImage) (*rl.Image, bool) {
	sound, ok := il.images[imgName]
	return sound, ok
}

func (il *ImageLoader) _load() {}
