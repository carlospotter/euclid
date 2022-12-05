package euclid

import "math"

func ManhattanDistance(StartPoint, EndPoint [2]float64) float64 {
	return math.Abs(EndPoint[0]-StartPoint[0]) + math.Abs(EndPoint[1]-StartPoint[1])
}

func EuclideanDistance(StartPoint, EndPoint [2]float64) float64 {
	return math.Sqrt(math.Pow(EndPoint[0]-StartPoint[0], 2) + math.Pow(EndPoint[1]-StartPoint[1], 2))
}
