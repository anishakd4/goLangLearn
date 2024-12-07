package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
CONCURRENCY

Concurrency is the ability to perform multiple tasks at the same time. Typically, our code is executed one line
at a time, one after the other. This is called sequential execution or synchronous execution.

If the computer we're running our code on has multiple cores, we can even execute multiple tasks at exactly the same
time. If we're running on a single core, a single core executes code at almost the same time by switching between tasks very
quickly. Either way, the code we write looks the same in Go and takes advantage of whatever resources are available.

HOW DOES CONCURRENCY WORK IN GO?
Go was designed to be concurrent, which is a trait fairly unique to Go. It excels at performing many tasks simultaneously
safely using a simple syntax.

Concurrency is as simple as using the go keyword when calling a function:

go doSomething()

In the example above, doSomething() will be executed concurrently with the rest of the code in the function. The go keyword
is used to spawn a new goroutine. Al we are saying using the go keyword is that this execution can happen in parallel.

*/

func sendEmail(message string){
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received %s\n", message)
	}()
	fmt.Printf("Email sent %s\n", message)
}

func printSendEmail(){
	sendEmail("Hello there Kaladin!")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

/*
CHANNELS

Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.

CREATE A CHANNEL

Like maps and slices, channels must be created before use. They also use the same make keyword:

ch := make(chan int)

SEND DATA TO A CHANNEL

ch <- 69

The <- operator is called the channel operator. Data flows in the direction of the arrow. This operation will block until 
another goroutine is ready to receive the value.

RECEIVE DATA FROM A CHANNEL

v := <-ch

This reads and removes a value from the channel and saves it into the variable v. This operation will block until there is a value 
in the channel to be read

BLOCKING AND DEADLOCKS
A deadlock is when a group of goroutines are all blocking so none of them can continue. This is a common bug that you need to watch
out for in concurrent programming.

*/

type email struct{
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	fmt.Println(isOld)
	return isOld
}

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

func printCheckEmailAge(){
	checkEmailAge([3]email{
		{
			body: "I have stolen princesses back from sleeping barrow kings.",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I burned down the town of Trebon",
			date: time.Date(2019, 6, 6, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I have spent the night with Felurian and left with both my sanity and my life.",
			date: time.Date(2022, 7, 0, 0, 0, 0, 0, time.UTC),
		},
	})
}

/*
Empty structs are often used as tokens in Go programs. In this context, a token is a unary value. In other words, we don't 
care what is passed through the channel. We care when and if it is passed.

We can block and wait until something is sent on a channel using the following syntax

<-ch

This will block until it pops a single item off the channel, then continue, discarding the item.
*/

func waitForDBs(numDBs int, dbChan chan struct{}){
	for i:=0; i<numDBs; i++{
		<-dbChan
	}
}

func getDBsChannel(numDBs int)(chan struct{}, int){
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, count
}

func printDBOnline(){
	dbChan, _ := getDBsChannel(3)
	waitForDBs(3, dbChan)
}

/*
BUFFERED CHANNELS
Channels can optionally be buffered.

CREATING A CHANNEL WITH A BUFFER
You can provide a buffer length as the second argument to make() to create a buffered channel:

ch := make(chan int, 100)

Sending on a buffered channel only blocks when the buffer is full.

Receiving blocks only when the buffer is empty.
*/

func addEmailsToQueue(emails []string) chan string{
	ch := make(chan string, len(emails))
	for i:=0; i<len(emails); i++ {
		ch <- emails[i]
	}
	return ch
}

func sendEmails(){
	fmt.Println("#####sendEmails#####")
	ch := addEmailsToQueue([]string{
		"It's life, Jim, but not as we know it.",
		"Infinite diversity in infinite combinations.",
		"Engage!",
	})
	fmt.Println(len(ch))
	fmt.Println("#####sendEmails#####")
}

/*
CLOSING CHANNELS IN GO
Channels can be explicitly closed by a sender:

ch := make(chan int)
// do some stuff with the channel
close(ch)

CHECKING IF A CHANNEL IS CLOSED
receivers can check the ok value when receiving from a channel to test if a channel was closed.

v, ok := <-ch

ok is false if the channel is empty and closed.

DON'T SEND ON A CLOSED CHANNEL
Sending on a closed channel will cause a panic. A panic on the main goroutine will cause the entire program to crash, 
and a panic in any other goroutine will cause that goroutine to crash.

Closing isn't necessary. There's nothing wrong with leaving channels open, they'll still be garbage collected if they're 
unused. You should close channels to indicate explicitly to a receiver that nothing else is going to come across.
*/

func countReports(numSentCh chan int) int {
	total := 0
	for{
		numSent, ok := <-numSentCh
		if !ok{
			break
		}

		total += numSent
	}
	return total
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

func testSendReports(){
	numSentCh := make(chan int)
	go sendReports(4, numSentCh)
	output := countReports(numSentCh)
	fmt.Println(output)
}

/*
RANGE
Similar to slices and maps, channels can be ranged over.

for item := range ch {
    // item is the next value received from the channel
}

*/

func concurrentFib(n int){
	chInts := make(chan int)
	go func(){
		fibonacci(n, chInts)
	}()

	for v := range chInts {
		fmt.Println(v)
	} 
}


func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func testFibonacci(){
	fmt.Println("######testFibonacci####")
	concurrentFib(7)
	fmt.Println("######testFibonacci####")
}

/*
SELECT
Sometimes we have a single goroutine listening to multiple channels and want to process data in the order it comes through each channel.

A select statement is used to listen to multiple channels at the same time. It is similar to a switch statement but for channels.

select {
case i, ok := <- chInts:
    fmt.Println(i)
case s, ok := <- chStrings:
    fmt.Println(s)
}

The first channel with a value ready to be received will fire and its body will execute. If multiple channels are ready at the same time 
one is chosen randomly. The ok variable in the example above refers to whether or not the channel has been closed by the sender yet.
*/

func logMessages(chEmails, chSms chan string) {
	for {
		select{
			case email, ok := <- chEmails:
                if!ok{
                    return
                }
                fmt.Println(email)
            case sms, ok := <- chSms:
                if!ok{
                    return
                }
                fmt.Println(sms)
		}
	}
}

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
	chSms = make(chan string)
	chEmails = make(chan string)
	go func() {
		for i := 0; i < len(sms) && i < len(emails); i++ {
			done := make(chan struct{})
			s := sms[i]
			e := emails[i]
			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
			go func() {
				time.Sleep(t1)
				chSms <- s
				done <- struct{}{}
			}()
			go func() {
				time.Sleep(t2)
				chEmails <- e
				done <- struct{}{}
			}()
			<-done
			<-done
			time.Sleep(10 * time.Millisecond)
		}
		close(chSms)
		close(chEmails)
	}()
	return chSms, chEmails
}

func testLogMessages(){
	fmt.Println("#####testLogMessages######")
	chSms, chEmails := sendToLogger([]string{
		"hi friend",
		"What's going on?",
		"Welcome to the business",
		"I'll pay you to be my friend",
	},
	[]string{
		"Will you make your appointment?",
		"Let's be friends",
		"What are you doing?",
		"I can't believe you've done this.",
	})
	logMessages(chEmails, chSms)
	fmt.Println("#####testLogMessages######")
}

/*
SELECT DEFAULT CASE

The default case in a select statement executes immediately if no other channel has a value ready. A default case stops 
the select statement from blocking.

select {
case v := <-ch:
    // use v
default:
    // receiving from ch would block
    // so do something else
}

TICKERS
time.Tick() is a standard library function that returns a channel that sends a value on a given interval.
time.After() sends a value once after the duration has passed.
time.Sleep() blocks the current goroutine for the specified amount of time.

READ-ONLY CHANNELS
A channel can be marked as read-only by casting it from a chan to a <-chan type. For example:

func main() {
    ch := make(chan int)
    readCh(ch)
}

func readCh(ch <-chan int) {
    // ch can only be read from
    // in this function
}

WRITE-ONLY CHANNELS
The same goes for write-only channels, but the arrow's position moves.

func writeCh(ch chan<- int) {
    // ch can only be written to
    // in this function
}

*/

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string){
	for{
		select{
			case <-snapshotTicker:
				takeSnapshot(logChan)
			case <-saveAfter:
				saveSnapshot(logChan)
				return
			default:
				waitForData(logChan)
				time.Sleep(time.Millisecond * 500)
		}
		
	}
}

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}

func testSaveBackup(){
	fmt.Println("######testSaveBackup######")
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	logChan := make(chan string)
	go saveBackups(snapshotTicker, saveAfter, logChan)

	for actualLog := range logChan {
		fmt.Println(actualLog)
	}

	fmt.Println("######testSaveBackup######")
}

/*
CHANNELS REVIEW

A SEND TO A NIL CHANNEL BLOCKS FOREVER

var c chan string // c is nil
c <- "let's get started" // blocks

A RECEIVE FROM A NIL CHANNEL BLOCKS FOREVER

var c chan string // c is nil
fmt.Println(<-c) // blocks

A SEND TO A CLOSED CHANNEL PANICS

var c = make(chan int, 100)
close(c)
c <- 1 // panic: send on closed channel

A RECEIVE FROM A CLOSED CHANNEL RETURNS THE ZERO VALUE IMMEDIATELY

var c = make(chan int, 100)
close(c)
fmt.Println(<-c) // 0

*/

func main() {
	printSendEmail()
	printCheckEmailAge()
	printDBOnline()

	sendEmails()

	testSendReports()

	testFibonacci()

	testLogMessages()

	testSaveBackup()
}