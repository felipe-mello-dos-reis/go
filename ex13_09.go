package main
import("fmt")

func textolegal() {
  fmt.Println("Este texto é muito massa")
}

func ouviralgolegal(f func()){
  fmt.Println("Lá vem coisa boa:")
  f()
}

func main(){
  ouviralgolegal(textolegal)
}
