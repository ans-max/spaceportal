package main

import (
	"fmt"
	"github.com/ans-max/spaceportal/apod"
)

func main() {
	fmt.Println("Starting APOD server at port :9090")
	apod.StartApod(":9090")
}
