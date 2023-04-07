package main

import("fmt")

func main() {
  array := []int{10,20,30,40,50,60,70,80,90,100}
  fmt.Println(array[:3])
  fmt.Println(array[4:])
  fmt.Println(array[1:7])
  fmt.Println(array[2:len(array)-1])
	fmt.Printf("%T\n", array)
}
