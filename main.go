package main

import (
	"flag"
	"fmt"
	"github.com/domino14/go_euler/fivethirtytwo"
)

func main() {
	var nip = flag.Int("numrobots", 3, "The number of robots")
	flag.Parse()
	// n = 1034 robots. Found this by trial and error. Definitely not
	// under a minute...
	n := *nip
	for true {
		l := fivethirtytwo.GetLengthForRobots(n)
		perRobot := l / float64(n)
		fmt.Println("Length for", n, "robots is", l, "per_robot", perRobot)
		if perRobot >= 1000 {
			fmt.Println("Total length is", l, "num robots", n)
			break
		}
		n += 1
	}
}
