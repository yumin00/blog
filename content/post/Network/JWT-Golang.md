---
title: "GO에서 JWT를 사용해보자"
date: 2023-11-13T22:03:33+09:00
draft: true
categories :
- Network
---

# GO에서 JWT를 사용해보자
# JWT-GO
Go에서 JWT를 발급 받고 검증받을 때 다음과 같은 라이브러리를 사용할 수 있다.

[golang-jwt](https://github.com/golang-jwt/jwt)

## 1. 토큰 발급
```go
package main

import (
"fmt"
"log"
"time"

jwt "github.com/dgrijalva/jwt-go"
)

// MyCustomClaims 사용자 정의 클레임 구조체
type MyCustomClaims struct {
	UserID string `json:"userid"`
	jwt.StandardClaims
}

// createToken JWT 생성 함수
func createToken(secretKey []byte) string {
	// 사용자 정의 클레임 설정
	claims := MyCustomClaims{
		UserID: "1234567890",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "test",
		},
	}

	// 새 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 토큰 서명
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

func main() {
	// 비밀키 (실제 사용 시 안전하게 보관)
	secretKey := []byte("your-256-bit-secret")

	// JWT 생성
	tokenString := createToken(secretKey)
	fmt.Println("Token:", tokenString)
}

```

이렇게, custom token claims 구조체를 만들고, 토큰을 생성하는 코드를 구현할 수 있다.

## 2. 토큰 검증
```go
package main

import (
	"fmt"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

// MyCustomClaims 사용자 정의 클레임 구조체
type MyCustomClaims struct {
	UserID string `json:"userid"`
	jwt.StandardClaims
}

// parseToken 토큰 검증 및 파싱 함수
func parseToken(tokenStr string, secretKey []byte) (*MyCustomClaims, error) {
	// 토큰 파싱 및 검증
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func main() {
	// 비밀키
	secretKey := []byte("your-256-bit-secret")

	// 예제 토큰 (실제 토큰 사용 필요)
	tokenString := "your.jwt.token"

	// 토큰 검증
	claims, err := parseToken(tokenString, secretKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("UserID: %s\n", claims.UserID)
}

```