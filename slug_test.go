package main

import (
	"fmt"
	"testing"
)

func TestSlugTooSmall(t *testing.T) {
	slug := &Slug{"", "", ""}
	if slug.ValidateSlugLength() != ErrInvalidSlugLength {
		t.Error("Expected ErrInvalidSlugLength for 0 character slug")
	}
}

func TestSlugTooLarge(t *testing.T) {
	slug := &Slug{"", "AxdHx0WzmLw3KJfELyUu0vFcgJMoHNa4Qopra2ejfm0HklzpYWYpdQTTsu4W3yOngfR01DcJfw89s5e5MHUANdId2H78MSgYlQZWx", ""}
	if slug.ValidateSlugLength() != ErrInvalidSlugLength {
		t.Error("Expected ErrInvalidSlugLength for 101 character slug")
	}
}

func TestSlugWithSpaces(t *testing.T) {
	slug := &Slug{"", "a-b c-d", ""}
	if slug.ValidateSlugFormat() != ErrInvalidSlugFormat {
		t.Error("Expected ErrInvalidSlugFormat for slug with spaces")
	}
}

func TestSlugWithNumericAndText(t *testing.T) {
	slug := &Slug{"", "a-b-c-123", ""}
	if slug.ValidateSlugFormat() != nil {
		t.Error(fmt.Sprintf("Expected %s to be a valid slug with numbers and characters", slug.Name))
	}
}
