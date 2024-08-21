package handlers

import (
	"net/http"
	"strconv"

	"example.com/goserver/views"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	hxReq, _ := strconv.ParseBool(r.Header.Get("HX-Request"))
	views.Home(hxReq).Render(r.Context(), w)
}
