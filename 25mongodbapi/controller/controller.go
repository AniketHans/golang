package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongodbapi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://aniket-hans:aniket1234@golangtest.dnsps.mongodb.net/?retryWrites=true&w=majority&appName=GolangTest"
const dbName = "Netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

//connect with mongoDb
func init(){
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client,err:= mongo.Connect(context.TODO(),clientOptions)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection 	= (*mongo.Collection)(client.Database(dbName).Collection(collectionName))

	fmt.Println("Collection instance is raedy")
}


func insertOneMovie(movie model.Netflix){
	res,err := collection.InsertOne(context.TODO(),movie)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie is DB with ID:-",res.InsertedID)
}

func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update := bson.M{"$set": bson.M{"watched":true}}
	res,err := collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("modified count :=",res.ModifiedCount)

}

func deleteOneRecord(movieID string){
	id, _ := primitive.ObjectIDFromHex(movieID)

	filter := bson.M{"_id":id}
	
	res,err := collection.DeleteOne(context.Background(),filter,nil)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Deleted movies :-",res.DeletedCount)

}


func deleteAllRecords() int64{
	filter := bson.D{{}}
	res ,err := collection.DeleteMany(context.Background(),filter,nil)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Number of  movies deleted :-",res.DeletedCount)
	return res.DeletedCount
}

func getAllMovies() []primitive.M{
	cursor, err := collection.Find(context.Background(),bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}
 
	var movies []primitive.M
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()){
		var movie bson.M
		err := cursor.Decode(&movie)
		if err!=nil{
			log.Fatal(err)
		}

		movies = append(movies, movie)

	}

	return movies
}


func GetAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	allMovies := getAllMovies()
	err := json.NewEncoder(w).Encode(allMovies)
	if err!=nil{
		panic(err)
	} 
}

func CreateAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err!=nil{
		log.Fatal(err)
	}
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["movieId"])
	json.NewEncoder(w).Encode("Data updated")
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneRecord(params["movieId"])
	json.NewEncoder(w).Encode("Deleted movie")
}


func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllRecords()
	json.NewEncoder(w).Encode(fmt.Sprintf("Deleted %v movies",count))
}
