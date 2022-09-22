package main

import "fmt"

func main() {
  //예제1
  //loop lable 해당 결로로 빠져나감.
  //Looplable 밑에는 continue 나 break를 사용해야함. 즉 for 문이 안나오면 애러
  Loop1:
    for i := 0; i < 5; i++ {
      for j := 0; j < 5; j++{
        if i == 2 && j == 4 {
          break Loop1
        }
        fmt.Println("ex1 : ", i, j)
      }
    }

    //예제2
    //continue
    for i := 0; i < 10; i++ {
      if i % 2 == 0{
        continue
      }
      fmt.Println("ex2 : ", i)
    }

    //애러 발생(루프 레이블 밑에 관련이 없는 소스코드시 애러)
/*
    Loop2:
    fmt.Println()
*/


Loop2:
  for i := 0; i <3; i++ {
    for j := 0; j < 3; j++ {
      if i == 1 && j == 2 {
        continue Loop2
      }
      fmt.Println("ex3 : ", i, j)
    }
  }

}
