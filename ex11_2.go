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
  
  pessoa2 :=pessoa{"Felipe","dos Reis",[]string{"Pistache","Chocomenta"}}


  familia := map[string]pessoa{}
  familia["dos Reis"] = pessoa2
  familia["Mello"] = pessoa1

  for key, value := range familia{
    fmt.Println(key,":")
    for _,sabor := range value.saboresfav{
      fmt.Println("\t",sabor)
    }
  }

}
