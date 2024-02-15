package main
 
import "fmt"
 
func main() {
	defer fmt.Println("Bye!")

	for ;true; {
		fmt.Println("Select operation:")
		fmt.Println("1) A + B")
		fmt.Println("2) A - B")
		fmt.Println("3) A * B")
		fmt.Println("4) A / B")

		switch InputNumber() {
		case 1:
			DoOperation(Add)
		case 2:
			DoOperation(Subtract)
		case 3:
			DoOperation(Multiply)
		case 4:
			DoOperation(Divide)
		default:
			continue;
		}
	}
}

func DoOperation(operation func(int, int) int){
	fmt.Println("Input A")
	a:=InputNumber()
	fmt.Println("Input B")
	b:=InputNumber()
	fmt.Printf("Result: %d \n",operation(a,b))
}
func InputNumber() int{
	var a int;
	a, err := fmt.Scan(&a)
	if err!=nil{
		return InputNumber()
	} else {
		return a;
	}
}
func Add(a int, b int) int{
	return a+b
}
func Subtract(a int, b int) int{
	return a-b
}
func Multiply(a int, b int) int{
	return a*b
}
func Divide(a int, b int) int{
	return a/b
}