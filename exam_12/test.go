package main

import "fmt"

func main()  {
  var a int = 10
  fmt.Printf("变量地址: %x\n", &a)

  arrayptr()
}

func nilptr()  {
  var nilptr *int
  fmt.Printf("指针缺省值: %x\n", nilptr)

  if nilptr == nil {
    fmt.Println("指针为null")
  }
}

const LEN int = 5

func arrayptr(){
  var args =  []int{1,2,3,4,5}
  var ptr [LEN]*int
  var i int = 0

  for i=0; i<LEN; i++ {
    ptr[i] = &args[i]
  }

  for i=0; i<LEN; i++ {
    fmt.Printf("%d\n", *ptr[i])
  }
}
