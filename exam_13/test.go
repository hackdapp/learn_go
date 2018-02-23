package main

import "fmt"

type Animal struct{
    _name string
    _weight float32
    _age int
}

func main()  {
  var dog Animal

  dog._name = "wang"
  dog._age = 2
  print(dog)

  modify(&dog)
  print(dog)

}

func print(animal Animal){
  fmt.Printf("age: %d\n", animal._age)
  fmt.Printf("_name: %s\n", animal._name)
}

func modify(animal *Animal){
  animal._name = "test_wang"
}
