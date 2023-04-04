package main

import("fmt")

var diadasemana string

func main(){
  diadasemana = "Sexta"
  if diadasemana == "Sexta"{
    fmt.Println("Sextou bebe")
  } else if diadasemana == "Sábado"{
    fmt.Println("Corre que ainda dá tempo de ir para a balada")
  } else {
    fmt.Println("Infelizmente vai ter que ficar em casa")
  }
}
