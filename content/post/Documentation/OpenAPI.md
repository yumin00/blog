---
title: "Swagger(OpenAPI)에 대해 알아보자"
date: 2023-11-24T00:08:56+09:00
draft: false
categories :
- Documentation
---

# Swagger(OpenAPI)
API 명세 구현 및 API 문서화를 위해 Swagger를 사용해보고자, Swagger(OpenAPI)에 대해 학습해보려고 한다.

## 개념
Swagger는 OpenAPI 명세(Specification)를 구현하는 도구 중 하나이다. OpenAPI는 RESTful 웹 서비스를 설계, 생성, 문서화 및 사용하기 위한 표준이다.

OpenAPI를 사용하면 개발자는 API의 모든 리소스와 메서드에 대한 정확한 정보를 명확하고 일관된 형식으로 제공할 수 있다.

## 특징
- 표준화된 문서화

OpenAPI를 사용하면 API의 구조, 사용 방법, 파라미터, 응답 형식 등을 명시적으로 문서화할 수 있다. 따라서 API 문서화와 관리를 쉽게 할 수 있다.

- 언어 중립적

OpenAPI 스펙은 YAML이나 JSON 형식으로 작성된다. 이러한 형식은 언어에 구애받지 않기 때문에 다양한 프로그래밍 언어와 플랫폼에서 사용할 수 있다.

- 클라이언트 및 서버 코드 생성

OpenAPI 스펙을 바탕으로 다양한 프로그래밍 언어에 대한 클라이언트 라이브러리 및 서버 스텁을 자동으로 생성할 수 있다.

- API 테스팅 및 인터랙션

OpenAPI를 기반으로 하는 도구(예: Swagger UI)를 사용하여 API를 직접 탐색하고 테스트할 수 있다. 사용자는 API의 엔드포인트를 쉽게 시험해보고, 요청을 
보내고, 응답을 받을 수 있다.

- API 보안 및 권한 부여

OpenAPI 스펙은 보안 정의 및 권한 부여 방법을 포함하여 API의 보안 관련 정보를 명시할 수 있게 해준다.

## 적용
그럼 이제 직접 Swagger를 직접 적용해보자!

Swagger 적용에 앞서 프로토콜 버퍼(ProtoBuf)와 연결이 필요하다.

### 1. 프로토콜 버퍼 정의 작성
- 정해진 API에 대하여 protobuf 작성이 필요하다.
```protobuf
option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
  info: {
    title: "Auth"
    description: "Auth module"
    version: "1.0"
    contact: {
      name: "gRPC-Gateway project"
      url: "https://github.com/grpc-ecosystem/grpc-gateway"
      email: "none@example.com"
    }
    license: {
      name: "BSD 3-Clause License"
      url: "https://github.com/grpc-ecosystem/grpc-gateway/blob/main/LICENSE"
    }
  }
  schemes: HTTPS;
  consumes: "application/json";
  produces: "application/json";
  host: "";
  base_path: "";
};
```
Swagger 문서의 기본 정보를 설정할 수 있다. 여기에는 문서의 제목, 설명, 버전, 연락처, 라이선스, 사용 프로토콜, 데이터 포맷, 호스트 및 기본 경로 정보가 포함된다.

```protobuf
service AuthData {
  rpc Login(LoginRequest) returns (LoginResponse) {
    option (google.api.http) = { post: "/v1/auth" };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      responses: {
        key: "200",
        value: {
          description: "OK",
          schema: {
            json_schema: {
              ref: ""
            }
          }
        }
      }
      parameters: {
        headers: [
          {
            name: "x-request-dtx-dst-protocol",
            type: STRING,
            description: "for network routing header. must set 'http'",
            required: true,
          },
          {
            name: "Authorization",
            type: STRING,
            description: "for authentication header",
            required: true,
          }
        ]
      }
    };
  }
}
```

해당 RPC 메서드에 대한 Swagger 문서 정보를 설정한다. 여기에는 응답, 파라미터, 헤더 등에 대한 정보가 포함된다.


```protobuf
    message LoginRequest {
      option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema) = {
        json_schema: {
          title: "LoginRequest"
          description: "LoginRequest"
          required: ["userId"]
        }
      };
      int32 userId = 1 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {description: "The UserID field."}];
    }
    
    message LoginResponse {
      int32 statusCode = 1;
      option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema) = {
        json_schema: {
          title: "LoginRequest"
          description: "LoginResponse"
        }
      };
    }
```
각 메시지 타입에 대한 Swagger 문서 정보를 설정한다.

- json schema: 메시지 타입의 json 스키마를 정의한다. 
  - title: 메시지 타입의 제목을 제공한다. ex) `LoginRequest`, `LoginResponse`
  - description: 메시지 타입에 대한 자세한 설명을 제공한다.
  - required: 필수 필드를 나타낸다 ex) `["userId"]`

### 2. Protobuf To Swagger
스웨거를 위한 프로토 파일 작성을 완료했다면, 이를 빌드시켜 json 파일을 만들어야 한다.

나는 프로토 파일 빌드를 위해 다음과 같은 파일을 만들었다.

```shell
#!/bin/zsh
cd ./proto
protoc -I. --openapiv2_out=.  \
auth/auth.proto
```

이는 protobuf 파일을 사용하여 swagger 스펙의 JSON 파일을 생성한다. `protoc` 컴파일러를 사용하여 .proto 파일을 Swagger JSON으로 변환한다.

- `cd ./proto`: build시켜야 하는 proto 파일이 proto 라는 폴더 안에 있었기 때문에 proto로 이동시켰다.
- `protoc`: 프로토콜 버퍼 파일을 다른 형식으로 변환하는 데 사용되는 컴파일러이다.
- `-I.`: Import 경로를 현재 디렉토리로 지정한다. 이는 protoc가 .proto 파일을 찾을 때 현재 디렉토리를 포함하여 검색하도록 지시하는 역할을 한다. 
- `--openapiv2_out=.`: Swagger 스펙의 JSON 파일을 현재 디렉토리에 출력하도록 지정한다.
- `auth/auth.proto`: 변환할 프로토콜 버퍼 파일의 경로를 지정한다.

해당 빌드 파일을 실행시키면 다음과 같은 swagger json 파일을 얻을 수 있다.

<img width="200" alt="image" src="https://github.com/yumin00/Swagger/assets/130362583/e51a9daf-3ac9-42a6-a3de-28bee26622e1">



### 3. swagger-ui docker 파일 작성
swagger json 파일까지 생성했다면, 이제 해당 json 파일을 웹으로 띄워야 한다. 이때, swagger-ui를 사용할 수 있다.

처음에는 swagger-ui 깃헙에서 해당 파일을 다운로드 받아 적용할 수도 있지만, 도커파일을 이용하는 방법이 있다.

나는 다음과 같이 Docker Compose를 이용하여 Swagger UI를 도커 컨테이너로 실행할 수 있는 파일을 만들엇다.
```yaml
version: '2'
services:
  swagger-ui:
    image: swaggerapi/swagger-ui
    container_name: "swagger-ui-container"
    ports:
      - "8082:8080"
    volumes:
      - ./proto/auth/auth.swagger.json:/swagger.json
    environment:
      SWAGGER_JSON: /swagger.json
```

- volume

호스트 시스템의 ./proto/auth/auth.swagger.json 파일을 컨테이너 내의 /swagger.json으로 마운트한다.
이는 컨테이너가 로컬 파일 시스템의 해당 파일에 접근할 수 있게 한다.

즉, 로컬에 있는 파일을 컨테이너로 복사하고, 그 경로 그대로 환경 변수로 저장하면 스웨거가 부팅하면서 해당 파일을 바라보고 실행이 된다.

- SWAGGER_JSON 

명세 파일을 해당 파일로 사용하겠다라는 의미이다.

### 4. 실행
아래 명령어를 실행하여 "locahost:8082"에서 swagger ui를 확인할 수 있다! 
```
docker compose up -d
```

나는 서버가 실행되면 함께 swagger ui를 확인할 수 있도록 구현하고자 main에서 실행시키는 코드를 구현해보았다.
```go
package main

import (
	"log"
	"os/exec"
)

func main() {
	// 'docker compose down' 명령 실행
	downCmd := exec.Command("docker-compose", "down")
	if err := downCmd.Run(); err != nil {
		log.Fatalf("Failed to execute 'docker-compose down': %s", err)
	}

	// 'docker compose up -d' 명령 실행
	upCmd := exec.Command("docker-compose", "up", "-d")
	if err := upCmd.Run(); err != nil {
		log.Fatalf("Failed to execute 'docker-compose up -d': %s", err)
	}

	log.Println("Docker Compose commands executed successfully.")
}
```

## 마치며
swagger에서는 다양한 옵션을 사용하여 request와 response를 구현할 수 있다. 이번에 auth 관련 API를 작성하면서 swagger의 옵션들에 대해 깊게 공부해보고자 한다.