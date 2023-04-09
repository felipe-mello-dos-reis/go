package main
import("fmt")

func funcao1() int{
  return 1
}

func funcao2() (int, string){
  return 1, "deu bom"
}

func main(){
  fmt.Println(funcao1())
  fmt.Println(funcao2())
}
