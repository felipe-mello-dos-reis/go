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

func mian() {
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

// Cap 7.9


