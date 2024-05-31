package helpers

import (
	"belimang/src/entities"
	"math"
)

func NearestNeighbor(listLocation []entities.Location) ([]entities.Location, float64) {
	if len(listLocation) == 0 {
		return nil, 0
	}

	visited := make([]bool, len(listLocation))
	trip := make([]entities.Location, 0, len(listLocation))

	currentLocation := listLocation[0]
	trip = append(trip, currentLocation)
	visited[0] = true
	totalDistance := 0.0

	for len(trip) < len(listLocation) {
		nearestLocationIndex := 0
		shortestDistance := math.MaxFloat64

		for i, location := range listLocation {
			if !visited[i] {
				distance := Haversine(currentLocation, location)
				if distance < shortestDistance {
					shortestDistance = distance
					nearestLocationIndex = i
				}
			}
		}
		visited[nearestLocationIndex] = true
		currentLocation = listLocation[nearestLocationIndex]
		trip = append(trip, currentLocation)

		totalDistance += shortestDistance
	}

	totalDistance += Haversine(currentLocation, listLocation[0])
	trip = append(trip, listLocation[0])

	return trip, totalDistance
}
