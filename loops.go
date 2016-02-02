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
  
  /*
    If/else if/else conditionals are similar to Javascript in Go.
    But like for loops, they don't require parenthesis.
  */
  
  if 7 % 2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }
  
  // We can also declare variables before/in a conditional:
  if num := 9; num < 0 {
    fmt.Println(num, " is negative")
  } else if num < 10 {
    fmt.Println(num, " has 1 digit")
  } else {
    fmt.Println(num, " has multiple digits")
  }
  
  // Switches are pretty similar
  switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
      fmt.Println("it's the weekend")
    default:
      fmt.Println("it's a weekday")
  }
  
  // But may also be written in some different ways:
  t := time.Now()
  switch {
    case t.Hour() < 12:
      fmt.Println("it's before noon")
    default:
      fmt.Println("it's after noon")
  }
}
