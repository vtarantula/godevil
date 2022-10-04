package app

import (
	"fmt"
	"godevil/src/envt"

	"github.com/shirou/gopsutil/disk"
)

func Run(l_config []string) {
	// s, _ := envt.GetKernelVersion()
	// fmt.Printf("%s\n", s)
	m, _ := disk.Partitions(false)
	fmt.Printf("%v\n", m)
	// a, _ := host.KernelVersion()
	// fmt.Printf("%s\n", a)
	a, _ := envt.GetKernelModules()
	fmt.Printf("%v\n", a)
}
