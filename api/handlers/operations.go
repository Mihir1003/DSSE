package handlers

import (
	"api/data"
	"api/dataStructure/BinaryTrees"
	"encoding/json"
	"time"

	"io"

	"api/dataStructure/Lists"

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
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")

	if r.Method == http.MethodGet {
		o.getOperations(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		o.receiveOperations(rw, r)
		return
	}

	if r.Method == "OPTIONS" {
		rw.WriteHeader(http.StatusOK)
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
	} else {
		rw.WriteHeader(http.StatusOK)

		//_,e:=rw.Write(strs)
		//if e !=nil {
		//	http.Error(rw,"Invalid request",http.StatusBadRequest)
		//}
		p.l.Println(op)
		parsed := parse(op)
		structures := compete(parsed, *p)

		b := structures.ToJSON(rw)
		p.l.Println(string(b))
		//if e!=nil{
		//	http.Error(rw,"operations invalid",http.StatusInternalServerError)
		//}

	}

}

func parse(ops *data.Operations) []string {
	var list []string
	for _, element := range *ops {
		switch element.Name {
		case "add":

			for i := 0; i < element.Frequency; i++ {
				list = append(list, "add")

			}
		case "delete":
			for i := 0; i < element.Frequency; i++ {
				list = append(list, "delete")
			}
		case "extractMin":
			for i := 0; i < element.Frequency; i++ {
				list = append(list, "extractMin")
			}
		case "extractMax":
			for i := 0; i < element.Frequency; i++ {
				list = append(list, "extractMax")
			}
		}
	}
	return list
}

type das interface {
	Add(int)
	Delete(int)
	ExtractMin() int
	ExtractMax() int
}

type dataStructure struct {
	Name      string `json:"name"`
	Time      string `json:"time"`
	structure das
}

func (o *dataStructures) ToJSON(w io.Writer) []byte {
	//e := json.NewEncoder(w)
	//return e.Encode(o)
	b, e := json.Marshal(o)
	if e == nil {
		_, e := w.Write(b)
		if e != nil {

		}
	}
	return b
}

type dataStructures []*dataStructure

func compete(ops []string, o Operations) dataStructures {
	listDataStructures := dataStructures{&dataStructure{
		Name: "list",
		Time: "initial",
		structure: &Lists.LinkedList{
			Val:  0,
			Next: nil,
		},
	},
		&dataStructure{
			Name: "sortedList",
			Time: "initial",
			structure: &Lists.SortedLinkedList{
				Val:  0,
				Next: nil,
			},
		},
		&dataStructure{
			Name: "binaryTree",
			Time: "initial",
			structure: &BinaryTrees.ItemBinarySearchTree{
				Root: &BinaryTrees.Node{Left: nil, Right: nil},
			},
		},
		&dataStructure{
			Name: "AVLTree",
			Time: "initial",
			structure: &BinaryTrees.AvlNode{
				Elem:   0,
				Left:   nil,
				Right:  nil,
				Height: 0,
			},
		},
	}
	for _, ds := range listDataStructures {
		start := time.Now()
		for i, o := range ops {
			switch o {
			case "add":
				ds.structure.Add(i)
			case "delete":
				ds.structure.Delete(i)
			case "extractMin":
				ds.structure.ExtractMax()
			case "extractMax":
				ds.structure.ExtractMax()
			}
		}

		ds.Time = time.Since(start).String()
		o.l.Println(ds)

	}

	return listDataStructures
}
