package main

import (
	"flag"
	"github.com/TraiOi/games"
)

var (
	daily = flag.Bool("daily", false, "Daily notification")
	cTime = flag.String("time", "", "Time notification")
)

func main() {
	flag.Parse()

	if *daily {
		games.Handler("daily", *cTime)
	 }
}
