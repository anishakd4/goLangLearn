package main

import (
	"errors"
	"fmt"
	"time"
)

/*
GENERICS IN GO
As we've mentioned, Go does not support classes. For a long time, that meant that Go code couldn't easily be reused in many cases. For example,
imagine some code that splits a slice into 2 equal parts. The code that splits the slice doesn't care about the types of values stored in the
slice. Before generics, we needed to write the same code for each type, which is a very un-DRY thing to do.

func splitIntSlice(s []int) ([]int, []int) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}

func splitStringSlice(s []string) ([]string, []string) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}

TYPE PARAMETERS
Put simply, generics allow us to use variables to refer to specific types. This is an amazing feature because it allows us to write
abstract functions that drastically reduce code duplication.

func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}

In the example above, T is the name of the type parameter for the splitAnySlice function, and we've said that it must match the any constraint,
which means it can be anything. This makes sense because the body of the function doesn't care about the types of things stored in the slice.


firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
fmt.Println(firstInts, secondInts)

*/

func getLast[T any](s []T) T{
	if len(s) == 0 {
		var zeroVal T
		return zeroVal
	}
	return s[len(s)-1]
}

func printGetLast() {
	fmt.Println(getLast([]int{1, 2, 3, 4}))
	fmt.Println(getLast([]int{}))
	fmt.Println(getLast([]string{"a", "b", "c", "d"}))
	fmt.Println(getLast([]bool{true, false, true, true, false}))
}

/*
CONSTRAINTS

Sometimes you need the logic in your generic function to know something about the types it operates on. The example we used in the first 
exercise didn't need to know anything about the types in the slice, so we used the built-in any constraint:

func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}


CREATING A CUSTOM CONSTRAINT

Let's take a look at the example of a concat function. It takes a slice of values and concatenates the values into a string. This should 
work with any type that can represent itself as a string, even if it's not a string under the hood. For example, a user struct can have 
a .String() that returns a string with the user's name and age.

type stringer interface {
    String() string
}

func concat[T stringer](vals []T) string {
    result := ""
    for _, val := range vals {
        // this is where the .String() method
        // is used. That's why we need a more specific
        // constraint instead of the any constraint
        result += val.String()
    }
    return result
}

*/

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error){
	newBalance := balance - newItem.GetCost()
	if(balance < 0.0){
		return nil, 0.0, errors.New("insufficient funds")
	}
	oldItems = append(oldItems, newItem)
	return oldItems, newBalance, nil
}

type lineItem interface{
	GetCost() float64
	GetName() string
}

type subscription struct{
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

func testChargeForLineItem(){
	fmt.Println("######testChargeForLineItem######")
	newItems, newBalance, err := chargeForLineItem(
		lineItem(subscription{
			userEmail: "geralt@rivia.com",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			interval:  "yearly",
		}), 
		[]lineItem{
			subscription{
				userEmail: "yen@vengerberg.com",
				startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				interval:  "monthly",
			},
			oneTimeUsagePlan{
				userEmail:        "triss@maribor",
				numEmailsAllowed: 100,
			},
		},
		1000.00)

	fmt.Println(newItems)
	fmt.Println(newBalance)
	fmt.Println(err)
	fmt.Println("######testChargeForLineItem######")
}

func main() {
	printGetLast()

	testChargeForLineItem()
}