package main

import("fmt")
var esporteFavorito string

func main(){
  esporteFavorito = "Volei"
  switch esporteFavorito{
  case "Natação":
    fmt.Println("Vc gosta de nadar")
  case "Futebol":
    fmt.Println("Vc gosta de correr")
  case "Volei":
    fmt.Println("Vc gosta do Cupido pq ele jogando voley de shortinho é irestível")
    default:
    fmt.Println("Vai procurar o que fazer para não morrer de osteoporose")
  }
}
