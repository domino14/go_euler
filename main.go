package main

import "github.com/domino14/go_euler/fivethirtytwo"
import "fmt"

func main() {
	n := 100

	// n = 1034 robots. Found this by trial and error. Definitely not
	// under a minute...
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
