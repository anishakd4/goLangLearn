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

/*
INTERFACE TYPE LISTS
When generics were released, a new way of writing interfaces was also released at the same time!

We can now simply list a bunch of types to get a new interface/constraint.

// Ordered is a type constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
        ~float32 | ~float64 |
        ~string
}

*/

/*
PARAMETRIC CONSTRAINTS

Your interface definitions, which can later be used as constraints, can accept type parameters as well.

*/

type store[P product] interface{
	Sell(P)
}

type product interface{
	Price() float64
	Name() string
}

type book struct{
	title  string
	author string
	price  float64
}

func (b book) Price() float64 {
	return b.price
}

func (b book) Name() string {
	return fmt.Sprintf("%s by %s", b.title, b.author)
}

type toy struct {
	name  string
	price float64
}

func (t toy) Price() float64 {
	return t.price
}

func (t toy) Name() string {
	return t.name
}

type bookStore struct {
	booksSold []book
}

func (bs *bookStore) Sell(b book) {
	bs.booksSold = append(bs.booksSold, b)
}

type toyStore struct {
	toysSold []toy
}

func (ts *toyStore) Sell(t toy) {
	ts.toysSold = append(ts.toysSold, t)
}

func sellProducts[P product](s store[P], products []P) {
	for _, p := range products {
		s.Sell(p)
	}
}

func testSellProducts() {
	fmt.Println("######testSellProducts######")
	bs := bookStore{
		booksSold: []book{},
	}

    // By passing in "book" as a type parameter, we can use the sellProducts function to sell books in a bookStore
	sellProducts[book](&bs, []book{
		{
			title:  "The Hobbit",
			author: "J.R.R. Tolkien",
			price:  10.0,
		},
		{
			title:  "The Lord of the Rings",
			author: "J.R.R. Tolkien",
			price:  20.0,
		},
	})
	fmt.Println(bs.booksSold)

    // We can then do the same for toys
	ts := toyStore{
		toysSold: []toy{},
	}
	sellProducts[toy](&ts, []toy{
		{
			name:  "Lego",
			price: 10.0,
		},
		{
			name:  "Barbie",
			price: 20.0,
		},
	})
	fmt.Println(ts.toysSold)
	fmt.Println("######testSellProducts######")
}

type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.Plan)
}

type orgBiller struct{
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

func testBiller(){
	b := orgBiller{Plan: "pro"}
	c := org{Admin: user{UserEmail: "fringilla@nilfgaard.com"}, Name: "Nilfgaard"}
	currentBill := b.Charge(c)
	fmt.Println(currentBill)

	d := userBiller{Plan: "pro"}
	e := user{UserEmail: "zoltan@mahakam.com"}
	currentBill = d.Charge(e)
	fmt.Println(currentBill)
}

/*
NAMING GENERIC TYPES

Let's look at this simple example again:

func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}

Remember, T is just a variable name, We could have named the type parameter anything. T happens to be a fairly common 
convention for a type variable, similar to how i is a convention for index variables in loops.

This is just as valid:

func splitAnySlice[MyAnyType any](s []MyAnyType) ([]MyAnyType, []MyAnyType) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
*/

func main() {
	printGetLast()

	testChargeForLineItem()

	testSellProducts()

	testBiller()
}