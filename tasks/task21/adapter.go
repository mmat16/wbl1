package main

import "fmt"

/*  типы и методы, которые требуют применения адаптера  */
/* ===================================================  */

// тип бесплатного видео-стриминга
type FreeStreaming struct{}

// метод бесплатного видео-стриминга, перечисляющий все доступные к просмотру
// фильмы
func (f *FreeStreaming) ListMovies() {
	fmt.Println("\nfree movies availible:\nулица разбитых фонарей\nмухтар\nглухарь")
}

// тип платного стриминга
type SubscribeStreaming struct{}

// метод платного стриминга, перечисляющий все доступные к просмотру фильмы, но
// требующий чтобы пользователь был авторизован
func (s *SubscribeStreaming) ListMovies(loggedIn bool) {
	if loggedIn {
		fmt.Println("\nmovies availible with subscription:\nLOTR\nMatrix\nGame of Thrones")
	}
}

// ===============================================

/*             адаптеры для стримингов            */
// ===============================================

// общий тип-интерфейс адаптера с методом AllMovies, перечисляющий все доступные
// к просмотру фильмы
type StreamingAdapter interface {
	AllMovies()
}

// Адаптер бесплатного стриминга, встраивает в себя экземпляр бесплатного
// стриминга
type FreeStreamingAdapter struct {
	*FreeStreaming
}

// метод адаптера бесплатного стриминга, реализующий интерфейс общего
// адаптера через вызов метода бесплатного стриминга, перечисляющий все
// доступные к просмотру фильмы
func (a *FreeStreamingAdapter) AllMovies() {
	a.ListMovies()
}

// конструктор адаптера бесплатного стриминга встраивает переданный экземпляр
// бесплатного стриминга в адаптер бесплатного стриминга и возвращает созданный
// экземпляр адаптера бесплатного стриминга как экземпляр общего интерфейса
// адаптера
func NewFreeStreamingAdapter(streaming *FreeStreaming) StreamingAdapter {
	return &FreeStreamingAdapter{streaming}
}

// адаптер платного стриминга, со встроенным экземпляром платного стриминга
type SubscribeStreamingAdapter struct {
	*SubscribeStreaming
}

// метод, реализующий интерфейс общего адаптера, вызывает метод встроенного
// платного стриминга, автоматически передавая флаг авторизованного пользователя
// для того чтобы сигнатура метода интерфейса не нарушалась
func (a *SubscribeStreamingAdapter) AllMovies() {
	a.ListMovies(true)
}

// конструктор адаптера платного стриминга, встраивающий переданный экземпляр
// платного стриминга в его адаптер и возвращающий новый экземпляр структуры
// адаптера платного стриминга, как интерфейс общего адаптера
func NewSubscribeStreamingAdapter(streaming *SubscribeStreaming) StreamingAdapter {
	return &SubscribeStreamingAdapter{streaming}
}

// ==========================================================================

// тип, реализующий интерфейс адаптера "по умолчанию"
// ==========================================================================

// тип агрегатора рейтингов фильмов
type IMDB struct{}

// метод агрегатора, перечисляющий фильмы и их оценки, а так же реализующий
// интерфейс общего адаптера стримингов
func (i *IMDB) AllMovies() {
	fmt.Println("\nmovies rating:\nLOTR-8.8\nMatrix-8.1\nGame of Thrones-8.5")
}

// ==========================================================================

/*
в main функции создаётся массив объектов, удовлетворяющих интерфейс общего
адаптера - FreeStreamingAdapter, SubscribeStreamingAdapter и IMDB. далее в цикле
через каждый из объектов вызывается метод интерфейса AllMovies, который
реализован в каждом из элементов массива.
*/
func main() {
	services := [3]StreamingAdapter{
		NewFreeStreamingAdapter(&FreeStreaming{}),
		NewSubscribeStreamingAdapter(&SubscribeStreaming{}),
		&IMDB{},
	}

	for _, service := range services {
		service.AllMovies()
	}
}
