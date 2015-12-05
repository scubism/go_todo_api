package utils

import (
	"."
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"testing"
)

func getId(n int) string {
	return "00000000000000000000000" + strconv.Itoa(n)
}

func generateChildren(nums []int) []bson.ObjectId {
	N := len(nums)
	children := make([]bson.ObjectId, N)
	for i := 0; i < N; i++ {
		children[i] = bson.ObjectIdHex(getId(nums[i]))
	}
	return children
}

func assertChildren(t *testing.T, actual []bson.ObjectId, expected []bson.ObjectId) {
	N := len(expected)
	for i := 0; i < N; i++ {
		if actual[i] != expected[i] {
			t.Errorf("got %v\nwant %v", actual, expected)
			return
		}
	}
}

func TestMoveInChildrenToForward(t *testing.T) {
	children := generateChildren([]int{0, 1, 2, 3, 4, 5})
	targetId := bson.ObjectIdHex(getId(4))
	priorSiblingId := getId(2)
	actual, _ := utils.MoveInChildren(children, targetId, priorSiblingId)
	expected := generateChildren([]int{0, 1, 2, 4, 3, 5})
	assertChildren(t, actual, expected)
}

func TestMoveInChildrenToForwardFromEnd(t *testing.T) {
	children := generateChildren([]int{0, 1, 2, 3, 4, 5})
	targetId := bson.ObjectIdHex(getId(5))
	priorSiblingId := getId(2)
	actual, _ := utils.MoveInChildren(children, targetId, priorSiblingId)
	expected := generateChildren([]int{0, 1, 2, 5, 3, 4})
	assertChildren(t, actual, expected)
}

func TestMoveInChildrenToFirst(t *testing.T) {
	children := generateChildren([]int{0, 1, 2, 3, 4, 5})
	targetId := bson.ObjectIdHex(getId(4))
	priorSiblingId := ""
	actual, _ := utils.MoveInChildren(children, targetId, priorSiblingId)
	expected := generateChildren([]int{4, 0, 1, 2, 3, 5})
	assertChildren(t, actual, expected)
}

func TestMoveInChildrenToBackward(t *testing.T) {
	children := generateChildren([]int{0, 1, 2, 3, 4, 5})
	targetId := bson.ObjectIdHex(getId(2))
	priorSiblingId := getId(4)
	actual, _ := utils.MoveInChildren(children, targetId, priorSiblingId)
	expected := generateChildren([]int{0, 1, 3, 4, 2, 5})
	assertChildren(t, actual, expected)
}

func TestMoveInChildrenToBackwardFromFirst(t *testing.T) {
	children := generateChildren([]int{0, 1, 2, 3, 4, 5})
	targetId := bson.ObjectIdHex(getId(0))
	priorSiblingId := getId(4)
	actual, _ := utils.MoveInChildren(children, targetId, priorSiblingId)
	expected := generateChildren([]int{1, 2, 3, 4, 0, 5})
	assertChildren(t, actual, expected)
}

func TestMoveInChildrenToEnd(t *testing.T) {
	children := generateChildren([]int{0, 1, 2, 3, 4, 5})
	targetId := bson.ObjectIdHex(getId(2))
	priorSiblingId := getId(5)
	actual, _ := utils.MoveInChildren(children, targetId, priorSiblingId)
	expected := generateChildren([]int{0, 1, 3, 4, 5, 2})
	assertChildren(t, actual, expected)
}

func TestMoveInChildrenForMissingTarget(t *testing.T) {
	children := generateChildren([]int{0, 1, 2, 3, 4, 5})
	targetId := bson.ObjectIdHex(getId(6))
	priorSiblingId := getId(2)
	actual, _ := utils.MoveInChildren(children, targetId, priorSiblingId)
	expected := generateChildren([]int{0, 1, 2, 6, 3, 4, 5})
	assertChildren(t, actual, expected)
}

func TestMoveInChildrenForMissingTargetToFirst(t *testing.T) {
	children := generateChildren([]int{0, 1, 2, 3, 4, 5})
	targetId := bson.ObjectIdHex(getId(6))
	priorSiblingId := ""
	actual, _ := utils.MoveInChildren(children, targetId, priorSiblingId)
	expected := generateChildren([]int{6, 0, 1, 2, 3, 4, 5})
	assertChildren(t, actual, expected)
}

func TestMoveInChildrenForMissingPriorSiblingId(t *testing.T) {
	children := generateChildren([]int{0, 1, 2, 3, 4, 5})
	targetId := bson.ObjectIdHex(getId(4))
	priorSiblingId := getId(6)
	actual, _ := utils.MoveInChildren(children, targetId, priorSiblingId)
	expected := generateChildren([]int{0, 1, 2, 3, 5, 6, 4})
	assertChildren(t, actual, expected)
}

func TestMoveInChildrenForMissingTargetAndPriorSiblingId(t *testing.T) {
	children := generateChildren([]int{0, 1, 2, 3, 4, 5})
	targetId := bson.ObjectIdHex(getId(6))
	priorSiblingId := getId(7)
	actual, _ := utils.MoveInChildren(children, targetId, priorSiblingId)
	expected := generateChildren([]int{0, 1, 2, 3, 4, 5, 7, 6})
	assertChildren(t, actual, expected)
}
