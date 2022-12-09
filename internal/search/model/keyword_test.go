package model_test

import (
	"testing"

	"github.com/sog01/solid-go/internal/search/model"
)

func TestKeyword_ValidateBadKeyword(t *testing.T) {
	tests := []struct {
		name    string
		k       model.Keyword
		wantErr bool
	}{
		{
			name:    "validate good keyword",
			k:       model.NewKeyword("good keyword"),
			wantErr: false,
		},
		{
			name:    "validate bad single keyword",
			k:       model.NewKeyword("bad single keyword crazy"),
			wantErr: true,
		},
		{
			name:    "validate bad multiple keyword",
			k:       model.NewKeyword("bad single keyword ugly, crazy"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.k.ValidateBadKeyword(); (err != nil) != tt.wantErr {
				t.Errorf("Keyword.ValidateBadKeyword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
