package pkg

import (
	"fmt"
	"image/color"
	"image/png"
	"os"

	sm "github.com/flopp/go-staticmaps"
	"github.com/golang/geo/s2"
)

const (
	ImageWidth  = 600
	ImageHeight = 600
)

func GenerateMapImage(locations []GeoLocation) error {
	ctx := sm.NewContext()
	ctx.SetSize(ImageWidth, ImageHeight)

	for _, location := range locations {
		lat, lon := location.Lat, location.Lng
		if lat == 0 && lon == 0 {
			continue
		}
		ctx.AddMarker(sm.NewMarker(s2.LatLngFromDegrees(lat, lon), color.RGBA{0xff, 0, 0, 0xff}, 16.0))
	}

	img, err := ctx.Render()
	if err != nil {
		return fmt.Errorf("could not render map: %v", err)
	}

	file, err := os.Create("./output/map.png")
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return fmt.Errorf("could not encode image: %v", err)
	}

	return nil
}
