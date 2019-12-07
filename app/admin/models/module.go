package models

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Module struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Icon  string             `bson:"icon"`
	Color string             `bson:"color"`
	Roles []Role             `bson:"roles"`
	Views []View             `bson:"views"`
}

func (m *Module) New(id primitive.ObjectID, name string, icon string, color string, roles []Role, views []View) Module {
	m.Id = id
	m.Name = name
	m.Icon = icon
	m.Color = color
	m.Roles = roles
	m.Views = views
	return *m
}

func (m *Module) ToBson() bson.M {
	return bson.M{
		"_id,omitempty": m.Id,
		"name":          m.Name,
		"icon":          m.Icon,
		"color":         m.Color,
	}
}

func ModuleToBson(moduleI interface{}) (*Module, *bson.M, error) {
	moduleM, ok := moduleI.(map[string]interface{})
	if !ok {
		return nil, nil, errors.New("error casting module")
	}
	id := primitive.NewObjectID()
	name := moduleM["name"].(string)
	icon := moduleM["icon"].(string)
	color := moduleM["color"].(string)
	moduleBson := bson.M{
		"_id":   id,
		"name":  name,
		"icon":  icon,
		"color": color,
		"roles": []Role{},
		"views": []View{},
	}
	return &Module{Id: id, Name: name, Icon: icon, Color: color, Roles: []Role{}, Views: []View{}}, &moduleBson, nil
}

func ModuleUpdateToBson(idI interface{}, moduleI interface{}) (*Module, *bson.M, error) {
	idM, ok := idI.(string)
	if !ok {
		return nil, nil, errors.New("error casting id module")
	}
	idm, _ := primitive.ObjectIDFromHex(idM)
	moduleM, ok := moduleI.(map[string]interface{})
	if !ok {
		return nil, nil, errors.New("error casting module")
	}
	moduleBson := bson.M{
		"name":  moduleM["name"].(string),
		"icon":  moduleM["icon"].(string),
		"color": moduleM["color"].(string),
	}
	return &Module{Id: idm}, &moduleBson, nil
}
