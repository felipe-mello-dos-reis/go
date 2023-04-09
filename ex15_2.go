package main
import("fmt")

type pessoa struct{
  nome  string
  sobrenome string
}

func mudeMe (p *pessoa, novo_nome string){
  (*p).nome = novo_nome
}

func main(){
  eu := pessoa{"Felipe","Mello"}
  fmt.Println(eu)
  mudeMe(&eu, "Fernanda")
  fmt.Println(eu)

}
