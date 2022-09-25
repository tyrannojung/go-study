package main

import (
  "math"
)


func main() {
  //오버플로우 에러 : uint8을 넘으므로 에러
  var n1 uint8 = math.MaxUint8 + 1
  var n2 uint16 = math.MaxUint16 + 1
  var n3 uint32 = math.MaxUint32 + 1
  var n4 uint64 = math.MaxUint64 + 1

  //오버플로우 범위 미만
  // 양수를 저장해야하지만 음수를 저장함.
  var n5 uint8 = -1
  var n6 uint16 = -1
  var n7 uint32 = -1
  var n8 uint64 = -1

}
