package main
import("fmt")

func soma(x ...int) int{
  val := 0
  for _,v := range x {
    val += v
  }
  return val
}

func soma_slice(x []int) int{
  val := 0
  for _,v := range x{
    val += v
  }
  return val
}

func main(){
  slice := []int{1,2,3,4,5,6,7,8,9,10}
  fmt.Println(soma(slice...))
  fmt.Println(soma_slice(slice))
}
