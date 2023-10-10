package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const GetPageEndpoint string = ""

var questionOrder map[string]string = map[string]string{
	"Q0": "Q1",
	"Q1": "Q2",
	"Q2": "Q3",
	"Q3": "Q4",
	"Q4": "Q5",
}

var pageContent map[string]Page = map[string]Page{
	"Q1": {
		Question: "question 1 text",
		Buttons: []Button{
			buildButton("question 1 - button 1 text", 1),
			buildButton("question 1 - button 2 text", 2),
		},
	},
	"Q2": {
		Question: "question 2 text",
		Buttons: []Button{
			buildButton("question 2 - button 1 text", 1),
			buildButton("question 2 - button 2 text", 2),
		},
	},
	"Q3": {
		Question: "question 3 text",
		Buttons: []Button{
			buildButton("question 3 - button 1 text", 1),
			buildButton("question 3 - button 2 text", 2),
		},
	},
	"Q4": {
		Question: "question 4 text",
		Buttons: []Button{
			buildButton("question 4 - button 1 text", 1),
			buildButton("question 4 - button 2 text", 2),
		},
	},
	"Q5": {
		Question: "question 5 text",
		Buttons: []Button{
			buildButton("question 5 - button 1 text", 1),
			buildButton("question 5 - button 2 text", 2),
		},
	},
}

var pointsSetup map[string]int = map[string]int{
	"Q1:OP1": 1,
	"Q1:OP2": 1,
	"Q2:OP1": 1,
	"Q2:OP2": 1,
	"Q3:OP1": 1,
	"Q3:OP2": 1,
	"Q4:OP1": 1,
	"Q4:OP2": 1,
	"Q5:OP1": 1,
	"Q5:OP2": 1,
}

func buildButton(text string, order int) Button {
	return Button{
		ButtonText:  &text,
		ButtonOrder: &order,
	}
}
func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Hello World!")

	response := events.APIGatewayProxyResponse{}

	// Get Next Page call
	if request.HTTPMethod == http.MethodGet {
		// get params
		nextPage := getNextPage(request.PathParameters["playerId"], request.PathParameters["question"])
		responseBody, err := json.Marshal(nextPage)
		if err != nil {
			response.StatusCode = 500
			response.Body = err.Error()
		} else {
			response.StatusCode = 200
			response.Body = string(responseBody)
		}

		return response, nil
	}

	// Save Gamer Response
	if request.HTTPMethod == http.MethodPost {
		response.StatusCode = 200
		return response, nil
	}

	response.StatusCode = 404
	response.Body = "hello World - You are at the wrong place"

	return response, nil
}

func getNextPage(playerId string, question string) Page {
	if question == "" {
		question = "Q0"
	}
	return pageContent[getNextQuestion(question)]
}

type GetPageRequest struct {
	PlayerId *string `json:"playerId"`
	Question *string `json:"question"`
}

func getNextQuestion(question string) string {
	return questionOrder[question]
}

func buildPageRequest(playerId string, question string) string {
	return GetPageEndpoint + "/" + playerId + "/" + question
}

type Page struct {
	Question string   `json:"question"`
	Buttons  []Button `json:"buttons"`
	Points   string   `json:"points"`
}

type Button struct {
	ButtonText  *string `json:"buttonText"`
	ButtonOrder *int    `json:"buttonOrder"`
}

type Player struct {
	PlayerId *string `json:"playerId"`
}
