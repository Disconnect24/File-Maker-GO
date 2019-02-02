package Forecast

/*
	Forecast Locations Parser - Go Edition
		Original Python Version by John Pansera (2017)
		Adapted to GoLang by CornierKhan1 (2019)
*/
import (
	"github.com/cevaris/ordered_map"
)

type value int

var filename = "forecast.bin"
var current = 0
var names = ordered_map.NewOrderedMap()

func seek(value) {
}
