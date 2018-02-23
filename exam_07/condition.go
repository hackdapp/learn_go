/*
# Go 语言条件语句
1. if
2. if else
3. switch
4. select
*/

package main

import "fmt"

func main()  {
  ifmain()
  switchmain()
  selectmain()
}

func ifmain(){
  var a = 10
  b := 20

  if a > 10 {
    if b > 1 {

    }
  } else {

  }
}

func switchmain(){
  var a int = 100
  switch a {
    case 10: println("A")
    case 100: println("B")
    case 200: println("C")
    default : println("default")
  }
  a = 300
  switch  {
    case a > 1: println("AA1")
    case a > 10: println("AA2")
    case a > 200: println("AA3")
    case a > 500: println("AA4")
  }

  var x   interface{

    };

  switch i := x.(type) {
    case nil:
         fmt.Printf(" x 的类型 :%T",i)
      case int:
         fmt.Printf("x 是 int 型")
      case float64:
         fmt.Printf("x 是 float64 型")
      case func(int) float64:
         fmt.Printf("x 是 func(int) 型")
      case bool, string:
         fmt.Printf("x 是 bool 或 string 型" )
      default:
         fmt.Printf("未知型")
  }
}

func selectmain(){
  var c1, c2, c3 chan int
   var i1, i2 int
   
   select {
      case i1 = <-c1:
         fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n")
   }
}
