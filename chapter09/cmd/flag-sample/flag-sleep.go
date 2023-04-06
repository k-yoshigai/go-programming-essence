package main

import (
	"flag"
	"time"
)

func main() {
	var duration time.Duration
	flag.DurationVar(&duration, "d", 0, "sleep duration")
	flag.Parse()
	println("sleeping for", duration, "[ns]")
	time.Sleep(duration)
}
