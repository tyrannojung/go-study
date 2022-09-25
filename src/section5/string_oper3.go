package main

import (
  "fmt"
  "strings"
)

func main() {
  //예제1(결합 : 일반연산) 2등
  // string은 한번 선언하면, 메모리에서 수정이 불가하다. 새로 계속 생성하게 된다. 비효율. 자바에서는 stringbuffer + append사용.
  str1 := "GO can make programming very productive because - any type can be given methods," +
          "which opens up interesting design" +
          "possibilities"
  str2 := "Can somebody please explain how can this feature increase productivity?"

  fmt.Println("ex1 : ", str1+str2)
  //예제2(결합 : Join) --> 추천(성능)
  strSet := []string{} // 슬라이스 선언
  strSet = append(strSet, str1)
  strSet = append(strSet, str2)

  fmt.Println("ex2 : ", strings.Join(strSet, ""))

}
