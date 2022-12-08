package model

import (
	"fmt"
	"strings"
)

type Keyword string

func (k Keyword) ValidateBadKeyword() error {
	errMessage := "consist of bad word '%s' isn't allowed"
	badWords := []string{"crazy", "ugly"}

	for _, badWord := range badWords {
		if strings.Contains(string(k), badWord) {
			return fmt.Errorf(errMessage, badWord)
		}
	}

	return nil
}

func NewKeyword(keyword string) Keyword {
	return Keyword(keyword)
}
