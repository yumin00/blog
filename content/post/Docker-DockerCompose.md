---
title: "Docker와 Docker Compose"
date: 2024-01-25T19:20:28+09:00
draft: true
---

Docker와 Docker Compose에 대해 알아보자

도커와 도커 컴포즈에 대해 공부하기 전, **컨테이너** 에 대한 이해가 필요하다.

> 컨테이너
> 
> 컨테이너란, 소프트웨어 서비스를 어디서든 실행할 수 있도록 소프트웨어 서비스에 필요한 특정 버전의 프로그래밍 언어 런타임 및 라이브러리와 같은 종속성 항목과 애플리케이션 코드를 함께 포함하는 경량 패키지이다.

# 도커
## 도커 이미지
도커 이미지는 컨테이너를 만드는 데 사용되는 읽기 전용의 템플릿이다.
컨테이너 실행에 필요한 파일, 설정값 등을 포함하고 있는 도커 파일을 만든 후 Docker 빌드하여 이미지를 만든다.

## 도커 컨테이너
도커 이미지를 실행한 상태를 말한다.

# 도커 컴포즈
도커 컴포즈란, 단일 서버에서 여러 개의 컨테이너를 하나의 서비스로 정의해 컨테이너 묶음으로 관리할 수 있는 작업 환경을 제공해주는 관리 도구이다.

여러 개의 컨테이너가 하나의 애플리케이션으로 동작할 때, 이를 테스트하려면 각 컨테이너를 하나씩 생성해야 한다. 현재 진행중인 프로젝트를 예시로 들면, 스웨거 / 서버 / DB 가 필요한데 테스트를 진행하려면 세 개의 컨테이너를 생성해야 한다.

도커 컴포즈를 사용하면 세 개의 컨테이너를 생성하지 않고, 단일 서버에 세 개의 컨테이너를 하나의 서비스로 정의하여 테스트할 수 있다.

## docker-compose.yml
도커 컴포즈를 사용하려면 컨테이너 설정을 저장해 놓은 `docker-compose.yml` 파일이 필요하다. 

## 도커 컴포즈 yaml 파일 옵션
### 버전 정의 `version`
도커 컴포즈 파일의 버전을 명시한다. 주로 3 또는 3.x를 사용한다.

### 서비스 정의 `services`
 `services` 는 생성할 컨테이너들을 정의하한다.

### 서비스 내 옵션
- `image`: 각 컨테이너의 이미지를 정의한다. 이 이미지는 Docker 에 이미 정의된 이미지를 사용할 수도 있고 `build`를 통해 Dockerfile 로부터 빌드된 이미지일 수도 있다.
- `container_name`: 각 컨테이너의 이름을 정의한다.
- `build`: 이미지를 빌드할 경우 도커파일의 경로를 명시한다.
  - `context`: 도커파일에서 사용할 context를 지정한다.
  - `dockerfile`: 빌드할 도커파일의 이름을 지정한다.
- `ports`: 호스트와 컨테이너 간의 포트 매핑을 정의한다. (`호스트_포트`:`컨테이너_포트`)
  - 호스트_포트 : 외부에서 컨테이너에 접근하기 위해 사용할 포트
  - 컨테이너_포트 : 컨테이너 내부에서 실행되는 애플리케이션이 사용하는 포트
- `environment`: 컨테이너 내부에서 사용할 환경변수를 지정한다.
- `volumes`: 볼륨을 통해 컨테이너의 데이터의 영속성을 보장할 수 있다.


이를 바탕으로 직접 도커 컴포즈 yaml 파일을 작성해보고자 한다.

실행이 필요한 컨테이너는 swagger / server / mysql 이 있다.

## docker-compose.yml 작성
### Swagger UI
```yaml
version: '3' // 버전 정의
services: // 서비스 정의
  swagger-ui: //  컨테이너 정의
    image: swaggerapi/swagger-ui // 이미지 정의
    container_name: "swagger-ui-container" // 컨테이너 이름 정의
    ports: // 포트 정의
      - "8082:8080" // 호스트_포트:컨테이너_포트
    volumes:
      - ./server-proto/proto/auth/auth.swagger.json:/swagger.json // 컨테이너에 주입할 데이터 정의
    environment:
      SWAGGER_JSON: /swagger.json // 환경 변수 지정
```
- `image`: 이미지는 도커에 이미 정의된 swaggerapi/swagger-ui 라는 이미지를 사용한다.
- `volumes`: Swagger UI에 주입할 데이터는 미리 정의해놓은 swagger json 파일을 사용하고자 한다.
- `environment`:
  - SWAGGER_JSON: Swagger UI는 SWAGGER_JSON 환경변수를 사용하여 Swagger 파일의 위치를 지정한다. 해당 환경변수가 설정되면 Swagger UI는 이 경로에서 Swagger 파일을 로드하고, API 문서로 표시한다.
  - /swagger.json: 컨테이너 내부의 파일 경로를 나타낸다. 즉, /swagger.json 경로의 파일을 Swagger UI 가 로드하도록 지시한다.

### server
```yaml
version: '3'
services:
  server:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: server
    ports:
      - "8080:8080"
```
- `ports`: Swagger UI를 통해 서버가 호출되어야 하기 때문에, Swagger UI에서 정의된 URL에 맞춰 호스트 포트를 정의한다.

### db
```yaml
version: '3'
services:
  mysql:
    image: mysql:8
    restart: always
    container_name: mysql
    ports:
      - "3306:3306"
    environment:
      MYSQL_ROOT_PASSWORD: ${MYSQL_PASSWORD}
    env_file:
      - app.env
    volumes:
      - ./db/mysql/data:/var/lib/mysql
      - ./db/mysql/init:/docker-entrypoint-initdb.d
```