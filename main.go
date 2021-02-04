package main

import (
	"fmt"
	"my-go-sandbox/models"
)

const (
	first  = 1
	second = iota
	third
	fourth = iota + 6
	shift  = 2 << iota
)

func main() {
	var i int
	i = 42
	fmt.Println(i)

	var ptif = 3.14
	fmt.Println(ptif)

	firstName := "Me"
	fmt.Println(firstName)

	c := complex(3, 4)
	fmt.Println(c)

	im, re := imag(c), real(c)

	fmt.Println(im, re)

	var name *string = new(string)
	*name = "Gabe"
	fmt.Println(*name)

	namePtr := &firstName
	fmt.Println(namePtr, *namePtr)

	const pi = 3.14
	fmt.Println(pi)
	fmt.Println(first, second, third, fourth, shift)

	//arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	//slices
	slice := []int{1, 2, 3}

	slice = append(slice, 4, 5, 6)

	s1 := slice[1:]
	fmt.Println(s1)

	//map
	m := map[string]int{"foo": 45}
	fmt.Println(m)

	var u models.User
	u.ID = 1
	u.Firstname = "ARt"
	u.LastName = "sda"
	fmt.Println(u)

	u2 := models.User{2,
		"dasd",
		"sadas",
	}
	fmt.Println(u2)
}

func startWebServer(port int) (int, error) {
	return port, nil
}
