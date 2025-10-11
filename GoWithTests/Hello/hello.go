package gowithtests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)


func repeat(char string) string{

	var final string;

	for i:=0;i<5;i++{
		final = final + char
	}
	return final
}

func Add1(x,y int)int{
	return 0
}

func Add(x,y int)int{
	return x+y
}

func Array_Sum(numbers [5]int)int{

	sum := 0
	for _,number := range numbers{
		sum += number
	}
	return sum
}

func Array_Sum2(num[]int) int{
	sum := 0
	for _,number := range num{
		sum += number
	}
	return sum
}

func sumAll(array ...[]int)[] int{
	lengthOfArray := len(array)
	sums := make([]int, lengthOfArray)

	for i, numbers := range array{
		sums[i] = Array_Sum2(numbers)
	}
	return sums
}

func Perimeter(l int, b int) int{
	return 2*(l+b)
}

func Area(l int,b int) int{
	return l*b
}

type Trapezium struct{
	parallel1 float64
	parallel2 float64
	height float64
}

func PerimeterOfTrapezium(t Trapezium) float64{
	return t.parallel1 + t.parallel2 + t.height
}

func AreaOfTrapezium(t Trapezium) float64{
	return (0.5)*(t.parallel1 + t.parallel2)* t.height
}

type Wallet struct{
	balance int
}

func (w *Wallet) Deposit(amount int){
	w.balance += amount
}

func (w *Wallet) Balance() int{
	return w.balance
}

func (w *Wallet) Withdraw(amount int){
	w.balance -= amount
}

func WordFind(dict map[string]string,word string)string{
	return dict[word]
}

type Dictionary map[string]string

func (d Dictionary) Search(word string)string{
	return d[word]
}

func (d Dictionary) Search2(word string)(string,error){
	return d[word],nil
}

func (d Dictionary) Add(word , definition string){
	d[word] = definition
}

func Greet(writer *bytes.Buffer,name string){
	fmt.Fprintf(writer, "Hello, %s", name)
}


func Countdown()string{
	var num int = 3
	var str string
	for i:=num;i>=0;i--{
		if i==0{
			str += "go !"
		}else{
			str += strconv.Itoa(i)+ "\n"
		}
		
	}
	return str
}

const countDown = 3
const finalWord = "go !"


type Sleeper interface{
	Sleep()
}

type SpySleeper struct{
	Calls int
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}

func Countdown2(out io.Writer, sleeper Sleeper) {
	for i := countDown; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := countDown; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, finalWord)
}

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker , urls []string) map[string]bool {
	res := make(map[string]bool)
	for _, url := range urls{
		res[url] = wc(url)
	}
	return res
}

func Racer(a,b string) (winner string){
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration{
		return a
	}
	return b
}