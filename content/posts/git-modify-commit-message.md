---
title: "[Git] Modify Commit Message"
date: 2023-07-12T20:26:39+09:00
draft: false
---

새로운 글을 포스팅하여 commit을 남기던 중, commit message 를 잘못 작성하였다.

local에서 push하기 전, commit message 수정은 간단하다.

## commit 을 PUSH 하기 전 Message 를 수정하고 싶을 때
````
git commit --amend
````

해당 명령어를 통해 이 전에 작성했던 commit message를 확인할 수 있다.

``i`` 를 눌러 `insert mode` 로 변경한 뒤, commit message를 수정하면 된다.
수정이 완료되었다면 `esc` 를 눌러 `insert mode` 에서 빠져나오고,
`:wq` 를 통해 저장을 한 후 다시 push를 하면 변경된 commit message로 push가 가능하다.