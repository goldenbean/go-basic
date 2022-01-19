package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
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
	getEnv := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("%s not set\n", key)
		} else {
			fmt.Printf("%s=%s\n", key, val)
		}
	}

	getEnv("EDITOR")
	getEnv("SHELL")
	fmt.Println("Hello World")

	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))
	fmt.Println("Hello World")

	cmd := exec.Command("mvn", "-version")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
