package utils

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

func MoveInChildren(
	children []bson.ObjectId,
	targetId bson.ObjectId,
	priorSiblingId string) ([]bson.ObjectId, error) {

	N := len(children)
	converted := make([]bson.ObjectId, N)
	var foundTargetId = false
	var foundPriorSiblingId = false

	ptr := 0
	if priorSiblingId == "" {
		converted[ptr] = targetId
		ptr++
		foundPriorSiblingId = true
		for i := 0; i < N; i++ {
			if children[i] == targetId {
				foundTargetId = true
			} else {
				converted[ptr] = children[i]
				ptr++
			}
		}
	} else {
		priorSiblingId := bson.ObjectIdHex(priorSiblingId)
		for i := 0; i < N; i++ {
			if children[i] == priorSiblingId {
				converted[ptr] = priorSiblingId
				ptr++
				converted[ptr] = targetId
				ptr++
				foundPriorSiblingId = true
			} else if children[i] == targetId {
				foundTargetId = true
			} else {
				converted[ptr] = children[i]
				ptr++
			}
		}
	}
	if ptr != N {
		return nil, errors.New("Unkown sort error")
	}
	if !foundTargetId {
		return nil, errors.New("Target id is not found")
	}
	if !foundPriorSiblingId {
		return nil, errors.New("Prior sibling id is not found")
	}
	return converted, nil
}
