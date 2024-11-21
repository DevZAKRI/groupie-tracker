package server

import (
	"fmt"
	"net/http"
)

func AssetsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/assets/" {
		fs := http.FileServer(http.Dir("./assets"))
		http.StripPrefix("/assets/", fs).ServeHTTP(w, r)
		fmt.Fprint(w, "Access Forbiden")
		//(w, http.StatusForbidden, "Access Forbidden")
	} else {
		fmt.Fprint(w, "Page Not Found")
		//ErrorPage(w, http.StatusNotFound, "Page Not Found")
	}
}
