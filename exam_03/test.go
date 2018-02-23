package main
/*
Go 语言数据类型

1. 布尔型
var b bool = true

2. 数字类型
int
float32
float64
----
int uint uintptr

uint8/16/32/64 取值范围: 0 ~ 2^8, ...

int8/16/32/64 取值范围: -2^8 ~ 2^8-1 ......
----
浮点型
float32/float64: IEEE-754 32， IEEE-754 64
complex64/complex128: 32位实数/虚数， 64位实数/虚数

----
其他数字类型
byte : 类似于uint8
rune 类 int32
uint 32/64
int 与 uint一样大小
uintptr 无符号整型，用于存放一个指针

补充说明:
go 1.9版本对于数字类型，无需定义int及float32、float64，系统会自动识别。
package main
import "fmt"

func main() {
   var a = 1.5
   var b =2
   fmt.Println(a,b)
}
-----------------------------------------------------
3. 字符串类型

4. 派生类型
  a) 指针类型
  b) 数据类型
  c) 结构化类型
  d) channel类型
  e) 函数类型
  f) 切片类型
  g) 接口类型
  h) Map类型
*/

func main()  {
  var flag bool = false;

}
