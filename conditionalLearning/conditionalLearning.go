package main

import "fmt"

/*
if statements in Go do not use parentheses around the condition:

if height > 4 {
    fmt.Println("You are tall enough!")
}

else if and else are supported as you might expect:

if height > 6 {
    fmt.Println("You are super tall!")
} else if height > 4 {
    fmt.Println("You are tall enough!")
} else {
    fmt.Println("You are not tall enough!")
}

*/
func ifLearning(){
	messageLen := 10
	maxMessageLen := 20

	if messageLen > maxMessageLen {
        fmt.Println("Message sent")
    } else {
        fmt.Println("Message not sent")
    }
}

/*
This is just some syntactic sugar that Go offers to shorten up code in some cases. For example, instead of writing:

length := getLength(email)
if length < 1 {
    fmt.Println("Email is invalid")
}

We can do:

if length := getLength(email); length < 1 {
    fmt.Println("Email is invalid")
}

Not only is this code a bit shorter, but it also removes length from the parent scope, which is convenient because we don't need it 
there - we only need access to it while checking a condition.
*/


func main() {
	ifLearning()
}