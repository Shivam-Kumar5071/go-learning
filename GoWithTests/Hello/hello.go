package gowithtests


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

