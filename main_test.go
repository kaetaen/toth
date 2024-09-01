package main

import (
	"reflect"
	"testing"

	"golang.design/x/clipboard"
)

// Tests the getImage function when the clipboard contains an image
func TestGetImage_ClipboardHasImage(t *testing.T) {
	expectedImage := []byte{1, 2, 3, 4, 5}
	clipboard.Write(clipboard.FmtImage, expectedImage)

	gotImage, gotNotFound := getImage()

	if !reflect.DeepEqual(gotImage, expectedImage) {
		t.Errorf("Expected image: %v, got: %v", expectedImage, gotImage)
	}
	if gotNotFound {
		t.Error("Expected not found to be false, got true")
	}
}

// Tests the getImage function when the clipboard not contains an image
func TestGetImage_ClipboardHasNoImage(t *testing.T) {
	clipboard.Write(clipboard.FmtImage, []byte{})
	gotImage, gotNotFound := getImage()

	if len(gotImage) != 0 {
		t.Errorf("Expected image to be empty, got: %v", gotImage)
	}
	if !gotNotFound {
		t.Error("Expected not found to be true, got false")
	}
}
