package fivethirtytwo

import (
	"fmt"
	"log"
	"math"
)

const (
	TinyStepSize = 1e-6
	Epsilon      = 1e-4
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
	d := geodesicDistance(n, r)
	if d < Epsilon {
		log.Println("Got there!")
		return false
	}
	sinD := math.Sin(d)
	sinLat, cosLat := math.Sincos(n.lat)
	sinLon, cosLon := math.Sincos(n.lon)
	sinRLat, cosRLat := math.Sincos(r.lat)
	sinRLon, cosRLon := math.Sincos(r.lon)
	distTraveled := TinyStepSize * d
	A := math.Sin(d-distTraveled) / sinD
	B := math.Sin(distTraveled) / sinD
	Acoslat := A * cosLat
	Bcosrlat := B * cosRLat
	x := Acoslat*cosLon + Bcosrlat*cosRLon
	y := Acoslat*sinLon + Bcosrlat*sinRLon
	z := A*sinLat + B*sinRLat
	n.lat = math.Atan2(z, math.Sqrt(x*x+y*y))
	n.lon = math.Atan2(y, x)
	n.traveled += distTraveled
	return true
}

func (n *Nanobot) String() string {
	return fmt.Sprintf("id: %v, lat: %v, lon: %v", n.id, n.lat, n.lon)
}

var HavMemo map[float64]float64

func haversine(angle float64) float64 {
	if val, ok := HavMemo[angle]; ok {
		return val
	} else {
		sin := math.Sin(angle / 2)
		HavMemo[angle] = sin * sin
		return HavMemo[angle]
	}
}

func geodesicDistance(n1 *Nanobot, n2 *Nanobot) float64 {
	deltaLat := n1.lat - n2.lat
	deltaLon := n1.lon - n2.lon
	d := haversine(deltaLat) +
		math.Cos(n1.lat)*math.Cos(n2.lat)*haversine(deltaLon)
	return 2 * math.Asin(math.Sqrt(d))
}

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
	if HavMemo == nil {
		HavMemo = make(map[float64]float64)
	}
	robots := place_robots(n)
	lon_diff := 2 * math.Pi / float64(n)
	traveling := true
	iterations := 0

	for traveling {
		iterations += 1
		traveling = robots[0].travel(robots[1])
		robots[1].lat = robots[0].lat
		robots[1].lon = robots[0].lon + lon_diff
	}
	return robots[0].traveled * float64(n)
}
