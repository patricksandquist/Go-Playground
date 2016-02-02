/*
  Some mathematics for example!
  Notice how I can write multiline comments, too.
  :-)
*/

package main

import "fmt"
import m "math" // m is the alias for the math package

// This is a constant
const s string = "immutable"

func main() {
  // This is a numerical constant. It has no type until explicitely stated/cast
  const n = 1000
  const d = 3e20 / n
  
  // Here we cast to force a type on the constant
  fmt.Println(int64(d))
  fmt.Println(m.Sin(n))
}
