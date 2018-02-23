/*
Go 语言常量
常量是一个简单值的标识符，在程序运行时，不会被修改的量。

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

常量的定义格式：const identifier [type] = value
  显式类型定义： const b string = "abc"
  隐式类型定义： const b = "abc"

多个相同类型的声明可以简写为：
  const c_name1, c_name2 = value1, value2
*/


package main

import "fmt"
import "unsafe"


const LENGTH = 10
const WIDTH = 20
var area int

//枚举
const (
  UNKNOWN = -1
  FEMALE = 0
  MALE = 1
)

const (
  //常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：
  a1 = "abc"
  b2 = len(a1)
  c3 = unsafe.Sizeof(a1)
)

func main()  {
  area = LENGTH * WIDTH
  fmt.Printf("area: %d\n", area)

  const a, b, c = 1, true, "hello world"
  println("")
  println(a, b, c)

  //枚举
  println(UNKNOWN, FEMALE, MALE)
  println(a1, b2, c3)

}
