package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var str = "100"
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(i, err, reflect.TypeOf(i))

}
