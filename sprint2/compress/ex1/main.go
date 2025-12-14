package main

import (
	"compress/flate"
	"compress/zlib"
	"io"
	"log"
	"net/http"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за gzip-сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

type zlibWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w zlibWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за zlib-сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

func defaultHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<html><body>"+strings.Repeat("Hello, world<br>", 20)+"</body></html>")
}

func gzipHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// проверяем, что клиент поддерживает gzip-сжатие
		// это упрощённый пример. В реальном приложении следует проверять все
		// значения r.Header.Values("Accept-Encoding") и разбирать строку
		// на составные части, чтобы избежать неожиданных результатов
		log.Printf("Accept-Encoding: %s", r.Header.Get("Accept-Encoding"))
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "deflate") {
			// если gzip не поддерживается, передаём управление
			// дальше без изменений
			next.ServeHTTP(w, r)
			return
		}

		// создаём gzip.Writer поверх текущего w
		flatew, err := zlib.NewWriterLevel(w, flate.BestCompression)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer flatew.Close()

		w.Header().Set("Content-Encoding", "deflate")
		// передаём обработчику страницы переменную типа gzipWriter для вывода данных
		next.ServeHTTP(zlibWriter{ResponseWriter: w, Writer: flatew}, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandle)
	http.ListenAndServe(":3000", gzipHandle(mux))
}
