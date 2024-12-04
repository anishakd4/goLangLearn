package main

import (
	"fmt"
	"strings"
)

/*
Pointers

As we have learned, a variable is a named location in memory that stores a value. We can manipulate
the value of a variable by assigning a new value to it or by performing operations on it. When we assign a
value to a variable, we are storing that value in a specific location in memory.

x := 42
// "x" is the name of a location in memory. That location is storing the integer value of 42

A POINTER IS A VARIABLE
A pointer is a variable that stores the memory address of another variable. This means that a pointer "points to"
the location of where the data is stored NOT the actual data itself.

The * syntax defines a pointer:
var p *int

A pointer's zero value is nil

The & operator generates a pointer to its operand.

myString := "hello"
myStringPtr := &myString

WHY ARE POINTERS USEFUL?
Pointers allow us to manipulate data in memory directly, without making copies or duplicating data. This can make programs
more efficient and allow us to do things that would be difficult or impossible without them.

The * dereferences a pointer to gain access to the value

fmt.Println(*myStringPtr) // read myString through the pointer
*myStringPtr = "world"    // set myString through the pointer

*/

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	messageValue := *message
	messageValue = strings.ReplaceAll(messageValue, "fubb", "****")
	messageValue = strings.ReplaceAll(messageValue, "shiz", "****")
	messageValue = strings.ReplaceAll(messageValue, "witch", "*****")
	*message = messageValue
}

func testRemoveProfanity(){
	m := "English, motherfubber, do you speak it?"
	removeProfanity(&m)
	fmt.Println(m)

	m = "Oh man I've seen some crazy ass shiz in my time..."
	removeProfanity(&m)
	fmt.Println(m)

	m = "Does he look like a witch?"
	removeProfanity(&m)
	fmt.Println(m)
}

/*NIL POINTERS 
If a pointer points to nothing (the zero value of the pointer type) then dereferencing it will cause a runtime error (a panic) that 
crashes the program. Generally speaking, whenever you're dealing with pointers you should check if it's nil before trying to dereference it.
*/

func testRemoveProfanityNilCase(){
	removeProfanity(nil)
}

//POINTER RECEIVERS
/*
A receiver type on a method can be a pointer.
Methods with pointer receivers can modify the value to which the receiver points. Since methods often need to modify their receiver, pointer receivers are more common than value receivers. However, methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.
*/
type car struct {
	color string
}

func (c *car) setColor(color string){
	c.color = color
}

func (c car) setColor2(color string) {
	c.color = color
}

func testSetColor(){
	c := car{
		color: "white",
	}
	fmt.Println(c)
	c.setColor("blue")
	fmt.Println(c)
	c.setColor2("green")
	fmt.Println(c)
}

func main(){
	testRemoveProfanity()

	testRemoveProfanityNilCase()

	testSetColor()
}