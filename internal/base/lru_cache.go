package base

/*
Метод добавляет значение в начало списка:
если размер списка превышает size, удаляет последний элемент
*/
func (l *LruCache) Put(key string, value string) {
	var newNode Node
	newNode.Prev = nil
	newNode.Next = nil
	newNode.Key = key
	newNode.Value = value
	if l.size == 0 {
		return
	}
	if l.Head == nil {
		l.Head = &newNode
		l.Tail = &newNode
		return
	}
	currentNode := *l.Head
	currentNode.Prev = &newNode
	newNode.Next = &currentNode
	l.Head = &newNode
	i := 1
	var tmpNode Node
	tmpNode = *l.Head
	for tmpNode.Next != nil {
		tmpNode = *tmpNode.Next
		i++
	}
	if i == l.size {
		l.Tail = tmpNode.Prev
		// по-хорошему, надо удалить объект из памяти
		tmpNode.Prev = nil
		tmpNode.Next = nil
	}
}

/*
Метод возвращает значение по ключу:
если значения нет — возвращает nil, если значение есть — перемещает узел в начало списка
*/
func (l *LruCache) Get(key string) *string {
	if l.size == 0 || l.Head == nil {
		return nil
	}
	currentNode := *l.Head
	var res *string
	for i := 0; i < l.size; i++ {
		if currentNode.Key == key {
			res = &currentNode.Value
			// если не первый или не единсственный элемент, следует переместить в начало
			if currentNode != *l.Head {
				// связка между окружающими элементами
				var nextNode, prevNode *Node
				nextNode = currentNode.Next
				if nextNode != nil {
					nextNode.Prev = currentNode.Prev
					prevNode = currentNode.Prev
					if prevNode != nil && prevNode.Next != nil {
						prevNode.Next = nextNode
					}
					if prevNode != nil && prevNode.Prev == nil {
						prevNode.Prev = &currentNode
					}
				}
				if nextNode == nil {
					prevNode = currentNode.Prev
					if prevNode.Next != nil {
						prevNode.Next = nil
					}
					l.Tail = prevNode
				}
				// в начало
				tmpNode := l.Head
				tmpNode.Prev = &currentNode
				currentNode.Prev = nil
				currentNode.Next = tmpNode
				l.Head = &currentNode
			}
			return res
		}
		if currentNode.Next != nil {
			currentNode = *currentNode.Next
		}
	}
	return res
}

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	size int
	Head *Node
	Tail *Node
}

func NewLruCache(size int) *LruCache {
	return &LruCache{
		size: size,
	}
}
