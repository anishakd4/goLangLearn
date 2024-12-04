package main

import "fmt"

/*
A programming language is said to have "first-class functions" when functions in that language are treated like any other variable.
For example, in such a language, a function can be passed as an argument to other functions, can be returned by another function and
can be assigned as a value to a variable.

first-class functions means those functions which can be passed as arguments.
A function that returns a function or accepts a function as input is called a Higher-Order Function.
*/
func add(a, b int) int{
	return (a+b)
}

func mul(a, b int) int{
	return (a*b)
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int{
	return arithmetic(arithmetic(a,b), c)
}

func printAggregate(a, b, c int, arithmetic func(int, int) int){
	x := aggregate(a, b, c, arithmetic)
	fmt.Println(x)
}

/*CURRYING
Function currying is the practice of writing a function that takes a function (or functions) as input, and returns a new function.
*/
func multiplyCurry(x, y int) int {
	return x * y
}

func addCurry(x, y int) int {
	return x + y
}

func selfMath(mathFunc func(int, int) int) func (int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

func printSelfMath(){
    squareFunc := selfMath(multiplyCurry)
	doubleFunc := selfMath(addCurry)
	
	fmt.Println(squareFunc(5))

	fmt.Println(doubleFunc(5))
}

//DEFER
/*
The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically 
just before its enclosing function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until 
the surrounding function returns.

Deferred functions are typically used to close database connections, file handlers and the like.

Defer is a great way to make sure that something happens at the end of a function, even if there are multiple return statements.
*/
const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

type user struct {
	name   string
	number int
	admin  bool
}

func logAndDelete(users map[string]user, name string) (log string){
	user,ok := users[name]
	if !ok {
		delete(users, name)
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	delete(users, name)
	return logDeleted
}

func printLogAndDelete(){
	mp := map[string]user{
		"laura": {name: "laura", number: 4355556023, admin: false},
		"dale":  {name: "dale", number: 8015558937, admin: true},
	}
	logAndDelete(mp, "laura")

	fmt.Println(mp)
}

/*
CLOSURES
A closure is a function that references variables from outside its own function body. The function may access 
and assign to the referenced variables.
*/
func concatter() func(string) string {
	doc := ""
	return func(str string) string {
		doc += str + " "
        return doc
	}
}

func printConcatter(){
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
}

/*
ANONYMOUS FUNCTIONS
Anonymous functions are true to form in that they have no name.
*/
func doMath(f func(int) int, nums []int) []int {
	var results []int
	for _, n := range nums {
		results = append(results, f(n))
	}
	return results
}

func printDoMath(){
	nums := []int{1, 2, 3, 4, 5}
	
    // Here we define an anonymous function that doubles an int
    // and pass it to doMath
	allNumsDoubled := doMath(func(x int) int {
		return x + x
	}, nums)
	
	fmt.Println(allNumsDoubled)
}

func main(){
	printAggregate(2,3,4, add)
	printAggregate(2,3,4, mul)

	printSelfMath()

	printLogAndDelete()

	printConcatter()

	printDoMath()
}