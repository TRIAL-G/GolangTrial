package main
import "fmt"
func main(){
	var Name string
	fmt.Println("Enter your name: ")
	fmt.Scan(&Name)
	fmt.Printf("Welcome %v to the Conference\n",Name)
}
