package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const PORT = 8080

const humanPlayer = "X"
const aiPlayer = "O"
var gameActive bool
var gameTable [9]string
var winConditions = [8][3]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
	{0, 4, 8}, {2, 4, 6},
}

func initHandler(w http.ResponseWriter, request *http.Request) {
	gameTable = [9]string{"", "", "", "", "", "", "", "", ""}
	gameActive = true

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w,"init game")
}

func actionHumanHandler(w http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id := query.Get(humanPlayer)
	idx, _ := strconv.ParseInt(id, 10, 10)
	if gameTable[idx] == "" && gameActive {
		gameTable[idx] = humanPlayer
	} else {
		idx = -1
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, fmt.Sprint(idx))
}

func actionAIHandler(w http.ResponseWriter, request *http.Request) {
	idx := -1
	for gameActive {
		idx = rand.Intn(9)
		if gameTable[idx] == "" {
			gameTable[idx] = aiPlayer
			break
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, fmt.Sprint(idx))
}

func statusGameHandler(w http.ResponseWriter, request *http.Request) {
	state := "active"
	for i := 0; i < len(winConditions); i++ {
		a := gameTable[winConditions[i][0]]
		b := gameTable[winConditions[i][1]]
		c := gameTable[winConditions[i][2]]
		if a == "" || b == "" || c == "" {
			continue
		}
		if a == b && b == c {
			gameActive = false
			state = a
			break
		}
	}

	if gameActive {
		isDraw := true
		for _, a := range gameTable {
			if a == "" {
				isDraw = false
				break
			}
		}
		if isDraw {
			gameActive = false
			state = "draw"
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, state)
}

func main() {
	http.HandleFunc("/init", initHandler)
	http.HandleFunc("/action/human", actionHumanHandler)
	http.HandleFunc("/action/ai", actionAIHandler)
	http.HandleFunc("/status", statusGameHandler)

	fmt.Println(fmt.Sprintf("Server running at port %d...", PORT))

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}
