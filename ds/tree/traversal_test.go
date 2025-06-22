package tree

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	bst := &BST{}
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	expected := []int{2, 3, 4, 5, 6, 7, 8}
	actual := bst.InorderTraversal()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestPreorderTraversal(t *testing.T) {
	bst := &BST{}
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	expected := []int{5, 3, 2, 4, 7, 6, 8}
	actual := bst.PreorderTraversal()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestPostorderTraversal(t *testing.T) {
	bst := &BST{}
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	expected := []int{2, 4, 3, 6, 8, 7, 5}
	actual := bst.PostorderTraversal()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
