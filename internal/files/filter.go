package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Filter struct {
	Name   string
	Filter string
}
type FilterList struct {
	Filters []Filter
}

func (fl FilterList) ByName(name string) *Filter {
	if idx := fl.IndexByName(name); idx > -1 {
		return &fl.Filters[idx]
	}

	return nil
}

func (fl FilterList) IndexByName(name string) int {
	for i := range fl.Filters {
		if fl.Filters[i].Name == name {
			return i
		}
	}

	return -1
}

func Filters() (*FilterList, error) {
	list, err := loadFilterList()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func NewFilter(name, filter string) error {
	list, err := loadFilterList()
	if err != nil {
		return err
	}

	f := Filter{
		Name:   name,
		Filter: filter,
	}

	if idx := list.IndexByName(name); idx > -1 {
		list.Filters[idx] = f
	} else {
		list.Filters = append(list.Filters, f)
	}

	return saveFilterList(list)
}

func DeleteFilter(name string) error {
	list, err := loadFilterList()
	if err != nil {
		return err
	}

	if idx := list.IndexByName(name); idx > -1 {
		list.Filters = append(list.Filters[:idx], list.Filters[idx+1:]...)

		return saveFilterList(list)
	}

	return fmt.Errorf("not found")

}

func loadFilterList() (*FilterList, error) {
	f, err := openFiltersFile()
	if err != nil {
		return nil, fmt.Errorf("cannot read filters file '%s': %v", FiltersFile(), err)
	}

	var filters FilterList
	err = json.NewDecoder(f).Decode(&filters)
	if err != nil {
		return nil, fmt.Errorf("cannot decode  filters file '%s': %v", FiltersFile(), err)
	}

	return &filters, nil
}

func saveFilterList(filters *FilterList) error {
	f, err := os.OpenFile(FiltersFile(), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0655)
	if err != nil {
		return fmt.Errorf("cannot write filters file '%s': %v", FiltersFile(), err)
	}

	err = json.NewEncoder(f).Encode(filters)
	if err != nil {
		return fmt.Errorf("cannot encode filters file '%s': %v", FiltersFile(), err)
	}

	return nil
}

func openFiltersFile() (*os.File, error) {
	file := FiltersFile()
	err := ensureFiltersFile(file)
	if err != nil {
		return nil, fmt.Errorf("cannot initially create filters file '%s': %v", file, err)
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("cannot read filters file '%s': %v", file, err)
	}

	return f, nil
}

func ensureFiltersFile(file string) error {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0655)
		if err != nil {
			return err
		}

		err = json.NewEncoder(f).Encode(FilterList{})
		if err != nil {
			return err
		}
	}

	return nil
}
