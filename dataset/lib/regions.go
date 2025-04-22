package lib

import (
	"github.com/yizhiqin809/ayizhiqin/common"
	"github.com/yizhiqin809/ayizhiqin/googlemaps"

	"math"
)

const ZOOM = 18

var regionMap map[string][4]float64 = map[string][4]float64{
	
	"hangzhou":[4]float64{30.273914, 120.170362,30.253914, 120.200362},
	"xiaoshan":[4]float64{30.185474, 120.265899,30.165474, 120.295899},
	"beijing":[4]float64{39.905250, 116.397281,39.885250, 116.427281},
	
	
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
		if name == "beijing" || name == "xiaoshan" || name == "jinghua" || name == "san diego" || name == "pittsburgh" || name == "montreal" || name == "vancouver" || name == "tokyo" || name == "saltlakecity" || name == "paris" || name == "amsterdam" {
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
