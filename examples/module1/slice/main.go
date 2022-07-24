package main

import (
	"fmt"
)

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:3]
	mySlice = append(mySlice, 8)
	mySlice = append(mySlice, 9)
	mySlice = append(mySlice, 10)
	mySlice = append(mySlice, 11)
	fmt.Printf("mySlice %+v\n", mySlice)
	fmt.Printf("myArray %+v\n", myArray)
	fullSlice := myArray[:]
	remove3rdItem := deleteItem(fullSlice, 2)
	fmt.Printf("remove3rdItem %+v\n", remove3rdItem)

	a := make([]int,1,1)
	b := make([]int,3,3)
	b[0] =1
	b[0] =2
	b[0] =3
	c := a
	a = append(b,4)
	fmt.Printf("a %v \nb %v \nc %v \n",a,b,c)

	d := new([]int)
	e := make([]int,0)
	var f []int
	fmt.Println(d)
	fmt.Printf("d %v %d %d\n",d,len(*d),cap(*d))
	fmt.Printf("e %v %d %d\n",e,len(e),cap(e))
	fmt.Printf("f %v %d %d\n",f,len(f),cap(f))

	x := []int{}
	fmt.Println(x,len(x),cap(x))



}

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
