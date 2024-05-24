package src

import (
	"belimang/src/drivers/db"
	"belimang/src/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	h := http.New(
		&http.Http{
			DB: db.InitDB(),
		},
	)
	defer db.InitDB().Close()

	h.Launch()
}
