---
title: "Unit Test"
date: 2024-09-04T18:36:54+09:00
draft: true
categories :
- GoLang
---


# 유닛테스트란?

유닛테스트란, 소프트웨어 개발 단계에서 초기 단계 테스트 방법으로, 유닛 단위(메서드, 클래스)가 제대로 동작하는지 확인하는 과정입니다.

유닛테스트의 주요 목적은 다음과 같습니다.

- 개별 코드가 정확히 작동하는지 검증
- 코드 수정이나 새로운 기능 추가 시 기존 코드에 영향을 주지 않는지 검증
- 시스템 내에서 발생할 수 있는 문제를 사전에 방지

# dta-wir-api Architecture

현재 `dta-wir-api` Go Clean Architecture로 Delivery - Usecase - Repository 로 구성되어 있습니다. 따라서 각 레이어별로 테스트가 필요합니다.

## Repository

Repository 레이어에서는 데이터베이스에 직접 접근하는 계층으로 의도한대로 CRUD Command가 제대로 실행되었는지 확인해야 합니다.

## Usecase

usecase 레이어는 비즈니스 로직을 구현하는 계층으로 로직에 좀 더 집중한 테스트가 필요합니다.

## Delivery

Delivery 레이어는 외부에서 올바른 요청이 들어왔는지 확인하고, Usecase 레이어로부터 반환한 데이터를 올바르게 처리하고 외부에 반환하는지를 테스트해야 합니다.

# 테스트 방법

## Mocking

유닛 테스트의 목적은 코드 유닛 자체의 동작을 검증하는 것이기 때문에 외부 의존성을 제거하고 레이어 별로 독립적인 테스트가 필요합니다. 모킹을 사용하면 외부 의존성을 제거하고, 테스트하려는 코드만 검증이 가능합니다.

예를 들어, 데이터베이스를 사용하는 코드를 테스트할 때 실제 데이터베이스에 접근하는 대신, 데이터베이스 호출을 모킹하여 테스트의 독립성을 보장할 수 있습니다.

이전에 hun께서 자동으로 mock 인터페이스를 만들 수 있는 스크립트를 작성해주셨습니다. https://github.com/weltcorp/gops-tools

`dta-wir-api` 의 `mock.sh` 스크립트를 사용하면 모듈 내 domain의 interface를 기준으로 mock을 생성할 수 있습니다.

```
#!/bin/sh

cd "$(dirname "$0")/.."
curl https://raw.githubusercontent.com/weltcorp/gops-tools/v1/mockery_installer.sh | sh -
```