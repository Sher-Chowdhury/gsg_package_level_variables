package main

import "fmt"

func thePointerDatatype() {
	fmt.Println("##### EG11 output #####")

	// A pointer is a variable which points to the memory location of another variable.
	// you can view memory location using '&'
	fullname := "Peter Parker"
	fmt.Println("memory location for 'fullname' variable is:", &fullname)  // memory location for 'fullname' variable is: 0xc0000782d0
	// Think of it like:
	// 'Peter Parker' is the content inside the box
	// &fullname - is the label attached to the box's lid
	// 'fullname' - is a human-friendly name we use to tell golang which box we're talking about. 


	// This initialises a pointer, which just points to memory location.
	// the rhs also ends up declaring this and allocates the memory needed to store a string object.
	var firstname *string = new(string)
	// new is a golang builtin funciton - https://golang.org/pkg/builtin/#new

	// To store anything inside the memory, you use asterisk
	*firstname = "peter"

	// this prints out the memory location
	fmt.Println(firstname) // 0xc0000782e0
	// this prints out what's stored in the memory location itself.
	fmt.Println(*firstname) // peter

	// With this "firstname" pointer, you essentially have a box that's content is just a memory location.
	// This memory-location points to a second box thats created by  the 'new(string)' function. This 2nd box
	// doesn't have a human-friendly name. It's just a box with some content and a memory location written on the
	// box lid. So if you want to see the content of the first box, do:

	// this shows the content of the first box
	fmt.Println(firstname) // 0xc0000782e0
	// this happens to be the what's written on the boxlid of the 2nd box. 
	// So 'firstname' is a humanfriendly way to refer to the first box. 

	// If you want to see the content of the second box that the first box's content is 'pointing' to, 
	// you need to use the special "*" syntax:
	fmt.Println(*firstname) // peter
	
	// This prints the what's written on the box-lid for the first box (memory-location of the pointer itself). 
	fmt.Println(&firstname) // 0xc00000e028

	
	// also see:
	// https://github.com/Sher-Chowdhury/gsg_pointers_datatype

}