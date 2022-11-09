package controller

import (
	"fmt"
	"godevil/src/disk"
	"strings"
)

func Run(l_config []string) {
	fmt.Printf("Gathering information about: %v\n", l_config)
	for _, v := range l_config {
		if strings.ToLower(v) == "storage" {
			d := disk.Get()
			for _, v := range d {
				fmt.Printf("%v\n", v)
			}
		}
	}
}
