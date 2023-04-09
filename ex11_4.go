package main
import("fmt")

func main(){
  aluno := struct{
    notas map[string]float32
    materias_preferidas []string
  }{
    notas:  map[string]float32{"Matemática": 9.5, "Física": 8.9, "Filosofia": 7.2},
    materias_preferidas:  []string{"Matemática","Física"},
  }
  fmt.Println(aluno)
}
