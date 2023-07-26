package product

import (
	"errors"
	"testing"

	mockRepository "base-gin-golang/mock/domain/repository"
	mockDataPkg "base-gin-golang/mock/pkg/data"

	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProductRepository := mockRepository.NewMockProductRepository(ctrl)
	mockDataService := mockDataPkg.NewMockDataService(ctrl)
	productUseCase := NewProductUseCase(mockProductRepository, mockDataService)
	mockDataService.EXPECT().Copy(gomock.Any(), gomock.Any()).Return(errors.New("Copy fail"))
	t.Run("Test copy fail", func(t *testing.T) {
		_, err := productUseCase.Create(&CreateProductInput{})
		if err != nil && err.Error() != "Copy fail" {
			t.Errorf("Test copy fail")
		}
	})
	mockDataService.EXPECT().Copy(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockProductRepository.EXPECT().Create(gomock.Any()).Return(nil, errors.New("Fail"))
	t.Run("Test create fail", func(t *testing.T) {
		_, err := productUseCase.Create(&CreateProductInput{})
		if err != nil && err.Error() != "Fail" {
			t.Errorf("Test create fail")
		}
	})
	mockProductRepository.EXPECT().Create(gomock.Any()).Return(nil, nil)
	_, err := productUseCase.Create(&CreateProductInput{})
	t.Run("Test create success", func(t *testing.T) {
		if err != nil {
			t.Errorf("Test create success fail")
		}
	})
}
