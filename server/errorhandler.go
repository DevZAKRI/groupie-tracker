package server

import (
	"fmt"
	"html/template"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, messgage string, code int) {
	tmp, err := template.ParseFiles("templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `
		<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Error 500 - Something Went Wrong</title>
    <link rel="stylesheet" href="../assets/errStyles.css">
</head>
<body>
  <div class="container">
    <h1>Oops!</h1>
    <div class="error-code"> ERROR 500</div>
    <p>Internal Server Error, Try again Later</p>
    <div class="buttons">
      <a href="/">Go Home</a>
    </div>
    <footer>
      &copy; 2024 Groupie Tracker | Need help? <a href="mailto:mostafazakri@gmail.com" style="color: #ffffff; text-decoration: underline;">Contact Support</a>
    </footer>
  </div>
</body>
</html>

		`)
		return
	}
	w.WriteHeader(code)

	errorData := Error{Code: code, Message: messgage}
	tmp.Execute(w, errorData)

}
