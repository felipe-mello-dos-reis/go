// Cap 2.2

- Package main diz o pacote principal
- Notação é sempre a mesma: pacote.identificador
- Em go, vocÊ não pode declarar uma variável e não usar ela
- Se vc nao quiser uma variavel, só descartar o valor com um _ (blank identifer)

package main

import (
  "fmt"
)
// aqui é onde vc comeca e onde vc termina
func main() {
  numrodebytes, _ = fmt.Println("Hello World")
  fmt.Println(numerodebytes)
}


// Cap 2.3

- Operador curto de declaração
- := parece uma marmota (mascote do go gopher)
  - serve para declarar variáveis com tipagem automática
  - Palavras reservadas (keywords): package, func, break, case, default, go, var, type, select, map, else, for, range, return, continue, const
  - operador de comparação, retorna um bool: ==, <, >
  - uma expressão gera um reultado, mas não uma ação
  - declaração (statement)

package main

import (
  "fmt"
)

// package level scope

var y := "olá bom dia"
func main() {
  // Escopo apenas dentro do bloco, a marmota só funciona dentro de um code block

  x := 10

  // %v fala do valor, enquanto %T fala do tipo
  fmt.Println("x : %v, %T\n", x, x)
  fmt.Println("y : %v, %T\n", y, y)

  // Se quiser mudar o valor da sua variável, vc deve usar apenas o =

  x = 20
  
  // Mas se vc for declarar pelo menos uma e tribuir outrasm, vc pode usar a marmota

  x, z := 30, 40
}


// Cap 2.4



package main

import (
  "fmt"
)

var z := 20

func main() {
  y := 10
  qualquercoisa(y)
}

func qualquercoisa(x int) {
  fmt.Println(x)
  // Funciona fmt.Println(z)
  // Não funciona fmt.Println(y)
}


// Cap 2.5

- Tipos são estáticos, ou seja, não é possível mudar

package main

import (
  "fmt"
)

var x = 10
var y int
// Não funciona y = 5

func main() {
  // Não funciona x = 20.5
  // Atribuição de valor de uma variavel já declarada sem valor, só pode ser feita dentro de um code block
  y = 15
}

// Cap 2.6

- Inicialização é a primeira vez que vc coloca um valor dentro da variável
- Atribuição é quando você edita o valor da varívael
- Valor zero é o valor que a variável tem antes de ser inicializada
  - ints: 0
  - floats: 0.0
  - booleans: false
  - strings: ""
  - pointers, functions, interfaces, slices, channels, maps: nil

// Cap 2.7

- Literal em ciência da computação, literal é uma notação para representar um valor fixo no código fonte
- Int, string, bool são literals
- Para strings, vc tem os raws e interpreted
- Format printing: documentação.
    - Grupo #1: Print → standard out // serve para exibir algo na tela
        - func Print(a ...interface{}) (n int, err error)
        - func Println(a ...interface{}) (n int, err error) // colona um \n no final
        - func Printf(format string, a ...interface{}) (n int, err error)
            - Format verbs. (%v %T)
    - Grupo #2: Print → string, pode ser usado como variável //retorna uma string e não coloca nada na tela
        - func Sprint(a ...interface{}) string // não adiciona espaco entre os operandos
        - func Sprintf(format string, a ...interface{}) string
        - func Sprintln(a ...interface{}) string
    - Grupo #3: Print → file, writer interface, e.g. arquivo ou resposta de servidor
        - func Fprint(w io.Writer, a ...interface{}) (n int, err error)
        - func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
        - func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

package main

import{
  "fmt"
}

func main() {
  // \" \n, \t são interpretados
  x := "oi, bom dia\ncomo vai?\tespero \"que\" tudo bem."
  fmt.Println(x)
  
  // Se quiser que seja tudo raw, do jeito que vc escrever, use ``

  y := `"oi, bom dia\ncomo vai?\tespero \"que\" 
  tudo bem."`
  fmt.Println(y)

  x1 := "oi"
  x2 := "bom dia"
  x3 := fmt.Sprint(x1,x2)
}


// Cap 2.8

- Quando criamos um novo tipo, podemos usar de maneiras diferentes
import main

import{
  "fmt"
}

type hotdog int
var b hotdog = 10

func main() {
  x := 10
  fmt.Printf("%v\n", x)
  fmt.Printf("%v", b)

  // Não funciona b = x
  // No capítulo seguinte veremos como fazer CONVERSÃO
}


// Cap 2.9

- Para contorna o problema do capítulo anterior, é possível corrigir
- x = int(b)

// Cap 3.1

- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
    1. 42
    2. "James Bond"
    3. true

// Cap 3.2

- Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
    1. Identificador "x" deverá ter tipo int
    2. Identificador "y" deverá ter tipo string
    3. Identificador "z" deverá ter tipo bool
- Na função main:
    1. Demonstre os valores de cada identificador
    2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?

// Cap 3.3

- Utilizando a solução do exercício anterior:
    1. Em package-level scope, atribua os seguintes valores às variáveis:
        1. para "x" atribua 42
        2. para "y" atribua "James Bond"
        3. para "z" atribua true
    2. Na função main:
        1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
        2. Demonstre a variável "s".


// Cap 3.4

- Crie um tipo. O tipo subjacente deve ser int.
- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
- Na função main:
    1. Demonstre o valor da variável "x"
    2. Demonstre o tipo da variável "x"
    3. Atribua 42 à variável "x" utilizando o operador "="
    4. Demonstre o valor da variável "x"


// Cap 3.5
- Utilizando a solução do exercício anterior:
    1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
    2. Na função main:
        1. Isto já deve estar feito:
            1. Demonstre o valor da variável "x"
            2. Demonstre o tipo da variável "x"
            3. Atribua 42 à variável "x" utilizando o operador "="
            4. Demonstre o valor da variável "x"
        2. Agora faça tambem:
            1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
            2. Demonstre o valor de "y"
            3. Demonstre o tipo de "y"


// Cap 3.6

- Prova (tirei 32/32)

// Cap 4.1

- Bool -> sempre são obtidos quando se usa operadores relacionados
- Usado em declarações if, switch, fluxo de controle, etc


package main

import("fmt")

var x bool

func main() {
  fmt.Println(x) // Zero value
  x = true
  fmt.Println(x) // Valor atribuido
  x = (101 <= 100)
  y := 10 =! 20 // parenteses não são obrigatórios, mas ajudam a ler o código
  fmt.Println(x,y)
}

// Cap 4.2

- Sistemas digitais possuem apenas dois estados, on e off
- Esquema de codificação: é possível transmitir 2^n mensagens com n booleans
- Binary Digits -> Bits
- ASCII vs UTF-8
- UTF-8 é um super set que começa com os ASCII e expande (incluindo letras com acento, kanjis, etc)
- 8 bit = byte
- 1000 bytes = kb
- 1000 kb = mb

// Cap 4.3

- byte = uint8
- rune = int32 (UFT8), strings são feitos de runes, cada rune é um caractere, podendo der até 4 bytes
- alguns usam 1 byte, outros usam 2,3 ou 4 bytes
- Default é int e float64
- Para ver o tipo do seu computador, use o package runtime
- O código em go é sempre em UFT8
- Uma string é uma "slice of bytes"


package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Println(runtime.GOOS)
  fmt.Println(runtime.GOARCH)
}

// Cap 4.4

- Overflow

package main

import("fmt")

func main() {
  var i uint16
  i = 65535
  fmt.Println(i)
  i++
  fmt.Println(i)
  i++
  fmt.Println(i)
}


// Cap 4.5

- Se quiser a formatação raw, basta usar backtip ``
- Slide de bytes -> vc tem o valor de cada byte da sua string
- para obter o valor em hexadecimal use "%#X"
package main

import("fmt")

func main(){
  s := `Hello, 
                playground


                      Hasld`
  fmt.Printf("%v\n%T", s, s)

  str := "Hello World"
  sb := []bytes(str)

  for _, v := range sb {
    fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
  }
}

// Cap 4.6

- Sistemas numéricos
- %d Decimal
- %b Binário
- #x Hexadecimal

package main

import("fmt")

func main() {
  a := 100
  fmt.Printf("%d\t%b\t%#x", a, a, a)
}

// Cap 4.7

- Interessante usar constantes para informações como endereço do banco de dados
- Quando declaradas com const, elas não tem tipo
- Tipo delas é definido apenas após o primeiro uso
- Pode declarar mais de uma const de uma vez

const(
  x = 10
  y = 20
  z = 30
)

// Cap 4.8

- Iota são constantes inteiras não tipadas
- Servem para situações que vc não se preocupa com o valor da constante,
- apenas que elas sejam diferentes da outra
- iota é atribuido em sequência. Usando o _, voce pode descartar alguns valores

package main

import("fmt")

const(
  a = iota // recebe o valor 0
  _ = iota // recebe o valor 1 
  c = iota // recebe o valor 2 
  _ = iota // recebe o valor 3 
  e = iota // recebe o valor 4 
)


func main() {
  fmt.Prinln(a,c,e)
}

// Cap 4.9

- Desocamento de bits pode ser da esquerda para diretita e vice versa
- Por exemplo, a sequencia abaixo contém o deslocamento
- Usefull para operações bitwise, são muito mais baratas (em tempo de máquina)
- do que multiplicação ou divisão
- Operador: >> ou <<
- Deslocar multiplica ou divide por 2^n
- Exemplo: "https://go.dev/play/p/7MOnbhx4R4"

011000
001100
000110
000011

package main

import("fmt")

func main() {
  x:=2 // em binario é 10
  y:= x >> 1 // desloca uma casa para a direita, ou seja, 1
  z:= x << 2 // desloca duas casas para a esquera, ou seja, 1000
}


// Cap 5.1

- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.

// Cap 5.2

- Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
- Demonstre estes valores.

// Cap 5.3

- Crie constantes tipadas e não-tipadas.
- Demonstre seus valores.

// Cap 5.4

- Crie um programa que:
    - Atribua um valor int a uma variável
    - Demonstre este valor em decimal, binário e hexadecimal
    - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    - Demonstre esta outra variável em decimal, binário e hexadecimal


// Cap 5.5

- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.

// Cap 6.1

- Fluxo de controle diz a maneira como o programa irá ler as linhas
- Existe o fluxo sequencial, condicional e repetição

// Cap 6.2

- Inicialização, condição e pós
- Ao final de todo statement em go, existe um ponto e vírgula, que vc pode omitir
- se trocar de linha
- Não existe o while em go

package main

import ("fmt")

func main(){
  for x := 0; x<10; x++{
    fmt.Println(x)
  }
}

// Cap 6.3

- Nested loops -> exemplo do relógio, cada vez que o ponteiro dos minutos da uma volta
- o ponteiro dos segundos deu 60 voltas
- o mesmo acontece em calendários, com dias, semanas, meses

package main

import("fmt")

func main () {
  for horas := 1; horas <= 12; horas++{
    fmt.Printf("Hora:\t%v\n", horas)
    for minutos := 0; minutos < 60; minutos ++{
      fmt.Print(" ", minutos)
    }
    fmt.Println()
  }
}

// Cap 6.4

- Para usar o while, basta usar o for apenas com o parametro de condicional
- break quebra qualquer loop

// Cap 6.5

- continue encerra a iteração atual do loop, e vai para a próxima

package main

import("fmt")

func main(){
  for i:=0; i<20; i++{
    if i%2 != 0{
      fmt.Println(i) // printa os impares
    }
  }

  for i:=0; i<20; i++
    if i%2 != 0{
      continue //vai pular os impares
    }
    fmt.Println(i) // printa os pares

}

// Cap 6.6

- Format printing:
    - Decimal       %d
    - Hexadecimal   %#x
    - Unicode       %#U
    - Tab           \t
    - Linha nova    \n
- Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.

package main

import("fmt")

func main() {
	for i := 33; i < 123; i++ {
    fmt.Printf("%v: %v\n", i, string(i))
	}
}

// Cap 6.7

- Pode usar o operador NOT (!) nas expressões lógicas
- Pode usar a condição de inicialização no if

// Cap 6.8

- if, else
- if, else if, else
- if, else if, else if, ..., else

//  Cap 6.9

- switch avalia expressões e escolhe a primeria que é true
- expression pode ser true, false, ou variável, ele vai comparar a condition com
- a expression do switch. Default é true
- se dentro de um case que FOI EXECUTADO tiver fallthrough
- ele vai executar o proximo case, independente da condition
- default serve caso nenhum case seja atendido
- switch expression {
case condition:
  
}

package main

import("fmt")

func main(){
  quem_esta_no_escritório_hoje := "zezinho"
  switch quem_esta_no_escritório_hoje{
    case "zezinho":
      fmt.Println("hoje quem está no escritório é o zezinho")
      fallthrough
    case "marquinhos":
      fmt.Println("sempre que o zezinho vem, o marquinhos vem")
    case "joana":
      fmt.Println("hoje quem está no escritório é a joana")
    case "zezinho", "zezao", "ze": // cases compostos (OR)
      fmt.Println("hoje quem está no escritório é o zezinho 2") // não será executado pq executa apenas o primeiro
    default:
      fmt.Println("não encontrei ninguẃm no escritório")
  }
}

// Cap 6.10

- switch pode avaliar tipos
- switch pode ter inicialização

package main

import("fmt")

var x interface{} // sem tipo, interface vazia

func main(){
  switch x = 3.14; x.(type){
  case int:
    fmt.Println("x é um int")
  case float64:
    fmt.Println("x é um float64")
  case bool:
    fmt.Println("x é um bool")
  case string:
    fmt.Println("x é um string")
    
  }
}

// Cap 6.11

- Operadores lógicos
- &&    conditional AND    p && q  is  "if p then q else false"
- ||    conditional OR     p || q  is  "if p then true else q"
- !     NOT                !p      is  "not p"

// Cap 7.1

- Põe na tela: todos os números de 1 a 10000.

// Cap 7.2

- Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
- Por exemplo:
    65
        U+0041 'A'
        U+0041 'A'
        U+0041 'A'
    66
        U+0042 'B'
        U+0042 'B'
        U+0042 'B' 
    ...e por aí vai.

// Cap 7.3

- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.

// Cap 7.4

- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.

// Cap 7.5

- Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

// Cap 7.6

- Crie um programa que demonstre o funcionamento da declaração if.

// Cap 7.7

- Utilizando a solução anterior, adicione as opções else if e else.

// Cap 7.8

- Crie um programa que utilize a declaração switch, sem switch statement especificado.

// Cap 8.1

- Arrays permitem guardar valores diferentes que podem ou não ser do mesmo tipo
- Tipos e tamanhos diferentes de arrays não são compativeis

package main

import("fmt")

var x [5]int

func main(){
  x[0] = 1
  x[1] = 10
  fmt.Println(x[0],x[1])
  fmt.Println(x) // Os demais es~tao com o valor zero do seu tipo
  fmt.Println(len(x))
}

// Cap 8.2

- Tipos de dados compostos, sem comprimento definido
- São feitos de arrays

package  main

import("fmt")

func main(){
  array := [5]int{1,2,3,4,5}
  fmt.Println(array)
  slice := []int{1,2,3,4,5}
  fmt.Println(slice)

  slice2 := append(slice,6)
  fmt.Println(slice2)

  smt.Prinln(slice[3])
  slice[3] = 34990
  fmt.Prinln(slice[3])

  // Proibido em slice pois ele vem de um array subadjacente: slice[2] = 1
}

// Cap 8.3

-

package main

import("fmt")

func main(){
  slice :=[]string{"banana", "maça", "jaca", "pessego"}

  // Se não quiser um destes outputs, só usar o _
  for indice, valor := range slice {
    fmt.Prinln("No indice", indice, "temos o valor:", valor)
  }
  
  // Não funciona: slice[4] = "melancia"
  slice = append(slice, "melancia")

  for indice, valor := range slice {
    fmt.Prinln("No indice", indice, "temos o valor:", valor)
  }
}

// Cap 8.4

- x[:], x[a:], x[:b], x[a:b]
- "a" é incluso
- "b" não é incluso
- você excluir um item adiconando tudo que tinha antes e tudo que tinha depois

package main

import("main")

func main(){
  sabores := []string{"pepperoni","mozzarela","abacaxi","4 queijos", "marguerita"}

  fatia := sabores[0:2]
  fmt.Prinln(fatia) // -> pepperoni e mozzarela

  fatia = sabores[2:len(sabores)] // -> abacaxi até o final
  sabores = append(sabores[:2], sabores[3:]...) // tirando a abacaxi
}


// Cap 8.5

- append aceita como argumentos uma slice e os tipos da sice

package main
import("fmt")

func main(){
  umaslice := []int{1,2,3,4}
  outra slice := []int{9,10,11,12}

  fmt.Println(umaslice)

  umaslice = append(umaslice, 5,6,7,8)
  // Os ... servem para "desenrolar" a slice, ele enumera os itens dessa slice
  // Precisa fazer iso pq o o argumento tem que ser int, e não slice of int

  umaslice = append(umaslice, outraslide...)
}

// Cap 8.6

- Sempre que um slice muda de tamanho, um novo array é criado
- Para evitar problemas de benchmark, use make([]T,len,cap)


package main
import("fmt")

func main(){
  slice := make([]int,5,10)

  slice[0], slice[1], slice[2], slice[3], slice[4] = 1,2,3,4,5

  slice = append(slice,6)
  slice = append(slice,7)
  slice = append(slice,8)
  slice = append(slice,9)
  slice = append(slice,10)

  fmt.Println(slice, len(slice), cap(slice)) // 10 elementos, 10, 10

  slice = append(slice,10)

  fmt.Println(slice,len(slice),cap(slice)) // 11 elementos, 11, 20
}

// Cap  8.7

- slices mmulti-dimensional (matrix)

package main
import("fmt")
func main(){
  ss := [][]int{[]int{1,2,3,4,5,6},[]int{7,8,9,10,11,12},[]int{13,14,15,16,17,18}}
  fmt.Prtinln(ss[2][4])
}

// Cap 8.8

- Quando vc quer usar slices de um slice, faça um for loop usando for loop
- Se usar slices diferentes com o msmo array subjacente, vc modifcará o array original

// Cap 8.9

- Map são key-values, não ordenados
- Exemplo: lista telefonica 
- Performance muito boa
- Use os comma ok

package main
import("fmt")
func main(){
  amigos := map[string]int{
    "alfredo": 5551234,
    "joana": 9996789
  }
  fmt.Print(amigos)
  fmt.Println(amigos["joana"])

  amigos["gopher"] = 4444444
  fmt.Println(amigos["gopher"])

  fmt.Prtinln(amigos["romario"]) // Vai devolver o Zero Value
  // Como diferenciar Zero Value de um valor 0
  // Comma ok

  será, ok := amigos["fantasma"]
  fmt.Prinln(será,ok) // 0 false (ok é um bool que mostra se tem ou não no map)

  if será, ok := amigos["fantasma"], !ok {
    fmt.Println("não tem!")
  } else {
    fmt.Prinln(será)
  }

}

// Cap 8.10

- Para usar range em map, use sempre kay, value, ou _
- Para deletar, asta usar delete(map,key)

package main
import("fmt")
func main(){
  qualquercoisa := map[int]string{
    123: "muito legal",
    98: "menos legal um pouquinho",
    19: "idade de ir para festa",
    983: "esse é massa"
  }

  fmt.Println(qualquercoisa)

  total := 0
  for key, value := range qualquercoisa { //lembre que não é ordenado
    total += key
  }

  delete(qualquercoisa, 123)
}

// Cap 9.1

- Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.

// Cap 9.2

Usando uma literal composta:
  - Crie uma slice de tipo int
  - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.

// Cap 9.3

- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
    - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
    - Do quinto ao último item do slice (incluindo o último item!)
    - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
    - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
    - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

// Cap 9.4

- Começando com a seguinte slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Anexe a ela o valor 52;
- Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
- Demonstre a slice;
- Anexe a ela a seguinte slice:
    - y := []int{56, 57, 58, 59, 60}
- Demonstre a slice x.

// Cap 9.5

- Começe com essa slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Utilizando slicing e append, crie uma slice y que contenha os valores:
    - [42, 43, 44, 48, 49, 50, 51]


// Cap 9.6

- Crie uma slice usando make que possa conter todos os estados do Brasil.
    - Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
- Demonstre o len e cap da slice.
- Demonstre todos os valores da slice sem utilizar range.

// Cap 9.7

- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.


// Cap 9.8

- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.


// Cap 10.1

- Permite armazenar dados de tipos diferentes

package main
import("fmt")

type cliente struct{
    nome      string
    sobrenome string
    fumante   bool
}
func main(){
  cliente1 := cliente{
    nome:       "João",
    sobrenome:  "da Silva",
    fumante:    false,
  }
  ciente2 := cliente{"Joana","Pereira", true}

  fmt.Println(cliente1)
  fmt.Println(cliente2)
}


// Cap 10.2

- Também é possível ter structs dentro de outro struct

package main
import("fmt")

type pessoa struct{
  nome  string
  idade int
}

type profissional struct{
  pessoa
  titulo  int
  salario int

func main(){
  pessoa1 := pessoa{
    nome: "Alfredo",
    idade:  30,
  }

  pessoa2 := pessoa{
    pessoa: pessoa{
      nome: "Maricota",
      idade:  31,
    },
    titulo: "Pizzaiola",
    salario:  10000,
  }
  fmt.Println(pessoa1)
  fmt.Println(pessoa2)

  pessoa3 := pessoa{"Mauricio",40}
  pessoa4 := profissional{pessoa{"Vanderlei",50}, "Político", 30000}

}

// Cap 10.3

- Para acessar apenas um campo, use .


  fmt.Println(pessoa1.nome)
  fmt.Println(pessoa2.pessoa.nome)
  // Caso não haja conflito de nome, é possível acessar o struct imbutido
  fmt.Println(pessoa2.pessoa)


// Cap 10.4

- Structs vc declara um tipo e vc faz quantos valores vc quiser desse tipo
- Structs anônimos são descartáveis e vc não declara tipos

package main
import("fmt")

func main(){
  x := struct{
    nome  string
    idade int
  }{
    nome: "Mario",
    idade:  50}
  
    fmt.Println(x)
}

// Cap 11.1

- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.


// Cap 11.2

- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.

// Cap 11.3

- Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.

// Cap 11.4

- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

// Cap 12.1

- Em go, quase tudo é pass by value, e não pass by reference (como é no python)
- func (receiver) nome(parametro) (output1, output2)
- output := func(argumento)
- outras maneiras de definir os paramtros
- (x, y, int) se forem todos int
- (x int, y string)
- Também é possível fazer uma funcão com uma quantidade variádica de elementos
- Parâmetros variádicos deve ser sempre o último (s tring, x ...int) // espera receber ints, opr isso, se tier um slice, use o operador ...

func soma (x,y int) int{
  return x+y
}

func basica () {
  fmt.Prinln("oi, bom dia")
}

func soma_variadica(x ... int) (int, int){
  soma := 0
  for _,v := range x{
    soma += v 
  }
  return soma, len(x)
}

// Cap 12.2

- Use o operador ... para desenrolar uma slice of intt  e passar como paramtros o tipo int, e não slice of int

si := []int{1,2,3,4,5,6}
soma_variadica(si...)

func soma_variadica(x ... int) int{
  soma := 0
  for _,v := range x{
    soma += v 
  }
  return soma
}

// Cap 12.3

- defer é aquilo que o brasileiro faz quando ele tiver tempo
- defer é um statement que será executado de última hora
- se tiver dois defer, o primeiro a ser adicionado é o último a ser executado (pilha)
- first in, last out
- Útil para abrir e fechar arquivos/conexões de rede, nesse caso:
- 1: abrir arquivo
- 2: defer fechar arquivo
- 3: editar arquivo
package main
import("fmt")

func main(){
  defer fmt.Println("com defer, veio primeiro") // é o 2 a ser impresso
  fmt.Println("sem defer, veio depois") // é  impresso antes
}

func main() {
  defer fmt.Println("1") // 4 a ser impresso
  defer fmt.Println("2") // 3 a ser impresso
  fmt.Println("3") // 1 a ser impresso
  fmt.Println("4") // 2 a ser impresso

}

// Cap 12.4

- Método é uma função anexada a um tipo (metodo)
- Receiver é exclusivo, a função somente pode  receber aquele tipo
- Um método adiciona uma função exclusivamente a um tipo, podendo suar a chamada por variavel.metodo() 
- Não é possível chamá-la com nome_func()

package main
import("fmt")

type pessoa struct{
  nome  string
  idade int
}

// sintaxe
// func (receiver) identifier(parameters) (returns) { code }

func (p pessoa) oibomdia() {
  fmt.Println(p.nome, "diz bom dia!")
}

func main() {
  mauricio := pessoa{"Maurício", 30}
  mauricio.oibomdia()
}

// Cap 12.5

- Interfaces e polimorfismo

- Em Go, valores podem ter mais que um tipo.
- Uma interface permite que um valor tenha mais que um tipo.
- Se um tipo tem os mesmos métodos que uma interface, ele percente a ela
- Declaração: keyword identifier type → type x interface
- Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.
- Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
- Esse tipo será o seu tipo e também o tipo da interface.

- Exemplos:
    - Os tipos profissão1 e profissão2 contem o tipo pessoa
    - Cada um tem seu método oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()
    - Implementam a interface gente
    - Ambos podem acessar o função serhumano() que chama o método oibomdia() de cada gente
    - Tambem podemos no método serhumano() tomar ações diferentes dependendo do tipo:
        switch pessoa.(type) { case profissão1: fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }* 
    - Go Playground pré-pronto: https://play.golang.org/p/VLbo_1uE-U
    https://play.golang.org/p/zGKr7cvTPF
    - Go Playground ao vivo: 
    https://play.golang.org/p/njiKbTT20Cr
- Onde se utiliza?
    - Área de formas geométricas (gobyexample.com)
    - Sort
    - DB
    - Writer interface: arquivos locais, http request/response
- Se isso estiver complicado, não se desespere. É foda mesmo. Com tempo e prática a fluência vem.


package main()
import("fmt")

type pessoa struct {
  nome  string
  sobrenome string
  idade int
}

type dentista struct{
  pessoa
  dentesarrancados int
  salario float64
}

type arquiteto struct{
  pessoa
  tipodeconstrução  string
  tamanhodaloucura  string
}

func (x dentista) oibomdia(){
  fmt.Println("Meu nome é", x.nome,"e eu já arranquei", x.dentesarrancados, "dentes")
}

func (x arquiteto) oibomdia(){
  fmt.Printlm("Meu nome é", x.nome,"e ouve só minha loucura: BOM DIA!")
}

type gente interface{
  oibomdia()
}

func serhumano(g gente){
  g.oibomdia()
}

func main(){
  mrdente := dentista{
    pessoa: pessoa{
      nome: "Mister Dente",
      sobrenome:  "da Silva",
      idade:  50,
    },
    dentesarrancados: 8912,
    salário: 252525,
  }

  mrprédio := arquiteto{
    pessoa: pessoa{"Mister Prédio","Afonso",51},
    tipodeconstrução: "Hospício",
    tamanhodaloucura: "mucholoko",
    }
  }

  // usando os métodos:
  mrdente.oibomdia()
  mrprédio.oibomdia()

  // via interface
  serhumano(mrdente)
  serhumano(mrprédio)
}

// Cap 12.6

- Funções anônimas não precisam de nome
- São descartáveis
- Útil para go routine
- Define e já usa ela
- Dentro do main

func main(){
  x := 341

  func(x int){
    fmt.Println(x,"vezes 12345 é:",x*12345)
  }(x)
}

// Cap 12.7

- É possível atribuir uma função como valor à uma variável

func main(){
  x := 341

  y := func(x int) int{
    return x*12345
  }

  fmt.Println(x,"vezes 12345 é:",y(x))
}

// Cap 12.8

- A função devolve uma função (que pode ter int como parâmtro e devolver int)
//                  parâmtetro
//                            devolve uma função cujo parâmtreo e int e valor de retorno int
func retornaumafunção() func(int) int{
  return func(i int) int{
    return i*10
  }
}

x := retornaumafunção()
y := x(3)
fmt.Println(y)

// se quiser ir  direto, faça assim

fmt.Println(retornaumafunção()(3))o


// Cap 12.9

- Callback é passar uma função como argumento de outra função 

package main

import (
	"fmt"
)

func main() {
	t := somenteImpares(soma, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)
	fmt.Println(t)
}

func soma(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func somenteImpares(f func(x ...int) int, y ...int) int {
	var impares []int
	for _, i := range y {
		if i%2 != 0 {
			impares = append(impares, i)
		}
	}
	return f(impares...)
}

// Cap 12.10

- Closure é cercar ou capturar um scope para que possamos utilizá-lo em outro contexto. Já vimos:
    - Package-level scope
    - Function-level scope
    - Code-block-in-code-block scope
- Exemplo de closure:
    - func i() func() int { x := 0; return func() int { x++; return x } }
    - Quando fizermos a := i() teremos um scope, um valor para x.
    - Quando fizermos b := i() teremos outro scope, e x terá um valor independente do x acima.
- Closures nos permitem salvar dados entre function calls e ao mesmo tempo isolar estes dados do resto do código.


package main
import("fmt")

func main(){
  a := i()

  fmt.Println(a()) // Printa 1
  fmt.Println(a()) // Printa 2
  fmt.Println(a()) // Printa 33
  fmt.Println(a()) // Printa 4

  b := i()

  fmt.Println(b()) // Printa 1
  fmt.Println(b()) // Printa 2
  fmt.Println(b()) // Printa 3
}

func i() func() int{
  x := 0
  return func() int {
    x++
    return x
  }
}


// Cap 12.11

- Prefira usar loops pois eles usam menos memoria e tem menos chance de quebrar


func fatorial(x int)  int{
  if x == 0 || x == 1 {
    return 1
  } else {
    return x * fatorial(x-1)
  }
}

// Revisão Cap 12

- Revisão:
    - Funções!
        - Servem para abstrair código
        - E para reutilizar código
    - A ordem das coisas é:
        - func (receiver) identifier (parameters) (returns) { code }
    - Parâmetros vs. argumentos
    - Funções variádicas
        - Múltiplos parâmetros
        - Múltiplos argumentos
    - Métodos
    - Interfaces & polimorfismo
    - Defer
        - "Deixa pra depois!"
    - Returns
        - Múltiplos returns
        - Returns com nome (blé!)
    - Funcs como expressões
        - Atribuindo uma função a uma variável
    - Callbacks
        - Passando uma função como argumento para outra função
    - Closure
        - Capturando um scope
        - Variáveis declaradas em scopes externos são visíveis em scopes internos
    - Recursividade
        - Uma função que chama a ela mesma
        - Fatoriais 

// Cap 13.1

- Exercício:
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados

// Cap 13.2

- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função

// Cap 13.3

- Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.

// Cap 13.4

- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.


// Cap 13.5

- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"

// Cap 13.6

- Crie e utilize uma função anônima.

// Cap 13.7

- Atribua uma função a uma variável.
- Chame essa função.

// Cap 13.8

- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.

// Cap 13.9

- Callback: passe uma função como argumento a outra função.

// Cap 13.10

- Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.

// Cap  14.1

- Todos os valores ficam armazenados na memória.
- Toda localização na memória possui um endereço.
- Um pointeiro se refere a esse endereço.
- Notações:
    - &variável mostra o endereço de uma variável
        - %T: variável vs. &variável
    - *variável faz de-reference, mostra o valor que consta nesse endereço
    - ????: *&var funciona!
    - *type é um tipo que contem o endereço de um valor do tipo type, nesse caso * não é um operador
- Exemplo: a := 0; b := &a; *b++

// Cap 14.2

- Ponteiros permitem compartilhar endereços de memória. Isso é útil quando:
    - Não queremos passar grandes volumes de dados pra lá e pra cá
    - Queremos mudar um valor em sua localização original (tudo em Go é pass by value!)
- Exemplos:
    - x := 0; funçãoquemudaovalordoargumentopra1(x); Print(x)
    - x := 0; funçãoquemudaovalordo*argumentopra1(&x); Print(x)

// Cap 15.1

- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória.

// Cap 15.2

- Crie um struct "pessoa"
- Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. Essa função deve mudar um valor armazenado no endereço *pessoa.
    - Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
    - Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
    - "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), 
    - → x.f is shorthand for (*x).f." ←
    - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
- Crie um valor do tipo pessoa;
- Use a função mudeMe e demonstre o resultado.
