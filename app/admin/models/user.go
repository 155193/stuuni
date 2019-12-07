package models

import (
	"encoding/hex"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/sha3"
)

type User struct {
	Id       primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Names    string               `bson:"names"`
	Ln       string               `bson:"ln"`  //last name
	Mln      string               `bson:"mln"` //mother last name
	Dni      string               `bson:"dni"`
	Photo    string               `bson:"photo"`
	Nick     string               `bson:"nick"`
	Email    string               `bson:"email"`
	Password string               `bson:"password"`
	Roles    []primitive.ObjectID `bson:"roles"`
}

//Constructor for User model
func (u *User) New(id primitive.ObjectID, Dni string,
	Nick string, Email string, Password string, Roles []primitive.ObjectID) User {
	u.Id = id
	u.Dni = Dni
	u.Nick = Nick
	u.Email = Email
	u.Password = Password
	u.Roles = Roles
	return *u
}

func (u *User) ToBson() bson.M {
	return bson.M{
		"_id,omitempty": u.Id,
		"names":         u.Names,
		"ln":            u.Ln,
		"mln":           u.Mln,
		"dni":           u.Dni,
		"photo":         u.Photo,
		"nick":          u.Nick,
		"email":         u.Email,
		"password":      u.Password,
	}
}

func Login(nickI interface{}, passI interface{}) (string, string, error) {
	nick, ok := nickI.(string)
	if !ok {
		return "", "", errors.New("error casting nick")
	}
	password, ok := passI.(string)
	if !ok {
		return "", "", errors.New("error casting password")
	}
	h := sha3.NewLegacyKeccak256()
	//h := sha3.NewLegacyKeccak512()
	h.Write([]byte(password))
	p := hex.EncodeToString(h.Sum(nil))
	return nick, p, nil
}

func UserToBson(userI interface{}) (*User, *bson.M, error) {
	userM, ok := userI.(map[string]interface{})
	if !ok {
		return nil, nil, errors.New("Error casting user")
	}
	id := primitive.NewObjectID()
	dni := userM["dni"].(string)
	_, p, err := Login("", dni) //password
	if err != nil {
		return nil, nil, err
	}
	names := userM["names"].(string)
	ln := userM["ln"].(string)
	mln := userM["mln"].(string)
	nick := userM["nick"].(string)
	email := userM["email"].(string)
	userBson := bson.M{
		"_id":      id,
		"names":    names,
		"ln":       ln,
		"mln":      mln,
		"dni":      dni,
		"photo":    "",
		"nick":     nick,
		"email":    email,
		"password": p,
		"roles":    []primitive.ObjectID{},
	}
	return &User{Id: id, Names: names, Ln: ln, Mln: mln, Nick: nick, Email: email, Roles: []primitive.ObjectID{}}, &userBson, nil
}

func UserUpdateToBson(idI interface{}, userI interface{}) (*User, *bson.M, error) {
	idU, ok := idI.(string)
	if !ok {
		return nil, nil, errors.New("Error casting idUser")
	}
	userM, ok := userI.(map[string]interface{})
	if !ok {
		return nil, nil, errors.New("Error casting User")
	}
	user := bson.M{
		"names": userM["names"].(string),
		"ln":    userM["ln"].(string),
		"mln":   userM["mln"].(string),
		"dni":   userM["dni"].(string),
		"nick":  userM["nick"].(string),
		"email": userM["email"].(string),
	}
	aux, _ := primitive.ObjectIDFromHex(idU)
	return &User{Id: aux}, &user, nil
}

func UserGenericToBson(userI interface{}) (*primitive.ObjectID, *bson.M, error) {
	userM, ok := userI.(map[string]interface{})
	if !ok {
		return nil, nil, errors.New("Error casting user")
	}
	//fmt.Println(userM)
	id := primitive.NewObjectID()
	userGenericBson := bson.M{}
	dni := userM["dni"].(string)
	_, p, err := Login("", dni) //password
	if err != nil {
		return nil, nil, err
	}
	names := userM["names"].(string)
	ln := userM["ln"].(string)
	mln := userM["mln"].(string)
	userGenericBson["names"] = names
	userGenericBson["ln"] = ln
	userGenericBson["mln"] = mln
	userGenericBson["nick"] = dni
	userGenericBson["dni"] = dni
	userGenericBson["password"] = p

	userGenericBson["_id"] = id
	//fmt.Println(userGenericBson)
	return &id, &userGenericBson, nil
}
