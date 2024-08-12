package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.news)
	mux.HandleFunc("GET /podcasts", app.podcasts)
	mux.HandleFunc("GET /videos", app.videos)

	mux.HandleFunc("POST /feed", app.addRSSFeed)

	return mux
}
