package main; import("fmt")


func main(){
  estados := make([]string, 26, 26)
  estados = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}
  fmt.Println(len(estados), cap(estados))
  for i:=0;i<len(estados);i++{
    fmt.Println(i, estados[i])
  }

}
