/*
#Go 语言数组
语法格式: var variable_name [SIZE] variable_type
exp. var balance [10] float32


*/
package main

import "fmt"

const (
  LEN = 10
)


func main(){
  var balance = [...]float32{1.9, 212.9, 33}
  fmt.Println(balance)

  balance[1] = 12312.123
  fmt.Println(balance)

  var ars [5]int;
  var i int
  for i=0; i<5; i++ {
    ars[i] = 100 + i
  }
  fmt.Println(ars)

 var a [3][4]int
  a = [3][4]int{
   {0, 1, 2, 3} ,   /*  第一行索引为 0 */
   {4, 5, 6, 7} ,   /*  第二行索引为 1 */
   {8, 9, 10, 11}   /*  第三行索引为 2 */}
  //var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
  fmt.Println(a)
  fmt.Println(a[1][2])


  var param = [10]int{1,2,3,4,5,6,7,8,9,10}
  fmt.Println(myFunction1(param))
  var param2 = []int{1,2,3,4,5,6,7,8,9,10}
  fmt.Println(myFunction2(param2, 10))

}

/**
# 向函数传递数组

*/
func myFunction1(param [10]int) int{
  var result, i int;

  for i=0; i<10; i++ {
    result += param[i]
  }
  return result
}

func myFunction2(param []int, size int) int{
  var result, i int;
  for i=0; i<size; i++ {
    result += param[i]
  }
  return result
}
