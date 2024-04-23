package main

// There is a biker going on a road trip. The road trip consists of n + 1 points at different altitudes. The biker starts his trip on point 0 with altitude equal 0.

// You are given an integer array gain of length n where gain[i] is the net gain in altitude between points i​​​​​​ and i + 1 for all (0 <= i < n). Return the highest altitude of a point.

// func largestAltitude(gain []int) int {
// 	highestAltitude := 0
// 	altitude := []int{0}
// 	altitude = append(altitude, gain[0])
// 	for i := 1; i < len(gain); i++ {
// 		altitude = append(altitude, altitude[i-1]+gain[i])
// 		highestAltitude = max(altitude[i], highestAltitude)
// 	}
// 	return highestAltitude
// }

func largestAltitude(gain []int) int {
	highestAltitude := 0
	currentAltitude := 0
	for _, altitudeGain := range gain {
		currentAltitude += altitudeGain
		highestAltitude = max(highestAltitude, currentAltitude)
	}
	return highestAltitude
}
