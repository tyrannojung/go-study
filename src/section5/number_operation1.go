package main

import (
  "fmt"
  "math"
)


func main() {
  //숫자 연산(산술, 비교)
  //타입이 같아야 가능
  //다른 타입끼리는 반드시 형 변환 후 연산
  //형 변환 없을 경우 예외(에러) 발생
  // +, -, *, %, /, <<, >>, &, ^

  //예제1
  var n1 uint8 = math.MaxUint8
  var n2 uint16 = math.MaxUint16
  var n3 uint32 = math.MaxUint32
  var n4 uint64 = math.MaxUint64

  fmt.Println("ex uint8 : ",  n1)
  fmt.Println("ex uint16 : ",  n2)
  fmt.Println("ex uint32 : ",  n3)
  fmt.Println("ex uint64 : ",  n4)

  fmt.Println("ex1 : ",  math.MaxInt8)
  fmt.Println("ex1 : ",  math.MaxInt16)
  fmt.Println("ex1 : ",  math.MaxInt32)
  fmt.Println("ex1 : ",  math.MaxInt64)
  fmt.Println("ex1 : ",  math.MaxFloat32)
  fmt.Println("ex1 : ",  math.MaxFloat64)

  n5 := 100000 // 기본값으로 int로 선언
  n6 := int16(10000) // 형변환이 일어난다.
  //n7 := uint8(300) // overflow넘치면 실행되지 않는다.
  n7 := uint8(100)

  //fmt.Println("ex2 : ", n5+n6); 예욉잘생. 형변환이 없을경우(동일하지않으면) 예외발생
  fmt.Println("ex2 : ", n5+int(n6))
  fmt.Println("ex2 : ", n6+int16(n7))

  fmt.Println("ex2 : ", n6 > int16(n7))
  fmt.Println("ex2 : ", n6-int16(n7) > 5000)


}
