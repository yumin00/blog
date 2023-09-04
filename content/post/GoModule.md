
---
title: "Go Module"
date: 2023-08-06T20:29:48+09:00
draft: false
---

## 모듈
Golang 1.11 이전에는 모듈이라는 기능이 존재하지 않았다. 1.13 때는 모듈은 선택사항이었지만, 1.16부터는 기본 사양이 되었다.

Go Module이란, library dependency를 관리해주는 것이다.
Golang에서 module은 Package의 모음으로, 한 개의 모듈은 다수의 패키지를 포함할 수 있다. 모듈을 통해 패키지들의 종속성을 관리할 수 있고, 모듈은 패키지 관리 시스템으로 활용이 된다.

모듈은 패키지를 트리 형식으로 관리하고, root 폴더에 `go.mod` 파일을 생성하여 모듈을 정의하고 종속성 정보를 관리하게 된다.

## 종속성 관리
종속성 관리란, 프로젝트가 의존하는 라이브러리 및 패키지의 버전을 관리하는 것이다.


## Go Mudule 사용 방법
### go mod init
```go
go mod init MODULE_NAME
```

주로 module name에는 Github 저장소 주소나 URL을 사용한다.

### go mod tidy
Golang에서 외부 패키지를 사용할 때, 해당 패키지를 다운받을 필요가 있다.
이때, `go mod tidy` 를 통해 외부 패키지 다운로드가 가능하다.