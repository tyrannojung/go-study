package main

import "fmt"

func main() {
  // 여러개 선언 정말 많이 쓰는것.
  // 이 기개가 높이, 무개, 작동하고 있는지 var로 묶어 쓰면 가독성이 좋다. 그룹을 묶어놓는 느낌
  var (
    name string = "machine"
    height int32
    weight float32
    isRunning bool
  )

  height = 250
  weight = 350.56
  isRunning = true

  fmt.Println("name : ", name, "height : ", height, "weight : ", weight, "isRunning : ", isRunning)

}
