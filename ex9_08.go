package main; import("fmt")

func main(){
  info := map[string][]string{
    "felipe_mello":[]string{"atletismo","genshin","maromba"},
    "joao_couto":[]string{"ping pong","coito","dirigir"},
    "deodado dados":[]string{"codar", "dormir","comer no rancho"},
  }
  for key, value := range(info){
    fmt.Println(key)
    for _,hobbie := range value{
      fmt.Println("\t",hobbie)
    }
  }
}
