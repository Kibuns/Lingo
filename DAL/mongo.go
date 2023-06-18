package DAL

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Kibuns/Lingo/Models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// global variable mongodb connection client
var client *mongo.Client
var err error
var sessionCollection *mongo.Collection

func init() {
	client, err = newClient()
	if err != nil {
		// Handle the error here, such as logging or exiting the program
		fmt.Println("Failed to connect to MongoDB:", err)
	} else{
		sessionCollection = client.Database("SessionDB").Collection("sessions")
	}

}


func StoreSession() (string, error) {
	if err != nil {
		return "", fmt.Errorf("could not connect to database")
	}
	//simplify detaileduser to user
	var session Models.Session
	session.ID = uuid.New().String()
	session.Guesses = 0
	session.IsComplete = false
	session.SecretWord = "test"
	session.Created = time.Now()

	
	_, err := sessionCollection.InsertOne(context.TODO(), session)
	if err != nil {
		return "", err
	}

	fmt.Println("started new session with id: " + session.ID)
	return session.ID, nil
}

func GetSession(id string) (Models.Session, error) {
	if err != nil {
		return Models.Session{}, fmt.Errorf("could not connect to database")
	}
	fmt.Println("getting session")

	// Create a filter to search for the document with the specified username
	filter := bson.M{"id": id}

	// Find the first document that matches the filter
	var result Models.Session
	err := sessionCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Models.Session{}, errors.New("session not found")
		}
		return Models.Session{}, err
	}

	// Display the retrieved document
	fmt.Println("Displaying the result from the search query")
	fmt.Println(result)

	return result, nil
}

func newClient() (client *mongo.Client, err error) {
	// err = godotenv.Load()
	// if err != nil {
	// 	return nil, err
	// }

	connectionString := os.Getenv("CONNECTION_STRING")
	fmt.Println("connectionstring: " + connectionString)

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}
