package services

import (
	"../../db"
	"../models"
	"context"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	Coluser = "users"
)

//Function to add an user model
func AddUser(user bson.M) error {
	c, errC := db.GetCollection(Coluser)

	if errC != nil {
		return errC
	}
	_, errC = c.InsertOne(db.Ctx, user)
	return errC
}

//Function to update an user by id
func UpdateUser(id primitive.ObjectID, user bson.M) error {
	c, errC := db.GetCollection(Coluser)

	if errC != nil {
		return errC
	}
	_, errC = c.UpdateOne(db.Ctx,
		bson.M{"_id": id},
		bson.M{"$set": user},
	)
	return errC
}

//Function to remove an user by id
func RemoveUser(id primitive.ObjectID) (bool, error) {
	c, errC := db.GetCollection(Coluser)

	if errC != nil {
		return false, errC
	}
	_, err := c.DeleteOne(db.Ctx,
		bson.M{"_id": id},
	)
	return err == nil, err
}

//Function to add a role to an existed user
func AddRole2User(idUser primitive.ObjectID, idRole primitive.ObjectID) (bool, error) {
	c, errC := db.GetCollection(Coluser)

	if errC != nil {
		return false, errC
	}
	_, err := c.UpdateOne(db.Ctx,
		bson.M{"_id": idUser},
		bson.M{
			"$addToSet": bson.M{
				"roles": idRole,
			},
		},
	)
	return err == nil, err
}

//Function to remove a role to an existed user
func RemoveRole2User(idUser primitive.ObjectID, idRole primitive.ObjectID) (bool, error) {
	c, errC := db.GetCollection(Coluser)

	if errC != nil {
		return false, errC
	}
	_, err := c.UpdateOne(db.Ctx,
		bson.M{"_id": idUser},
		bson.M{
			"$pull": bson.M{
				"roles": idRole,
			},
		},
	)
	return err == nil, err
}

//Function to get user by Id
func GetUserById(idUser primitive.ObjectID) (*models.User, error) {
	c, errC := db.GetCollection(Coluser)

	if errC != nil {
		return nil, errC
	}
	var result models.User
	var err = c.FindOne(db.Ctx, bson.M{"_id": idUser}).Decode(&result)
	return &result, err
}

func GetUserByDNI(dni string) (*models.User, error) {
	c, errC := db.GetCollection(Coluser)

	if errC != nil {
		return nil, errC
	}
	var result models.User
	var err = c.FindOne(db.Ctx, bson.M{"dni": dni}).Decode(&result)
	return &result, err
}

//Function to get a collection of all users
func GetUsers() (*[]models.User, error) {
	c, errC := db.GetCollection(Coluser)

	if errC != nil {
		return nil, errC
	}
	var results []models.User
	cursor, err := c.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &results)
	return &results, err
}

//Function to login by nick and password
func Login(nick string, password string) (*models.User, error) {
	c, errC := db.GetCollection(Coluser)

	if errC != nil {
		return nil, errC
	}
	result := models.User{}
	var err = c.FindOne(db.Ctx, bson.M{"nick": nick, "password": password}).Decode(&result)
	return &result, err
}

func LoginT(id string) (*models.User, error) {
	c, errC := db.GetCollection(Coluser)

	if errC != nil {
		return nil, errC
	}
	result := models.User{}
	ida, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = c.FindOne(db.Ctx, bson.M{"_id": ida}).Decode(&result)
	return &result, err
}

type UserTkn struct {
	Id string `json:"id"`
}

type UserToken struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func CreateTokenString(idUser string) (string, error) {
	if idUser == "" {
		return "", nil
	}
	//expirationTime := time.Now().Add(5 * time.Minute)
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &UserToken{
		Id: idUser,
		//StandardClaims: jwt.StandardClaims{
		//	// In JWT, the expiry time is expressed as unix milliseconds
		//	ExpiresAt: expirationTime.Unix(),
		//},
	})
	tokenstring, err := token.SignedString([]byte("**.38//9.--23%CZM%@**Z..X"))
	return tokenstring, err
}

func DecodeTokenString(myToken string) (string, error) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("**.38//9.--23%CZM%@**Z..X"), nil
	})

	if err == nil && token.Valid {
		user := UserToken{}
		token, err = jwt.ParseWithClaims(myToken, &user, func(token *jwt.Token) (interface{}, error) {
			return []byte("**.38//9.--23%CZM%@**Z..X"), err
		})
		return user.Id, err
	} else {
		return "", err
	}
}

func RefreshPasswordUser(idUser primitive.ObjectID) (*models.User, error) {
	user, err := GetUserById(idUser)
	if err != nil {
		return nil, err
	}
	_, p, err := models.Login("", user.Dni)
	if err != nil {
		return nil, err
	}
	c, errC := db.GetCollection(Coluser)

	if errC != nil {
		return nil, errC
	}
	_, err = c.UpdateOne(db.Ctx,
		bson.M{"_id": idUser},
		bson.M{"$set": bson.M{
			"password": p,
		}},
	)
	return user, err
}

func UpdatePasswordUser(idUser primitive.ObjectID, password string, newPassword string) (*models.User, error) {
	user, err := GetUserById(idUser)
	if err != nil {
		return nil, err
	}
	_, p, err := models.Login("", password)
	if err != nil {
		return nil, err
	}
	_, p2, err := models.Login("", newPassword)
	if err != nil {
		return nil, err
	}
	c, errC := db.GetCollection(Coluser)

	if errC != nil {
		return nil, errC
	}
	_, err = c.UpdateOne(db.Ctx,
		bson.M{"_id": idUser, "password": p},
		bson.M{"$set": bson.M{
			"password": p2,
		}},
	)
	if err != nil {
		return nil, err
	}
	return user, err
}
