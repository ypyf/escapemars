package escapemars

import . "math"

const (
	// 火星坐标采用的SK-42参考系
	// Krasovsky 1940
	//
	// a = 6378245.0, 1/f = 298.3
	// b = a * (1 - f)
	// ee = (a^2 - b^2) / a^2;
	A  = 6378245.0
	EE = 0.00669342162296594323
)

func Wgs84ToMars(wgLat, wgLon float64) (mgLat, mgLon float64) {
	if out_of_china(wgLat, wgLon) {
		return wgLat, wgLon
	}
	dLat := transform_lat(wgLon-105.0, wgLat-35.0)
	dLon := transform_lon(wgLon-105.0, wgLat-35.0)
	radLat := wgLat / 180.0 * Pi
	magic := Sin(radLat)
	magic = 1 - EE*magic*magic
	sqrtMagic := Sqrt(magic)
	dLat = (dLat * 180.0) / ((A * (1 - EE)) / (magic * sqrtMagic) * Pi)
	dLon = (dLon * 180.0) / (A / sqrtMagic * Cos(radLat) * Pi)
	mgLat = wgLat + dLat
	mgLon = wgLon + dLon
	return
}

func out_of_china(lat, lon float64) bool {
	if lon < 72.004 || lon > 137.8347 {
		return true
	}
	if lat < 0.8293 || lat > 55.8271 {
		return true
	}
	return false
}

func transform_lat(x, y float64) (lat float64) {
	lat = -100.0 + 2.0*x + 3.0*y + 0.2*y*y + 0.1*x*y + 0.2*Sqrt(Abs(x))
	lat += (20.0*Sin(6.0*x*Pi) + 20.0*Sin(2.0*x*Pi)) * 2.0 / 3.0
	lat += (20.0*Sin(y*Pi) + 40.0*Sin(y/3.0*Pi)) * 2.0 / 3.0
	lat += (160.0*Sin(y/12.0*Pi) + 320*Sin(y*Pi/30.0)) * 2.0 / 3.0
	return
}

func transform_lon(x, y float64) (lon float64) {
	lon = 300.0 + x + 2.0*y + 0.1*x*x + 0.1*x*y + 0.1*Sqrt(Abs(x))
	lon += (20.0*Sin(6.0*x*Pi) + 20.0*Sin(2.0*x*Pi)) * 2.0 / 3.0
	lon += (20.0*Sin(x*Pi) + 40.0*Sin(x/3.0*Pi)) * 2.0 / 3.0
	lon += (150.0*Sin(x/12.0*Pi) + 300.0*Sin(x/30.0*Pi)) * 2.0 / 3.0
	return
}
