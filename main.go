package main

import (
	"fmt"
	"os"

	"github.com/hongminhcbg/mint-gen/service"
)

func gen(serviceName string) {
	err := os.Mkdir(serviceName, 0777)
	if err != nil {
		panic(err)
	}

	gen := service.New("templates/service-gin-template", serviceName, serviceName)
	err = gen.Gen(serviceName)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Hello world")
	fmt.Println("Only for unix")
	gen("sms")
}
