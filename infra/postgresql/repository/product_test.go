package repository

import (
	"context"
	"errors"
	"testing"

	"base-gin-golang/domain/entity"
	"base-gin-golang/infra/postgresql/model"
	mockPostgreSQL "base-gin-golang/mock/infra/postgresql"
	mockDataPkg "base-gin-golang/mock/pkg/data"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
)

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDataService := mockDataPkg.NewMockService(ctrl)
	mockDB, errConnect := mockPostgreSQL.ConnectPostgresql()
	if errConnect != nil {
		t.Errorf("connect db fail")
	}
	mockDB.Migrator().DropTable(&model.Product{})
	mockDB.DB.AutoMigrate(&model.Product{})
	productRepository := NewProductRepository(mockDB, mockDataService)
	mockDataService.EXPECT().Copy(gomock.Any(), gomock.Any()).Return(errors.New("Copy failed"))
	t.Run("Test copy fail", func(t *testing.T) {
		_, err := productRepository.Create(context.Background(), &entity.Product{})
		if err != nil && err.Error() != "Copy failed" {
			t.Errorf("Test copy fail")
		}
	})
	input := &entity.Product{
		ProductCode: "123",
		ProductName: "test",
		Price:       123,
	}
	product := &model.Product{}
	copier.Copy(product, input)
	mockDataService.EXPECT().Copy(gomock.Any(), gomock.Any()).Return(nil).SetArg(0, *product)
	mockDataService.EXPECT().Copy(gomock.Any(), gomock.Any()).Return(errors.New("Copy failed"))
	t.Run("Test copy fail", func(t *testing.T) {
		_, err := productRepository.Create(context.Background(), &entity.Product{})
		if err != nil && err.Error() != "Copy failed" {
			t.Errorf("Test copy fail")
		}
	})
	mockDataService.EXPECT().Copy(gomock.Any(), gomock.Any()).Return(nil).SetArg(0, *product)
	product.ID = 1
	mockDataService.EXPECT().Copy(gomock.Any(), gomock.Any()).Return(nil).SetArg(1, *product)
	t.Run("Test all", func(t *testing.T) {
		_, err := productRepository.Create(context.Background(), &entity.Product{})
		if err != nil {
			t.Errorf("Test all")
		}
	})
}
