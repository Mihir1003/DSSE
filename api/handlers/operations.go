package handlers

import (
	"api/data"
	"github.com/mihir1003/DSSE/dataStructures/BinaryTrees"
	"github.com/mihir1003/DSSE/dataStructure/Lists"

	"log"
	"time"
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
	}else {
		rw.WriteHeader(http.StatusOK)

		//_,e:=rw.Write(strs)
		//if e !=nil {
		//	http.Error(rw,"Invalid request",http.StatusBadRequest)
		//}

	}



}

func parse(ops *data.Operations) []string {
	var list []string
	for _, element := range *ops {
		switch element.Name {
		case "add":
			list = append(list, "add")
		case "delete":
			list = append(list, "delete")
		case "extractMin":
			list = append(list, "extractMin")
		case "extractMax":
			list=append(list,"extractMax")
		}
	}
	return list
}

type empty interface {

}

type dataStructure struct {
	name string
	time time.Duration
	structure empty

}

type dataStructures []dataStructure

func compete(ops []string) dataStructures {
	listDataStructures:=dataStructures{dataStructure{
		name: "list",
		time: 0,
		structure: Lists.LinkedList{
			Val:  0,
			Next: nil,
		},
	},
	dataStructure{
		name: "sortedList",
		time: 0,
		structure: Lists.SortedLinkedList{
		Val:  0,
		Next: nil,
	},
	},
	dataStructure{
		name: "binaryTree",
		time: 0,
		structure: BinaryTrees.ItemBinarySearchTree{root:nill}
	},
	}
	for _,ds := range listDataStructures {
		for _,o := ops{
			switch o {
			case "add":
				ds.add(0)
			case "delete":
				list = append(list, "delete")
			case "extractMin":
				list = append(list, "extractMin")
			case "extractMax":
				list=append(list,"extractMax")
			}
		}
	}


}

