---
title: "Fiber"
date: 2023-08-07T07:55:52+09:00
draft: false
---

# Fiber
Fiber란, golang의 `fasthttp`를 탑재한 웹 프레임워크다. 이때, 웹 프레임워크란 무엇일까?

## 웹 프레임워크
개발자의 경험이나 스타일에 따라 소스는 모두 다르다. 여러 사람이 공동작업하는 경우, 상대방의 스타일을 이해하는데 시간이 필요하며 개발 스킬의 차이가 클 경우 이해가 힘들 수 있다.

프레임워크는 구조화된 스크립트를 통해 개발자의 스크립트 패턴을 정형화할 수 있다. 또한 개발자가 반복적으로 해야하는 공통 부분을 최소호할 수 있도록 설계되어 있다.

프레임워크 장점
- 체계적인 코드관리로 유지보수가 용이하다
- 기본설계 및 기능 라이브러리를 제공하여 개발 생산성이 높다
- 코드 재사용성이 높다
- 추상화된 코드 제공을 통해 확장성이 좋다
> 추상화된 코드
> 
> 코드를 읽기 쉽도록 있던 코드를 분리하면서, 새로운 함수를 만들어 나가는 과정


다양한 웹 프레임워크들은 `MVC` 혹은 `MVT` 패턴을 지원하고 있다.


## MVC
MVC란, Model View Controller로 디자인 패턴 중 하나로 프로젝트를 구성할 때 세 가지 역할로 구분한 패턴이다.
> 디자인 패턴?
> 
> 소프트웨어를 설계할 때, 특정 맥락에서 자주 발생하는 고질적인 문제를이 또 발생했을 때 재사용할 수 있는 해결책

- Model : 어플리케이션 정보, 데이터를 의미한다. 또한 이러한 데이터와 정보의 가공을 책임지는 컴포넌트를 말한다.
- View : 데이터 및 객체의 입출력을 담당한다. 데이터를 기반으로 사람들이 볼 수 있는 화면이다.
- Controller : 사용자와 데이터를 잇는 다리 역할아다. 사용자가 데이터를 클릭하고, 입력, 수정하는 것에 대한 이벤트 처리를 담당한다.


## MVT
MVT는 MVC의 C가 Template의 T로 바뀐 것이다. 하지만 둘의 역할을 거의 비슷하다.

## Fiber
Fiber는 zero memory allocation 과 performance 를 염두해 빠른 개발을 위해 디자인되었다.

### 설치
```shell
$ go get -u github.com/gofiber/fiber/v2
```

### 예제
```go
func main() {
app := fiber.New()

app.Get("/", func(c *fiber.Ctx) error {
return c.SendString("Hello, World!")
})

app.Listen(":3000")
}

```

#### Basic Routing 예제
```go
app.Method(path string, ...func(*fiber.Ctx) error)
```

더 많은 코드 예제는 [API 문서](https://docs.gofiber.io/) 에서 참고할 수 있다.

### Zero Allocation
`*fiber.Ctx`로 부터 반환되는 일부 값들은 기본적으로 값이 변경될 수 있다.
따라서, 오직 `handler` 안에서만 context value를 사용해야하고, 어떠한 참조도 유지하지 말아야한다.

`handler`에서 return을 하면, context로 부터 얻은 어떠한 값들도 미래의 requests에 다시 재사용될 것이고 사용자의 제어에 벗어나 변경될 것이다.


### Basic Routing
routing은 어떻게 application이 특정 endpoint에 대한 client request에 반응할 것인지 에 대해 참조한다. endpoint는 URI(or path)와 특정 HTTP request method들로 구성된다.(GET, POST, PUT, ...)

### Fiber API
새로운 fiber app을 만들기 위해서 New 메서드를 사용하면 된다. 여기에 optional config를 추가할 수 있다.

```go
func New(config ...Config) *App
```

````go
// Custom config
app := fiber.New(fiber.Config{
    Prefork:       true,
    CaseSensitive: true,
    StrictRouting: true,
    ServerHeader:  "Fiber",
    AppName: "Test App v1.0.1"
})
// ...
````

## NewError
```go
func NewError(code int, message ...string) *Error
app.Get("/", func(c *fiber.Ctx) error {
    return fiber.NewError(782, "Custom error message")
})
```

## Tests
````go
func (app *App) Test(req *http.Request, msTimeout ...int) (*http.Response, error)
````

```go
// Create route with GET method for test:
app.Get("/", func(c *fiber.Ctx) error {
fmt.Println(c.BaseURL())              // => http://google.com
fmt.Println(c.Get("X-Custom-Header")) // => hi

return c.SendString("hello, World!")
})

// http.Request
req := httptest.NewRequest("GET", "http://google.com", nil)
req.Header.Set("X-Custom-Header", "hi")

// http.Response
resp, _ := app.Test(req)

// Do something with results:
if resp.StatusCode == fiber.StatusOK {
body, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(body)) // => Hello, World!
}
```

다음은 fiber 포스트에서는, fiber를 사용하여 코드를 작성하고 이에 대해 정리해보고자 한다.