package models

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Access struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	View     primitive.ObjectID `bson:"view"`
	Order    int                `bson:"order"`
	Position string             `bson:"position"`
}

func (a *Access) ToBson() bson.M {
	return bson.M{
		"_id,omitempty": a.Id,
		"view":          a.View,
		"order":         a.Order,
		"position":      a.Position,
	}
}

//cast accesses to list of bson.M
func AccessAddToBson(idM interface{}, idR interface{}, accessI interface{}) (*Module, *Role, *Access, *bson.M, error) {
	idm, ok := idM.(string)
	if !ok {
		return nil, nil, nil, nil, errors.New("error casting id module")
	}
	idr, ok := idR.(string)
	if !ok {
		return nil, nil, nil, nil, errors.New("error casting id role")
	}
	idma, _ := primitive.ObjectIDFromHex(idm)
	idra, _ := primitive.ObjectIDFromHex(idr)
	accessM, ok := accessI.(map[string]interface{})
	if !ok {
		return nil, nil, nil, nil, errors.New("error casting access")
	}
	id := primitive.NewObjectID()
	view, _ := primitive.ObjectIDFromHex(accessM["view"].(string))
	order := accessM["order"].(int)
	position := accessM["position"].(string)
	accessBson := bson.M{
		"_id":      id,
		"view":     view,
		"order":    order,
		"position": position,
	}
	return &Module{Id: idma}, &Role{Id: idra}, &Access{Id: id, View: view, Order: order, Position: position}, &accessBson, nil
}

func AccessUpdateToBson(nA interface{}, idA interface{}, accessI interface{}) (primitive.ObjectID, primitive.ObjectID, primitive.ObjectID, int, string, error) {
	idRole, ok := nA.(string)
	aux := primitive.ObjectID{}
	if !ok {
		return aux, aux, aux, -1, "", errors.New("error on id role")
	}
	idAccess, ok := idA.(string)
	if !ok {
		return aux, aux, aux, -1, "", errors.New("error on id access")
	}
	accessM, ok := accessI.(map[string]interface{})
	if !ok {
		return aux, aux, aux, -1, "", errors.New("error on data access")
	}
	idR, _ := primitive.ObjectIDFromHex(idRole)
	id, _ := primitive.ObjectIDFromHex(idAccess)
	view, _ := primitive.ObjectIDFromHex(accessM["view"].(string))
	order := accessM["order"].(int)
	position := accessM["position"].(string)
	return idR, id, view, order, position, nil
}
