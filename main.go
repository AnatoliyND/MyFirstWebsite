package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct { //создание структуры, которая описывает нашу таблицу
	Id                     uint16
	Title, Anons, FullText string
}

var posts = []Article{} //создаем слайс(список) с типом данных Article, в который будем сохранять новые посты

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `articles`") //делаем выборку данных. SELECT * FROM `articles` - команда позволяет вытянуть все данные из +
	//таблички `articles`, либо вместо * можно указать конкретные поля

	if err != nil {
		panic(err)
	}

	posts = []Article{} //перед циклом обращаемся к постам и указываем, что это пустой список(необходимо, чтобы посты не дублировались +
	//при обновлении страницы) каждый раз когда будем попадать на главную страницу список будут назначаться пустым(ранее добавленые данные сохраняются)
	for res.Next() { //перебираем все res. Метод Next возвращает нам либо true - если есть следующая строка, которую можно обработать,+
		// либо false - если нет строк, которые можно обработать
		var post Article                                                   //создаем объект на основе структуры Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText) //убеждаемся существуют ли какие-либо данные в ряде, который рассматриваем и вытягиваем их
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)

	}

	t.ExecuteTemplate(w, "index", posts) //используем ExecuteTemplate(), т.к. внутри шаблонов будем создавать динамическое подключение
}

func save_article(w http.ResponseWriter, r *http.Request) { //метод для обработки данных и переадресации пользователя на какую-либо страницу
	//создаем переменные для получения данных из заполняемой на сайте формы
	title := r.FormValue("title") // в метод r.FormValue передаем название того поля из которого хотим получить значение
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" { //делаем проверку, что указанные поля заполнены
		fmt.Fprintf(w, "Не все данные введены!")
	} else {

		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang") //подключение к базе данных sql. "mysql" - указывает+
		// к какой субд подключаемся. root:root - логин пароль. @tcp(127.0.0.1:3306) - сетевой адрес БД. golang - незвание БД к которой подключаемся
		if err != nil {
			panic(err)
		}
		defer db.Close()

		//установка данных
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles`(`title`, `anons`, `full_text`) VALUES('%s', '%s', '%s')", title, anons, full_text)) // команда sql+
		//для добавления новой записи в таблицу `articles` в поля `title`, `anons`, `full_text`. Добавляем значения VALUES, перечисляем добавляемые значения
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther) // метод http.Redirect переадресовывает нас на страницу. http.StatusSeeOther позволяет делать +
		// переадресацию с верным кодом ответа
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil) //используем ExecuteTemplate(), т.к. внутри шаблонов будем создавать динамическое подключение
}

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", save_article)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
