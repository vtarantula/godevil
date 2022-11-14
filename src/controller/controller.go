package controller

import (
	"fmt"
	"godevil/src/hardware/disk"
	"godevil/src/server"
	"strings"
	"sync"
)

func getStorage(wg *sync.WaitGroup) {
	d := disk.Get()
	for _, v := range d {
		fmt.Printf("%v\n", v)
	}
	wg.Done()
}

func runHTTPServer(wg *sync.WaitGroup) {
	fmt.Printf("Starting web server...\n")
	err := server.NewHttp("0.0.0.0", 8560)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	wg.Done()
}

func Run(l_config []string) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go runHTTPServer(wg)

	fmt.Printf("Gathering information about: %v\n", l_config)
	//TODO: Add support for n/w
	for _, v := range l_config {
		if strings.ToLower(v) == "storage" {
			wg.Add(1)
			go getStorage(wg)
		}
	}

	wg.Wait()
	fmt.Printf("Ending application...\n")
}
