data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./infra/postgresql/migrations",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15/dev"
  migration {
    dir    = "file://infra/postgresql/migrations"
    format = golang-migrate
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
