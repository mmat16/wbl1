package main

/*
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

применение данной функции приводит к утечке памяти, так как строки в Go являются
неизменяемыми, и создание подстроки из первых 100 символов ориганальной строки
не позволяет сборщику мусора освободить память, выделенную под оригинальную
строку.
*/

var justString string

func createHugeString(size uint64) string {
	// some code
	return ""
}

// someFunc создаёт "огромную" строку и копирует в justString первые 100 байт
// этой строки при помощи функции copy
func someFunc() {
	v := createHugeString(1 << 10)
	justString = ""
	copy([]byte(justString), v[:100])
}
