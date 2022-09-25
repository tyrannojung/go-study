package main

import (
  "fmt"
)

func main() {
  //데이터 타입 : 숫자형
  //정수, 실수, 복소수
  //32bit, 64bit, unsigned(양수)
  //정수 : 8진수(0), 16진수(0x), 10진수
  //자료형을 엄격히 하면서 가독성이 좋아짐


  // 내가 데이터 타입을 정확히 알고있으면 32, 64를 써도되는데, 보편적으로는 int만 쓰는게 속 편하다.(내부적으로 최적화됨)
  var num1 int = 17
  var num2 int = -68
  var num3 int = 0631 //8진수
  var num4 int = 0x32fa2c75 //16진수

  fmt.Println("ex1 : ", num1)
  fmt.Println("ex2 : ", num2)
  fmt.Println("ex3 : ", num3)
  fmt.Println("ex4 : ", num4)



}
