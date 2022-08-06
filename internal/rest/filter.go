package rest

import (
	"encoding/json"
	"fmt"
	"github.com/Oppodelldog/webdiff/internal/files"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Filter struct {
	Name string `json:"name"`
	Def  string `json:"def"`
	Type string `json:"type"`
}
type FilterList struct {
	Filters []Filter `json:"filters"`
}

func allFiltersHandler() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		list, err := files.Filters()
		if err != nil {
			http.Error(writer, "error", http.StatusInternalServerError)

			fmt.Println(err)
			return
		}

		err = json.NewEncoder(writer).Encode(outputFilters(list))
		if err != nil {
			fmt.Printf("error sending response to client: %v", err)
		}
	}
}

func upsertFilterHandler() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var newFilter Filter
		err := json.NewDecoder(request.Body).Decode(&newFilter)
		if err != nil {
			http.Error(writer, "cannot read request", http.StatusBadRequest)

			return
		}

		err = files.NewFilter(newFilter.Name, newFilter.Def, newFilter.Type)
		if err != nil {
			http.Error(writer, "cannot add filter", http.StatusInternalServerError)

			return
		}

		writer.WriteHeader(http.StatusNoContent)
	}
}

func outputFilters(list *files.FilterList) FilterList {
	var res = FilterList{Filters: []Filter{}}
	for _, filter := range list.Filters {
		res.Filters = append(res.Filters, Filter(filter))
	}

	return res
}

func deleteFilterHandler() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var name = params.ByName("name")

		if len(name) == 0 {
			http.Error(writer, "filter name must not be empty", http.StatusNotFound)

			return
		}

		err := files.DeleteFilter(name)
		if err != nil {
			http.Error(writer, "not found", http.StatusNotFound)

			return
		}

		writer.WriteHeader(http.StatusNoContent)
	}
}
