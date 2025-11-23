package main

import (
	"fmt"
	"log"
	"net/http"
)

// middlware — это функция-посредник (middleware), которая добавляет к HTTP-ответу
// заголовки перед передачей управления следующему обработчику в цепочке.
//
// Добавляемые заголовки:
//   - "name": устанавливается в значение "ivan"
//   - "content-type": принудительно устанавливается в "application/json"
//
// Параметры:
//   - next: следующий HTTP-обработчик в цепочке, который будет вызван после
//     установки заголовков.
//
// Возвращаемое значение:
//   - Объект http.Handler, который можно использовать как промежуточный слой
//     в цепочке обработки HTTP-запросов.
//
// Пример использования:
//
//	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	    fmt.Fprintf(w, `{"message": "hello"}`)
//	})
//	wrapped := middleware(handler)
//	http.Handle("/api", wrapped)
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Добавляем заголовок "name" со значением "ivan".
		// Если заголовок уже существует, значение будет добавлено (Add).
		w.Header().Add("name", "ivan")

		// Устанавливаем тип содержимого в JSON.
		// Любой предыдущий "content-type" будет заменён (Set).
		w.Header().Set("content-type", "application/json")

		// Передаём управление следующему обработчику в цепочке.
		next.ServeHTTP(w, r)
	})
}


func mainPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	body += "Query parameters ===============\r\n"
	if err := req.ParseForm(); err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	for k, v := range req.Form {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	res.Write([]byte(body))
}

func main() {
	http.Handle("/", middleware(http.HandlerFunc(mainPage)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
