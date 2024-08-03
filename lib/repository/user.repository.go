package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ShoreLab/shorelab-backend/lib/dto"
	"github.com/ShoreLab/shorelab-backend/lib/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) CreateUser(user *dto.UserCreateRequest) error {
	u := model.User{}
	r.db.MongoDBDatabase.Collection("users").
		FindOne(r.ctx, bson.M{"email": user.Email}).Decode(&u)

	if u.Email != "" {
		return errors.New("user already exists")
	}

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Default().Println(err)
		return err
	}

	u.Name = user.Name
	u.Email = user.Email
	u.Password = hashedPw

	ctx, cancel := context.WithTimeout(r.ctx, 30*time.Second)
	defer cancel()

	res, err := r.db.MongoDBDatabase.Collection("users").InsertOne(ctx, u)
	if err != nil {
		log.Default().Println(err)
		return err
	}
	log.Default().Println("Object created: ", res.InsertedID)
	return nil
}
