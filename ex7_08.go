package main

import("fmt")

func main(){
  x := 4
  switch{
  case x < 0:
    fmt.Println("x é negativo")
  case x > 10:
    fmt.Println("x maior que 10")
  case x < 4 && x > 0:
    fmt.Println("x é positivo, mas menor que 4")
    default:
    fmt.Println("x está entre 4 e 10")
  }
}
