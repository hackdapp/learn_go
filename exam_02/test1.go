//这是一个包名
package main

/*行注释*/
import "fmt"

/**
 * 块注释
 */
func main()  {
  //标识符： 一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。
  var _name string;
  var age int;

  
  _name = "this is a string"
  fmt.Printf("%s", _name)
}
