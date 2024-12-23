package main

import (
	"./lib"
	"github.com/yizhiqin809/ayizhiqin/common"
	"github.com/yizhiqin809/ayizhiqin/googlemaps"

	"fmt"
	"os"
	"strconv"
)

const ZOOM = 18

func main() {
	rname := os.Args[1]
	mode := os.Args[2]
	x, _ := strconv.ParseFloat(os.Args[3], 64)
	y, _ := strconv.ParseFloat(os.Args[4], 64)
	var region lib.Region
	for _, r := range lib.GetRegions() {
		if r.Name == rname {
			region = r
		}
	}
	p := common.Point{x, y}
	if mode == "frompix" {
		fmt.Println(googlemaps.PixelToLonLat(p, region.CenterGPS, ZOOM))
	} else if mode == "topix" {
		fmt.Println(googlemaps.LonLatToPixel(p, region.CenterGPS, ZOOM))
	}
}
