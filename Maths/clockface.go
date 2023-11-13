package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

type Point struct {
	X float64
	Y float64
}

func hoursInRadian(t time.Time) float64 {
	return (math.Pi/(hoursInHalfClock/float64(t.Hour()%12)) + minutesInRadians(t)/hoursInClock)
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadian(t))
}

func minutesInRadians(t time.Time) float64 {
	return (math.Pi/(minutesInHalfClock/float64(t.Minute())) + secondsInRadians(t)/minutesInClock)
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func roughlyEqualFloat64(a, b float64) bool {
	const equlityThreshold = 1e-7
	return math.Abs(a-b) < equlityThreshold
}

func roughlyEqualPoints(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
