package main

import (
	"fmt"
	"net/http"
	helper "user-login-register/helpers"
)

func main() {

	username, email, password, comfirm := "", "", "", ""

	mux := http.NewServeMux()
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		username = r.FormValue("username")
		email = r.FormValue("email")
		password = r.FormValue("password")
		comfirm = r.FormValue("comfirm")

		usernameCheck := helper.IsEmpty(username)
		emailCheck := helper.IsEmpty(email)
		passwordCheck := helper.IsEmpty(password)
		confirmCheck := helper.IsEmpty(comfirm)

		if usernameCheck || emailCheck || passwordCheck || confirmCheck {
			fmt.Fprintf(w, "Error code is -10 : There is empty data.")
			return
		}

		if password == comfirm {
			// Veritabanina kaydet
			fmt.Fprintf(w, "Registration succesful.")
		} else {
			fmt.Fprintf(w, "Password information must be the same")
		}

		//for key, value := range r.Form {
		//	fmt.Print("%s = %s\n", key, value)
		//}

	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email = r.FormValue("username")
		password = r.FormValue("password")

		emailCheck := helper.IsEmpty(email)
		passwordCheck := helper.IsEmpty(password)

		if emailCheck || passwordCheck {
			fmt.Fprintf(w, "There is empty data")
		}
		dbPwd := "123"
		dbUsername := "zafer"

		if dbPwd == password && dbUsername == email {
			fmt.Fprintf(w, "Login Succesful")
		} else {
			fmt.Fprintf(w, "login failed")
		}

	})

	http.ListenAndServe(":8080", mux)
}
