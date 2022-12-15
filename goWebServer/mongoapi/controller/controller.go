package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Deepanshu276/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://############3:#####@cluster0.bwahgx4.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"    //Name of our Database
const colName = "watchList" //Name of our Collection

// Reference of the Database so that we can access it any time
var collection *mongo.Collection

//Connect with mongoDB

// init method execute only once at the starting
func init() {
	//client options, that is setting up the connection

	clientOption := options.Client().ApplyURI(connectionString)

	//Connect to mongodb

	//Context is used to tell for how long i can make a request and what happens when the request goes off,
	//Hence i used todo() context type as i don't want my request to keeps my connection alive
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	//To reach inside a Database and inside a collection
	collection = client.Database(dbName).Collection(colName)

	//collection reference

	fmt.Println("Collection reference is ready")

}

//Mongodb Helper file

// insert 1 record in database
func insertOneMovie(movie model.Netflix) {
	//used InsertOne to insert one record in a database, , collection act as a
	inserted, err := collection.InsertOne(context.TODO(), movie)

	if err != nil {
		log.Fatal(err)
	}
	//InsertedID will print the id of the record which we stored in our Database
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

//Update 1 record

//First Provided the filter based on which data is filtered out like which one should be updated
// pass the flag and updated information

func updateOneMovie(movieId string) {
	//Converting the Movie id into mongoDB understandable id
	//objectIDFromHex-------> Convert string into objectId which is acceptable by the MongoDb
	id, _ := primitive.ObjectIDFromHex(movieId)

	//bson.M is used while performing filter when the shorter and Clearer result is required
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	collection.UpdateOne(context.TODO(), filter, update)

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified Count:", result.ModifiedCount)

}

//delete One record

func deleteOneMovie(movieId string) {
	//Converting the Movie id into mongoDB understandable id
	//objectIDFromHex-------> Convert string into objectId which is acceptable by the MongoDb
	id, _ := primitive.ObjectIDFromHex(movieId)

	//bson.M is used while performing filter when the shorter and Clearer result is required
	filter := bson.M{"_id": id}
	// deleteCount------> count of record deleted
	deleteCount, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count:", deleteCount)

}

// Delete all records from mongodb

func deleteAllMovie() int64 {
	//bsonD-----> when the result needed be in sorted Order
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	//deleteResult will return the key value pair
	fmt.Println("Number of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount

}

//Get all movies from Mongodb

func getAllMovies() []primitive.M {
	//bson.D used for selecting the whole data without any filter in the ordered manner
	//cur is object which contain lots of data and looping thogh it will ogive us the required data
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	//array type of primitive.M is used which is used to store the data fetched from the database
	var movies []primitive.M
	//traversing the data returned by cur
	for cur.Next(context.TODO()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)

	}
	//closing the connection
	defer cur.Close(context.TODO())
	return movies

}

// Getting all the movies from Database
// this function accepts a ResponseWriter to send a response back and a Request object that provides more information about the request itself.
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	//Need to send the data in json format , hence used json encoder
	json.NewEncoder(w).Encode(allMovies)

}

// inserting the movie to the database
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	//Set sets the header entries associated with key to the single element value.
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Methos", "POST")

	var movie model.Netflix
	//decoding the json data
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

//Mark movie as watched

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Methos", "PUT")
	//mux is used to get the unique value of the movie from the url so that we make update
	//Params will have the key value pair
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

// Deleting a single movie
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Methos", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

// Deleting all movie
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Methos", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)

}
