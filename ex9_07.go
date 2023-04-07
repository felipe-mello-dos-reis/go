package main; import("fmt")

func main(){
  ss := [][]string{
          []string{"Felipe","Mello","Atletismo"},
          []string{"Joao","Couto","Ping Pong"},
          []string{"Deodado","Dados","Codar"},
  }
  for _, v := range ss{
    fmt.Println(v[0])
    for _, item := range v{
      fmt.Println("\t",item)
    }
  }
}
