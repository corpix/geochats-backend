package harvesine

import (
	"math"
)

// see: http://bluemm.blogspot.ru/2007/01/excel-formula-to-calculate-distance.html
// for explanation

// radius of Earth on kilometers
const r = 6371

func haversin(theta float64) float64 {
	return 0.5 * (1 - math.Cos(theta))
}

func toRadians(deg float64) float64 {
	return deg * (math.Pi / 180)
}

// Distance calculates a distance between two points
func Distance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	phi1 := toRadians(lat1)
	phi2 := toRadians(lat2)
	lambda1 := toRadians(lon1)
	lambda2 := toRadians(lon2)

	return 2 * r * math.Asin(
		math.Sqrt(
			haversin(phi2-phi1)+math.Cos(phi1)*math.Cos(phi2)*haversin(lambda2-lambda1),
		),
	)
}
