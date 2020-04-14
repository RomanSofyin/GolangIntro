package main

import "testing"

/*
   Чтобы добавить тесты в go, просто добавьте файл main_test.go.
   Добавим функцию проверяющую метод sayHi.
   Сигнатура методов тестирования должна быть TestXxx(t *testing.T).
   Первый X всегда должен быть в верхнем регистре.

   Запустить все тесты из всех файлов и подпапок в проекте:
   go test ./...
*/

func TestSayHi(t *testing.T) {
    expected := "Hi Marco!"
    greeting := sayHi("Marco")
    if greeting != expected {
        t.Errorf("Greeting was incorrect, got: '%s', want: '%s'", greeting, expected)
    }
}