package main

import (
	"dataLinkedList/linkedList"
	"fmt"
)

func main() {
	list := linkedList.New(3)

	list.Add(8)

	list.Pop()

	list.DeleteFrom(1)

	list.UpdateAt(1, 55)

	list.AddAt(0, 34)
	list.AddAt(2, 33)

	list.Output()

	fmt.Println("Node[1] = ", list.At(1))
	fmt.Println("Size is", list.Size())
}
