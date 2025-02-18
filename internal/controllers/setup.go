package controllers

import "net/http"

func (c *CartController) Start(addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: c.ginEngine,
	}

	srv.ListenAndServe()
}
