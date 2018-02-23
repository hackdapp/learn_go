/*
# 值类型与引用类型
所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值：

你可以通过 &i 来获取变量 i 的内存地址，

注意事项:
a := 10 只能允许出现在函数中, 且只能初始化一次。但是赋值是可以 a=20
*/

package main

import "fmt"

//全局变量是允许声明但不使用
//同一类型的多个变量可以声明在同一行
var a, b, c int;

func main()  {
  a, b, c = 1, 2, 3

  var d string = "ssdf";
  fmt.Printf("test %s", d)
}
