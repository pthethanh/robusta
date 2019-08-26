package bsonutil

import "github.com/globalsign/mgo/bson"

// ToBSONMap convert v to bson.M
func ToBSONMap(v interface{}) (bson.M, error) {
	b, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}
	var m bson.M
	if err := bson.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}
