package main

import "fmt"

//Arrays are fixed-size groups of variables of the same type.
//The type [n]T is an array of n values of type T
//var myInts [10]int
//primes := [6]int{2, 3, 5, 7, 11, 13}
func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	return [3]string{primary, secondary, tertiary},[3]int{5, 9, 6}
}

func printMessageWithRetries(primary, secondary, tertiary string){
	x,y :=getMessageWithRetries(primary, secondary, tertiary)
	fmt.Println(x)
	fmt.Println(y)
}

//Slices
//A slice is a dynamically-sized, flexible view of the elements of an array.
//Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array. 
//If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller,
func printSlices(){
	fmt.Println("#####printSlices#####")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println(primes[0: 4])
	fmt.Println(primes[2: 5])
	fmt.Println(primes[1: 5])
	fmt.Println(primes[:])
	fmt.Println("#####printSlices#####")
}

/*Most of the time we don't need to think about the underlying array of a slice. We can create a new slice using the make function:
// func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
mySlice := make([]int, 5)

Slices created with make will be filled with the zero value of the type.

If we want to create a slice with a specific set of values, we can use a slice literal:

mySlice := []string{"I", "love", "go"}

Note that the array brackets do not have a 3 in them. If they did, you'd have an array instead of a slice.

LENGTH
The length of a slice is simply the number of elements it contains. It is accessed using the built-in len() function:

mySlice := []string{"I", "love", "go"}
fmt.Println(len(mySlice)) // 3

CAPACITY
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. It is accessed 
using the built-in cap() function:

mySlice := []string{"I", "love", "go"}
fmt.Println(cap(mySlice)) // 3

Generally speaking, unless you're hyper-optimizing the memory usage of your program, you don't need to worry about the capacity of a slice 
because it will automatically grow as needed.
*/
func getMessageCosts(messages []string) []float64 {
	costs :=  make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
        costs[i] = float64(len(messages[i])) * 0.01
    }
	return costs
}

func printMessageCosts(messages []string){
    x := getMessageCosts(messages)
    fmt.Println(x)
}

/*
LEN AND CAP REVIEW
The length of a slice may be changed as long as it still fits within the limits of the underlying array

The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume. Here is a function 
to append data to a slice. If the data exceeds the capacity, the slice is reallocated.

The function uses the fact that len and cap are legal when applied to the nil slice, and return 0.
*/

//VARIADIC

func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}

func checkVariadic(){
	final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
}

//SPREAD OPERATOR
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func checkSpreadOperator(){
	names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}

/*
APPEND
The built-in append function is used to dynamically add elements to a slice:

If the underlying array is not large enough, append() will create a new underlying array and point the slice to it.
*/
type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	fmt.Println(len(costsByDay))
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		fmt.Println(len(costsByDay))
        for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
    }
	return costsByDay
}

func checkAppendFunction(){
	fmt.Println("######checkAppendFunction######")
	costs := []cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{5, 2.5},
		{2, 3.6},
		{1, 2.7},
		{1, 3.3},
	}
	y := getCostsByDay(costs)
	fmt.Println(y)
	fmt.Println("######checkAppendFunction######")
}

/*SLICE OF SLICES
Slices can hold other slices, effectively creating a matrix, or a 2D slice.
rows := [][]int{}
*/
func createMatrix(rows, cols int) [][]int{
	matrix := make([][]int, 0)
    for i := 0; i < rows; i++ {
        row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
    }
    return matrix
}

func printMatrix(rows, cols int){
	fmt.Println("######printMatrix######")
	y := createMatrix(rows, cols)
	fmt.Println(y)
	fmt.Println("######printMatrix######")
}

/*
TRICKY SLICES
The append() function changes the underlying array of its parameter AND returns a new slice. This means that using append() on anything 
other than itself is usually a BAD idea.

// dont do this!
someSlice = append(otherSlice, element)
*/

//Range
func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

func printBadWordsIndex(msg []string, badWords []string) {
	fmt.Println("######printBadWordsIndex######")
	x := indexOfFirstBadWord(msg, badWords)
	fmt.Println(x)
	fmt.Println("######printBadWordsIndex######")
}


func main() {
	printMessageWithRetries("Anish", "Kumar", "Dubey")
	printSlices()

	printMessageCosts([]string{"Welcome to the movies!", "Enjoy your popcorn!"})

	checkVariadic()

	checkSpreadOperator()

	checkAppendFunction()

	printMatrix(10, 10)

	printBadWordsIndex([]string{"ugh", "oh", "my", "frick"}, []string{"crap", "shoot", "frick", "dang"})
}