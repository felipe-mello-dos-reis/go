package main
import(
  "fmt"
  "math"
)

type quadrado struct{
  lado  float64
}
type círculo struct{
  raio  float64
}

func (l quadrado) área() float64{return l.lado*l.lado}
func (r círculo) área() float64{return math.Pi*r.raio*r.raio}

type figura interface{
  área() float64
}

func info(f figura) float64{
  return f.área()
}

func main(){
  figura1 := quadrado{12.0}
  figura2 := círculo{1.0}
  fmt.Println(info(figura1))
  fmt.Println(info(figura2))
}
