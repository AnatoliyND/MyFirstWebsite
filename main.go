package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) { // Передаем два параметра. 1 за счет первого параметра (http.ResponseWriter) мы +
	//можем обращаться к указаной транице и показывать что-либо с помощью этой страницы пользователю. Второй параметр (http.Request) это запрос,+
	// тот параметр, который всегда передается
	fmt.Fprintf(w, "Go is very nice!") //создаем форматированную строку, та строка, в которую можно вставлять подстраиваемые значения, переменные

}

func main() {
	http.HandleFunc("/", home_page) //функция принимает два параметра. Функция позволяет отследить переход по определенному URL адресу +
	//и при переходе по этому адресу вызвать какой-либо метод, который будет показывать что-либо пользователю. +
	//Первым параметром передается тот URL адрес, который будет отслеживаться("/" - означает, что будем отслеживать главную страницу). +
	//Вторым параметром передаем метод который будет вызываться при переходе на главную страницу

	http.ListenAndServe(":8080", nil) //Метод принимает два параметра. Первый параметр это порт, по которому будем слушать локальный сервер.+
	// Во втором параметре передаются параметры настройки самого сервера

}
