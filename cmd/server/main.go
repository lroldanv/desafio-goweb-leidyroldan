package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"desafio-go-web-leidyroldan/cmd/server/handler"
	"desafio-go-web-leidyroldan/internal/domain"
	"desafio-go-web-leidyroldan/internal/tickets"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load csv file.
	list, err := LoadTicketsFromFile("../../tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	repository := tickets.NewRepository(list)
	service := tickets.NewService(repository)
	ts := handler.NewService(service)

	r := gin.Default()
	api := r.Group("api/v1")

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	ticket := api.Group("/ticket")
	{
		ticket.GET("/getByCountry/:dest", ts.GetTicketsByCountry())
		ticket.GET("/getAverage/:dest", ts.GetAverageDestination())
	}

	if err := r.Run(); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
