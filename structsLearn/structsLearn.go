package main

import "fmt"

//Structs can be nested to represent more complex entities.
//We use structs in Go to represent structured data. It's often convenient to group different types of variables together.
//For example, if we want to represent a car we could do the following:

type car struct {
	maker string
	model string
	doors int
	mileage int
	frontWheel wheel
	backWheel wheel
}

type wheel struct {
	radius int
	material string
}

func setCar() {
	//all the values inside of the struct will be initialized with their default values
	myCar := car{}
	myCar.frontWheel.radius = 5
}

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true
}

//Go is not an object-oriented language. However, embedded structs provide a kind of data-only inheritance that can be 
//useful at times. Keep in mind, Go doesn't support classes or inheritance in the complete sense, 
//but embedded structs are a way to elevate and share fields between struct definitions.
//EMBEDDED STRUCTS
type sender struct {
	rateLimit int
	user1
}

type user1 struct {
	name   string
	number int
}

func getSenderLog(s sender) string {
	return fmt.Sprintf(`
====================================
Sender name: %v
Sender number: %v
Sender rateLimit: %v
====================================
`, s.name, s.number, s.rateLimit)
}


func main() {
	setCar()
	x := canSendMessage(messageToSend{
		message:   "you have an birthday tomorrow",
		sender:    user{name: "Jason Bjorn", number: 16545550987},
		recipient: user{name: "Jim Bond"},
	})
	fmt.Println(x)

	output := getSenderLog(sender{
		rateLimit: 45,
		user1: user1{
			name: "Anish",
			number: 80,
		},
	})
	fmt.Printf(output)
}