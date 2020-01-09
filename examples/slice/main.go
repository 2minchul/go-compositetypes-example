package main

import "fmt"

func main() {
	{
		// 리터럴로 슬라이스 선언
		s := []string{"a", "b", "c", "d"}
		fmt.Printf("%#v\n", s) // []string{"안녕하세요", "감사해요", "잘있어요", "다시만나요"}
	}

	{
		// make([]T, len, cap) 로 슬라이스 선언
		s := make([]byte, 5, 5)
		fmt.Printf("%#v\n", s) // []byte{0x0, 0x0, 0x0, 0x0, 0x0}
	}

	{
		// cap 생략 가능 - 생략시 len 과 같음
		s1 := make([]byte, 5, 5)
		s2 := make([]byte, 5)
		fmt.Printf("len:%d, cap:%d\n", len(s1), cap(s2)) // len:5, cap:5
	}

	{
		// empty slice 생성
		// empty slice 는 nil 이 아니다
		// 내부 배열에 원소를 갖고있지 않기 때문에 저장소를 할당하지 않는다
		emptySlice := make([]byte, 0, 0)
		fmt.Printf("%#v\n", emptySlice) // []byte{}
		fmt.Printf("%#v\n", []byte{})   // []byte{}
		fmt.Println(emptySlice == nil)  // false
	}

	{
		// slice 의 zero-value 는 nil slice 이다.
		// nil slice 는 nil 과 같다
		var nilSlice []byte
		fmt.Printf("%#v\n", nilSlice)   // []byte(nil)
		fmt.Println(nilSlice == nil)    // true
		fmt.Println([]byte(nil) == nil) // true
	}

	{
		// nil slice 와 empty slice 의 len 과 cap 은 모두 0 이다.
		fmt.Println(len([]byte{}))    // 0
		fmt.Println(len([]byte(nil))) // 0
		fmt.Println(cap([]byte{}))    // 0
		fmt.Println(cap([]byte(nil))) // 0
	}

	{
		// slice 는 배열 또는 슬라이스를 '슬라이싱' 하여 만들 수 있다.
		arr := [...]string{"g", "o", "l", "a", "n", "g"} // array
		s := []string{"g", "o", "l", "a", "n", "g"}      // slice
		fmt.Printf("%#v\n", arr[1:4])                    // []string{"o", "l", "a"}
		fmt.Printf("%#v\n", s[1:4])                      // []string{"o", "l", "a"}
	}

	{
		// 슬라이싱 표현의 start 와 end 는 생략 가능하다
		// start 의 기본값은 0
		// end 의 기본값은 len(slice)
		arr := [6]string{"g", "o", "l", "a", "n", "g"}
		fmt.Printf("%#v\n", arr[:4]) // []string{"g", "o", "l", "a"}
		fmt.Printf("%#v\n", arr[1:]) // []string{"o", "l", "a", "n", "g"}
		fmt.Printf("%#v\n", arr[:])  // []string{"g", "o", "l", "a", "n", "g"}
		// arr[:] 는 array 를 slice 로 만들때 자주 쓰인다.
	}

	{
		// 슬라이스는 같은 배열을 공유 한다.
		s := []string{"r", "o", "a", "d"}
		ad := s[1:] // 원래 배열을 가르키는 새로운 slice value 를 만듦
		ad[0] = "e"
		fmt.Printf("%#v\n", s) // []string{"r", "e", "a", "d"}

		// 슬라이스를 함수에 전달하면 함수가 내부 배열의 원소를 변경 할 수 있다.
		modify := func(slice []string) {
			slice[0] = "l"
		}
		modify(s)
		fmt.Printf("%#v\n", s) // []string{"l", "e", "a", "d"}
	}

	{
		// 슬라이스는 cap 만큼 다시 확장 가능하다.
		// - cap 이상으로 늘릴 수 없다
		// - 이전 요소로 다시 접근 할 수 없다 (0 이하로 다시 줄일 수 없다)
		s := make([]byte, 5) // len:5, cap:5
		s = s[2:4]           // len:2, cap:3
		s = s[:3]            // len:3, cap:3
	}

	{
		// 슬라이스를 함수의 파라미터로 넘기면 slice value 가 복사 된다. (call by value)
		someFunc := func(s []int) {
			// fmt.Printf("%p", s) 는 s 의 첫번째 원소의 포인터를 출력한다 <fmt.Printf("%p", &s[0]) 과 같음>
			// fmt.Printf("%p", &s) 는 slice value 자체의 포인터를 출력한다
			fmt.Printf("s: %p, &s: %p\n", s, &s)
		}

		a := []int{1, 2, 3}
		someFunc(a) // 실행할 때마다 slice value 의 포인터가 매번 달라짐
		someFunc(a)
		someFunc(a)

		// 함수 안에서 len 을 수정해도 함수 밖의 slice value 에 영향을 미치지 못한다
		modifyLen := func(s []int) {
			s = s[:2]
			fmt.Printf("len: %d\n", len(s))
		}
		modifyLen(a)                    // len: 2
		fmt.Printf("len: %d\n", len(a)) // len: 3

	}

	{
		// copy(dst, src []T) 로 슬라이스의 값을 복사 할 수 있다.
		// 길이가 다른 슬라이스의 복사도 가능하다. (단 dst 가 src 보다 같거나 커야함)

		s1 := []int{1, 2, 3}
		s2 := make([]int, 3)
		copy(s2, s1)
		fmt.Println(s2) // [1 2 3]
	}

	{
		// append(s []T, x ...T) 로 슬라이스의 끝에 원소를 추가 할 수 있습니다.

		// append 는 슬라이스 내부 배열의 끝에 원소를 추가 하고, len 이 확장된 새로운 slice value 를 return 합니다
		s := make([]int, 1, 3)
		fmt.Printf("%v len:%d, cap:%d\n", s, len(s), cap(s)) // [0] len:1, cap:3

		s = append(s, 10)
		fmt.Printf("%v len:%d, cap:%d\n", s, len(s), cap(s)) // [0 10] len:2, cap:3

		s = append(s, 100)
		fmt.Printf("%v len:%d, cap:%d\n", s, len(s), cap(s)) // [0 10 100] len:3, cap:3

		// cap 보다 더 큰 용량이 필요한경우 더 큰 새로운 배열을 할당합니다. (기존의 값은 복사)
		s = append(s, 1000)
		fmt.Printf("%v len:%d, cap:%d\n", s, len(s), cap(s)) // [0 10 100 1000] len:4, cap:6

		// cap 이 1000 보다 작은 경우에는 두배로 확장, 1000개를 넘어서면 25% 씩 확장
		// 확장 알고리즘은 변경될 수 있다
	}

	{
		// nil slice 와 empty slice 도 append 함수를 사용 할 수 있다.
		ns := []int(nil)
		ns = append(ns, 1)
		fmt.Println(ns) // [1]

		es := make([]int, 0, 0)
		es = append(es, 1)
		fmt.Println(es) // [1]
	}

	{
		// append 할 내부 배열의 자리에 이미 원소가 존재한다면 append 한 값으로 덮어씌워진다.
		arr := [...]int{1, 2, 3, 4, 5}

		s := arr[:2]
		fmt.Printf("%v len:%d, cap:%d\n", s, len(s), cap(s)) // [1 2] len:2, cap:5

		s = append(s, 999)
		fmt.Printf("%v len:%d, cap:%d\n", s, len(s), cap(s)) // [1 2 999] len:3, cap:5

		fmt.Println(arr) // [1 2 999 4 5]
	}

	{
		// 슬라이싱 표현의 세번째 인덱스 옵션으로 cap 을 제한 할 수 있다.
		// slice[i:j:k] -> 최대로 확장 가능한 용량을 k 번째 인덱스 까지로 제한 < cap(slice) = k - i >
		// 원본 슬라이스의 용량보다 큰 값의 용량이 되도록 지정하면 런타임 에러가 발생한다
		arr := [...]int{1, 2, 3, 4, 5}

		s := arr[:2:2]
		fmt.Printf("%v len:%d, cap:%d\n", s, len(s), cap(s)) // [1 2] len:2, cap:2

		s = append(s, 999)
		fmt.Printf("%v len:%d, cap:%d\n", s, len(s), cap(s)) // [1 2 999] len:3, cap:4

		fmt.Println(arr) // [1 2 3 4 5]
	}

	{
		// 슬라이스도 배열처럼 range 키워드로 순회 할 수 있다
		s := []string{"hi", "hello"}
		for i, v := range s {
			fmt.Println(i, v)
		}
	}

}
