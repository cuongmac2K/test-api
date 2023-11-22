  sudo docker run -d --restart=always -p 27017:27017 --name my_mongo -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=123456 mongo
