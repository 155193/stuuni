package models

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Accesses    []Access           `bson:"accesses"`
}

func (r *Role) New(id primitive.ObjectID, name string, description string, access []Access) Role {
	r.Id = id
	r.Name = name
	r.Description = description
	r.Accesses = access
	return *r
}

func (r *Role) ToBsonM() bson.M {
	return bson.M{
		"_id,omitempty": r.Id,
		"name":          r.Name,
		"description":   r.Description,
		//"accesses":      r.Accesses,
	}
}

func RoleAddToBson(idM interface{}, RoleI interface{}) (*Module, *Role, *bson.M, error) {
	idm, ok := idM.(string)
	if !ok {
		return nil, nil, nil, errors.New("error casting id Module")
	}
	idma, err := primitive.ObjectIDFromHex(idm)
	if err != nil {
		return nil, nil, nil, errors.New("error at idRole")
	}
	roleM, ok := RoleI.(map[string]interface{})
	if !ok {
		return nil, nil, nil, errors.New("error casting Role")
	}
	id := primitive.NewObjectID()
	name := roleM["name"].(string)
	description := roleM["description"].(string)
	role := bson.M{
		"_id":         id,
		"name":        name,
		"description": description,
		"accesses":    []Access{},
	}
	return &Module{Id: idma}, &Role{Id: id, Name: name, Description: description, Accesses: []Access{}}, &role, nil
}

func RoleUpdateToBson(idR interface{}, RoleI interface{}) (*Role, error) {
	id, ok := idR.(string)
	if !ok {
		return nil, errors.New("error casting id Role")
	}
	idr, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	roleM, ok := RoleI.(map[string]interface{})
	if !ok {
		return nil, errors.New("error casting Role")
	}
	return &Role{Id: idr, Name: roleM["name"].(string), Description: roleM["description"].(string)}, nil
}
