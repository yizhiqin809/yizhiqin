package lib

import (
	"github.com/yizhiqin809/ayizhiqin/common"
	"github.com/yizhiqin809/ayizhiqin/googlemaps"

	"math"
)

const ZOOM = 18

var regionMap map[string][4]float64 = map[string][4]float64{
	"denver": [4]float64{39.775325, -105.031729, 39.730041, -104.975639},
	"kansas city": [4]float64{39.114856, -94.602818, 39.080144, -94.565256},
	"san diego": [4]float64{32.729028, -117.177619, 32.704259, -117.145797},
	"pittsburgh": [4]float64{40.449522, -80.016203, 40.429688, -79.982987},
	"montreal": [4]float64{45.511932, -73.583147, 45.485183, -73.555746},
	"vancouver": [4]float64{49.292083, -123.137753, 49.267390, -123.108410},
	"tokyo": [4]float64{35.671386, 139.722466, 35.624368, 139.755382},
	"saltlakecity": [4]float64{40.769193, -111.905448, 40.749489, -111.876258},
	"paris": [4]float64{48.860293, 2.317015, 48.838124, 2.346284},
	"amsterdam": [4]float64{52.376805, 4.882880, 52.363139, 4.904190},
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
