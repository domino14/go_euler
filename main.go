package main

import (
	"flag"
	"fmt"
	"github.com/domino14/go_euler/fivethirtytwo"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	var nip = flag.Int("numrobots", 3, "The number of robots")
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// n = 1034 robots. Found this by trial and error. Definitely not
	// under a minute...
	n := *nip
	// for true {
	for n = 3; n < 8; n++ {
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
