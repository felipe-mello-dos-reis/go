package main; import("fmt")

type veiculo struct{
  portas  int
  cor string
}

type caminhonete struct{
  veiculo veiculo
  quatroRodas bool
}

type sedan struct{
  veiculo veiculo
  modeloLuxo  bool
}

func main(){
  carro1 := caminhonete{
    veiculo:  veiculo{
      portas: 4,
      cor:  "azul",
    },
    quatroRodas: true,
  }
  
  carro2 := sedan{veiculo{4,"prata"},false}
  
  fmt.Println("caminhonete:")
  fmt.Println("\tPortas:",carro1.veiculo.portas)
  fmt.Println("\tCor:",carro1.veiculo.cor)
  fmt.Println("\tTração nas 4 rodas:", carro1.quatroRodas)
  fmt.Println("sedan:")
  fmt.Println("\tPortas:",carro2.veiculo.portas)
  fmt.Println("\tCor:",carro2.veiculo.cor)
  fmt.Println("\tModelo de luxo:", carro2.modeloLuxo)
}
