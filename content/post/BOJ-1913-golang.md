---
title: "[BOJ] 1913 : 달팽이 (Golang)"
date: 2023-07-22T20:30:45+09:00
draft: false
---

## 문제
- BOJ 1913: 달팽이 https://www.acmicpc.net/problem/1913

## 풀이
- (0, 0)에 n*n 값을 넣고, 1씩 줄여나가며 1이 되면 멈추는 코드를 구현하고자 했다.

1. `n`값과 좌표를 알고자 하는 `location` 값을 입력 받는다.
2. `snail`을 2차원 배열로 만들고 length를 n으로 설정해준다.
3. `snail`의 x좌표와 y좌표를 기본 0으로 세팅한다.
4. 1씩 줄여나가면서 `snail`의 좌표가 좌표 범위를 벗어나거나 값이 이미 채워져있으면 방향을 틀어야하기 때문에 방향을 전환할 수 있는 나침반과 같은 배열(dx, dy)을 만든다.
5. 달팽이가 움직이는 방향은 `아래, 오른쪽, 위, 왼쪽` 을 반복하기 때문에 `dx`와 `dy` 값을 미리 설정해놓는다.
6. `dx`와 `dy`의 방향을 정해주는 `idx` 값을 0으로 설정하고 4보다 작을 때까지 무한 루프로 `snail`에 값을 넣어준다.
7. 다음 좌표가 범위 안에 있고, `nx`와 `ny`가 0보다 크거나 같고, 아직 값을 넣지 않은 위치라면 이전 좌표의 수보다 1 작은 값을 넣어준다.
8. 만약 해당 값이 1이라면 멈춘다.
9. 다음 좌표를 찾기 위해서 값을 갱신해준다.
10. 한 변이 다 채워지면 방향을 틀어야하기 때문에 `idx` 값을 늘려준다.
11. 아직 1이 채워지지 않았지만 `idx`가 4면 for 루프가 끝나기 때문에, 다시 0으로 돌아가 `아래, 오른쪽, 위, 왼쪽`을 반복해준다.
12. `snail`에 `location`와 같은 값이 채워진다면 해당 좌표를 `locationX`와 `locationY`에 저장해놓는다.

## 첫 번째 시도 코드
```cgo
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	var location int
	
	//1. n 값 입력받기
	fmt.Fscan(r, &n)
	
	//1. location 값 입력 받기
	fmt.Fscan(r, &location)
	
	//2. snail을 2차원 배열로 만들고, length를 n으로 설정하기
	snail := make([][]int, n)
	for i := 0; i < n; i++ {
		snail[i] = make([]int, n)
	}

	var locationX int
	var locationY int
	
	//3. snail의 x,y 좌표를 0으로 세팅
	x := 0
	y := 0
	
	//4,5. 나침반 배열 dx, dy 생성 및 값 설정
	dx := make([]int, 4)
	dx[0] = 1  //아래
	dx[1] = 0  //오른쪽
	dx[2] = -1 // 위
	dx[3] = 0  //왼쪽

	dy := make([]int, 4)
	dy[0] = 0  //아래
	dy[1] = 1  //오른쪽
	dy[2] = 0  //위
	dy[3] = -1 //왼쪽

	snail[x][y] = n * n
	
	//6. 방향을 정해주는 idx 값을 0으로 설정하고, 4보다 작을 때까지 무한 루프로 돌리기
	idx := 0
	for idx < 4 {
		nx := x + dx[idx]
		ny := y + dy[idx]
		
		//7. 다음 좌표가 범위 안에 있고, nx와 ny가 0보다 크거나 같고, 아직 값을 넣지 않은 위치라면
		if nx >= 0 && ny >= 0 && nx < n && ny < n && snail[nx][ny] == 0 {
		    
		    //7. 이전 좌표 값보다 1 작은 값을 넣어준다.
		    snail[nx][ny] = snail[x][y] - 1
		    
		    //12. snail의 값이 location과 같을 때, 해당 좌표를 locationX와 locationY에 저장해놓는다.
			if snail[nx][ny] == location {
				locationX = nx + 1
				locationY = ny + 1
			}
			
			//8. 만약 해당 값이 1이 됐다면 멈춘다
			if snail[nx][ny] == 1 {
				break
			}
			
			//9. 다음 좌표를 찾기 위해서 x와 y를 갱신해준다.
			x = nx
			y = ny
		} else {
		    //10. 한 변이 다 채워지면 방향을 틀어야하기 때문에 idx 값을 늘려준다
			idx++
		}
		
        //11. 아직 값이 다 채워지지 않았지만 idx가 4면 루프가 for 루프가 끝나기 때문에 다시 idx를 0으로 설정해준다.
		if idx == 4 {
			idx = 0
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprintf(w, "%d ", snail[i][j])
		}
		fmt.Fprintln(w)
	}
	fmt.Fprint(w, locationX, locationY)

	w.Flush()

}

```

## 오류 해결
2번의 실패가 있었다.

1. n이 1인 경우
n이 1인 경우, 기존 코드로 진행했을 때 break 하는 부분이 없었기 때문에 실패했다. 그래서 다음 코드를 추가했다.
```cgo
// ...

	for idx < 4 {
		if snail[x][y] == 1 {
			break
		}
		
// ...
```

특이 케이스로 n이 1인 경우에는 (0, 0)까지만 저장해놓고 바로 멈출 수 있도록 코드를 수정했다.

2. `location`이 n^2 인 경우

`location`이 n^2 인 경우, location과 같은 값을 찾을 때 `(1,0)`부터 시작하기 때문에, `locationX`와 `locationY` 값이 설정되지 않아 자동으로 `0 0` 으로 출력되는 것이 문제였다.

따라서 `locationX`와 `locationY` 를 출력할 때, 해당 값에 1을 더해줄 수 있도록 다음과 같이 코드를 수정했다.
```cgo
// ...

    if nx >= 0 && ny >= 0 && nx < n && ny < n && snail[nx][ny] == 0 {
			snail[nx][ny] = snail[x][y] - 1
			if snail[nx][ny] == location {
				locationX = nx
				locationY = ny
			}
// ...

	fmt.Fprint(w, locationX+1, locationY+1)
	
//...
```

## 정답 코드
해당 특이 케이스를 반영할 수 있도록 코드를 수정하여 결국 성공할 수 있었다!
```cgo
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	var location int
	
	//1. n 값 입력받기
	fmt.Fscan(r, &n)
	
	//1. location 값 입력 받기
	fmt.Fscan(r, &location)
	
	//2. snail을 2차원 배열로 만들고, length를 n으로 설정하기
	snail := make([][]int, n)
	for i := 0; i < n; i++ {
		snail[i] = make([]int, n)
	}

	var locationX int
	var locationY int
	
	//3. snail의 x,y 좌표를 0으로 세팅
	x := 0
	y := 0
	
	//4,5. 나침반 배열 dx, dy 생성 및 값 설정
	dx := make([]int, 4)
	dx[0] = 1  //아래
	dx[1] = 0  //오른쪽
	dx[2] = -1 // 위
	dx[3] = 0  //왼쪽

	dy := make([]int, 4)
	dy[0] = 0  //아래
	dy[1] = 1  //오른쪽
	dy[2] = 0  //위
	dy[3] = -1 //왼쪽

	snail[x][y] = n * n
	
	//6. 방향을 정해주는 idx 값을 0으로 설정하고, 4보다 작을 때까지 무한 루프로 돌리기
	idx := 0
	for idx < 4 {
	    // ** n이 1인 경우 바로 중단
		if snail[x][y] == 1 {
			break
		}
		
		nx := x + dx[idx]
		ny := y + dy[idx]
		
		//7. 다음 좌표가 범위 안에 있고, nx와 ny가 0보다 크거나 같고, 아직 값을 넣지 않은 위치라면
		if nx >= 0 && ny >= 0 && nx < n && ny < n && snail[nx][ny] == 0 {
		    
		    //7. 이전 좌표 값보다 1 작은 값을 넣어준다.
		    snail[nx][ny] = snail[x][y] - 1
		    
		    //12. snail의 값이 location과 같을 때, 해당 좌표를 locationX와 locationY에 저장해놓는다.
			if snail[nx][ny] == location {
				locationX = nx
				locationY = ny
			}
			
			//8. 만약 해당 값이 1이 됐다면 멈춘다
			if snail[nx][ny] == 1 {
				break
			}
			
			//9. 다음 좌표를 찾기 위해서 x와 y를 갱신해준다.
			x = nx
			y = ny
		} else {
		    //10. 한 변이 다 채워지면 방향을 틀어야하기 때문에 idx 값을 늘려준다
			idx++
		}
		
        //11. 아직 값이 다 채워지지 않았지만 idx가 4면 루프가 for 루프가 끝나기 때문에 다시 idx를 0으로 설정해준다.
		if idx == 4 {
			idx = 0
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprintf(w, "%d ", snail[i][j])
		}
		fmt.Fprintln(w)
	}
	
	//** locaiton n^2인 경우를 위해 좌표 값은 아래에서 1을 더해준다.
	fmt.Fprint(w, locationX+1, locationY+1)

	w.Flush()

}
```
