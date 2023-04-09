package main
import("fmt")

func pot(n int) func(x int) int{
  magica := func(x int) int{
    i := 1
    for j := 0; j < n; j++{
      i *= x
    }
    return i
  }
  return magica
}


func main(){
  quadrado := pot(2)
  cubo := pot(3)
  fmt.Println("Quadrado de 8:", quadrado(8))
  fmt.Println("Cubo de 8:", cubo(8))
}
