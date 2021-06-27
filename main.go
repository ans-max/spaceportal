package main

import (
	"fmt"
	//"io/ioutil"
	//"log"

	"github.com/ans-max/spaceportal/apod"
)

func main() {
	//	files, err := ioutil.ReadDir("/app")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	for _, f := range files {
	//		fmt.Println(f.Name())
	//	}
	fmt.Println("Starting APOD server at port :9090")
	apod.StartApod(":9090")
}
