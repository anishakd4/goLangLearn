package main

import "fmt"

/*
Every file of a go program has a package declaration at the top
Here we have a package main simple because this program builds into a executable go program meaning we can this code as a stand-alone.
*/

/*
main function is a entry point to a go program. main does not take any parameters and does return anything
*/
func main() {
	fmt.Println("Anish kumar")
}


/*
Any go code is a human readable code.

Our CPU only understands binary

So we need some process which will convert human readable code to  machine code that can be executed by the computer's hardware and that's all 
compilation is.

main.go -> go build -> executable file

this executable can be run directly on the operating system without ever having to use the go tool chain again, we can run this executable 
on another system without using go tool chain again. This is very different from python where we have to use python interpreter each time.

With an interpreted language as we run the program, the interpreter is reading the human code and kind of converting this to machine code at 
runtime

CPU is designed buy the manufacturer to run a specific format of binary.
*/