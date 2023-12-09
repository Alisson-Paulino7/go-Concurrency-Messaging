package main

import (
	// "encoding/json"
	"net/http"

	"github.com/Alisson-Paulino7/go-concurrence-messaging/internal/entity"
	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/middleware"
	// Framework Web services
	"github.com/labstack/echo/v4"
)

func main() {

	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/order", OrderHandler)

	// // http.HandleFunc("/order", OrderHandler)

	// http.ListenAndServe(":8080", r)

	e := echo.New()
	e.GET("/order", OrderHandler)
	// Incorpora o Logger e o servidor numa única linha
	e.Logger.Fatal(e.Start(":8080"))

}

func OrderHandler(c echo.Context) error {

	order, _ := entity.NewOrder("1", 10.0, 1.0)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err)
	// }
	err := order.CalculateFinalPrice()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// Dar o encode na order e jogar na resposta (w)
	return c.JSON(http.StatusOK, order)

}

// func OrderHandler(w http.ResponseWriter, r *http.Request) {


// 	if r.Method != http.MethodGet {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}

// 	order, err := entity.NewOrder("1", 10.0, 1.0)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	err = order.CalculateFinalPrice()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	// Dar o encode na order e jogar na resposta (w)
// 	json.NewEncoder(w).Encode(order)

// }