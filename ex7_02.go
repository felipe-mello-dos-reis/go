package main

import("fmt")

func main(){
  for i := 65; i <= 65+25; i++{
    fmt.Println(i)
    for x:=0;x<3;x++{
     fmt.Printf("\t%#U\n", i)
   }
  }
}
