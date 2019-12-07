package models

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type View struct {
	Id          primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name        string               `bson:"name"`
	Description string               `bson:"description"`
	Url         string               `bson:"url"`
	Icon        string               `bson:"icon"`
	Operations  []primitive.ObjectID `bson:"operations"`
}

//Constructor for View model
func (v *View) New(id primitive.ObjectID, Name string, Description string, Url string, Icon string) View {
	v.Id = id
	v.Name = Name
	v.Description = Description
	v.Url = Url
	v.Icon = Icon
	v.Operations = []primitive.ObjectID{}
	return *v
}

func (v *View) ToBsonM() bson.M {
	return bson.M{
		"_id,omitempty": v.Id,
		"name":          v.Name,
		"description":   v.Description,
		"url":           v.Url,
		"icon":          v.Icon,
	}
}

func ViewToBson(idM interface{}, viewI interface{}) (*Module, *View, *bson.M, error) {
	idm, ok := idM.(string)
	if !ok {
		return nil, nil, nil, errors.New("error casting Id module")
	}
	idma, _ := primitive.ObjectIDFromHex(idm)
	viewM, ok := viewI.(map[string]interface{})
	if !ok {
		return nil, nil, nil, errors.New("error casting")
	}
	id := primitive.NewObjectID()
	name := viewM["name"].(string)
	des := viewM["description"].(string)
	url := viewM["url"].(string)
	icon := viewM["icon"].(string)
	//operations := viewM["operations"].([]string)
	viewBson := bson.M{
		"_id":         id,
		"name":        name,
		"description": des,
		"url":         url,
		"icon":        icon,
		//"operations":  operations,
	}
	return &Module{Id: idma}, &View{Id: id, Name: name, Description: des, Url: url, Icon: icon}, &viewBson, nil
}

func ViewUpdateToBson(idV interface{}, viewI interface{}) (*View, error) {
	id, ok := idV.(string)
	if !ok {
		return nil, errors.New("error casting Id view")
	}
	idv, _ := primitive.ObjectIDFromHex(id)
	viewM, ok := viewI.(map[string]interface{})
	if !ok {
		return nil, errors.New("error casting view")
	}
	name := viewM["name"].(string)
	description := viewM["description"].(string)
	url := viewM["url"].(string)
	icon := viewM["icon"].(string)
	return &View{Id: idv, Name: name, Description: description, Url: url, Icon: icon,}, nil
}
