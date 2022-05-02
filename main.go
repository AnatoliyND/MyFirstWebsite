package main

import (
	"fmt"
	"html/template" //пакет для вывода полноценных html страниц
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User Name is: %s. He is %d Age and "+
		"he has Money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	ben := User{"Ben", 24, -105, 4.5, 0.9, []string{"Football", "Skate", "Dance"}}
	/* 	fmt.Fprintf(w, `<h1>Main Text</h1>
	   	<b>Main Taxt</b1>`) //формат для вывода HTML страниц(приметивный)
	*/
	tmpl, _ := template.ParseFiles("templates/home_page.html") //создаем переменную, которая будет хранить шаблон для html страницы
	tmpl.Execute(w, ben)                                       //отображение шаблона на странице
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
