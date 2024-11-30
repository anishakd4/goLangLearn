package main

import "fmt"

/*
Every file of a go program has a package declaration at the top
Here we have a package main simple because this program builds into a executable go program meaning we can run this code as a stand-alone.
*/

/*
main function is a entry point to a go program. main does not take any parameters and does not return anything
*/
func main() {
	fmt.Println("Anish kumar")
}


/*
Any go code is a human readable code.

Our CPU only understands binary

So we need some process which will convert human readable code to  machine code that can be executed by the computer's hardware and that's all compilation is.

main.go -> go build -> executable file

this executable can be run directly on the operating system without ever having to use the go tool chain again. 

This is very different from python where we have to use python interpreter each time.

With an interpreted language as we run the program, the interpreter is reading the human code and kind of converting this to machine code at runtime

CPU is designed buy the manufacturer to run a specific format of binary.

Go is strongly and statically typed.

A string variable like "hello world" can not be changed to an int, such as the number 3.

When we are talking about the performance of a programming language or an application, then we really care about how it performs across 2 
different axes, 1 is speed how fast it can do computations which is measured in CPU cycles and then we also have memory consumption which 
is how bloaty program is, how much data it has to store in memory to be able to do such computations.

In programming languages like Rust, C memory management is effectively manual.

Java is a garbage collected language which essentially means memory management is automated, and in JAVA is done by java virtual machine.
Every time we run a java program, we are actually creating a mini virtual machine that your java byte code runs within and JVM is what 
takes care of allocating and de-allocating all the memory we use.This create a overhead meaning java programs use quite a bit of more 
memory than Rust or C program.

GO is in a interesting sort of in-between world where go is a garbage collected language like java so it has a automated memory management, but it doesn't really have any JVM. GO executable runs, just like Rust and C we get 1 binary. The difference is that go includes a runtime within every single binary built in go programming language. We can think of it as kind of like a sidecar that is compiled alongside your code. So your code has this a little of extra code that is added to it and that little code is what handles garbage collection and automated memory management. So its a little bit more bloaty than Rust and C but it is not nearly as expensive in terms of memory overhead as a language like java or C-Sharp.

GO PROGRAMS ARE EASY ON MEMORY
Go programs are fairly lightweight. Each program includes a small amount of "extra" code that's included in the executable binary. 
This extra code is called the Go Runtime. One of the purposes of the Go runtime is to clean up unused memory at runtime.

In other words, the Go compiler includes a small amount of extra logic in every Go program to make it easier for developers 
to write code that's memory efficient.

COMPARISON
As a general rule, Java programs use more memory than comparable Go programs. There are several reasons for this, but one of 
them is that Java uses an entire virtual machine to interpret bytecode at runtime. Go programs are compiled into machine code, 
and the overhead of the Go runtime is typically less than the overhead of the Java virtual machine.

On the other hand, Rust and C programs use slightly less memory than Go programs because more control is given to the developer 
to optimize the memory usage of the program. The Go runtime just handles it for us automatically.


*/