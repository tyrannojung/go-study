package main

import "fmt"

func main() {
  //짧은 선언
  // go에만 있는것
  // 함수 안에서만 사용 (전역 x). 선언 후 할당 예외 발생
  // 1회성만 사용  메소드가 끝나면 메모리에서 지워진다.
  // 주로 제한된 범위의 함수내에서 사용 할 경우 코드 가독성을 높일 수 있따

  shortVar1 := 3
  shortVar2 := "Test"
  shortVar3 := false

  // shortVar1 := 2  초기화 후 재 할당을 할 수 없다. : 예외발생

  fmt.Println("shortVar1 : ", shortVar1, "shortVar2 : ", shortVar2, "shortVar3 : ", shortVar3)

  //Example
  // 제한된 범위의 함수 뿐 아니라 조건문에서도 많이 사용
  if i := 10; i < 11 {
    fmt.Println("Short Variable Test Success");

  }

}
