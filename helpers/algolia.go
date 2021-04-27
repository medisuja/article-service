package helpers

import (
	"encoding/json"

	"article-service/config"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

// algolia client
func AlgoliaClient(algoliaIndex string) algoliasearch.Index {
	client := algoliasearch.NewClient(config.Config.Algolia.APPID, config.Config.Algolia.APPKEY)
	index := client.InitIndex(algoliaIndex)

	return index
}

func AlgoliaClientCreate(algoliaIndex string) algoliasearch.Index {
	client := algoliasearch.NewClient(config.Config.Algolia.APPID, config.Config.Algolia.APPKEYCREATE)
	index := client.InitIndex(algoliaIndex)

	return index
}

func AddBatchObject(algoliaIndex string, data interface{}) (bool, error) {
	var (
		objects []algoliasearch.Object
	)
	index := AlgoliaClientCreate(algoliaIndex)

	buff, _ := json.Marshal(data)
	if err := json.Unmarshal(buff, &objects); err != nil {
		return false, err
	}

	_, errAdd := index.AddObjects(objects)

	if errAdd != nil {
		return false, errAdd
	}

	return true, nil
}

func AddUpdateObject(algoliaIndex string, data interface{}) bool {
	var (
		object algoliasearch.Object
	)
	index := AlgoliaClientCreate(algoliaIndex)

	buff, _ := json.Marshal(data)
	if err := json.Unmarshal(buff, &object); err != nil {
		return false
	}

	_, err := index.AddObject(object)
	if err != nil {
		return false
	}

	return true
}

func DeleteObject(algoliaIndex string, objectID string) bool {
	index := AlgoliaClient(algoliaIndex)

	_, err := index.DeleteObject(objectID)
	if err != nil {
		return false
	}

	return true
}

func SearchAlgolia(algoliaIndex string, keyword string, limit int) (algoliasearch.QueryRes, error) {
	index := AlgoliaClient(algoliaIndex)
	var res algoliasearch.QueryRes

	params := algoliasearch.Map{
		"hitsPerPage":   limit,
		"typoTolerance": false,
	}

	res, err := index.Search(keyword, params)

	if err != nil {
		return res, err
	}

	return res, nil
}
