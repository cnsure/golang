package main

import "fmt"

func main() {
	arr := [5]string{"I","am","stupid","and","weak"}
	fmt.Println(arr)
	for index := range arr {
		if index== 2 {
			arr[index] = "smart"
		}else if index == 4 {
			arr[index] = "strong"
		}
	}
	fmt.Println(arr)
}
