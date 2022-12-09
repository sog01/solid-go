// +build integration

package service_test

import (
	"testing"
	"time"

	"github.com/sog01/solid-go/internal/search/model"
	"github.com/sog01/solid-go/internal/search/repository"
	"github.com/sog01/solid-go/internal/search/service"
	"github.com/stretchr/testify/assert"
)

var repo repository.Repository

func setup() {
	repo = repository.NewElasticSearch("http://localhost:9500")
	repo.SeedingData(1, 4)
	service.NewConstructService(repo).CreateIndex()
}

func TestMain(m *testing.M) {
	setup()
	time.Sleep(1 * time.Second) // add sleep to ensure all data is already saved
	m.Run()                     // running the tests
	teardown()
}

func TestSearchService_SearchEmployees(t *testing.T) {
	type args struct {
		keywordString string
	}
	tests := []struct {
		name    string
		args    args
		want    []*model.Employee
		wantErr bool
	}{
		{
			name: "Search Exists Employee",
			args: args{
				keywordString: "person",
			},
			want: []*model.Employee{
				{
					Id:      1,
					Name:    "person 1",
					Address: "address 1",
					Salary:  100,
				},
				{
					Id:      2,
					Name:    "person 2",
					Address: "address 2",
					Salary:  200,
				},
				{
					Id:      3,
					Name:    "person 3",
					Address: "address 3",
					Salary:  300,
				},
			},
			wantErr: false,
		},
		{
			name: "Search Not Exists Employee",
			args: args{
				keywordString: "personnotexists",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Search a Bad Word",
			args: args{
				keywordString: "personnotexists crazy",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service.NewSearchService(repo)
			got, err := s.SearchEmployees(tt.args.keywordString)
			assert.Equal(t, tt.want, got)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func teardown() {
	for i := 1; i < 4; i++ {
		repo.DeleteData(i)
	}
}
