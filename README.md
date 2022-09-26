## 💡 Introduction
- GO STUDY

## DAY6

- 배열(Array)
  - 길이(len)와 용량(cap)이 같다.
  - 길이가 고정되 있으며, 값이 복사 전달된다. 또한 비교연산자를 사용 할 수 있다.
  - 배열은 개수와, 안에 들어가는 데이터 타입이 맞아야 한다.
    - ex)  `arr := [5]int{1, 2, 3, 4, 5}`
    - ex)  `arr6 := [...]int{1, 2, 3, 4, 5}` // 길이에 확신이 없을때 사용한다.
    
- 슬라이스(Slice)
  - 길이가 가변, 래퍼런스타입(참조 값을 전달), 비교연산자 사용 불가
  - 길이와 용량을 동적으로 할당 가능
  - 선언방식
    - 배열처럼 선언
      - ex) `slice := []int{}` // 배열은 크기를 선언했지만,  slice는 선언하지 않는다. 
    - make함수(자료형, 길이, 용량[생략시, 길이])
      - ex) `slice []int = make([]int, 5, 10)` // int타입, 길이5, 용량10
      - ![image](https://user-images.githubusercontent.com/58019931/192328542-fc47703f-1af4-48b7-b408-0a712af97ba5.png)   
      - ex) `slice := make([]int, 5)` // int타입, 길이5, 용량5
      - 만약 용량을 처음 설정보다 초과하면 재할당이 일어나 성능의 저하가 일어나며, 너무 적게쓰면 메모리의 낭비가 이루어짐. 길이와 용량을 `엔지니어 입장에서 성능을 고려해서 디테일`하게 사용할 수 있게 제공해준다.
 
- 맵(Map)
  - 레퍼런스 타입(참조 값 전달), 비교 연산자 사용 불가능
  - 참조타입(Key)로 사용 불가능, 값(Value)으로 모든 타입 사용가능
  - Make 함수 및 축약(리터럴)로 초기화 가능 & 순서없음
 
- 포인터(Pointer)
  - \*(에스터리스크)로 사용 & nil 로 초기화(nil == 0)
  - 

## DAY5

- Boolean
  - 조건부 논리 연산자랑 주로 사용 : !(not), ||(or), &&(and)
    - 암묵적 형변환이 안된다.(0 Nil --> false, 1 -->ture 로 변환 없음)
    
- 숫자형 기초
  - 정수, 실수, 복소수
    - 32bit, 64bit, unsigned(양수)
    - 8진수(0), 10진수, 16진수(0x)
    - 내가 데이터 타입을 정확히 알고 있으면 32,64를 써도 되는데, 보편적으로 int만쓰는게 속편함(내부적으로 최적화 진행함)
    
- 숫자형 연산
  - 산술 비교 : 타입이 같아야 가능
    - 다른타입끼리는 반드시 형 변환 후 연산
    - 형변환 없을경우 예외(에러) 발생
    - 형변환시, 자료형이 넘치는 경우 해당자료형의 max값 출력 
    - 연산시 overflow(연산 후, 자료형에 넘치는경우 실행 x)
    
- 문자열 기초
  - 큰따옴표(""), Backtick백틱(``), 작은따옴표('')
  - Golang은 char타입이 존재하지 않는다. --> rune문자 코드 값으로 표현
    - ex) var str1 string = "c:\\go_study\\src\\" // --> c:\go_study\src\
    - ex) str2 := &#96;c:\go_study\src\ &#96; // excape사용하지 않더라도 그대로 사용()
  - 문자열 길이구하기(len, RuneCountInString, rune) 
    - `var str4 string = "안녕하세요."`
    - ex) len : `len(str4)` = 16출력 // byte
    - ex) RuneCountInString : `utf8.RuneCountInString(str4)` = 6출력
    - ex) rune : `len([]rune(str4))` = 6출력
  - 문자열은 배열로 취급 --> 반복문에서 순회 및 인덱스로 접근 가능.
    - 인덱스 접근시 문자의 아스키코드 출력, printf 사용시 실제 문자 출력
    
- 문자열 연산
  - 추출
    -  var str1 string = "Golang"
    -  ex) `str1[0:2]` : Go, `str1[0]` : 71, `str1[3:]` : ang, `str1[:4]` : Gola, `str1[1:3]` : ol
  - 비교
    - ==, !=
    - 주의)  >, < 사용가능 but `Golang, World` 두개를 비교했을때, 바이트의 비교가 아닌, 아스키코드를 비교. World가 더 큼
  - 조합(결합)
    -`str1 := "GO can make"; str2 := " programming very productive"` 
    - 일반연산 :
      - ex) `str1+str2` : GO can make programming very productive
    - 결합(join) :
      - ex) strSet := []string{}; strSet = append(strSet, str1); strSet = append(strSet, str2); strings.Join(strSet, "") : GO can make programming very productive
    - 문자열은 한번 선언하면 메모리에서 수정이 불가하다. 새로 계속 생성해서 비효율적. --> go에서는 join 자바에서는 stringbuffer가 효율적으로 개선된 함수다.
    
## DAY4

- 패키지
  - 코드 구조화 및 재사용, 모든 언어들은 발전하면서 api, component화가 잘 되어있다. 재사용성↑(결합도↓ 응집도↑)
  - Go에서는 작은 패키지들을 결합해서 프로그래밍을 작성할 것을 권고
  - fmt와 같은 기본 패키지들은 GO ROOT에 있음
  - main패키지는 특별하게 인식(공유 라이브러리라고 인식x, 프로그램 시작점 즉 start point로 인식)
  - 패키지이름 : 디렉토리 이름 (같은 패키지내 소스파일은 디렉토리명을 패키지명으로 사용)
  - 네이밍규칙
    - 소문자 : private
    - 대문자 : public

- 접근제어 및 Alias
  - 변수, 상수, 함수, 메서드, 구조체 등 식별자
  - 소문자 : 패키지 외부 접근 불가(패키지 내에서만 접근 가능)
  - 대문자 : 패키지 외부에서 접근 가능
  - 별칭 및 빈(_) 식별자 사용가능

- 초기화 메소드(init)
  - 패키지 로드시 가장 먼저 호출
  ![화면 캡처 2022-09-25 162939](https://user-images.githubusercontent.com/58019931/192134623-7efd3575-1606-4872-a232-5b03c1737553.png)


## DAY3

- GO 특징
  - 모호하거나, 애매한 문법 금지(go철학)
  - 후치 연산자 허용 i++ but 전치 연산자 비허용 ++i x 문법애러
  - 증감연산 반환값 x sum := i++ 사용불가
  - 포인터(Pointer 허용, 연산 비허용)
  - 문장 끝 세미콜론(;) 주의 (자동으로 끝에 세미콜론 삽입)
  
- GO 코드서식지정
  - 터미널명령어 : gofmt -w ~~.go --> 코드가 뒤죽박죽이라도 자동 변환 

- 조건문(If)
  - 반드시 Boolean 검사 -> 1, 0 (사용불가 : 자동 형 변환 불가)
  - 조건 소괄호 사용안됨 && 조건 한문장일때 괄호생략 안됨
  
- 조건문(Switch) 
  - switch 뒤 표현식(Expression) 사용 가능 & case 뒤 표현식(Expression) 사용 가능
  - break가 default이므로, fallthough 존재
  - case 여러가지 나열가능
  
- 반복문(For)
  - Go에서 제공되는 '유일한' 반복문
  - section3 for 여러 사례들을 한번씩 보자
  

## DAY2

- GO 장점
  - 간결한 문법 및 단순함(코드정리 지원)
  - 병행 프로그래밍 지원
    - go의 핵심 장점
  - 정적 타입 및 동적 실행
  - 간편한 협업 지원
  - 컴파일 및 실행속도가 빠름
  - 제네릭 및 예외처리 미지원
    - 장점이자 단점이 될 수 있음

- 변수
  - var
    - 숫자 첫글자x, 대소문자 구분o, 문자, 숫자, 밑줄, 특수기호 사용 가능
    - 선언했는데 사용하지 않으면 애러가 나옴
    - var a string = "Test1"
    - var b = "Test2" // go 내부적으로 자동 컨버팅
    - 선언을 했는데, 초기화 하지 않으면 기본값 int = 0 string = null
  - 짧은선언
    - ex) shortVar1 := 3 
    - 함수 안에서는 사용 (전역x) 1회성 사용. - 메소드가 끝나면 메모리에서 지워진다.
    - 제한된 범위의 함수 내에서 사용 할 경우 코드 가독성↑
  - const
    - 사용 초기화, 한 번 선언 후 값 변경 금지, 고정 된 값 관리 용
  - iota
    - 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음
    - 규칙적으로 증가. 중간값 생략가능. 규칙성은 유지.


## DAY1

- 기본다운로드 
  - https://go.dev/, https://atom.io/

- ATOM Package 설정
  - go-plus : https://atom.io/packages/go-plus
  - terminal : https://atom.io/packages/platformio-ide-terminal
  - script : https://atom.io/packages/script/

- 기본폴더구조
  - bin
    - 소스코드를 컴파일해서, 실행가능한 파일을 만들면, bin폴더에 저장이 자동으로 된다.
  - pkg
    - 외부 패키지 라이브러리
  - src
    - 소스파일

- 기본명령어
  - go run .go파일(ctrl + shift + b)
    - 빌드하지 않고 바로시작(단위테스트 느낌)

  - go build .go파일
    - 실행가능한 파일을 만드는것(테스트개념)

  - go install
    - bin폴더를 가면 session1.exe파일이 만들어짐, go는 폴더이름을 프로젝트이름으로 인식, 개발이 완료됬다고 판단 배포하는 명령어 느낌(패키지에 만약 라이브러리가 있으면, 의존관계에 있는 것 까지 다 빌드후에 인스톨을 해줌)
