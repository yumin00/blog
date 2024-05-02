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