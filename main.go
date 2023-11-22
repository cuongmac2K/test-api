package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"test_api/controller"
	"test_api/database/mongodb"
	"test_api/env"
)

func main() {
	// set truong env vao de check moi
	os.Setenv("env", "dev")

	// check env de xem moi truong nao
	e := os.Getenv("env")
	var dataYaml = env.DataYaml{}
	if e == "dev" {
		yamlFile, _ := ioutil.ReadFile("./env/dev_env.yaml")
		_ = yaml.Unmarshal(yamlFile, &dataYaml)
		// ket noi mongo
		mongodb.ConnectToMongoDB(dataYaml.ConfigMysql.Username, dataYaml.ConfigMysql.Password)
	}
	http.HandleFunc("/add-many-record", controller.AddManyRecord)
	http.HandleFunc("/read-all-record", controller.ReadAllRecord)
	http.ListenAndServe(":3000", nil)
}
