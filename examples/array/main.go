package main

import "fmt"

func passArray(arr [4]int) {
	arr[0] = 999
}

func main() {
	{
		/*
			C와는 달리 배열은 value 이다.
		*/
		arr := [4]int{1, 2, 3, 4}
		passArray(arr)           // call by value
		fmt.Printf("%#v\n", arr) // [4]int{1, 2, 3, 4}
	}

	{
		/*
			배열은 원소의 zero-value 로 초기화 되어 있다.
		*/
		fmt.Printf("%#v\n", [5]int{})    // [5]int{0, 0, 0, 0, 0}
		fmt.Printf("%#v\n", [5]bool{})   // [5]bool{false, false, false, false, false}
		fmt.Printf("%#v\n", [5]string{}) // [5]string{"", "", "", "", ""}
	}

	{
		/*
			컴파일러에게 배열의 길이를 계산하도록 [...] 를 사용할 수 있다.
			! 컴파일러가 숫자를 대신 넣어 줄 뿐, 여전히 배열은 고정크기이다.
		*/
		arr := [...]int{1, 2, 3}
		fmt.Printf("%#v\n", arr) // [3]int{1, 2, 3}
	}

	{
		/*
			배열은 == 으로 비교 될 수 있다.
		*/
		arr := [4]int{1, 2, 3, 4}
		fmt.Println(arr == [4]int{1, 2, 3, 4}) // true
		fmt.Println(arr == [4]int{1, 2, 3, 0}) // false
	}

	{
		/*
			배열의 타입은 길이와 원소의 타입에 의해 결정된다.
			- 다른 타입의 배열은 비교 될 수 없다.
			- 다른 타입의 배열은 파라미터로 넘길 수 없다.
		*/
		//arr := [3]int{}
		//arr == [5]int{}     // invalid operation: (mismatched types [3]int and [5]int)
		//arr == [3]float32{} // invalid operation: (mismatched types [4]int and [3]float32)
		//passArray(arr)      // cannot use arr (type [3]int) as type [4]int in argument to passArray
	}

	{
		/*
			배열의 원소는 메모리에 순차적으로 저장되어 있다.
		*/
		byteArr := [4]byte{} // byte: 1 byte
		fmt.Printf("&byteArr[0]: %d\n", &byteArr[0])
		fmt.Printf("&byteArr[1]: %d\n", &byteArr[1])
		fmt.Printf("&byteArr[2]: %d\n", &byteArr[2])
		fmt.Printf("&byteArr[3]: %d\n", &byteArr[3])
		intArr := [4]int32{} // int32: 4 byte
		fmt.Printf("&intArr[0]: %d\n", &intArr[0])
		fmt.Printf("&intArr[1]: %d\n", &intArr[1])
		fmt.Printf("&intArr[2]: %d\n", &intArr[2])
		fmt.Printf("&intArr[3]: %d\n", &intArr[3])
	}

}
