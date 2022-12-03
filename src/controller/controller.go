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

func Run(l_config []string, httpserver bool) {
	wg := &sync.WaitGroup{}

	fmt.Printf("Received HTTP server setting to: %v\n", httpserver)
	if httpserver {
		wg.Add(1)
		go runHTTPServer(wg)
	}

	//TODO: Add support for n/w and memory
	fmt.Printf("Gathering information about: %v\n", l_config)
	for _, v := range l_config {
		if strings.ToLower(v) == "storage" {
			wg.Add(1)
			go getStorage(wg)
		}
	}

	// s1 := set.New[int]()
	// s2 := set.New[int]()
	// s1.Add(1)
	// s2.Add(1)
	// s2.Add(2)
	// fmt.Printf("Union: %v\n", s1.Union(s2))
	// fmt.Printf("Intersection: %v\n", s1.Intersect(s2))

	// nw, _ := netutil.GetNetworkIPv4("172.24.32.66")
	// fmt.Printf("Network: %v\n", nw)

	wg.Wait()
	fmt.Printf("Ending application...\n")
}
