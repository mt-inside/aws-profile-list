package main

import (
	"fmt"
	"strings"

	"github.com/mt-inside/go-usvc"
	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load(usvc.HomePath("/.aws/config"))
	if err != nil {
		panic(err)
	}

	for _, s := range cfg.Sections() {
		if s.Name() == "DEFAULT" {
			continue
		}
		fmt.Println(strings.TrimPrefix(s.Name(), "profile "))
	}
}
