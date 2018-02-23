package main

import "fmt"

func main()  {
  var count int= 0
   fmt.Println(t1(15, &count))
   fmt.Println(count)
}

func t1(n int, count *int) int{
  (*count)++
  if n > 1 {
    return n * t1(n-1, count)
  }
  return 1
}
