package main

import (
  "fmt"
)

func main() {
  //맵(Map)
  //맵 조회 할 경우에 주의 할 점

  //예제1
  map1 := map[string]int{
    "apple":15,
    "banana":115,
    "orange":1115,
    "lemon":0,
  }

  value1 := map1["lemon"]
  value2, ok1 := map1["kiwi"] //없을경우 value에 대한 기본 초기화값이 나온다. int 이므로 0(예외발생x) but 해당 값이 없는지, 값인지 구분이안됨.
  value3, ok := map1["kiwi"]//그러므로 두번째 리턴값은 존재 유뮤의 true false를 return한다.

  fmt.Println("ex1 : ", value1)
  fmt.Println("ex1 : ", value2, ok1)
  fmt.Println("ex1 : ", value3, ok)//두 번째 리턴 값으로 키 존재 유무 확인

  //예제2
  if value, ok := map1["kiwi"]; ok {
    fmt.Println("ex2 : ", value)
  } else {
    fmt.Println("ex2 : kiwi is not exist!")
  }

  if value, ok := map1["banana"]; ok {
    fmt.Println("ex2 : ", value)
  } else {
    fmt.Println("ex2 : banana is not exist!")
  }

  if _, ok := map1["kiwi"]; !ok {
    fmt.Println("ex2 : kiwi is not exist!")
  }

}
