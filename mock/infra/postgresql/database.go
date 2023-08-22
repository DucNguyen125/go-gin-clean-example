package postgresql

import (
	"base-gin-golang/infra/postgresql"
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var mockDb *gorm.DB

func ConnectPostgresql() (*postgresql.Database, error) {
	if mockDb == nil {
		dbContainer, err := testcontainers.GenericContainer(
			context.Background(),
			testcontainers.GenericContainerRequest{
				ContainerRequest: testcontainers.ContainerRequest{
					Image:        "postgres:15",
					ExposedPorts: []string{"5432/tcp"},
					WaitingFor:   wait.ForListeningPort("5432/tcp"),
					Env: map[string]string{
						"POSTGRES_DB":       "db",
						"POSTGRES_PASSWORD": "postgres",
						"POSTGRES_USER":     "postgres",
					},
				},
				Started: true,
			})
		if err != nil {
			return nil, err
		}
		host, _ := dbContainer.Host(context.Background())
		port, _ := dbContainer.MappedPort(context.Background(), "5432")
		dbURI := fmt.Sprintf("postgres://postgres:postgres@%v:%v/db", host, port.Port())
		mockDb, _ = gorm.Open(postgres.Open(dbURI), &gorm.Config{
			// Logger: &postgresql.Logger{
			// 	SkipErrRecordNotFound: true,
			// 	Debug:                 true,
			// },
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	}
	return &postgresql.Database{DB: mockDb}, nil
}
