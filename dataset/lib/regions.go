package lib

import (
	"github.com/yizhiqin809/ayizhiqin/common"
	"github.com/yizhiqin809/ayizhiqin/googlemaps"

	"math"
)

const ZOOM = 18

var regionMap map[string][4]float64 = map[string][4]float64{
	"louisville": [4]float64{38.266052, -85.773983, 38.237431, -85.716176},
	"columbus": [4]float64{39.985654, -83.033001, 39.945536, -82.974917},
	"chicago": [4]float64{41.903117, -87.679072, 41.827860, -87.602447},
	"milwaukee": [4]float64{43.057586, -87.979920, 43.014847, -87.894497},
	"minneapolis": [4]float64{44.999464, -93.296841, 44.960608, -93.240171},
	"seattle": [4]float64{47.626297, -122.353130, 47.591382, -122.316180},
	"portland": [4]float64{45.532231, -122.695860, 45.497663, -122.653470},
	"sf": [4]float64{37.804770, -122.429278, 37.761525, -122.379400},
	"san jose": [4]float64{37.343597, -121.905612, 37.319472, -121.884616},
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
