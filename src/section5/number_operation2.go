package main

import (
  "fmt"
)


func main() {

  //데이터 타입이 같을때 다양한 연산자 사용 가능.
  var n1 uint8 = 125
  var n2 uint8 = 90

  fmt.Println("ex1 : ", n1+n2)
  fmt.Println("ex1 : ", n1-n2)
  fmt.Println("ex1 : ", n1*n2)
  fmt.Println("ex1 : ", n1/n2)
  fmt.Println("ex1 : ", n1%n2)
  fmt.Println("ex1 : ", n1<<2)
  fmt.Println("ex1 : ", n1>>2)
  fmt.Println("ex1 : ", ^n1)

  var n3 int = 12
  var n4 float32 = 8.2
  var n5 uint16 = 1024
  var n6 uint32 = 120000

  //fmt.Println("ex3 : ", n3 + n4)

  fmt.Println("ex3 : ", float32(n3) + n4)
  fmt.Println("ex3 : ", n3+int(n4))
  // 정수형과 실수형을 계산할때는 실수형쪽으로 맞춰서 계산하는게 좋다.

  fmt.Println("ex3 : ", n5+uint16(n6)) // 데이터 타입이 넘쳐, uint16의 max값으로 계산됨.
  // 데이터타입으로 인해 손실된 값들을 생각해서 계산해야한다.

  

}
