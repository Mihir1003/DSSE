package data

import (
	"encoding/json"
	"io"
)

// Product defines the structure for an API product
type Operation struct {
	Name      string `json:"name"`
	Frequency int    `json:"frequency"`
}

// Products is a collection of Product
type Operations []*Operation

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (o *Operations) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}
func (o *Operations) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(o)
}

// GetProducts returns a list of products
func GetOperations() Operations {
	return operationList
}

// productList is a hard coded list of products for this
// example data source
var operationList = []*Operation{
	&Operation{

		Name:      "add",
		Frequency: 0,
	},
	&Operation{

		Name:      "delete",
		Frequency: 0,
	},
	&Operation{

		Name:      "extractMin",
		Frequency: 0,
	},
	&Operation{
		Name:      "extractMax",
		Frequency: 0,
	},
}
