package main

import (
	"fmt"
	"io"
	"time"
)

/*
Interfaces are just collections of method signatures. A type "implements" an interface if it has methods that match
the interface's method signatures.
When a type implements an interface, it can then be used as that interface type.

INTERFACES ARE IMPLEMENTED IMPLICITLY

A type implements an interface by implementing its methods. Unlike in many other languages,
there is no explicit declaration of intent, there is no "implements" keyword.
A type can implement any number of interfaces in Go. For example, the empty interface, interface{}, is always
implemented by every type because it has no requirements.
*/
type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func sendMessage(msg message){
	x:= msg.getMessage()
	fmt.Println(x)
}

/*
NAME YOUR INTERFACE ARGUMENTS
Consider the following interface:

type Copier interface {
	Copy(string, string) int
}

Based on the code alone, can you deduce what kinds of strings you should pass into the Copy function?

We know the function signature expects 2 string types, but what are they? Filenames? URLs? Raw string data? For that matter, 
what the heck is that int that's being returned?

Let's add some named arguments and return data to make it more clear.

type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

Much better. We can see what the expectations are now. The first argument is the sourceFile, the second argument is the 
destinationFile, and bytesCopied, an integer, is returned.
*/

//TYPE ASSERTIONS IN GO
type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()
	}
	s, ok := e.(sms)
	if ok {
		return s.toPhoneNumber, s.cost()
	}
	return "", 0.0
}

func printExpenseReport(e expense) {
	x, y := getExpenseReport(e)
	fmt.Println(x, y)
}

/*
TYPE SWITCHES
A type switch makes it easy to do several type assertions in a series.
*/
func getExpenseReport2(e expense) (string, float64) {
	switch v:= e.(type) {
		case email:
            return v.toAddress, v.cost()
        case sms:
            return v.toPhoneNumber, v.cost()
        default:
            return "", 0.0
	}
}

func printExpenseReport2(e expense) {
	x, y := getExpenseReport(e)
	fmt.Println(x, y)
}

func main() {
	sendMessage(birthdayMessage{time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC), "John Doe"})
	sendMessage(sendingReport{"First Report", 10})
	printExpenseReport(email{
		isSubscribed: false,
		body:         "It is I, Arthur, son of Uther Pendragon, from the castle of Camelot. King of the Britons, defeator 
		of the Saxons, sovereign of all England!",
		toAddress:    "soldier@monty.com",
	})
	printExpenseReport(sms{
		isSubscribed:  true,
		body:          "I am. And this my trusty servant Patsy.",
		toPhoneNumber: "+155555509832",
	})

	printExpenseReport2(email{
		isSubscribed: false,
		body:         "It is I, Arthur, son of Uther Pendragon, from the castle of Camelot. King of the Britons, defeator 
		of the Saxons, sovereign of all England!",
		toAddress:    "soldier@monty.com",
	})
	printExpenseReport2(sms{
		isSubscribed:  true,
		body:          "I am. And this my trusty servant Patsy.",
		toPhoneNumber: "+155555509832",
	})
}

/*
1. KEEP INTERFACES SMALL
If there is only one piece of advice that you take away from this article, make it this: keep interfaces small! Interfaces are meant to define 
the minimal behavior necessary to accurately represent an idea or concept.

2. INTERFACES SHOULD HAVE NO KNOWLEDGE OF SATISFYING TYPES
An interface should define what is necessary for other types to classify as a member of that interface. They shouldn’t be aware of any 
types that happen to satisfy the interface at design time.

For example, let’s assume we are building an interface to describe the components necessary to define a car.

type car interface {
	Color() string
	Speed() int
	IsFiretruck() bool
}
Copy icon
Color() and Speed() make perfect sense, they are methods confined to the scope of a car. IsFiretruck() is an anti-pattern. We are 
forcing all cars to declare whether or not they are firetrucks. In order for this pattern to make any amount of sense, we would 
need a whole list of possible subtypes. IsPickup(), IsSedan(), IsTank()… where does it end??

3. INTERFACES ARE NOT CLASSES
Interfaces are not classes, they are slimmer.
Interfaces don’t have constructors or deconstructors that require that data is created or destroyed.
Interfaces aren’t hierarchical by nature, though there is syntactic sugar to create interfaces that happen 
to be supersets of other interfaces.
Interfaces define function signatures, but not underlying behavior. Making an interface often won’t DRY up your 
code in regards to struct methods. For example, if five types satisfy the fmt.Stringer interface, they all need 
their own version of the String() function.
*/