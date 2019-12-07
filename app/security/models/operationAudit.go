package models

import (
	adminServices "../../admin/services"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OperationAudit struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Operation  primitive.ObjectID `json:"operation" bson:"operation"`
	Permission bool               `json:"permission" bson:"permission"`
	NowStruct  bson.M             `json:"nowStruct" bson:"nowStruct"`
	NewStruct  bson.M             `json:"newStruct" bson:"newStruct"`
	Error      string             `json:"error" bson:"error"`
	Result     bool               `json:"result" bson:"result"`
	TimeB      time.Time          `json:"timeB" bson:"timeB"`
	TimeE      time.Time          `json:"timeE" bson:"timeE"`
}

func OperationAuditToBson(ope string, permission bool, no interface{}, ne interface{}, err error, context context.Context) bson.M {
	timeB := context.Value("current").(Values).Get("time").(time.Time)
	now, _ := json.Marshal(no)
	new_, _ := json.Marshal(ne)
	errs := ""
	result := true
	if err != nil {
		errs = err.Error()
		result = false
	}
	audit := bson.M{
		"_id":        primitive.NewObjectID(),
		"permission": permission,
		"nowStruct":  string(now),
		"newStruct":  string(new_),
		"error":      errs,
		"result":     result,
		"timeB":      timeB,
		"timeE":      time.Now(),
	}
	oper, errOper := adminServices.GetOperationByName(ope)
	if errOper == nil {
		audit["operation"] = oper.Id
	}
	return audit
}
