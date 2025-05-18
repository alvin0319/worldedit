package item

import (
	"bytes"
	_ "embed"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
)

//go:embed wand.png
var imageBytes []byte

type Wand struct {
	world.CustomItem
}

func (w Wand) Texture() image.Image {
	return loadImage(imageBytes)
}

func (w Wand) Name() string {
	return "Wand"
}

func (w Wand) Category() category.Category {
	return category.Items()
}

func (w Wand) EncodeItem() (string, int16) {
	return "worldedit:wand", 0
}

func (w Wand) MaxCount() int {
	return 1
}

func loadImage(imageBytes []byte) image.Image {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		panic(err)
	}
	return img
}
