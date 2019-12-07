package db

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"reflect"
	"testing"
	"time"
)

/* function get String Data */
func GetStringData(data map[string]interface{}, bsonData bson.M, field string) error {
	temp, ok := data[field].(string)
	if ok {
		bsonData[field] = temp
		return nil
	} else {
		return errors.New("Not found " + field)
	}
}

/* function get Float64 Data */
func GetFloat64Data(data map[string]interface{}, bsonData bson.M, field string) error {
	temp, ok := data[field].(float64)
	if ok {
		bsonData[field] = temp
		return nil
	} else {
		return errors.New("Not found " + field)
	}
}

func GetIntData(data map[string]interface{}, bsonData bson.M, field string) error {
	temp, ok := data[field].(int)
	if ok {
		bsonData[field] = temp
		return nil
	} else {
		return errors.New("Not found " + field)
	}
}

func GetTimeData(data map[string]interface{}, bsonData bson.M, field string) error {
	temp, ok := data[field].(string)
	if ok {
		timeStp, err := time.Parse(time.RFC3339, temp)
		if err != nil {
			return err
		}
		bsonData[field] = timeStp
		return nil
	} else {
		return errors.New("Not found " + field)
	}
}

/* function get Boolean Data */
func GetBooleanData(data map[string]interface{}, bsonData bson.M, field string) error {
	temp, ok := data[field].(bool)
	if ok {
		bsonData[field] = temp
		return nil
	} else {
		return errors.New("Not found " + field)
	}
}

func GetIdData(data map[string]interface{}, bsonData bson.M, field string) error {
	temp, ok := data[field].(string)
	if ok {
		bsonData[field], _ = primitive.ObjectIDFromHex(temp)
		return nil
	} else {
		return errors.New("Not found " + field)
	}
}

/* function for setup test case */
func SetupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

/* function for setup Sub test */
func SetupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("setup sub test")
	return func(t *testing.T) {
		t.Log("teardown sub test")
	}
}

/*Function for get root*/
func GetFloatMatrix(data map[string]interface{}, data2Set bson.M, field string) error {
	matrixData, ok := data[field].([][]float64)
	if ok {
		data2Set[field] = matrixData
		return nil
	} else {
		return errors.New("Not found " + field)
	}
}

/*Function for get root*/
func GetListData(data map[string]interface{}, data2Set bson.M, field string) error {
	_, ok := data[field]
	if ok {
		//var List []primitive.ObjectID
		////for i := 0; i < list; i++ {
		////	List = append(List, bson.ObjectIdHex())
		////}
		//data2Set[field] = List
		return nil
	} else {

		return errors.New("Not found " + field)
	}
}

/* test structure*/
type CaseTest struct {
	Name         string
	Id           primitive.ObjectID
	Idaux        primitive.ObjectID
	NAccess      int
	Order        int
	Position     string
	Nam          string
	Des          string
	Url          string
	Icon         string
	Dni          string
	Period       string
	Address      string
	Days         []string
	Start        time.Time
	End          time.Time
	Model        bson.M
	IdList       []primitive.ObjectID
	ObjectList   []interface{}
	Thereiserror bool
	Action       bool
}
