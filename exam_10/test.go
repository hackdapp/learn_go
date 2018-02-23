/*
Go 语言变量作用域
作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。

Go 语言中变量可以在三个地方声明：
    函数内定义的变量称为局部变量
    函数外定义的变量称为全局变量
    函数定义中的变量称为形式参数
*/
package main

import "fmt"
//全局变量
var CFG_V1 = 10

//形参
func person(name string, age int) (string, int) {

  //局部变量
  /*Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑
    不同类型的局部和全局变量默认值为：
      int : 0
      float32 : 0
      pointer : nil
  */
  var _name string;
  var _age int;

  return _name, _age
}

func main(){
  var ptr *int

  fmt.Print(ptr)
}
