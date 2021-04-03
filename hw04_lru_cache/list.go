package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	Length int
	First  *ListItem
	Last   *ListItem
}

func NewList() List {
	return new(list)
}

func (l list) Len() int {
	return l.Length
}

func (l list) Front() *ListItem {
	return l.First
}

func (l list) Back() *ListItem {
	return l.Last
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{v, l.First, nil}

	if l.Len() > 0 {
		l.First.Prev, l.First = item, item
	} else {
		l.First, l.Last = item, item
	}

	l.Length++
	return l.First
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{v, nil, l.Last}

	if l.Len() > 0 {
		l.Last.Next, l.Last = item, item
	} else {
		l.First, l.Last = item, item
	}

	l.Length++
	return l.Last
}

func (l *list) Remove(i *ListItem) {
	// Если элемент один
	if l.Length == 1 {
		emptyItem := &ListItem{nil, nil, nil}
		l.First, l.Last = emptyItem, emptyItem
	} else {
		switch i {
		case l.First:
			// Если элемент первый делаем первым следующий
			i.Next.Prev, l.First = nil, i.Next
		case l.Last:
			// Если элемент последний делаем последним предпоследний
			i.Prev.Next, l.Last = nil, i.Prev
		default:
			// Ссылаем предыдущий и следующий элементы друг на друга
			i.Prev.Next, i.Next.Prev = i.Next, i.Prev
		}
	}

	l.Length--
}

func (l *list) MoveToFront(i *ListItem) {
	// Можно было использовать PushFront и Remove, но
	// 1. Решение ниже делает то же самое, но за меньшее количество итераций
	// 2. Если придется добавлять логирование действий (или какой-то дополнительный функционал в PushFront и Remove)
	// то функция перемещения станет создавать лишний шум

	// Если элемент и так первый ничего не делаем
	if i == l.First {
		return
	}

	// Если элемент последний меняем Last
	if i == l.Last {
		l.Last = i.Prev
	}

	// Ссылаем предыдущий и следующий (если есть) элементы друг на друга
	i.Prev.Next = i.Next
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	// Делаем переносимый элемент первым
	i.Next, i.Prev = l.First, nil
	l.First.Prev, l.First = i, i
}
