package config

import "flag"

var Port string
var Mode string

func init() {
	flag.StringVar(&Port, "p", "9608", "Listening port")
	flag.StringVar(&Mode, "m", "release", "release or debug")
	flag.Parse()
}
