package main
import("fmt")

func main(){
  x := 5
  x = func (x int) int{
    return x*x
  }(x)
  fmt.Println(x)
}
