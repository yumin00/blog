---
title: "SQL JOIN 에 대해 알아보자"
date: 2024-04-26T17:44:01+09:00
draft: false
categories :
- DataBase
- DataEngineering
---

# SQL JOIN
SQL에서 JOIN 은 테이블의 column을 기반으로 두 개 이상의 테이블을 결합하는 것을 말한다.

## INNER JOIN
![image](https://github.com/yumin00/blog/assets/130362583/c2f9f4fc-85a5-4ece-92e9-f5ef92eb09ad)
![image](https://github.com/yumin00/blog/assets/130362583/8442cdb1-b128-469e-a5e6-503b5215b533)

INNER JOIN 은 두 개 이상의 테이블에서 column을 기반으로 일치하는 값이 있는 데이터를 반환한다.

## LEFT JOIN (LEFT OUTER JOIN)
![image](https://github.com/yumin00/blog/assets/130362583/be016125-7aea-4fab-a642-936b7e9e62a7)
![image](https://github.com/yumin00/blog/assets/130362583/b26986f8-36c0-4c59-acdc-06513b780f29)

LEF JOIN(LEFT OUTER JOIN) 은 column을 기반으로 일치하는 값이 있는 데이터뿐만 아니라 왼쪽 테이블의 데이터를 반환한다.
일치하는 항목이 없으면 Null을 반환한다.

## RIGHT JOIN (RIGHT OUTER JOIN)
![image](https://github.com/yumin00/blog/assets/130362583/fcaa7035-2c1c-4fe0-b963-7bbb395e78e9)
![image](https://github.com/yumin00/blog/assets/130362583/a984eac8-aa67-4aa0-ad8a-97458a8b1205)

RIGHT JOIN(RIGHT OUTER JOIN) 은 column을 기반으로 일치하는 값이 있는 데이터뿐만 아니라 오른쪽 테이블의 데이터를 반환한다.
일치하는 항목이 없으면 Null을 반환한다.

## FULL OUTER JOIN
![image](https://github.com/yumin00/blog/assets/130362583/417f21a6-bba4-4c9a-b56c-45b8346960f7)
![image](https://github.com/yumin00/blog/assets/130362583/277a43fc-2b9e-41b7-ba72-67e557c21ff2)
FULL OUTER JOIN 은 column을 기반으로 일치하는 값이 있는 데이터뿐만 아니라 왼쪽, 오른쪽 테이블의 데이터를 반환한다.
일치하는 항목이 없으면 Null을 반환한다.

## CROSS JOIN
![image](https://github.com/yumin00/blog/assets/130362583/c0c331b7-5f86-4007-b0e0-4b009d1ca8b9)
![image](https://github.com/yumin00/blog/assets/130362583/0ddf34fa-65be-414f-b221-00b774ca48d4)

CROSS JOIN(상호 조인)은 한쪽 테이블의 모든 행과 다른쪽 테이블의 모든 행의 데이터를 반환한다.
결합되지 않고, 각 테이블의 모든 행이 다른 테이블의 모든 행과 결합되어, 상호 조인의 결과값은 각 행의 개수를 곱한 수만큼 된다.
카테시안 곱(CARTESIAN PRODUCT) 이라고도 한다.

## 차집합

SQL JOIN에서 차집합은 LEFT JOIN || RIGHT JOIN 을 사용하여 두 테이블을 결합하고,
오른쪽 || 왼쪽의 열이 NULL 인 행을 필터링하여 데이터를 반환할 수 있다. 예시는 다음과 같다.

```sql
LEFT JOIN Managers M ON E.EmployeeID = M.EmployeeID
WHERE M.EmployeeID IS NULL;
```