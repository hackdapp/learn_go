/*
# Go 语言函数
函数格式:
func function_name( [parameter list] ) [return_types] {
   函数体
}

默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数


*/

package main

import "fmt"

func main()  {
  var result = max(2, 10)
  println(result)

  var c,d = sort(2,11)
  println(c, d)

  //语言函数作为值
  getMaxVal := func(a, b int) int  {
    if a > b {
      return a
    } else {
      return b
    }
  }

  println(getMaxVal(100, 49))

  //Go 语言函数闭包
  /*
    Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

    以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量，代码如下：
  */
  nextNumber := getSequence()
  /* 调用 nextNumber 函数，i 变量自增 1 并返回 */
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())

   /* 创建新的函数 nextNumber1，并查看结果 */
   nextNumber1 := getSequence()
   fmt.Println(nextNumber1())
   fmt.Println(nextNumber1())

   add_func := add(1, 2)
   fmt.Println(add_func())
   fmt.Println(add_func())
   fmt.Println(add_func())
}

func max(a, b int) int  {
  if a > b {
    return a
  } else {
    return b
  }
}

func sort(a, b int) (int, int)  {
  if a > b {
  } else {
    a, b = b, a
  }
  return a, b
}

func getSequence() func() int  {
  i := 0
  return func() int {
    i += 1
    return i
  }
}

func add(a , b int) func() (int, int){
  i := 0
  return func() (int, int){
    i++
    return i, a+b
  }
}
