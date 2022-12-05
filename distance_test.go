package euclid

import "testing"

func TestManhattanDistance(t *testing.T) {
	testCases := []struct {
		name           string
		startPoint     [2]float64
		endPoint       [2]float64
		expectedResult float64
	}{
		{
			name:           "distance from origin",
			startPoint:     [2]float64{0, 0},
			endPoint:       [2]float64{3, 4},
			expectedResult: 7,
		},
		{
			name:           "distance equal points",
			startPoint:     [2]float64{0, 0},
			endPoint:       [2]float64{0, 0},
			expectedResult: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ManhattanDistance(tc.startPoint, tc.endPoint); got != tc.expectedResult {
				t.Errorf("expected: %f, got: %f", tc.expectedResult, got)
			}
		})
	}
}

func TestEuclideanDistance(t *testing.T) {
	testCases := []struct {
		name           string
		startPoint     [2]float64
		endPoint       [2]float64
		expectedResult float64
	}{
		{
			name:           "distance from origin",
			startPoint:     [2]float64{0, 0},
			endPoint:       [2]float64{3, 4},
			expectedResult: 5,
		},
		{
			name:           "distance equal points",
			startPoint:     [2]float64{0, 0},
			endPoint:       [2]float64{0, 0},
			expectedResult: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := EuclideanDistance(tc.startPoint, tc.endPoint); got != tc.expectedResult {
				t.Errorf("expected: %f, got: %f", tc.expectedResult, got)
			}
		})
	}
}
