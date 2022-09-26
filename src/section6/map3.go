package main

import (
  "fmt"
)

func main() {
  //맵(Map)
  //맵 값 변경 및 삭제

  //예제1
  map1 := map[string]string {
    "daum": "kakao",
    "naver": "nhn",
    "google": "american",
    "home": "xi",
  }

  fmt.Println("ex1 : ", map1 )

  map1["home2"] = "docksan" //추가
  fmt.Println("ex1 : ", map1 )

  map1["home2"] = "daesung" // 수정
  fmt.Println("ex1 : ", map1 )

  delete(map1, "home2") // 삭제
  fmt.Println("ex1 : ", map1 )

}
