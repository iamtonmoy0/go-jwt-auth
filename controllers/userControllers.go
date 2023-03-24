package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/iamtonmoy0/go-jwt-auth/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPass()

func VerifyPass()

func Signup()

func Login()

func GetUsers()

func GetUser()
