package large_number
//import "fmt"
func largestoftwonumbers(a int, b int) int{
if a > b {
return a
} else if b > a{
return b
} else if b == a {
return a
} else {
return 0
}
}

/* func main(){
ans := largestoftwonumbers(10, 20)
fmt.Println("the largest number is ", ans)
}
*/
/* func main(){
var a int
var b int
fmt.Println("Enter the number", a)
fmt.Scan(&a)
fmt.Println("enter the number", b)
fmt.Scan(&b)
if a > b {
   fmt.Println("A is the largest", a)
} else if b > a {
   fmt.Println("B is the largest", b)
} else if a == b {
   fmt.Println("Both the values are equal", a)
} else {
   fmt.Println("Invalid input")
 }
}
*/
