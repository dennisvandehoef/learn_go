package main

import "fmt"

func main() {
	var a int = 20 /* actual variable declaration */
	var ip *int    /* pointer variable declaration */

	ip = &a /* store address of a in pointer variable*/

	fmt.Printf("Address of a variable: %x\n", &a)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	/* access the value using the pointer */
	fmt.Printf("Value of a variable: %d\n", a)
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	// nil pointer
	var ptr *int
	if ptr == nil {
		fmt.Printf("The ptr == nil, and its value is: %x\n", ptr)
	} else {
		// We will not reach this code
		fmt.Printf("The value of ptr is: %x\n", ptr)
	}
}
