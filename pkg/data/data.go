package data

import (
	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

type Service interface {
	Copy(to interface{}, from interface{}) error
}

type dataService struct{}

var ProviderSet = wire.NewSet(NewDataService)

func NewDataService() Service {
	return &dataService{}
}

func (s *dataService) Copy(to interface{}, from interface{}) error {
	err := copier.CopyWithOption(to, from, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	return err
}
