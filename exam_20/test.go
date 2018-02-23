package main

import "fmt"

type error interface {
    Error() string
}

type DivError struct{
}
func (divError DivError) Error() string{
  return "除数不允许为0"
}

func div(a, b int) (result float32, errorMsg string) {
  if b == 0 {
    err := new(DivError)
    return float32(0), err.Error()
  }
  return float32(a)/float32(b), ""
}
func main(){

    if _, errorMsg := div(1,0); errorMsg != "" {
      fmt.Println("errorMsg is: ", errorMsg)
    }

    if result, errorMsg := div(8,2); errorMsg == "" {
      fmt.Println("8/2:", result)
    }
}
