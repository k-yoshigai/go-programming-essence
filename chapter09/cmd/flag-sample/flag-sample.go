package main

import "flag"

func main() {
	var name string
	var max int
	flag.IntVar(&max, "max", 255, "max value")
	flag.StringVar(&name, "name", "", "my name")
	flag.Parse()

	println("name:", name)
	println("max:", max)

	for _, arg := range flag.Args() {
		println(arg)
	}
}
