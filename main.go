package main

import (
	"flag"
	"github.com/TraiOi/games"
	"github.com/TraiOi/tpl"
	"github.com/TraiOi/telegram"
)

var (
	app = flag.Bool("app", false, "App")
	daily = flag.Bool("daily", false, "Daily notification")
	cTime = flag.String("time", "", "Time notification")
)

func main() {
	flag.Parse()

	if *daily {
		tpl.TPLDaily(*cTime, games.Handler("daily", *cTime))
	} else if *app {
		telegram.Handler()
	}
}

