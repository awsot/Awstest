packages user

import (
  _ "github.com/go-sql-driver/mysql"
  "github.com/flosch/pongo2"
  "net/http"
  models/models
)

func login(w http.ResponseWriter, r *http.Request)  {
  db.First(&user)
}
