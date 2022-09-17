package main

import (
	"errors"
	"flag"
	"fmt"
	"godevil/src/app"
	"os"
	"runtime"
)

func setFlags() {
	var intvar = flag.Int("num", 0, "some integer")
	// var listvar = flag.String("st", "unknown", "some string")
	fmt.Printf("Args: %d\n", *intvar)
}

func cleanup() {
	return_code := app.EXIT_SUCCESS_CODE
	if r := recover(); r != nil {
		return_code = app.EXIT_ERROR_CODE
		fmt.Printf("%v\n", r)
	}
	os.Exit(return_code)
}

func main() {
	defer cleanup()

	if runtime.GOOS != "linux" {
		err_msg := fmt.Sprintf("unsupported platform: %s", runtime.GOOS)
		panic(errors.New(err_msg))
	}

	setFlags()
}
