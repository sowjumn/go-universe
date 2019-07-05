// Prints the Operating System the code is running on
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch myOs := runtime.GOOS; myOs {
	case "darwin":
		fmt.Println("OS x.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s. \n", myOs)
	}
}
