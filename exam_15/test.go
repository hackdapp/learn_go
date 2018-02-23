/*
# 范围
Go 语言中 range 关键字用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。
*/

package main

import "fmt"

func main(){
  var arrays = []int{21,33,1,2}
  var sum int
  for _, num := range arrays {
      sum += num
  }
  fmt.Printf("汇总: %d\n", sum)

  sum = 0
  for i, num := range arrays {
      sum += num
      fmt.Printf("index:%d, sum:%d\n", i, sum)
  }
  fmt.Printf("汇总: %d\n", sum)

  var unk = map[string]string{"a":"apple", "b":"banana"}
  for k,v := range unk{
    fmt.Printf("%s -> %s\n", k, v)
  }

  for i, c := range "ABab" {
         fmt.Println(i, c)
  }
}
