package main

import (
	"fmt"
	stack "lab1"
)

const (
	cap   = "Введите вместимость (cap) стека:"
	title = "Выберите один из предложенных вариантов"
	n     = "1 - Проверка пустоты стека"
	tw    = "2 - Проверка заполненности стекового массива"
	th    = "3 - Добавление элемента в вершину стека"
	fr    = "4 - Удаление элемента из вершины стека"
	fv    = "5 - Вывод текущего состояния стека на экран"
	rw    = "Введите элемент:"
)

func main() {

	var cp int

	fmt.Printf("%s\n", cap)

	if _, err := fmt.Scan(&cp); err != nil {
		fmt.Println(err.Error())
	}

	csStack := stack.NewStack(cp)

mt:
	for {
		var ch int
		var elm int

		fmt.Printf("%s:\n%s\n%s\n%s\n%s\n%s\n", title, n, tw, th, fr, fv)

		if _, err := fmt.Scan(&ch); err != nil {
			fmt.Println(err.Error())
			goto mt
		}

		switch ch {
		case 1:
			fmt.Println(csStack.IsEmpty())
		case 2:
			fmt.Println(csStack.IsFull())
		case 3:

			fmt.Printf("%s\n", rw)

			if _, err := fmt.Scan(&elm); err != nil {
				fmt.Println(err.Error())
				goto mt
			}

			if err := csStack.Push(elm); err != nil {
				fmt.Println(err.Error())
				goto mt
			}
		case 4:
			fmt.Println(csStack.Pop())
		case 5:
			fmt.Println(csStack.Statistic())
		}
	}
}
