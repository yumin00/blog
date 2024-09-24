---
title: "테라폼(Terraform)에 대해 알아보자"
date: 2024-09-24T13:38:49+09:00
draft: false
categories :
- Infra
---

사내에서 Pub/Sub을 관리하기 위해 인프라를 코드로 관리할 수 있는 테라폼을 활용해보고자 한다. 이전에 먼저 테라폼이 무엇인지에 대해 좀 더 자세히 공부해보고자 한다.

# 테라폼
- 인프라를 코드로 관리하는 도구
- HCL(HashiCorp Configuration Language) 이라는 테라폼의 고유한 언어를 사용하여 리소르를 생성할 수 있음
- 테라폼은 현재 인프라 상태를 추적하여, 변경된 부분만 효율적으로 업데이트할 수 있음
- 인프라 코드를 모듈화하여 재사용할 수 있도록 지원
- `terraform plan`: 변경 사항을 미리 확인
- `terraform apply`: 변경 사항을 인프라에 배포

# 테라폼 동작 방식
테라폼의 동작 방식은 설정 파일 작성, 플랜 생성, 적용, 상태 관리 단계로 나눌 수 있다.

## 1. 설정 파일 작성
인프라를 정의하는 코드를 작성하는 단계이다. `.tf` 확장자를 사용하고, HCL 혹은 JSON 형식으로 작성한다.

- 프로바이더: 사용할 클라우드 플랫폼 정의
- 리소스: 실제로 생성할 인프라 리소스
- 변수: 코드 재사용을 위한 값들

pub/sub에서 스키마를 생성하는 예시 코드를 작성해보았다.

```
provider "google" {
  project = var.project
}

resource "google_pubsub_schema" "test_schema_yumin" {
  name = "projects/${var.project}/schemas/test_schema_yumin"
  type = "AVRO"

  definition = <<EOF
{
  "type": "record",
  "name": "User",
  "fields": [
    {"name": "name", "type": "string"},
    {"name": "age", "type": "int"}
  ]
}
EOF
}
```

## 2. 플랜 생성 `terraform plan`
설정 파일 작성 후, 플랜을 생성한다. 사용자가 정의한 코드와 현재 인프라의 상태를 비교하여 어떤 리소스를 생성/변경/삭제할 것인지 예측할 수 있다.

설정 파일을 작성할 때 variables 을 사용했다면, `terraform plan` 의 출력 예시는 다음과 같다.

![image](https://github.com/user-attachments/assets/d97ae874-b19c-48ce-869e-09b314c601ed)

여기에서 사용하고자 하는 variables 의 값을 입력하면 된다. 나는 project를 variable 로 설정했기에 아래와 같이 입력해주었다.

![image](https://github.com/user-attachments/assets/a54c618a-886e-4d39-885f-77fd72485db9)

![image](https://github.com/user-attachments/assets/092d3249-6a60-49ad-a455-68947d846a4c)

사진과 같이 현재 인프라 상태와 변경하고자 하는 상태의 다른점을 확인할 수 있다.

## 3. 적용 `terraform apply`
생성한 플랜을 바탕으로 실제 인프라에 적용하는 단계이다.

variables 를 사용하면 플랜 생성 단계와 동일하게 값을 입력해아 하고, 결과는 다음과 같이 나온다.

![image](https://github.com/user-attachments/assets/bfac9543-e633-4b03-9ccc-c3e70d9ba07d)

실제로 Pub/Sub에 토픽과 구독자가 생성된 것을 확인할 수 있다!

## 4. 상태 관리
테라폼은 리소스를 생성할 때 상태 파일(State file)을 생성하고 이를 통해 인프라의 현재 상태를 추적하고 기록한다.
상태 파일은 인프라의 실제 상태와 코드를 동기화하고 관리하는 데 매우 중요한 역할을 한다.

terraform apply 를 실행하면 상태 파일인 `.tfstate`와 `.tfsate.backup` 파일이 생성 혹은 갱신된다. 여기에서 인프라의 상태를 관리하는 것이다.
![image](https://github.com/user-attachments/assets/4e60bbaa-bb12-4ffc-98b2-3e7727f27f89)

- `.tfstate`: 현재 관리 중인 인프라의 실제 상태
- `.tfsate.backup`: 이전 실행에서의 상태 파일 백업

만약 여러 개발자가 한번에 각자의 로컬에서 한 테라폼 파일을 통해 배포한다고 가정해보자. 그러면 인프라 상태가 동기화가 되지 않거나 충돌이 발생할 수 있다.

따라서 GCS와 같은 중앙 집중식 서버를 사용하면 상태 파일이 클라우드에 저장되고, 여러 사용자가 동시에 테라폼을 사용해도 상태 파일이 동기화되고 충돌을 방지할 수 있다!

그러면 동시에 요청이 들어왔을 때 테라폼은 이를 어떻게 관리할까?

테라폼에서 apply 를 할 때, Lock을 만들어서 해당 세션만 접근 가능하게 한다. apply를 하여 Lock을 만들면, 해당 세션은 gcs에 접근할 수 있고, apply가 완료되면 요청하여 lock을 없애는다.

그러면 다음 테라폼 apply에서 lock을 만들어 apply하는 과정을 반복한다!