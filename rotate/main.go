package main

import "fmt"

// 기존의 슬라이스 내부 배열을 수정
// rotateOne 을 d번 반복
// time: O(n * d), space: O(1)
func rotate1(slice []int, k int) {
	for ; 0 < k; k-- {
		rotateOne(slice)
	}
}

func rotateOne(slice []int) {
	end := len(slice) - 1
	t := slice[0]
	i := 0
	for ; i < end; i++ {
		slice[i] = slice[i+1]
	}
	slice[i] = t
}

// 새로운 배열의 슬라이스를 return
// 원본 배열은 바뀌지 않음
// time: O(n), space: O(d)
func rotate2(slice []int, k int) []int {
	return append(slice[k:], slice[:k]...)
}

func main() {
	s := []int{0, 1, 2, 3}
	r1 := rotate2(s, 1)
	fmt.Println(r1)
	r2 := rotate2(s, 2)
	fmt.Println(r2)

	rotate1(s, 2)
	fmt.Println(s)
}
