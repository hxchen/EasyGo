package main

func main() {
	var a int = 20
	var ip *int

	ip = &a

	println("Address of a variable: ", &a)
	println("Address stored in ip variable: ", ip)
	println("Value of *ip variable: ", *ip)
}
