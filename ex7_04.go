package main

import("fmt")

var x int

func main(){
  x = 2002
  for {
    if x < 2024{
      fmt.Println(x)
    } else {
      break
    }
    x++
  }
}
