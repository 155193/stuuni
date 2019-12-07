package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuditEvent struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	User       primitive.ObjectID `json:"user" bson:"user,omitempty"`
	IpLocal    string             `json:"ipLocal" bson:"ipLocal"`
	IpRemote   string             `json:"ipRemote" bson:"ipRemote"`
	Agent      string             `json:"agent" bson:"agent"`
	Auth       bool               `json:"auth" bson:"auth"`
	Token      string             `json:"token" bson:"token"`
	Consult    bson.M             `json:"consult" bson:"consult"`
	Operations []OperationAudit   `json:"operations" bson:"operations"`
}

type RequestOptions struct {
	Query         string                 `json:"query" url:"query" schema:"query"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
	OperationName string                 `json:"operationName" url:"operationName" schema:"operationName"`
}

type Values struct {
	M map[string]interface{} `json:"m" url:"m" schema:"m"`
}

func (v Values) Get(key string) interface{} {
	return v.M[key]
}

func AuditUserToBson(idAudit primitive.ObjectID, remote interface{}, local interface{}, agent interface{}, user interface{}, auth interface{}, token interface{}, consult interface{}) bson.M {
	audit := bson.M{
		"_id":        idAudit,
		"ipLocal":    local.(string),
		"ipRemote":   remote.(string),
		"agent":      agent.(string),
		"auth":       auth.(bool),
		"token":      token.(string),
		"consult":    consult,
		"operations": []OperationAudit{},
	}
	if user != "" {
		audit["user"], _ = primitive.ObjectIDFromHex(user.(string))
	} else {
		audit["user"] = ""
	}
	return audit
}
