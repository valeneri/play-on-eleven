package server

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	*mongo.Database
}

func CreateDb(config Configuration) (*Database, error) {

	// cred.AuthSource = YourAuthSource
	cred := options.Credential{
		Username: config.Db.Username,
		Password: config.Db.Password,
	}

	// set client options
	poolSize := config.Db.PoolSize
	timeout := time.Duration(config.Db.Timeout) * time.Second
	dbName := config.Db.DatabaseName
	dbUrl := config.Db.DatabaseUrl
	fmt.Println("db url: ", dbUrl)
	// connect with params
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl).SetSocketTimeout(timeout).
		SetMaxPoolSize(poolSize).SetAuth(cred))

	// check error
	if err != nil {
		return nil, err
	}

	// ping to ensure connection is ok
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	fmt.Println("Database connected")
	// return &Database{
	// 	database,
	// }, nil
	return &Database{db}, nil
}
