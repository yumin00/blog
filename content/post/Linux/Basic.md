---
title: "Linux 기초 명령어에 대해 알아보자"
date: 2024-08-08T19:13:13+09:00
draft: true
categories :
---

### touch
빈 파일 생성

### cat
파일 출력

### head
파일의 앞 부분을 출력

ex) head -n 3 hello.txt : hello.txt 파일의 3번째 줄까지 출력

### tail
파일의 마지막 부분을 출력

### echo "hi" > hello.txt
hello.txt 파일에 hi 붙여넣기

### echo "bye" >> hello.txt
hello.txt 파일의 마지막 줄에 bye 붙여넣기

### cp
복사

ex) cp hello.txt hello2.txt : hello.txt 를 hello2.txt 로 복사

### comm
파일 비교

ex) comm hello.txt hello2.txt: hello.txt와 hello2.txt 파일을 비교한 내용 출력

