package gowithtests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
	"sync"
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

func Racer2(a,b string) (winner string,err error){
	select{
	case <-ping(a):
		return a,nil
	case <-ping(b):
		return b,nil
	case <-time.After(10 * time.Second):
		return "",fmt.Errorf("waiting for %s and %s",a,b)
	}
}

func ping(a string) chan struct{}{
	ch := make(chan struct{})
	go func ()  {
		http.Get(a)
		close(ch)
	}()
	return ch
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}


func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}

type Counter struct{
	val int
}

func(c *Counter) Inc(){
	c.val++
}

func (c *Counter) getValue() int{
	return c.val
}

type Counter2 struct{
	sync.Mutex
	val int
}

func(c *Counter2) Inc2(){
	c.Lock()
	defer c.Unlock()
	c.val++

}
type RomanNumber struct{
	value int
	romanVal string
}

var AllRomanVal = []RomanNumber{
	{10,"X"},
	{9,"IX"},
	{5,"V"},
	{4,"IV"},
	{1,"I"},
}

func ConvertToRoman(num int) string{
	var result strings.Builder

	for _,numeralVal := range AllRomanVal{
		for num >= numeralVal.value{
			result.WriteString(numeralVal.romanVal)
			num -= numeralVal.value
		}
	}
	return result.String()
}

var AllRomanVal2 = []RomanNumber{
	{100,"C"},
	{90,"XC"},
	{50,"L"},
	{40,"XL"},
	{10,"X"},
	{9,"IX"},
	{5,"V"},
	{4,"IV"},
	{1,"I"},
}

func convertRomanToArabic(rom string)int{

	var result = 0
	for _,roman := range AllRomanVal2{
		for strings.HasPrefix(rom,roman.romanVal){
			result += roman.value
			rom = strings.TrimPrefix(rom,roman.romanVal)
		}
	}
	return result
}