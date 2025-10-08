package gowithtests

import (
	"testing"
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

























