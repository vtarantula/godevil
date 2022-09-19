package app

import (
	"fmt"
	"godevil/src/envt"
)

func Run(l_config []string) {
	s, _ := envt.GetKernelVersion()
	fmt.Printf("%s\n", s)
}
