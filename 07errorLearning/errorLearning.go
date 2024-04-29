package main

import (
	"errors"
	"fmt"
)

/*
An Error is any type that implements the simple built-in error interface:

type error interface {
    Error() string
}
*/
func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costForCustomer, err := sendSMS(msgToCustomer)
	if err!= nil {
        return 0, err
    }
	costForSpouse, err := sendSMS(msgToSpouse)
	if err!= nil {
        return 0, err
    }
	return costForCustomer + costForSpouse, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func printSMSCost(msgToCustomer, msgToSpouse string){
	ans, _ := sendSMSToCouple(msgToCustomer, msgToSpouse)
	fmt.Println(ans)
}

/*
FORMATTING STRINGS REVIEW
A convenient way to format strings in Go is by using the standard library's fmt.Sprintf() function. 
*/
func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs %.2f to be sent to '%v'can not be sent", cost, recipient)
}

func printSMSErrorString(cost float64, recipient string){
	fmt.Println(getSMSErrorString(cost, recipient))
}

/*
THE ERROR INTERFACE
Because errors are just interfaces, you can build your own custom types that implement the error interface.
*/

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
    return fmt.Sprintf("can't divide by %v", de.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func printDivide(x, y float64){
	ans, err := divide(x, y)
    if err!= nil {
        fmt.Println(err)
        return
    }
    fmt.Println(ans)
}

/*
THE ERRORS PACKAGE
The Go standard library provides an "errors" package that makes it easy to deal with errors.

var err error = errors.New("something went wrong")
*/


func divide2(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("No dividing by zero")
	}
	return x / y, nil
}

func printDivide2(x, y float64){
	ans, err := divide2(x, y)
    if err!= nil {
        fmt.Println(err)
    } else {
        fmt.Println(ans)
    }
}


func main(){
	printSMSCost("Thanks for coming in to our flower shop today!", "We hope you enjoyed your gift.")
	printSMSCost("Thanks for joining us!", "Have a good day.")
	printSMSErrorString(32.1, "+1 (801) 555 7456")

	printDivide(10, 0)
	printDivide(15, 0)

	printDivide2(10, 0)
	printDivide2(15, 0)
}