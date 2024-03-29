package product

import (
	"context"
	"errors"
	"testing"

	"base-gin-golang/constants"
	mockRepository "base-gin-golang/mock/domain/repository"
	mockPostgreSQL "base-gin-golang/mock/infra/postgresql"
	mockDataPkg "base-gin-golang/mock/pkg/data"

	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.WithValue(context.Background(), constants.ContextKeyProcessID, "processID")
	mockDB, errConnect := mockPostgreSQL.ConnectPostgresql()
	if errConnect != nil {
		t.Errorf("connect db fail")
	}
	mockProductRepository := mockRepository.NewMockProductRepository(ctrl)
	mockDataService := mockDataPkg.NewMockService(ctrl)
	productUseCase := NewProductUseCase(mockProductRepository, mockDataService, mockDB)
	mockDataService.EXPECT().Copy(gomock.Any(), gomock.Any()).Return(errors.New("Copy fail"))
	t.Run("Test copy fail", func(t *testing.T) {
		_, err := productUseCase.Create(ctx, &CreateProductInput{})
		if err != nil && err.Error() != "Copy fail" {
			t.Errorf("Test copy fail")
		}
	})
	mockDataService.EXPECT().Copy(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockProductRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("Fail"))
	t.Run("Test create fail", func(t *testing.T) {
		_, err := productUseCase.Create(ctx, &CreateProductInput{})
		if err != nil && err.Error() != "Fail" {
			t.Errorf("Test create fail")
		}
	})
	mockProductRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, nil)
	_, err := productUseCase.Create(ctx, &CreateProductInput{})
	t.Run("Test create success", func(t *testing.T) {
		if err != nil {
			t.Errorf("Test create success fail")
		}
	})
}
