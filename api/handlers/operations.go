package handlers

import (
	"api/data"
	"github.com/mihir1003/dsse/compete"
	"log"
)

import "net/http"

type Operations struct {
	l *log.Logger
}

func NewOperations(l *log.Logger) *Operations {
	return &Operations{l}
}

func (o *Operations) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		o.getOperations(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		o.receiveOperations(rw, r)
		return
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Operations) getOperations(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Operations")

	// fetch the products from the datastore
	lp := data.GetOperations()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Operations) receiveOperations(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	op := &data.Operations{}

	err := op.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	compete.compete

}
