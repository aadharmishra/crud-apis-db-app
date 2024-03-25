package common

import (
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func ConvertStructToBson[T any](data []T) []bson.M {
	result := make([]bson.M, len(data))
	for i, doc := range data {
		result[i] = bson.M{}
		typeOfStruct := reflect.TypeOf(doc)
		for j := 0; j < typeOfStruct.NumField(); j++ {
			field := typeOfStruct.Field(j)
			value := reflect.ValueOf(doc).Field(j)
			fieldName := strings.TrimSuffix(strings.SplitN(field.Tag.Get("bson"), ",", 2)[0], ",")
			if fieldName == "" {
				fieldName = field.Name
			}
			result[i][fieldName] = value.Interface()
		}
	}
	return result
}
