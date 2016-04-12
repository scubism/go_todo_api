package utils

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

var ErrInvalidIdFormat = errors.New("Invalid id format")

func ValidateObjectIdHex(ids ...string) error {
	for _, id := range ids {
		if !bson.IsObjectIdHex(id) {
			return ErrInvalidIdFormat
		}
	}
	return nil
}
