package linkedList

import (
	"fmt"
)

type node struct { //элемент списка
	value int
	next  *node
}

type LinkedList struct { //связный список
	Head *node
}

func New(len int) *LinkedList { // инициализация
	if len != 0 {
		list := LinkedList{
			Head: &node{
				value: 0,
				next:  nil,
			},
		}
		var cur = list.Head
		for i := 1; i < len; i++ {
			cur.next = &node{
				value: 0,
				next:  nil,
			}
			cur = cur.next
		}
		return &list
	} else {
		list := LinkedList{
			Head: nil,
		}
		return &list
	}
}

func (list LinkedList) Output() { // вывод списка в консоль
	fmt.Println("List: ")
	var cur = list.Head
	var i int = 0
	for cur != nil {
		fmt.Println("Node [", i, "] = ", cur.value)
		cur = cur.next
		i++
	}

}

func (list *LinkedList) Add(val int) { // добавление нового элемента в конец
	temp := &node{
		value: val,
		next:  nil,
	}
	cur := list.Head
	if cur != nil {
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = temp
	} else {
		list.Head = temp
	}
}

func (list *LinkedList) Pop() { // удаление значения из конца
	cur := list.Head
	if cur != nil && cur.next != nil {
		for cur.next.next != nil {
			cur = cur.next
		}
		cur.next = nil
	} else if cur.next == nil { // только 1 элемент в списке
		list.Head = nil
	}
}

func (list LinkedList) Size() int { // возвращает размерность списка
	var counter int = 0
	temp := list.Head
	for temp != nil {
		counter++
		temp = temp.next
	}
	return counter
}

func (list *LinkedList) DeleteFrom(pos int) { // удаление элемента с позиции pos

	cur := list.Head
	switch {
	case pos == 0:
		if cur != nil {
			list.Head = cur.next
		}
	case pos == list.Size():
		list.Pop()
	case pos < 0 || pos > list.Size():
		fmt.Println("Wrong position")
		return
	default:
		for i := 0; i < pos-1; i++ { //доходим до элемента, стоящего перед удаляемым
			cur = cur.next
		}
		temp := cur.next
		cur.next = temp.next
	}

}

func (list *LinkedList) AddAt(pos int, val int) { // добавление нового элемента на позицию pos
	temp := &node{
		value: val,
		next:  nil,
	}
	cur := list.Head

	switch {
	case pos == 0:
		temp.next = list.Head
		list.Head = temp
	case pos == list.Size():
		list.Add(val)
	case pos < 0 || pos > list.Size():
		fmt.Println("Wrong position")
		return
	default:
		for i := 0; i < pos-1; i++ {
			cur = cur.next
		}
		temp.next = cur.next
		cur.next = temp
	}
}

func (list *LinkedList) UpdateAt(pos int, val int) { // оюновление значения на позиции pos
	cur := list.Head
	for i := 0; i < pos; i++ {
		cur = cur.next
	}
	cur.value = val
}

func (list *LinkedList) At(pos int) int { // получение значения на позиции pos
	cur := list.Head
	for i := 0; i < pos; i++ {
		cur = cur.next
	}
	return cur.value
}
