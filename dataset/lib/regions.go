package lib

import (
	"github.com/yizhiqin809/ayizhiqin/common"
	"github.com/yizhiqin809/ayizhiqin/googlemaps"

	"math"
)

const ZOOM = 18

var regionMap map[string][4]float64 = map[string][4]float64{
	"la": [4]float64{34.070393, -118.274908, 34.027270, -118.238398},
	"vegas": [4]float64{36.177588, -115.189791, 36.080612, -115.145384},
	"phoenix": [4]float64{33.464677, -112.088921, 33.441412, -112.061852},
	"dallas": [4]float64{32.798202, -96.817109, 32.772460, -96.776930},
	"austin": [4]float64{30.277550, -97.755393, 30.256969, -97.734987},
	"san antonio": [4]float64{29.440799, -98.503676, 29.418709, -98.475019},
	"houston": [4]float64{29.772755, -95.382410, 29.745711, -95.356047},
	"miami": [4]float64{25.791081, -80.212756, 25.755609, -80.184534},
	"tampa": [4]float64{27.959343, -82.468698, 27.940616, -82.446693},
	
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
