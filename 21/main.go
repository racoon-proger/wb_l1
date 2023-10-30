package main

import "log"

type newsService interface {
	GetTheMostPopular() (output PieceOfNews)
}

// в нашем сервисе используется интерфейс-адаптер newService с определенным методом
// который возвращает самую популярную новость.
// часто источники новостей меняются, и вместо того чтобы каждый раз
// менять код в месте получения новости (в данный момент это функция main, но это для примера)
// мы создали адаптер newService который инкапсулирует в себе конкретные источники новостей
func main() {
	var svc newsService = NewService(NewFirstChannelNewsService())
	n := svc.GetTheMostPopular()
	log.Println(n)
}
