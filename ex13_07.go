package main
import("fmt")

func main(){
  magica := func (x int) int{
    return x*x
  }
  fmt.Println(magica(7))
}
