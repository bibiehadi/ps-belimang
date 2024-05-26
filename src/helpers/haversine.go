package helpers

import (
	"math"
)

// haversine function calculates the distance between two points on the Earth
// given their latitude and longitude in decimal degrees.
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadiusKm = 6371.0 // Earth's radius in kilometers

	// Convert latitude and longitude from degrees to radians
	dLat := DegreesToRadians(lat2 - lat1)
	dLon := DegreesToRadians(lon2 - lon1)
	rLat1 := DegreesToRadians(lat1)
	rLat2 := DegreesToRadians(lat2)

	// Haversine formula
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(rLat1)*math.Cos(rLat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusKm * c
}

// degreesToRadians converts degrees to radians.
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
