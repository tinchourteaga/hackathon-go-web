package main

import (
	"github.com/gin-gonic/gin"

	"github.com/tinchourteaga-ml/desafio-go-web-martin-urteaga/cmd/server/handler"
	"github.com/tinchourteaga-ml/desafio-go-web-martin-urteaga/internal/tickets"
	"github.com/tinchourteaga-ml/desafio-go-web-martin-urteaga/pkg/store"
)

func main() {

	// Cargo csv.
	ticketList, err := store.LoadTicketsFromFile("./tickets.csv")

	if err != nil {
		panic("Couldn't load tickets")
	}

	repo := tickets.NewRepository(ticketList)
	service := tickets.NewService(repo)
	ticket := handler.NewService(service)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	tickets := router.Group("/tickets")
	{
		tickets.GET("/getByCountry/:dest", ticket.GetTicketsByCountry())
		tickets.GET("/getAverage/:dest", ticket.AverageDestination())
	}

	if err := router.Run(); err != nil {
		panic(err)
	}

}
