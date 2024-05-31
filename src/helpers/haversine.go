package helpers

import (
	"belimang/src/entities"
	"math"
)

// haversine function calculates the distance between two points on the Earth
// given their latitude and longitude in decimal degrees.
func Haversine(location, location2 entities.Location) float64 {
	const earthRadiusKm = 6371.0 // Earth's radius in kilometers

	// Convert latitude and longitude from degrees to radians
	lat1 := DegreesToRadians(location.Lat)
	lat2 := DegreesToRadians(location2.Lat)
	lon1 := DegreesToRadians(location.Long)
	lon2 := DegreesToRadians(location2.Long)
	dLat := lat2 - lat1
	dLon := lon2 - lon1

	// Haversine formula
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusKm * c
}

// degreesToRadians converts degrees to radians.
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
