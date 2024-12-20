package functions

import "fmt"

//Functions in Go can take zero or more arguments
func concat(s1 string, s2 string) string {
	return s1 + s2
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

/*
MULTIPLE PARAMETERS
When multiple arguments are of the same type, and are next to each other in the function signature, the type only needs to 
be declared after the last argument.

Here are some examples:

func addToDatabase(hp, damage int) {
  // ...
}

func addToDatabase(hp, damage int, name string) {
  // ?
}

func addToDatabase(hp, damage int, name string, level int) {
  // ?
}
*/
func incrementSends(sendsSoFar, sendsToAdd int) int{
	sendsSoFar += sendsToAdd
	return sendsSoFar
}

//in go variables are passed by value not by reference
func passingByValue(){
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println(sendsSoFar)
}

/*
Go supports the ability to return early from a function. This is a powerful feature that can clean up code, especially when used as guard clauses.

Guard Clauses leverage the ability to return early from a function (or continue through a loop) to make nested conditionals one-dimensional. 

Instead of using if/else chains, we just return early from the function at the end of each conditional block.

in case of multiple return values from a function we need to wrap all return values in parentheses
*/
func getProductInfo(tier string) (string, string, string) {
	if tier == "basic" {
		return "1,000 texts per month", "$30 per month", "most popular"
	} else if tier == "premium" {
		return "50,000 texts per month", "$60 per month", "best value"
	} else if tier == "enterprise" {
		return "unlimited texts per month", "$100 per month", "customizable"
	} else {
		return "", "", ""
	}
}

/*
A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore: _
The Go compiler will throw an error if you have any unused variable declarations in your code, 
so you need to ignore anything you don't intend to use.
underscore is not a conventional name we ignore. Compiler completely removes it from the code
*/
func getProductMessage(tier string) string {
	quantityMsg, priceMsg, _ := getProductInfo(tier)
	return "You get " + quantityMsg + " for " + priceMsg + "."
}

/*
NAMED RETURN VALUES
func getCoords() (x, y int){
  // x and y are initialized with zero values

  return // automatically returns x and y
}
Copy icon
Is the same as:

func getCoords() (int, int){
	var x int
	var y int
	return x, y
}

Named return values are mostly used if we have multiple return values from the same function and we want to document the return values.
*/

func getCoords() (x, y int){
	//by naming x and y are initialized with zero values
	x = x * 5
	y = y * 8
	return // automatically returns x and y
}

//Even though a function has named return values, we can still explicitly return values if we want to.
//EXPLICIT RETURNS
func getCoords2() (x, y int){
	return x, y // this is explicit
}

//EXPLICIT RETURNS
func getCoords3() (x, y int){
	return 5, 6 // this is explicit, x and y are NOT returned
}

func main() {
    test("Lane,", " happy birthday!")
	passingByValue()
	getProductMessage("basic")
	getCoords()
	getCoords2()
	getCoords3()
}