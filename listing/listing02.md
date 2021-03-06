Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ: Оператор defer добавляет в стек вызов функции после ключевого слова defer. Отложенная функция вызывается непосредственно перед возвратом из нее в обратном порядке их вызова (LIFO). Если в функции имеется именованный параметр результата, то defer до возврата из функции может получить доступ и изменить параметр результата. 
В функции **test** есть именнованный параметр результата **x**, которому присваивается **1**. Оператор return устанавливает параметр результата перед выполнением отложенной функции. Отложенная функция инкерментирует его.
В функции **anotherTest** создается переменная **x**. Инициализируется отложенная функция, где захватывается **x**. Оператор return, устанавливает параметр результата перед выполнением отложенной функции, отложенная функция инкрементирует переменную, которая уже установлена в результат возврата.

```
out: 
2
1
```
