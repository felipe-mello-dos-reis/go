package main
import("fmt")

func main(){
  for i:=1; i<5; i++{
    if i==1{
      defer fmt.Println(i)
    } else{
      fmt.Println(i)
    }
  }
}
