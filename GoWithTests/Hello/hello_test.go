package gowithtests

import (
	"bytes"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_HelloPass(t *testing.T) {
	repeated := repeat("a")
	actual := "aaaaa"
	assert.Equal(t,actual,repeated,"Actual and repeated should be equal")
	

}

func Test_HelloFail(t *testing.T){
	repeated := repeat("a")
	actual := "aaaaa"

	assert.Equal(t,repeated,actual,"Equal")
}

func Test_Add(t *testing.T){
	sum := Add(4,3)
	expected := 7

	assert.Equal(t,expected,sum,"Sum and expected should be equal")
}


func Test_Array_Sum(t *testing.T){
	num := [5]int{1,2,3,4,5}
	got := Array_Sum(num)
	expected := 15

	assert.Equal(t,expected,got,"Sum and expected should be equal")
}

func Test_Array_Sum_Variations(t *testing.T){

	t.Run("Collection of 5 numbers : ",func(t *testing.T){
		num := [5]int{1,2,3,4,5}
		expected := 15
		sum := Array_Sum(num)
		assert.Equal(t,expected,sum,"Sum and expected should be equal")
	})

	t.Run("Collection of any size : ",func(t *testing.T){
		num1 := []int{1,2,3}
		expected := 6
		sum := Array_Sum2(num1)
		assert.Equal(t,expected,sum,"Sum and expected should be equal")
	})
}

func Test_SumAll(t *testing.T) {

	got := sumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	assert.Equal(t,got,want,"Equality")
}


func Test_Rectangle(t *testing.T){

	t.Run("Perimeter of rectangle ",func (t *testing.T){

	got := Perimeter(2,3)
	want := 10

	assert.Equal(t,got,want,"Perimeter is there correct")
	})

	t.Run("Area of rectangle ",func (t *testing.T){
		got := Area(2,3)
		want := 6

		assert.Equal(t,got,want,"Area is correct")
	})
}

//struct test cases
func Test_Trapezium(t *testing.T){

	t.Run("Perimeter of Trapezium: ", func(t *testing.T){
		trap := Trapezium{12.0,6.0,8.0}
		got := PerimeterOfTrapezium(trap)
		want := 26.0
		assert.Equal(t,got,want,"Equal Perimeter of Trapezium")
	})

	t.Run("Area of Trapezium: ",func(t *testing.T){
		trap := Trapezium{12.0,8.0,3.0}
		got := AreaOfTrapezium(trap)
		want := 30.0
		assert.Equal(t,got,want,"Equal Area of Trapezium")
	})
}

//Test cases for pointers

func Test_Wallet_Deposit(t *testing.T){
	w := Wallet{}
	w.Deposit(120)
	got := w.Balance()
	want := 120
	assert.Equal(t,got,want,"Same as deposited amount")
}

func Test_Wallet_Withdraw(t *testing.T){
	w := Wallet{}
	w.Deposit(120)
	w.Withdraw(100)
	got := w.Balance()
	want := 20
	assert.Equal(t,got,want,"Same as after withdraw")
}

func Test_WordFind(t *testing.T){
	dict := map[string]string{"test":"this is a test"}
	want := "this is a test"
	word := "test"
	got := WordFind(dict,word)

	assert.Equal(t,want,got,"Maps are working fine")
}

func Test_WordFind2(t *testing.T){
	dict := Dictionary{"test":"this is a test"}
	got := dict.Search("test")
	want := "this is a test"
	assert.Equal(t,got,want,"Pass with map")
}

func Test_WordAdd(t *testing.T){
	dict := Dictionary{}
	dict.Add("test","this is a test")
	got,err := dict.Search2("test")
	want := "this is a test"

	if err != nil{
		t.Fatal("Should find added word",got)
	}

	assert.Equal(t,want,got,"Add function in dictionary is working fine")
}

func Test_WordDict(t *testing.T){
	mp1 := map[string]string{"test1" : "this is test1"}
	var word string = "test1"
	got := WordFind(mp1,word)
	want := "this is test1"

	if got != want{
		t.Errorf("got %s want %s ",got,want)
	}
}

func Test_Greet(t *testing.T){
	buffer := bytes.Buffer{}
	name := "Chris"
	Greet(&buffer,name)
	got := buffer.String()
	want := "Hello, Chris"

	assert.Equal(t,got,want,"Dependency Injection")

}

func Test_Countdown(t *testing.T){
	got := Countdown()
	want := `3
2
1
go !`
		
	assert.Equal(t,got,want,"Countdown test passed")

}

func Test_Countdown2(t *testing.T){
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown2(buffer,spySleeper)

	got := buffer.String()

	want := `3
2
1
go !`

	assert.Equal(t,want,got)

	if spySleeper.Calls != 3 {
		t.Errorf("expected 3 but got %d", spySleeper.Calls)
	}
}

func mockWebsiteHelper(url string) bool{
	return url != "www.happy.com"
}

func Test_WebsiteCheck(t *testing.T){
	websites := []string{
		"www.google.com",
		"www.gmail.com",
		"www.razorpay.com",
		"www.moneycontrol.com",
	}

	want := map[string]bool{
		"www.google.com": true,
		"www.gmail.com": true,
		"www.razorpay.com":true,
		"www.moneycontrol.com":true,
	}

	got := CheckWebsites(mockWebsiteHelper,websites)

	assert.Equal(t,got,want,"These mocked are equal")
}

func Test_URL(t *testing.T){
	slowUrl := "www.facebook.com"
	fastUrl := "www.quii.dev"

	want := fastUrl
	got := Racer(slowUrl,fastUrl)

	// assert.Equal(t,want,got,"good")

	if got != want{
		t.Errorf("got %s and want %s",got,want)
	}

}

func Test_Racer(t *testing.T){
	t.Run("Testing something",func(t *testing.T){
		// making slow and fast server
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		//closing the server at the end
		defer slowServer.Close()
		defer fastServer.Close()

		//now making the URL of the server for comparison
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got,_ := Racer2(slowURL,fastURL)

		assert.Equal(t,want,got,"Testing is done for the ends")

	})

	t.Run("Expecting an error ", func(t *testing.T) {
		slowServer := makeDelayedServer(11 * time.Millisecond)
		fastServer := makeDelayedServer(10 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		_ , err := Racer2(slowURL,fastURL)

		if err != nil{
			t.Errorf("Expected an error but did not get any error")
		}
	})
}


func TestWalker(t *testing.T){
	want := "Chris"
	var got []string

	x := struct{
		Name string
	}{want}

	walk(x , func(input string){
		got = append(got,input)
	})

	if len(got) != 1{
		t.Errorf("wrong number of function calls as we want %d and got %d" , 1,len(got))
	}

	if got[0] != want{
		t.Errorf("got %s and want %s",got,want)
	}

}

func TestCounter(t *testing.T){
	t.Run("Checking the logic of counter",func(t* testing.T){
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		finalValue := counter.getValue()
		if finalValue != 3{
			t.Errorf("got %d but wanted %d",finalValue,3)
		}
	})

	t.Run("it checks for concurrency ",func(t *testing.T){
		count := 1000
		counter2 := Counter2{}

		var wg sync.WaitGroup
		wg.Add(count)

		for i := 0;i<count;i++{
			go func(){
				counter2.Inc2()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t,&counter2,count)

	})
}

func assertCounter(t *testing.T,got *Counter2 ,want int){
	t.Helper()
	if got.val != want{
		t.Errorf("got %d and want %d",got,want)
	}
}

func Test_IntToRoman(t *testing.T){

	t.Run("1 is converted to I", func(t *testing.T){
		got := ConvertToRoman(1)
		want := "I"
		if got != want{
			t.Errorf("got %s and want %s",got ,want)
		}
	})

	t.Run("2 is converted to II",func(t *testing.T){
		got := ConvertToRoman(2)
		want := "II"
		if got != want {
			t.Errorf("Got %s and want %s",got,want)
		}
	})	

	t.Run("25 is coverted XV",func(t *testing.T){
		got := ConvertToRoman(25)
		want := "XXV"

		assert.Equal(t,got,want,"so conversion is good ")
	})

	t.Run("37 is coverted to XXXVII",func(t *testing.T){
		got := ConvertToRoman(37)
		want := "XXXVII"
		assert.Equal(t,got,want,"so, this conversion is good")
	})

}

func Test_RomanToInt(t *testing.T){
	t.Run("IV is converted to 4",func(t *testing.T){
		got := convertRomanToArabic("IV")
		want := 4
		assert.Equal(t,got,want)
	})

	t.Run("Converting 77 to LXXVII",func(t *testing.T){
		got := convertRomanToArabic("LXXVII")
		want := 77
		assert.Equal(t,got,want)
	})
}

func Test_Functions(t *testing.T){
	roman := ConvertToRoman(21)
	arabic := convertRomanToArabic(roman)
	check := 21
	assert.Equal(t,roman,"XXI")
	assert.Equal(t,arabic,check)
}


//In this interface helps us to make an easier approach where we can pass any type of data whether it is string , int , bool , etc. -> Generics
func Test_Generics(t *testing.T){
	assertEqual(t,"Hello","Hello")
	assertNotEqual(t,false,2)
}

func assertNotEqual(t *testing.T,want , got interface{}){
	t.Helper()
	if got == want{
		t.Errorf("want %v but got %v",want,got)
	}
}

func assertEqual(t *testing.T , want , got interface{}){
	t.Helper()
	if got != want{
		t.Errorf("want %v but got %v",want,got)
	}
}


func assertTrue(t *testing.T , got bool){
	t.Helper()
	if !got{
		t.Errorf("got %t ",got)
	}
}

func assertFalse(t *testing.T , got bool){
	t.Helper()
	if got{
		t.Errorf("got %t",got)
	}
}

func Test_StackInt(t *testing.T){

	stackofInt := new(StackOfInts)
	assertTrue(t,stackofInt.IsEmpty())

	stackofInt.push(420)
	assertFalse(t,stackofInt.IsEmpty())

	// stackofInt.push(210)
	// stackofInt.push(120)
	stackofInt.pop()
	// stackofInt.push(420)
	assertTrue(t,stackofInt.IsEmpty())


}



































