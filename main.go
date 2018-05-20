package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/domino14/go_euler/fivethirtytwo"
	"github.com/domino14/go_euler/sixtyone"
)

func mainRobots532() {
	var nip = flag.Int("numrobots", 3, "The number of robots")
	var cpuprofile = flag.String("cpuprofile", "", "Write cpu profile to file")
	var cont = flag.Bool("continue", false, "Loop until solution is found.")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// n = 830 robots. Found this by trial and error. Definitely not
	// under a minute...
	n := *nip
	for {
		l := fivethirtytwo.GetLengthForRobots(n)
		perRobot := l / float64(n)
		fmt.Println("Length for", n, "robots is", l, "per_robot", perRobot)
		if perRobot >= 1000 {
			fmt.Println("Total length is", l, "num robots", n)
			break
		}
		n += 1
		if *cont == false {
			break
		}
	}
}

func main() {
	sixtyone.Solve()
}
