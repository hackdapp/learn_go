/*
# Go 语言运算符
运算符用于在程序运行时执行数学或逻辑运算。

Go 语言内置的运算符有：

  算术运算符
  关系运算符
  逻辑运算符
  位运算符
  赋值运算符
  其他运算符
*/

package main

import "fmt"

func main()  {
  var a = 4
  var b int32
  var c float32

  var ptr *int

  fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
  fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
  fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );
  fmt.Printf("第 3 行 - c 变量类型为 = %T\n", ptr );

  ptr = &a;
  fmt.Printf("%s\n", ptr)
  fmt.Printf("%d\n", *ptr)

  //运算符优先级
  /*
  有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：
  7: ^ !
  6: * / % << >> && &^
  5: +- | ^
  4: == != < <= >= >
  3: <-
  2: &&
  1: ||

  */
  cal()
}

func cal(){
  var a int = 20
   var b int = 10
   var c int = 15
   var d int = 5
   var e int;

   e = (a + b) * c / d;      // ( 30 * 15 ) / 5
   fmt.Printf("(a + b) * c / d 的值为 : %d\n",  e );

   e = ((a + b) * c) / d;    // (30 * 15 ) / 5
   fmt.Printf("((a + b) * c) / d 的值为  : %d\n" ,  e );

   e = (a + b) * (c / d);   // (30) * (15/5)
   fmt.Printf("(a + b) * (c / d) 的值为  : %d\n",  e );

   e = a + (b * c) / d;     //  20 + (150/5)
   fmt.Printf("a + (b * c) / d 的值为  : %d\n" ,  e );
}
