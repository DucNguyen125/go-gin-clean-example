package elasticsearch

import (
	"base-gin-golang/config"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	log "github.com/sirupsen/logrus"
)

type Database struct {
	*elasticsearch.Client
}

func ConnectElasticsearch(cfg *config.Environment) (*Database, error) {
	esCfg := elasticsearch.Config{
		Addresses: []string{
			cfg.ElasticSearchURI,
		},
	}
	es, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		log.Info("error when create client elastic search")
		return nil, err
	}
	res, err := es.Info()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return &Database{es}, nil
}
