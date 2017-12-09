package kubernetes

import "math"

// round positive and negative numbers
// https://gist.github.com/siddontang/1806573b9a8574989ccb
func round(f float64) float64 {
	v, frac := math.Modf(f)
	if f > 0.0 {
		if frac > 0.5 || (frac == 0.5 && uint64(v)%2 != 0) {
			v += 1.0
		}
	} else {
		if frac < -0.5 || (frac == -0.5 && uint64(v)%2 != 0) {
			v -= 1.0
		}
	}

	return v
}
