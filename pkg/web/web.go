package web

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"github.com/Jyolando/link_shortener_bot/pkg/database"
)

var db *sql.DB

func getRoot(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) > 5 {
		if r.URL.Path != "/" && r.URL.Path[len(r.URL.Path)-4:len(r.URL.Path)] != ".ico" {
			path := r.URL.Path[1:len(r.URL.Path)]
			link := database.GetLink(db, path)
			http.Redirect(w, r, link, http.StatusSeeOther)
		}
	} else if r.URL.Path == "/" {
		http.Redirect(w, r, "https://telegram.me/short_my_link_bot", http.StatusSeeOther)
	} else {
		io.WriteString(w, "Not valid token")
	}
}

func Init() {
	http.HandleFunc("/", getRoot)
	db = database.Init()
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}
