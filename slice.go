package main
import "fmt"

func main(){
	slice1:= [4]int64{1,2,3,4}
	arr:= slice1[0:4]
	fmt.Println(arr)
}
