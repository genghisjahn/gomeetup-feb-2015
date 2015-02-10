package main

import (
	"flag"

	"code.google.com/p/gcfg"
)

func main() {

}

func NamedFlags1() {
	namePtr := flag.String("name", "George", "The name of a person")
	agePtr := flag.Int("age", 40, "The age of a person")
	honestPtr := flag.Bool("honest", false, "Bool for honesty")
	flag.Parse() //If you don't call this, flags aren't assigned.
}

func NamedFlags2() {
	var name string
	flag.StringVar(&name, "name", "George", "The name of a person")
	flag.Parse() //If you don't call this, flags aren't assigned.
}

func PosArgs() {
	posArgs := flag.Args()
	flag.Parse()
}

func Config1() {
	type Config struct {
		Section struct {
			Name string
			Flag bool
		}
	}

	var cfg Config
	err := gcfg.ReadFileInto(&cfg, "myconfig.gcfg")
}
