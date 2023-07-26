package data

import "github.com/jinzhu/copier"

type Service interface {
	Copy(to interface{}, from interface{}) error
}

type dataService struct{}

func NewDataService() Service {
	return &dataService{}
}

func (s *dataService) Copy(to interface{}, from interface{}) error {
	err := copier.Copy(to, from)
	return err
}
