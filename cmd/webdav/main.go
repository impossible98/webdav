package main

import (
	// import built-in packages
	"fmt"
	// import local packages
	"webdav/constant"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(constant.APP_NAME + ": " + constant.APP_VERSION)
}
