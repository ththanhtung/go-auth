package mongodbframework

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lab5.cmo/internal/application/core"
)

type Adapter struct {
	db *mongo.Client
}

func NewAdapter() *Adapter {
	client, err := mongo.NewClient(options.Client(), options.Client().ApplyURI("mongodb://localhost:27017/lab5"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return &Adapter{
		db: client,
	}
}

func (db Adapter) openCollection(collectionName string) *mongo.Collection {
	collection := db.db.Database("goauth").Collection(collectionName)
	return collection
}

func (db Adapter) CreateUser(username, password, firstname, lastname, email, dob, avatar, address string) (core.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userCollection := db.openCollection("user")

	count, _ := userCollection.CountDocuments(ctx, bson.M{
		"username": username,
	})
	defer cancel()

	if count > 0 {
		return core.User{},errors.New("user already existed")
	}

	newUser := core.User{
		ID: primitive.NewObjectID(),
		Username:  username,
		Password:  password,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		DoB:       dob,
		Avatar:    avatar,
		Address:   address,
	}

	newUser.UserId = newUser.ID.Hex()

	result, err := userCollection.InsertOne(ctx, newUser)
	defer cancel()

	if err!= nil {
		return core.User{}, err 
	}

	insertedID := result.InsertedID.(primitive.ObjectID)

	err = userCollection.FindOne(ctx, bson.M{"_id": insertedID}).Decode(&newUser)

	if err!= nil {
		return core.User{}, err 
	}

	return newUser, nil
}
func (db Adapter) UpdateUser(username, password string, updatecontent core.User) (core.User, error) {
	userCollection := db.openCollection("user")
	_, err := db.Login(username, password)
	
	if err != nil {
		return core.User{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{
		"username": username,
	}

	updateProperty := bson.M{
		"$set": bson.M{
			"username":  username,
			"password":  password,
			"firstname": updatecontent.Firstname,
			"lastname":  updatecontent.Lastname,
			"email":     updatecontent.Email,
			"dob":       updatecontent.DoB,
			"avatar":    updatecontent.Avatar,
			"address":   updatecontent.Address,
		},
	}

	var updatedUser core.User
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = userCollection.FindOneAndUpdate(ctx, filter, updateProperty, options).Decode(&updatedUser)
	defer cancel()

	if err != nil {
		return core.User{}, err
	}

	return updatedUser, nil
}

func (db Adapter) Login(username, password string) (core.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userCollection := db.openCollection("user")
	countUser, err := userCollection.CountDocuments(ctx, bson.M{
		"username": username,
	})
	defer cancel()
	if err != nil {
		return core.User{}, err
	}
	if countUser == 0 {
		return core.User{}, errors.New("incorrect username or password, name")
	}
	countPassword, err := userCollection.CountDocuments(ctx, bson.M{
		"password": password,
	})
	defer cancel()
	if err != nil {
		log.Println(err.Error())
		return core.User{}, err
	}
	if countPassword == 0 {
		return core.User{}, errors.New("incorrect username or password, pass")
	}
	var user core.User
	userCollection.FindOne(ctx, bson.M{
		"username": username,
		"password": password,
	}).Decode(&user)

	return user, nil
}
func (db Adapter) GetUsers() ([]core.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userCollection := db.openCollection("user")
	cursor, err := userCollection.Find(ctx, bson.M{})
	defer cancel()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var users []core.User
	for cursor.Next(ctx) {
		var user core.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	defer cancel()
	return users, nil
}
func (db Adapter) DeleteUser(username, password string) (core.User, error) {
	userCollection := db.openCollection("user")
	_, err := db.Login(username, password)
	
	if err != nil {
		return core.User{}, err
	}
	
	filter := bson.M{
		"username": username,
		"password": password,
	}
	
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var deletedUser core.User
	err = userCollection.FindOneAndDelete(ctx, filter).Decode(&deletedUser)
	defer cancel()

	if err != nil {
		return core.User{}, err
	}

	return deletedUser, nil
}
