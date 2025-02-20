package controllers

import "net/http"

const sessionIDKey = "session_id"

func (c *CartController) Start(addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: c.ginEngine,
	}

	srv.ListenAndServe()
}
