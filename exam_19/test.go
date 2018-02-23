package main

import (
  "fmt"
)

type Phone interface {
  call()
}

type SmartPhone struct {
}
func (smartPhone SmartPhone) call(){
  fmt.Println(" this is a smartPhone")
}

type ApplePhone struct{

}
func (applePhone ApplePhone) call(){
  fmt.Println(" this is a applePhone")
}

func main()  {
  var phone Phone

  phone = new(SmartPhone)
  phone.call();
  phone = new(ApplePhone)
  phone.call();
}
