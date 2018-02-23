package main

import "fmt"

func main(){
    s := []int{1,2,3,4}

    //startindex:endindex-1
    s1 := s[1:3]

    println(s1[0], s1[1])

    s2 := make([]int, 3, 30)
    println("size:", len(s2))
    println("cap:", cap(s2))

    s2 = nil
    println(s2)

    main1()
}




func main1() {
   var numbers []int
   printSlice(numbers)

   /* 允许追加空切片 */
   numbers = append(numbers, 0)
   printSlice(numbers)

   /* 向切片添加一个元素 */
   numbers = append(numbers, 1)
   printSlice(numbers)

   numbers = append(numbers, 2)
   printSlice(numbers)

   /* 同时添加多个元素 */
   numbers = append(numbers, 3)
   printSlice(numbers)

   numbers = append(numbers, 4)
   printSlice(numbers)

   numbers = append(numbers, 4, 2)
   printSlice(numbers)
   numbers = append(numbers, 2)
   printSlice(numbers)
   numbers = append(numbers, 4,5)
   printSlice(numbers)
   /* 创建切片 numbers1 是之前切片的两倍容量*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)

   /* 拷贝 numbers 的内容到 numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
