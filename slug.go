package main

import (
	"errors"
	"regexp"
)

var (
	SlugFormat         = "^[a-zA-Z0-9-]+$"
	SlugFormatRegex, _ = regexp.Compile(SlugFormat)
	SlugMinLength      = 1
	SlugMaxLength      = 100

	ErrInvalidSlugFormat = errors.New("Slug contained improper format")
	ErrInvalidSlugLength = errors.New("Slug was not of valid length")
)

type Slug struct {
	ID       string `bson:"_id,omitempty"`
	Name     string `bson:"name,omitempty"`
	Redirect string `bson:"redirect,omitempty"`
}

func (s *Slug) ValidateSlugFormat() error {
	if !SlugFormatRegex.Match([]byte(s.Name)) {
		return ErrInvalidSlugFormat
	}
	return nil
}

func (s *Slug) ValidateSlugLength() error {
	length := len([]byte(s.Name))
	if length < SlugMinLength || length > SlugMaxLength {
		return ErrInvalidSlugLength
	}
	return nil
}
