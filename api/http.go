package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/smaxwellstewart/sentiment/analyze"
	"github.com/smaxwellstewart/sentiment/api/response"
	"github.com/smaxwellstewart/sentiment/word"
)

const (
	DefaultOrder = "desc"
	DefaultLimit = 10
)

type Handler struct {
	Anlyzr analyze.Analyzer
}

type Payload struct {
	Content string `json:"content"`
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}
	// Decode posted json to Body struct
	decoder := json.NewDecoder(r.Body)
	var p Payload
	err := decoder.Decode(&p)
	if err != nil {
		response.InternalError(w)
	}
	defer r.Body.Close()

	// analyse
	ws := ContentToWords(p.Content)
	// TODO! perform error check
	scores, _ := h.Anlyzr.Analyze(ws)

	// TODO! turn into function and test
	order := strings.ToLower(r.URL.Query().Get("order"))
	if order == "" {
		order = DefaultOrder
	}

	switch order {
	case "asc":
		scores = word.OrderAsc(scores)
	case "desc":
		scores = word.OrderDesc(scores)
	default:
		response.BadRequest(w, "Order param must be <asc|desc>.")
	}

	// TODO! turn into function and test
	limit := r.URL.Query().Get("limit")
	var n int
	if limit == "" {
		n = DefaultLimit
	} else {
		n, err = strconv.Atoi(limit)
		if err != nil || n < 1 {
			response.BadRequest(w, "Limit param must be an integer greater than zero.")
			return
		}
	}
	if n > len(scores) {
		n = len(scores)
	}

	scores = scores[:n]

	response.JSON(w, scores)
}
