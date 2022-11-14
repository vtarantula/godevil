package main

// #cgo CFLAGS: -I${SRCDIR}/../src/clib
// #cgo LDFLAGS: ${SRCDIR}/../src/clib/clib.a
// #include <stdlib.h>
// #include <clib.h>
// import "C"
import (
	"errors"
	"flag"
	"fmt"
	"godevil/src/app"
	"godevil/src/controller"
	"os"
	"runtime"
)

func getConfigs(config_type string) []string {
	var l_config []string
	switch config_type {
	case "net":
		l_config = append(l_config, "net")
	case "storage":
		l_config = append(l_config, "storage")
	case "all":
		l_config = append(l_config, "net")
		l_config = append(l_config, "storage")
	default:
		panic(errors.New("invalid option: " + config_type))
	}
	return l_config
}

func parseFlags() {
	var configType = flag.String("type", "all", "Configuration type. Can be one of net and storage")
	flag.Parse()
	fmt.Printf("Args: %s\n", *configType)
	l_config := getConfigs(*configType)
	controller.Run(l_config)
}

func cleanup() {
	return_code := app.EXIT_SUCCESS_CODE
	if r := recover(); r != nil {
		return_code = app.EXIT_ERROR_CODE
		fmt.Printf("%v\n", r)
	}
	os.Exit(return_code)
}

// func RunC() {
// 	mess := C.CString("This is another string!")
// 	defer C.free(unsafe.Pointer(mess))
// 	C.cPrintMessage(mess)
// 	C.cGoBroker()
// }

func main() {
	defer cleanup()

	if runtime.GOOS != "linux" {
		err_msg := fmt.Sprintf("unsupported platform: %s", runtime.GOOS)
		panic(errors.New(err_msg))
	}
	// RunC()
	parseFlags()
}
