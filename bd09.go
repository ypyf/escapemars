package escapemars

import . "math"

const (
	X_Pi = Pi * 3000.0 / 180.0
)

// gcj02 to baidu
func MarsToBaidu(gg_lat, gg_lon float64) (bd_lat, bd_lon float64) {
	x := gg_lon
	y := gg_lat
	z := Sqrt(x*x+y*y) + 0.00002*Sin(y*X_Pi)
	theta := Atan2(y, x) + 0.000003*Cos(x*X_Pi)
	bd_lon = z*Cos(theta) + 0.0065
	bd_lat = z*Sin(theta) + 0.006
	return
}

// baidu to gcj02
func BaiduToMars(bd_lat, bd_lon float64) (gg_lat, gg_lon float64) {
	x := bd_lon - 0.0065
	y := bd_lat - 0.006
	z := Sqrt(x*x+y*y) - 0.00002*Sin(y*X_Pi)
	theta := Atan2(y, x) - 0.000003*Cos(x*X_Pi)
	gg_lon = z * Cos(theta)
	gg_lat = z * Sin(theta)
	return
}
