package main
import "fmt"
func main() {
    for i := 1; i <= 100; i++ {
   if i % 3 == 0 {
            fmt.Println(i, "divisible by 3")
        } else if  i % 5==0 {
            fmt.Println(i, " divisible by 5")
        } else if i % 4==0 {
        	fmt.Println(i, "divisible by 4")
        }
    }
}
