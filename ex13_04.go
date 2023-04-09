package main
import("fmt")

type pessoa struct{
  nome  string
  sobrenome string
  idade int
}

func (p pessoa) apresentação(){
 fmt.Println("Olá, meu nome é",p.nome,p.sobrenome,"e tenho",p.idade,"anos")
} 


func main(){
  pessoa1 := pessoa{
    nome: "Felipe",
    sobrenome: "Mello",
    idade:  20,
  }
  pessoa1.apresentação()
}
