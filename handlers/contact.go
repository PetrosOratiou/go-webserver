package handlers

import (
	"net/http"
	"strconv"

	"example.com/goserver/views"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	hxReq, _ := strconv.ParseBool(r.Header.Get("HX-Request"))
	views.Contact(hxReq).Render(r.Context(), w)
}
