/*
iota
iota，特殊常量，可以认为是一个可以被编译器修改的常量。

在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。

iota 可以被用作枚举值：
*/

package main

import "fmt"
import "unsafe"

const (
  a = iota
  b
  c = "123"
  d = iota
)

const (
    i=1<<iota //左移0位
    j=3<<iota //左移1位
    k //左移2位
    l //左移3位

    /**
    0, 1
    1, 11
    2, 1111
    3, 1111111
    */


)




func main()  {
  println(a,b,c, d)
  fmt.Println("i=",i)
  fmt.Println("j=",j)
  fmt.Println("k=",k)
  fmt.Println("l=",l)
  println()

  var str = "123";
  //字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
  println("string byte size: ",unsafe.Sizeof(str))
  var invar = 0;
  println("int byte size: ", unsafe.Sizeof(invar))
}
