package main

import (
	"fmt"
	"reflect"
)

/*

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // under the hood byte is just alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

generally speaking numbers fall into 4 different buckets.

We have integers, unsigned integers, floats and complex numbers.

uint8 and uint16 both represent unsigned integers but uint16 has twice as much room for data within it. uint16 has 16 bits whereas uint8 has
only 8 bits
*/

/*
VARIABLE DECLARATION

Variables are declared using the var keyword. For example, to declare a variable called number of type int, you would write:

var number int

To declare a variable called pi to be of type float64 with a value of 3.14159, you would write:

var pi float64 = 3.14159

The value of a declared variable with no assignment will be its zero value.

*/

func variableDeclaration(){
	var num int = 0
	fmt.Println(num)

	var hasPermission bool = false
	fmt.Println(hasPermission)

	var myName string = "Anish"
	fmt.Println(myName)

	var floatValue float64
	fmt.Println(floatValue)
}

/*
SHORT VARIABLE DECLARATION

the := short assignment statement can be used in place of a var declaration. The := operator infers the type of the new variable 
based on the value. It's called the walrus operator because it looks like a walrus... sort of.

These two lines of code are equivalent:

var empty string
empty := ""


numCars := 10 // inferred as an integer. it can be int32 or int64 depending on your computer's architecture
temperature := 0.0 // temperature is inferred as a float because it has a decimal
var isFunny = true // inferred as a boolean
*/

func shortVariableDeclaration(){
	congrats := "Happy birthday mutki"
	fmt.Println(congrats)
}

/*
TYPE INFERENCE
*/

func typeInference(){
	x := 56.0
	fmt.Printf("type inference is %T\n", x)
	fmt.Println(reflect.TypeOf(x))
}

/*
SAME LINE DECLARATIONS
You can declare multiple variables on the same line:

mileage, company := 80276, "Tesla"

The above is the same as:

mileage := 80276
company := "Tesla"

*/

func sameLineDeclarations(){
    mileage, company := 80276, "Tesla"
    fmt.Println(mileage)
    fmt.Println(company)
}

/*
Some types can be converted like this:

temperatureFloat := 88.26
temperatureInt := int64(temperatureFloat)

Casting a float to an integer in this way truncates the floating point portion.
*/

func typeConversion(){
    temperatureFloat := 88.26
    temperatureInt := int64(temperatureFloat)
    fmt.Println(temperatureInt)
}

/*
CONSTANTS

Constants are declared with the const keyword. They can't use the := short declaration syntax.

const pi = 3.14159

Constants can be character, string, boolean, or numeric values. They can not be more complex types like slices, maps and structs, 
which are types we will explain later.

As the name implies, the value of a constant can't be changed after it has been declared
*/

/*
COMPUTED CONSTANTS
Constants must be known at compile time. They are usually declared with a static value:

const myInt = 15

However, constants can be computed as long as the computation can happen at compile time.

For example, this is valid:

const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName

That said, you cannot declare a constant that can only be computed at run-time like you can in JavaScript. This breaks:

// the current time can only be known when the program is running
const currentTime = time.Now()
*/

func constantVariable() {
	const premiumPlanName = "premium"
	const basicPlanName = "basic"

	//premiumPlanName = "Sdfsdf" //const variable can not be re-assigned. //cannot assign to premiumPlanName (neither addressable nor a map index expression)

	fmt.Println(premiumPlanName)
	fmt.Println(basicPlanName)
}

/*
FORMATTING STRINGS IN GO

Go follows the printf tradition from the C language.

fmt.Printf - Prints a formatted string to standard out.
fmt.Sprintf() - Returns the formatted string

DEFAULT REPRESENTATION
The %v variant prints the Go syntax representation of a value, it's a nice default.

s := fmt.Sprintf("I am %v years old", 10)
// I am 10 years old

s := fmt.Sprintf("I am %v years old", "way too many")
// I am way too many years old

If you want to print in a more specific way, you can use the following formatting verbs:

STRING
s := fmt.Sprintf("I am %s years old", "way too many")
// I am way too many years old

INTEGER
s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old

FLOAT
s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old
// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old


*/

func formattingString(){
	const name = "anish"
	const openrate = 3.4

	msg := fmt.Sprintf("%s %.1f", name, openrate)

	fmt.Println(msg)
}

func main() {
	variableDeclaration()
	shortVariableDeclaration()
	typeInference()
	sameLineDeclarations()
	typeConversion()
	constantVariable()
	formattingString()
}