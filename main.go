package main

import (
	"flag"
	"github.com/TraiOi/games"
	"github.com/TraiOi/tpl"
)

var (
	daily = flag.Bool("daily", false, "Daily notification")
	cTime = flag.String("time", "", "Time notification")
)

func main() {
	flag.Parse()

	if *daily {
		tpl.TPLDaily(*cTime, games.Handler("daily", *cTime))
	 }
}

