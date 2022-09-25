package main

import (
  "fmt"
)

func main() {
  //예제2(비교)
  str1 := "Golang"
  str2 := "World"

  fmt.Println("ex1 : ", str1 == str2)
  fmt.Println("ex2 : ", str1 != str2)
  fmt.Println("ex3 : ", str1 > str2)// 바이트를 비교한다고 생각하면, str1이 클거같음
  fmt.Println("ex4 : ", str1 < str2)// Go문자열 -->아스키 코드에 대한 사전식 비교
  // 소문자 abc, ABC 크기를 비교할때 대문자 ABC가 더큼. 아스키코드가 너 크므로.

}
