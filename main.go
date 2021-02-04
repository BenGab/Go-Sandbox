package main

func main() {
	var i int
	i = 42
	println(i)

	var pi float32 = 3.14
	println(pi)

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
}
