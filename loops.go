/*
  Go only has one way of looping through things: for.
  There are three different ways of writing for, though.
*/

package main
import "fmt"
func main() {
  // The first is the basic conditional
  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }
  
  // The other is the classic initial; condition; increment
  for j := 1; j <= 3; j++ {
    fmt.Println(j)
  }
  
  // The last is an infinite loop
  for {
    fmt.Println("again!")
    break
  }
}
