package main

import "fmt"

func main(){
	var marks int
	fmt.Println("Enter the marks:")
	fmt.Scan(&marks)
	switch marks{
	case 80:
	fmt.Println("Your grade is A")
	case 70:
	fmt.Println("Your grade is B")
	case 60:
	fmt.Println("Your grade is C")
	case 50:
	fmt.Println("Your grade is C")
	default:
	fmt.Println("Your grade is D")
	return
}
}
