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

	var foundTargetId = false
	for i := 0; i < N; i++ {
		if children[i] == targetId {
			foundTargetId = true
			break
		}
	}

	if !foundTargetId {
		return nil, errors.New("Target id is not found")
	}

	converted := make([]bson.ObjectId, N)

	ptr := 0
	if priorSiblingId == "" {
		converted[ptr] = targetId
		ptr++
		for i := 0; i < N; i++ {
			if children[i] != targetId {
				converted[ptr] = children[i]
				ptr++
			}
		}
	} else {
		priorSiblingId := bson.ObjectIdHex(priorSiblingId)
		var foundPriorSiblingId = false
		for i := 0; i < N; i++ {
			if children[i] == priorSiblingId {
				converted[ptr] = priorSiblingId
				ptr++
				converted[ptr] = targetId
				ptr++
				foundPriorSiblingId = true
			} else if children[i] != targetId {
				converted[ptr] = children[i]
				ptr++
			}
		}

		if !foundPriorSiblingId {
			return nil, errors.New("Prior sibling id is not found")
		}
	}

	return converted, nil
}
