package main; import("fmt")

func main(){
  info := map[string][]string{
    "felipe_mello":[]string{"atletismo","genshin","maromba"},
    "joao_couto":[]string{"ping pong","coito","dirigir"},
    "deodado dados":[]string{"codar", "dormir","comer no rancho"},
  }
  info["marcelo_dedeus"] = []string{"lolzin","animes","garotas 2d"}
  delete(info,"joao_couto")
  for key, value := range(info){
    fmt.Println(key)
    for index,hobbie := range value{
      fmt.Println("\t",index+1,":",hobbie)
    }
    fmt.Println()
  }
}
