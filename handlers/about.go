package handlers

import (
	"net/http"
	"strconv"

	"example.com/goserver/views"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	hxReq, _ := strconv.ParseBool(r.Header.Get("HX-Request"))
	views.About(hxReq).Render(r.Context(), w)
}
