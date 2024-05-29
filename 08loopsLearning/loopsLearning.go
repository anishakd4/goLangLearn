package main

import "fmt"

func bulkSend(numMessages int) float64 {
	totalCost := 0.0

	for i := 0; i < numMessages; i++ {
        totalCost += 1.0 + (0.01*float64(i))
    }

	return totalCost
}

func printTotalCost(numMessages int) {
	x := bulkSend(numMessages)

	fmt.Println(x)
}

//OMITTING CONDITIONS FROM A FOR LOOP IN GO
//Loops in Go can omit sections of a for loop. For example, the CONDITION (middle part) can be omitted which causes the loop to run forever.
func maxMessages(thresh float64) int {
	totalCost := 0.0

	for i:=0 ; ; i++ {
		totalCost += 1.0 + (0.01*float64(i))

        if totalCost > thresh {
            return i
        }
	}
}

func printMaxMessages(thresh float64) {
	y := maxMessages(thresh)
	fmt.Println(y)
}

//THERE IS NO WHILE LOOP IN GO
//Because Go allows for the omission of sections of a for loop, a while loop is just a for loop that only has a CONDITION.
func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0
	for actualCostInPennies <= float64(maxCostInPennies) {
		actualCostInPennies *= costMultiplier
		maxMessagesToSend++
	}
	return maxMessagesToSend
}

func printMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) {
    y := getMaxMessagesToSend(costMultiplier, maxCostInPennies)
    fmt.Println(y)
}

func fizzbuzz(maxInt int){
	fmt.Println("##########FizzBuzz###########")
	for i := 1; i <= maxInt; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
	fmt.Println("##########FizzBuzz###########")
}

func printPrimes(max int) {
	fmt.Println("##########printPrimes###########")
	for n:=2; n<max; n++ {
		if n==2 {
			fmt.Println(n)
			continue
		}
		if n % 2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i <= n; i += 2 {
			if n % i == 0 {
                isPrime = false
                break
            }
		}
		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
	fmt.Println("##########printPrimes###########")
}

func main() {
	printTotalCost(4)
	printMaxMessages(10)
	printMaxMessagesToSend(1.1, 5)
	printMaxMessagesToSend(1.3, 9)

	fizzbuzz(10)
	printPrimes(10)
}