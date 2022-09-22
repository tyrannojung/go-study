package main

import "fmt"

func main() {
  //예제1
   sum1 := 0
   for i := 0; i <= 100; i++ {
     sum1 += i
   }
   fmt.Println("ex1 : ", sum1)


   //예제 2
   // 위보다 가독성이 훨씬 좋아진다. (외국 코드는 해당 패턴이 훨씬 많다.)
   sum2, i := 0, 0
   for i <= 100 {
     sum2 += i
     i++
   }
   fmt.Println("ex2 : ", sum2)


   //예제 3
   //while 형태랑 비슷
   sum3, i := 0, 0
   for {
     if i > 100 {
       break
     }
     sum3 += i
     i++
   }
   fmt.Println("ex3 : ", sum3)

   //예제4
   //예제1 응용형 여러변수의 값을 다양하게 활용가능
   for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
     fmt.Println("ex4 : ", i, j)
   }

   //에러 발생
/*
   // 후치연산은 반환값이 없으므로 애러 i++은 반환값이 없다.
   for i, j := 0, 0; i <= 10; i++, j+=10 {
     fmt.Println("ex5 : ", i, j)
   }
*/



}
