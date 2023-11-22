package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"test_api/database/mongodb"
)

func ReadAllRecord(writer http.ResponseWriter, request *http.Request) {
	cursor, err := mongodb.MongoClient.Database("11").Collection("12").
		Find(context.Background(), map[string]interface{}{})
	a := []map[string]interface{}{}
	if err == nil {
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var data map[string]interface{}
			_ = cursor.Decode(&data)
			a = append(a, data)
		}
	}
	b, _ := json.Marshal(a)
	writer.Write(b)
	return
}
