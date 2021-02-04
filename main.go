package main

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
	println(i)

	var ptif = 3.14
	println(ptif)

	firstName := "Me"
	println(firstName)

	c := complex(3, 4)
	println(c)

	im, re := imag(c), real(c)

	println(im, re)

	var name *string = new(string)
	*name = "Gabe"
	println(*name)

	namePtr := &firstName
	println(namePtr, *namePtr)

	const pi = 3.14
	print(pi)
	println(first, second, third, fourth, shift)

	//arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	println(arr[0], arr[1], arr[2])

	arr2 := [3]int{1, 2, 3}
	println(arr2[0])
}
