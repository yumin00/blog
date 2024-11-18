---
title: "Cloud Functions 배포 자동화하기"
date: 2024-11-18T19:23:32+09:00
draft: false
categories :
- DevOps
---

이번 문서에는 Cloud Functions(지금은 Cloud Run Functions라고 명칭을 변경했다고 한다.) 배포 자동화 프로세스를 구축한 경험에 대해 작성해보려고 한다.

## Cloud Functions 배포 자동화 계기
![image](https://github.com/user-attachments/assets/e8d0dd75-7d27-44e9-915b-0b584d23066a)

원래 사내에서 Cloud Functions를 배포할 때 위 사진과 같이 함수를 작성한 뒤 로컬에서 gcloud CLI 명령어를 사용한 스크립트를 실행하여 배포했다. 이 방식으로 배포할 경우 다음과 같은 문제가 발생할 수 있다.

- 버전 관리 어려움: 로컬 스크립트는 변경 사항을 관리하기 어렵고, 코드의 특정 버전으로 롤백하기 힘들다. 
- 협업과 추적성 부족: 배포 기록이 남지 않아, 특정 배포 시점에 발생한 이슈를 추적하기 어렵다.
- 환경 의존성: 로컬 환경에 의존적이어서 환경별 설정이 통일되지 않을 수 있다.

이러한 문제를 해결하기 위해서 Cloud Functions 배포 자동화 프로세스를 구축해보았다!

## Cloud Functions 배포 자동화
Cloud Functions 배포 자동화 구축을 생각하면서 중요하게 생각한 것은 세가지 정도가 있다. 

- 배포를 기록한다.
- 배포 전 리뷰를 진행한다.
- 개발자는 코드만 작성한다.

배포를 기록하기 위해서 인프라를 코드로 관리하는 Terraform을 사용하고자 했다. 그리고 이를 Github에서 관리하여 Atlantis도 함께 도입하여 배포 전 리뷰도 진행될 수 있도록 하고자 했다.
추가로, 개발자가 코드 작성 이후에 배포를 위한 시간을 최소화하고 배포 과정에 대해 정확히 이해하고 있지 않더라도 env 값만 잘 작성하면 자동으로 배포가 되는 프로세스를 만들고 싶었다.

배포 프로세스 큰 틀에 대해서 먼저 고민했다.

<img width="1038" alt="image" src="https://github.com/user-attachments/assets/5b9fed8b-6852-420e-a911-034098de72f2">

이러한 구조를 생각해보았는데 요약하자면 다음과 같다. (진행할수록 생각한 것과 다르거나 보완하면 좋을 것들이 보여서 점점 개선해나갔다! 마지막에 최종 결과물이 함께 나온다 😀)

1. 개발자가 소스코드 작업 후 이를 Commit 및 PR을 작성한다.
2. PR Action에 대해 Atlantis가 작동하여 변경점을 확인하다.
3. Apply하면 이를 Cloud Functions에 배포한다.

### Terraform
먼저 테라폼을 작성해보았다. 이때 terraform [공식 문서](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloudfunctions_function)를 참고했다.

문서를 참고하여 작성하다보니 Cloud Functions는 배포를 할 때 소스코드를 .zip 파일로 Cloud Storage에 업로드한 후 이를 참조하여 배포한다고 한다. Cloud Storage는 Cloud Functions 코드의 배포와 버전 관리를 위한 저장소 역할을 하게 되는 것이다.

그래서 테라폼을 다음과 같이 작성해보았다.

```terraform
provider "google" {
  project = var.project_id
  region  = var.region
}

data "archive_file" "function_zip" {
  type        = "zip"
  source_dir  = var.source_dir
  output_path = "${path.module}/function-source.zip"
}

resource "google_storage_bucket_object" "version_folder" {
  name   = "${var.function_name}/${var.version_name}/"
  content = " "
  bucket = var.storage_bucket
}

resource "google_storage_bucket_object" "function_zip_upload" {
  name   = var.object_name
  bucket = google_storage_bucket_object.version_folder.bucket
  source = data.archive_file.function_zip.output_path
}

resource "google_cloudfunctions2_function" "function" {
  name        = var.function_name
  location    = var.region
  build_config {
    runtime = var.runtime
    entry_point = var.entry_point
    source {
      storage_source {
        bucket = var.storage_bucket
        object = google_storage_bucket_object.function_zip_upload.name
      }
    }
  }
  service_config {
    service_account_email = var.service_account
    environment_variables = var.environment_variables
  }
  event_trigger {
    trigger_region = var.region
    event_type = var.event_type
    pubsub_topic   = var.topic_name
    retry_policy = "RETRY_POLICY_RETRY"
  }
}
```

하나씩 뜯어보자면,

```terraform
data "archive_file" "function_zip" {
  type        = "zip"
  source_dir  = var.source_dir
  output_path = "${path.module}/function-source.zip"
}

resource "google_storage_bucket_object" "version_folder" {
  name   = "${var.function_name}/${var.version_name}/"
  content = " "
  bucket = var.storage_bucket
}

resource "google_storage_bucket_object" "function_zip_upload" {
  name   = var.object_name
  bucket = google_storage_bucket_object.version_folder.bucket
  source = data.archive_file.function_zip.output_path
}

```
이 부분은 개발자가 작성한 소스코드를 .zip 파일로 압축하고 Cloud Storage에 업로드하는 과정이다. Cloud Functions는 .zip 파일을 참조하여 배포하기 때문에 이를 반영하여 추가했다.
(이때 순차적으로 실행될 수 있도록 테라폼의 암묵적 의존성 규칙을 사용했다.)

> 암묵적 의존성
> 
> Terraform은 리소스의 속성에서 다른 리소스의 값을 참조하면 이를 기반으로 암묵적으로 의존 관계를 파악하고 리소스가 순서대로 생성되도록 한다.
> 
> 예를 들어, resource google_storage_bucket_object.function_zip_upload 리소스에서 bucket 속성에 google_storage_bucket_object.version_folder의 속성을 참조함으로써 해당 Cloud Storage에 version_folder가 먼저 생성된 후 파일이 업로드된다!

![image](https://github.com/user-attachments/assets/0125525e-f3bb-454c-831a-391d19cb4c49)

실제로 해당 파일을 실행시키면 설정된 Cloud Storage에 .zip 파일이 생성된다!

```terraform
resource "google_cloudfunctions2_function" "function" {
  name        = var.function_name
  location    = var.region
  build_config {
    runtime = var.runtime
    entry_point = var.entry_point
    source {
      storage_source {
        bucket = var.storage_bucket
        object = google_storage_bucket_object.function_zip_upload.name
      }
    }
  }
  service_config {
    service_account_email = var.service_account
    environment_variables = var.environment_variables
  }
  event_trigger {
    trigger_region = var.region
    event_type = var.event_type
    pubsub_topic   = var.topic_name
    retry_policy = "RETRY_POLICY_RETRY"
  }
}
```

이 부분이 실제로 Cloud Functions를 배포하는 코드이다. google_storage_bucket_object.function_zip_upload 리소스를 통해 업로드한 .zip 파일을 참조하여 실제로 배포가 된다!
(참고로 해당 코드는 Pub/Sub을 트리거로 할 때만 사용될 수 있다.)


cloud functions를 관리할 폴더의 구조는 다음과 같다.
```
cloudrun-functions
├── function_A
│   └── dev                  # DEV 배포 파일
│       └── main.tf
│       └── variables.tf
│   └── prod                  # PROD 배포 파일
│       └── main.tf
│       └── variables.tf
│   └── main.py              # 소스 코드
│   └── requirements.txt     # 소스 코드
│   └── README.md
├── function_B
│   └── dev                  # DEV 배포 파일
│       └── main.tf
│       └── variables.tf
│   └── prod                 # PROD 배포 파일
│       └── main.tf
│       └── variables.tf
│   └── main.py              # 소스 코드
│   └── requirements.txt     # 소스 코드
│   └── README.md
├── function_C
│   └── dev                  # DEV 배포 파일
│       └── main.tf
│       └── variables.tf
│   └── prod                 # PROD 배포 파일
│       └── main.tf
│       └── variables.tf
│   └── main.py              # 소스 코드
│   └── requirements.txt     # 소스 코드
│   └── README.md
```

이렇게 구조를 잡고 보니 각 function마다 환경마다 테라폼 파일이 각각 관리되어야 한다는 문제점이 발생했다. 실제로 테라폼 파일은 동일한데 이렇게 각각 관리하다 보면 문제가 발생할 수 있다.
그래서 테라폼을 모듈로 따로 빼는 방법을 생각했다.

테라폼 모듈을 따로 배서 이 테라폼 모듈을 가져와 사용하는 `terragurn.hcl`을 사용하면 function에서는 테라폼 모듈이 아닌 테라그런트를 사용하여 variable만 작성해주면 된다!

그런데 테라폼 모듈을 공유하면 또 발생하는 문제가 있다. 내가 처음에 작성한 테라폼 모듈은 Pub/Sub 트리거만 가능하다는 것이다. Pub/Sub뿐만 아니라 Firestore 트리거도 가능한 하나의 테라폼 모듈을 만들어야했다. 그래서 테라폼의 dynamic block을 사용해보았다.

```terraform
resource "google_cloudfunctions2_function" "function" {
  name        = var.function_name
  location    = var.region

  build_config {
    runtime = var.runtime
    entry_point = var.entry_point
    source {
        storage_source {
          bucket = var.storage_bucket
          object = google_storage_bucket_object.function_zip_upload.name
        }
    }
  }

  service_config {
    service_account_email = var.service_account
    environment_variables = var.environment_variables
  }
  
  dynamic "event_trigger" {
    for_each = var.trigger_type == "pubsub" ? [1] : []
    content {
      trigger_region = var.region
      event_type    = var.event_type
      pubsub_topic  = var.topic_name
      retry_policy  = "RETRY_POLICY_RETRY"
    }
  }

  dynamic "event_trigger" {
    for_each = var.trigger_type == "firestore" ? [1] : []
    content {
      trigger_region = var.region
      event_type    = var.event_type
      retry_policy  = "RETRY_POLICY_RETRY"
      dynamic "event_filters" {
        for_each = [
          {
            attribute = "database"
            value     = var.firestore_database
            operator  = null
          },
          {
            attribute = "document"
            value     = var.firestore_document_path
            operator  = "match-path-pattern"
          }
        ]
        content {
          attribute = event_filters.value.attribute
          value     = event_filters.value.value
          operator  = event_filters.value.operator
        }
      }
    }
  }
}
```

내용은 거의 비슷하지만, event_trigger 부분에서 dynamic block을 사용했다. 어떤 trigger를 사용하느냐에 따라 서로 다른 content를 넣어서 공용으로 사용할 수 있게 수정했다.


### Terragrunt
공용으로 사용할 수 있는 테라폼 모듈을 만들었으니 그 다음으로 각 function에서 사용할 Terragurn.hcl을 작성해보았다. 여기에서 내가 처음으로 말한 3번 개발자는 코드만 작성한다. 라는 것이 적용될 수 있도록 노력했다.
최대한 모양을 잡아놓고 개발자는 여기에 function을 배포할 cloud, env 값들을 정의하면 쉽게 배포할 수 있도록 하고 싶었다.

```terraform
terraform {
  source = "git::git@github.com:terraform-module-repo"
}

locals {
  project = "dev"
  region = "asia-northeast3"
  runtime = "python311"
  trigger_type = "pubsub"
  event_type = "google.cloud.pubsub.topic.v1.messagePublished"

  function_name = "a_function_trigger_by_pub_sub"
  topic_name = "projects/dev/topics/pub_sub_name"
  entry_point = "main"
  environment_variables = {
    A_ENV = 1
    B_ENV    = "HI"
    C_ENV = 200
  }
  version = "1.0.0"
}

inputs = {
  object_name="cloudfunctions/${local.function_name}/${local.version}/function-source.zip"
  project_id   = local.project
  region         = local.region
  trigger_type = local.trigger_type
  function_name = local.function_name
  runtime = local.runtime
  topic_name     = local.topic_name
  entry_point = local.entry_point
  event_type= local.event_type
  environment_variables = local.environment_variables
  version_name = local.version
  source_dir = "${get_terragrunt_dir()}/../"
}
```

이런식으로 terragunt.hcl 파일에는 소스로 사용하는 테라폼에 필요한 inputs 값을 넣어줘야 한다. 각 function마다 버전도 같이 관리할 수 있도록 version variable도 추가해주었다.

이때 내가 중요하게 생각한 것은 inputs에 최대한 값을 많이 넣어놓고 각 코드마다 다른 값들만 작성하면 배포할 수 있게 하려고 했다. 그래서 실제로 locals에 있는 값들만 자신이 만든 functions에 맞게 수정하면 된다!

### Github Actions & Atlantis
![image](https://github.com/user-attachments/assets/95f3c07f-cb7c-407e-83f6-078652d815e4)

Github에서는 PR을 트리거 삼아 Atlantis Webhook이 실행되도록 했다. 내가 설정한 branch로 PR이 작성되면 `atlantis plan`이 실행되고 terragurn.hcl 파일의 변경점에 대해 알려준다.
개발자는 변경된 이력을 보고 문제가 없다면 `atlantis apply`를 통해 배포를 진행할 수 있다.

## 결과
<img width="1065" alt="image" src="https://github.com/user-attachments/assets/7ffac2a4-abca-4dc4-a715-53c2758d22b4">
결과적으로 결과물은 위와 같다.

1. 개발자가 소스코드를 작성하여 Cloud Functions를 관리하는 Repo(`cloud-functions-repo`)에 Commit 및 PR을 생성한다.
2. `cloud-functions-repo`에서는 PR을 트리거 삼아 Atalantis가 실행된다.
3. 개발자는 변경된 이력을 보고 문제가 없다면 `atlantis apply`를 통해 배포를 진행한다. 
   1. 상태 파일과 소스코드가 Cloud Storage에 저장된다.
   2. .zip 파일을 참조하여 Cloud Functions가 배포된다.

Cloud Functions 배포 자동화를 통해 다음과 같은 이점을 기대할 수 있다.

- 버전 관리 및 추적성 강화 
  - 모든 배포 코드를 GitHub에 관리하여 코드 변경사항을 체계적으로 추적할 수 있다.
  - 특정 버전으로 손쉽게 롤백이 가능하다. 
- 자동화된 협업 환경 제공
  - Atlantis를 통해 GitHub Pull Request 단위로 테라폼 배포 작업을 실행하여, 코드 리뷰와 배포를 하나의 프로세스로 통합할 수 있다.

배포 자동화 프로세스를 진행하면서 해당 프로세스의 아쉬운 점이 있다. 현재 환경 변수는 `terragurn.hcl` 파일에 직접 입력하고 있는데 Google Secret Manager를 사용한다면 환경 변수를 더욱 안전하게 관리할 수 있을 것 같다!

그리고 해당 프로세스는 배포만 진행하다보니 만약 테스트가 진행되지 않으면 오류가 있는 코드가 그대로 배포될 수 있다는 문제가 있다. Github Action에서 Atalantis를 통해 Diff를 확인하기 전에 먼저 Unit Test를 자동으로 실행한다면 문제 있는 코드가 실수로 배포되는 현상은 막을 수 있을 것 같다.