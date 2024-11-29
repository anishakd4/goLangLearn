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

/*
ANONYMOUS STRUCTS IN GO
An anonymous struct is just like a normal struct, but it is defined without a name and therefore cannot be referenced elsewhere in the code.

To create an anonymous struct, just instantiate the instance immediately using a second pair of brackets after declaring the type:

myCar := struct {
	maker string
	model string
} {
	maker: "tesla",
	model: "model 3",
}

You can even nest anonymous structs as fields within other structs:

type car struct {
	maker string
	model string
	doors int
	mileage int
	// wheel is a field containing an anonymous struct
	wheel struct {
		radius int
		material string
	}
}

WHEN SHOULD YOU USE AN ANONYMOUS STRUCT?
In general, prefer named structs. Named structs make it easier to read and understand your code, and they have the nice side-effect 
of being reusable. I sometimes use anonymous structs when I know I won't ever need to use a struct again. For example, sometimes 
I'll use one to create the shape of some JSON data in HTTP handlers.

If a struct is only meant to be used once, then it makes sense to declare it in such a way that developers down the road won’t be 
tempted to accidentally use it again.
*/

type car2 struct {
	maker bool
	model string
	doors int
	mileage int
	mecar struct{
		makerr int
		modell int
	}
}

func anonymousStruct(){
	myCar := struct {
		maker string
		model string
	} {
		maker: "tesla",
		model: "model 3",
	}

	fmt.Println(myCar.maker)
	fmt.Println(myCar.model)

	myCar2 := car2{}
	fmt.Println(myCar2.maker)
	fmt.Println(myCar2.mecar.makerr)
}

/*
Go is not an object-oriented language. However, embedded structs provide a kind of data-only inheritance that can be 
useful at times. Keep in mind, Go doesn't support classes or inheritance in the complete sense, 
but embedded structs are a way to elevate and share fields between struct definitions.
EMBEDDED STRUCTS

EMBEDDED VS NESTED
Unlike nested structs, an embedded struct's fields are accessed at the top level like normal fields.
Like nested structs, you assign the promoted fields with the embedded struct in a composite literal.
lanesTruck := truck{
	bedSize: 10,
	car: car{
		maker: "toyota",
		model: "camry",
	},
}

fmt.Println(lanesTruck.bedSize)
fmt.Println(lanesTruck.maker)
fmt.Println(lanesTruck.model)
*/


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

/*
STRUCT METHODS IN GO
While Go is not object-oriented, it does support methods that can be defined on structs. Methods are just 
functions that have a receiver. A receiver is a special parameter that syntactically goes before the name of the function.
A receiver is just a special kind of function parameter. Receivers are important because they will, as you'll 
learn in the exercises to come, allow us to define interfaces that our structs (and other types) can implement.
*/
type rect struct {
	width int
	height int
}

// area has a receiver of (r rect)
func (r rect) area() int {
	return r.width * r.height
}

func getAreaRect(){
	var r = rect{
		width: 5,
		height: 10,
	}

	fmt.Println(r.area())
}

func main() {

	setCar()
	x := canSendMessage(messageToSend{
		message:   "you have an birthday tomorrow",
		sender:    user{name: "Jason Bjorn", number: 16545550987},
		recipient: user{name: "Jim Bond"},
	})
	fmt.Println(x)

	fmt.Println("----anonymousStruct----")
	anonymousStruct()
	fmt.Println("----anonymousStruct----")

	//Unlike nested structs, an embedded struct's fields are accessed at the top level like normal fields.
	output := getSenderLog(sender{
		rateLimit: 45,
		user1: user1{
			name: "Anish",
			number: 80,
		},
	})
	fmt.Printf(output)

	getAreaRect()

}