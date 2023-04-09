package main
import("fmt")

type pessoa struct{
  nome  string
  sobrenome string
  saboresfav []string
}
func main(){
  pessoa1 := pessoa{
    nome: "Fernanda",
    sobrenome:  "Mello",
    saboresfav: []string{"Morango","Chocolate","Creme"},
  }
  
  pessoa2 :=pessoa{"Felipe","Mello",[]string{"Pistache","Chocomenta"}}

  familia := []pessoa{pessoa1,pessoa2}
  for _,membro := range familia{
      fmt.Println(membro.nome,":")
      for _,sabor := range membro.saboresfav{
        fmt.Println("\t",sabor)
      }
    }
}
