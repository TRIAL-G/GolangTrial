//feet to metres 
package main
import "fmt"
func main(){
	var feet float64
	feet=123
	var metres float64
	metres= feet* 0.3048
	fmt.Println(metres)
}
//or

package main
import "fmt"

func main(){
    fmt.Println("enter feet")
    var input float64
    fmt.Scanf("%f",&input)
    output:= input / 0.3048
    fmt.Println(output)

}