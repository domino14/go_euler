package fivethirtytwo

import (
	"fmt"
	"log"
	"math"
)

const (
	TinyStepSize = 1e-7
	Epsilon      = 1e-6
	PiOver2      = math.Pi / 2
	// TinyStepSize = 0.00048828125
	// Epsilon      = 0.00390625
)

type Nanobot struct {
	lat        float64
	lon        float64
	id         int
	traveled   float64
	lat_backup float64
	lon_backup float64
}

// travel towards nanobot r
func (n *Nanobot) travel(r *Nanobot) bool {
	if PiOver2-n.lat < Epsilon {
		log.Println("Got there!", n)
		return false
	}
	sinLat, cosLat := math.Sincos(n.lat)
	// log.Println("n.lat", n.lat)
	// A simplified version of the geodesic distance equation.
	sqrt_intermed := cosLat * sqrt_hav_londiff
	// d := 2 * math.Asin(sqrt_intermed)
	// sin 2x = 2 sinx cosx
	// so sin of d is sin of 2 * x
	// = 2 sin (Asin(sqrt_intermed)) cos(Asin(sqrt_intermed))
	// = 2 * sqrt_intermed * cos(Asin(sqrt_intermed))
	// = 2 * sqrt_intermed * sqrt(1 - sqrt_intermed*sqrt_intermed)

	intermed := sqrt_intermed * sqrt_intermed
	// sinD := math.Sin(d)
	sinD := 2 * sqrt_intermed * math.Sqrt(1-intermed)
	// log.Println("1-intermed", 1-intermed)
	cosD := math.Sqrt(1 - sinD*sinD)
	// log.Println("1-sinDsinD", 1-sinD*sinD)
	// sinD * sinD = 4 * intermed * (1 - intermed)
	// = 4I * (1 - I)
	// then math.Sqrt( 1 -(4I (1-I))) ==
	// 1 - (4I - 4II) = 1-4I + 4II = (2I - 1)^2
	// so cosD is just 2 I -1
	//cosD := 2*intermed - 1

	// longitude is always reset back to 0 after every iteration.
	// latitude of all robots is the same by symmetry.

	// See: http://williams.best.vwh.net/avform.htm#Intermediate
	A := cos_stepsize - cosD*sin_stepsize/sinD
	B := sin_stepsize / sinD

	Acoslat := A * cosLat
	Bcosrlat := B * cosLat
	x := Acoslat + Bcosrlat*cos_londiff
	y := Bcosrlat * sin_londiff
	z := sinLat * (A + B)
	// log.Println("Atan2args", z, x*x+y*y)
	n.lat = math.Atan2(z, math.Sqrt(x*x+y*y))

	// fmt.Println("Bearing", math.Atan2(lon_diff*cosLat,
	// 	cosLat*sinLat*(1-cos_londiff)))

	// Rotate sphere back to longitude 0. This problem is symmetric.
	n.lon = 0
	n.traveled += TinyStepSize
	return true
}

func (n *Nanobot) String() string {
	return fmt.Sprintf("id: %v, lat: %v, lon: %v", n.id, n.lat, n.lon)
}

var sin_londiff float64
var cos_londiff float64
var hav_londiff float64
var lon_diff float64
var sqrt_hav_londiff float64
var cos_stepsize float64
var sin_stepsize float64

// This function is only used once, does not need optimization.
func smallCircleLatitude(radius float64) float64 {
	// Use pythagoras
	h := math.Sqrt(1 - radius*radius)
	return math.Asin(h / math.Sqrt(radius*radius+h*h))
}

func place_robots(n int) []*Nanobot {
	robots := make([]*Nanobot, n)
	lat := smallCircleLatitude(0.999)
	for i := 0; i < n; i++ {
		lon := float64(i) * (2 * math.Pi / float64(n))
		robots[i] = &Nanobot{lat: lat, lon: lon, id: i}
	}
	return robots
}

func GetLengthForRobots(n int) float64 {
	robots := place_robots(n)
	lon_diff = 2 * math.Pi / float64(n)
	sin_londiff, cos_londiff = math.Sincos(lon_diff)
	hav_londiff = math.Pow(math.Sin(lon_diff/2), 2.0)
	sqrt_hav_londiff = math.Sin(lon_diff / 2)
	sin_stepsize, cos_stepsize = math.Sincos(TinyStepSize)
	traveling := true
	iterations := 0

	for traveling {
		iterations += 1
		traveling = robots[0].travel(robots[1])
		robots[1].lat = robots[0].lat
		robots[1].lon = robots[0].lon + lon_diff
	}
	log.Println("Iterations:", iterations)
	return robots[0].traveled * float64(n)
}
