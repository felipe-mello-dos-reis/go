package main
import("fmt")

func devolver() func() int{
  x := 0
  return func() int{
    x++
    return x
  }
}


func main(){
  i := devolver()
  fmt.Println(i())
  fmt.Println(i())
  fmt.Println(i())
  a := devolver()
  fmt.Println(a())
  fmt.Println(a())
  fmt.Println(a())
  fmt.Println(a())
  fmt.Println(a())
}

