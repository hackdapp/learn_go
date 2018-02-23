/*
# Go 语言变量

Go 语言变量名由字母、数字、下划线组成，其中首个字母不能为数字。
var identifier type

第一种，指定变量类型，声明后若不赋值，使用默认值。
var str bool;

第二种, 根据值自行判定变量类型
var str = "123";

第三种, 省略var, 注意:=左侧的变量不应该是已经声明过的，否则会导致编译错误
_name := value

exp.
var a int = 10
var a = 10
a := 10
以上三种赋值等效

*/
package main

import "fmt"
func main()  {
  //单变量声明
  var a int;
  var b = "this a string";
  c := "teststr";
  fmt.Println(a, b, c)

  //多变量声明, 非全局变量
  var _name1, _name2, _name3 = 1, 2, 3;
  fmt.Println(_name1, _name2, _name3)

  _name4, _name5, _name6 := 4, 5, 6
  fmt.Println(_name4, _name5, _name6)

  // 这种因式分解关键字的写法一般用于声明全局变量
  var {
    vname1 string
    vname2 int
  }

  //这种不带声明格式的只能在函数体中出现
  //g, h := 123, "hello"
};
