package lib

import (
	"github.com/yizhiqin809/ayizhiqin/common"
	"github.com/yizhiqin809/ayizhiqin/googlemaps"

	"math"
)

const ZOOM = 18

var regionMap map[string][4]float64 = map[string][4]float64{
	"orlando": [4]float64{28.545016, -81.393588, 28.527108, -81.375800},
	"atlanta": [4]float64{33.765592, -84.403740, 33.738920, -84.373764},
	"st louis": [4]float64{38.636738, -90.223322, 38.614568, -90.181919},
	"nashville": [4]float64{36.175205, -86.803698, 36.145061, -86.757950},
	"dc": [4]float64{38.911075, -77.046454, 38.893809, -77.025790},
	"baltimore": [4]float64{39.297646, -76.623223, 39.282575, -76.603782},
	"philadelphia": [4]float64{39.961156, -75.177383, 39.946988, -75.138210},
	"new york": [4]float64{40.722929, -74.019735, 40.689795, -73.984963},
	"london": [4]float64{51.510040, -0.027977, 51.494132, -0.008300},
	"toronto": [4]float64{43.669194, -79.413479, 43.637359, -79.359827},
}

type Region struct {
	Name string
	RadiusX int
	RadiusY int
	CenterGPS common.Point
	CenterWorld common.Point
}

func GetRegions() []Region {
	var regions []Region
	for name, array := range regionMap {
		centerGPS := common.Point{
			(array[1] + array[3]) / 2,
			(array[0] + array[2]) / 2,
		}
		extreme := googlemaps.LonLatToPixel(common.Point{array[1], array[0]}, centerGPS, ZOOM)
		radiusX := int(math.Ceil(math.Abs(extreme.X) / 4096))
		radiusY := int(math.Ceil(math.Abs(extreme.Y) / 4096))
		if name == "denver" || name == "kansas city" || name == "san diego" || name == "pittsburgh" || name == "montreal" || name == "vancouver" || name == "tokyo" || name == "saltlakecity" || name == "paris" || name == "amsterdam" {
			radiusX = 1
			radiusY = 1
		}
		regions = append(regions, Region{
			Name: name,
			RadiusX: radiusX,
			RadiusY: radiusY,
			CenterGPS: centerGPS,
			CenterWorld: googlemaps.LonLatToMeters(centerGPS),
		})
	}
	return regions
}
