package main

import (
  "fmt"
  "imartins/modulesimple/simplemodule/simplemath"
)

func main(){
  fmt.Println("access my module...")
  var c int
  c = simplemath.Add(5,8)
  fmt.Printf("add()=%d\n", c)

  c = simplemath.Subtract(5,8)
  fmt.Printf("Sub()=%d\n", c)

  c = simplemath.Multiply(5,3)
  fmt.Printf("mult()=%d\n", c)
}
