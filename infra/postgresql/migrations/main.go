package main

import (
	"fmt"
	"io"
	"os"

	"base-gin-golang/infra/postgresql/model"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(
		&model.User{},
		&model.Product{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts) //nolint:errcheck // library
}
