package main

import "fmt"

func main() {
  // Concatinating strings like a boss
  fmt.Println("Patrick" + " " + "Sandquist" + " knows Go!")
  
  // Arithmetic
  fmt.Println("1 + 1 =", 1 + 1)
  fmt.Println("7.0 / 3.0 =", 7.0 / 3.0)
  fmt.Println("7 / 3 =", 7 / 3)
  
  // Booleans and whatnot
  fmt.Println(true && false)
  fmt.Println(true || false)
  fmt.Println(!true)
  
  // Declaring some variables and their types
  var a string = "testing!"
  
  // Two at a time!
  var b, c int = 3, 4
  fmt.Println(a)
  fmt.Println(b + c)
  
  // Can be lazy and leave off type
  var hungry = true
  
  // Or can leave off value
  var e int
  fmt.Println(e) // -> 0
  
  // But let's stick to shorthand:
  f := "speedy"
  fmt.Println(f)
}
