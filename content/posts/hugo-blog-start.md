---
title: "Hugo Devlog 제작기"
date: 2023-07-09T18:27:45+09:00
draft: true
---

개발 공부를 하며 학습한 내용을 남기기 위해 Hugo와 Github Page를 통해 Devlog를 만들기로 했다.


### Hugo 선택 이유
Github Page를 이용하여 블로그를 작성할 때 사용할 수 있는 다양한 generator가 존재한다.

github에서는 github의 코파운더가 개발한 Jekyll을 권장하고 있다. Jekyll은 github에서 권장하고 있는 generator이기 때문에 많은 사람들이 사용하여 자료가 다양하지만,
포스팅이 늘어날수록 (20개 이하임에도 불구하고) build가 오래 걸린다는 issue가 있었다.

그래서 build가 빠르다는 hugo를 선택하게 되었다. 현재, go 언어를 주로 사용하고 있기 때문에 customizing이 용이할 것이라는 생각도 hugo 선택에 많은 영향을 미치기도 했다.

### Hugo 프로젝트 생성
Hugo 공식 문서의 [Quick Start](https://gohugo.io/getting-started/quick-start/)를 참고하여 프로젝트 생성을 진행했다.
```cgo
hugo new site quickstart
cd quickstart
git init
git submodule add https://github.com/kakawait/hugo-tranquilpeak-theme.git themes/tranquilpeak
echo "theme = 'tranquilpeak'" >> hugo.toml
hugo server
```
git init을 한 뒤 hugo themes 중, 카테고리와 태그를 쉽게 구별할 수 있을 것 같은 테마를 선택하여 submodule로 추가해주었다.
- submodule이란, 다른 git 레포지토리를 나의 현재 레포지토리에 import하여 사용하는 것
- git clone을 하면 root에 있는 git은 하위에 있는 git을 관리할 수 없지만, submodule을 사용하면 root에 있는 git이 하위의 git을 관리할 수 있다.

### content 작성 방법
hugo는 content 생성이 간단하다.
```cgo
hugo new posts/Hugo-blog.md
```
hugo new를 통해 `archetypes`의 `default.md`를 보고 땡겨 title, date, draft의 타입을 잡아 자동으로 생성해준다.
예를 들어, `default.md`가 다음과 같을 때,
```cgo
---
title: "{{ replace .Name "-" " " | title }}"
date: {{ .Date }}
draft: true
---
```
`hugo new post/hugo-blog.md`를 진행하면 자동으로 기본 타입이 생성되며 수정이 가능하다.
```cgo
---
title: "Hugo blog"
date: 2023-07-09T18:27:45+09:00
draft: true
---
```

hugo new를 통해 content가 생성되었다면 글은 mark down 형식으로 작성할 수 있다. 


### 블로그 꾸미기
Hugo 공식 문서대로 진행하여 hugo server를 build하면, 처음 보고 선택한 테마와 많이 달랐다.

hugo-tranquilpeak-theme 처럼 꾸미기 위해 [user.md](https://github.com/kakawait/hugo-tranquilpeak-theme/blob/master/docs/user.md)를 참고했다.