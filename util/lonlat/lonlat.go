package lonlat

import (
	"github.com/jftuga/geodist"
	"github.com/qichengzx/coordtransform"
)

type Coord struct {
	Lon float64
	Lat float64
}

type FullLonLat struct {
	WgsLon float64
	WgsLat float64
	GcjLon float64
	GcjLat float64
	BdLon  float64
	BdLat  float64
}

func WGS84ToFull(lon, lat float64) FullLonLat {
	full := FullLonLat{WgsLon: lon, WgsLat: lat}
	full.GcjLon, full.GcjLat = coordtransform.WGS84toGCJ02(full.WgsLon, full.WgsLat)
	full.BdLon, full.BdLat = coordtransform.GCJ02toBD09(full.GcjLon, full.GcjLat)
	return full
}

func BD09ToFull(lon, lat float64) FullLonLat {
	full := FullLonLat{BdLon: lon, BdLat: lat}
	full.GcjLon, full.GcjLat = coordtransform.BD09toGCJ02(full.BdLon, full.BdLat)
	full.WgsLon, full.WgsLat = coordtransform.GCJ02toWGS84(full.GcjLon, full.GcjLat)
	return full
}

func GpsDistance(lat1, lng1, lat2, lng2 float64) float64 {
	_, km := geodist.HaversineDistance(geodist.Coord{
		Lat: lat1,
		Lon: lng1,
	}, geodist.Coord{
		Lat: lat2,
		Lon: lng2,
	})

	return km * 1000
}
