package main

import (
  "fmt"
)

func sum(i int, f func(int, int)){
  f(i, 10)
}

//매개변수가 둘다 int 경우 생략가능
func add(a, b int) {
  fmt.Println("ex1 : ", a+b)
}

func multi_value(i int) {
  i = i +10
}

func multi_reference(i *int){
  *i *= 10// *i = *i * 10
}


func main() {
  //함수(콜백) 전달, 참조 전달(call by reference), 값 전달(call by value)
  //예제1
  sum(100, add) //매개변수로 함수전/

  //예제2
  // 주소참조가 아니므로, 값에 대한 변경이 없다.
  a := 100
  multi_value(a)
  fmt.Println("ex2 : ", a)

  //예제3
  //reference by value
  b := 100
  multi_reference(&b)
  fmt.Println("ex3 : ", b)


}
