package controller

import (
	"context"
	"net/http"
	"strconv"
	"test_api/database/mongodb"
)

func AddManyRecord(writer http.ResponseWriter, request *http.Request) {
	// insert many to mongo
	for i := 0; i < 100; i++ {
		mongodb.MongoClient.Database("11").Collection("12").
			InsertOne(context.Background(), map[string]interface{}{
				"name": "name lan " + strconv.Itoa(i),
			})
	}
	writer.Write([]byte("insert success"))
	return
}
