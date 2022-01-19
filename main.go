package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	env  *string
	port *int
)

func init() {
	env = flag.String("env", "development", "current environment")
	port = flag.Int("port", 3000, "port number")
}

func main() {
	flag.Parse()

	fmt.Println("env:", *env)
	fmt.Println("port:", *port)

	getenvironment := func(data []string, getkeyval func(item string) (key, val string)) map[string]string {
		items := make(map[string]string)
		for _, item := range data {
			key, val := getkeyval(item)
			items[key] = val
		}
		return items
	}
	environment := getenvironment(os.Environ(), func(item string) (key, val string) {
		splits := strings.Split(item, "=")
		key = splits[0]
		val = splits[1]
		return
	})
	fmt.Println(environment["PATH"])
}
