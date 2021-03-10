package main

import (

	"fmt"
	"github.com/gothello/go-scanner/port"
	"strconv"
)


func main() {

	//wait := make(chan bool, 50)
	Gen()
}

func Gen() {
	for i := 0; i < 999; i++ {
		array := []rune("177.72.81.")

		str := strconv.Itoa(i)
		r := []rune(str)

		for _, value := range r {

			array = append(array, value)

			fmt.Println(string(array))

			port.Run(string(array))
		}
		
	}
}

