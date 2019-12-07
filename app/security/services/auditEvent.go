package services

import (
	"../../db"
	"../models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net"
	"net/http"

	adminServices "../../admin/services"
)

var (
	Colaudit        = "auditEvents"
	errorPermission = errors.New("does not have permissions")
)

func AddAuditEvent(audit bson.M) error {
	c, errC := db.GetCollection(Colaudit)

	if errC != nil {
		return errC
	}
	_, errC = c.InsertOne(db.Ctx, audit)
	return errC
}

func AddOperationAudit(params graphql.ResolveParams, no string, ne string, errExecOper error) {
	contextTemp := params.Context
	idAudit := contextTemp.Value("current").(models.Values).Get("currentAudit").(primitive.ObjectID)
	data2set := models.OperationAuditToBson(params.Info.FieldName, true, no, ne, errExecOper, params.Context)
	c, errC := db.GetCollection(Colaudit)

	if errC != nil {
		return
	}
	_, err := c.UpdateOne(db.Ctx,
		bson.M{"_id": idAudit},
		bson.M{
			"$addToSet": bson.M{
				"operations": data2set,
			},
		},
	)
	if err != nil {
		fmt.Println("error on AddOperationAudit")
	}
}

func GetAuditEvents() (*[]models.AuditEvent, error) {
	c, errC := db.GetCollection(Colaudit)

	if errC != nil {
		return nil, errC
	}
	var results []models.AuditEvent
	cursor, err := c.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &results)
	return &results, err
}

func GetConsultGraphql(r *http.Request) (*models.RequestOptions, error) {
	var opts models.RequestOptions
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &opts, err
	}
	err = json.Unmarshal(body, &opts)
	if err != nil {
		return &opts, err
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return &opts, nil
}

func GetRemoteIp(r *http.Request) (string, error) {
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		return "", errors.New("Invalid user IP")
	}
	return ip + ":" + port, nil
}

func ValidatePermissionToOperation(params graphql.ResolveParams) error {
	contextTemp := params.Context
	operation := params.Info.FieldName
	idUser, ok := contextTemp.Value("current").(models.Values).Get("idUser").(string)
	if !ok {
		AddOperationAudit(params, "", "", errorPermission)
		return errors.New("does not have permissions - user nil")
	}
	idu, _ := primitive.ObjectIDFromHex(idUser)
	user, okGU := adminServices.GetUserById(idu)
	if okGU != nil {
		AddOperationAudit(params, "", "", errorPermission)
		return errors.New("does not have permissions - user not found")
	}
	oper, okGO := adminServices.GetOperationByName(operation)
	if okGO != nil {
		AddOperationAudit(params, "", "", errorPermission)
		return errors.New("does not have permissions - operation not found")
	}
	modules, okMO := adminServices.ModulesOperationByUser(oper.Id, user.Roles)
	if okMO != nil || len(*modules) == 0 {
		AddOperationAudit(params, "", "", errorPermission)
		return errors.New("does not have permissions")
	}
	return nil
}
