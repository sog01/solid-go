package service

import (
	"log"

	"github.com/sog01/solid-go/internal/search/model"
	"github.com/sog01/solid-go/internal/search/repository"
)

// interface area
type Search interface {
	SearchEmployees(keywordString string) ([]*model.Employee, error)
}

type Sync interface {
	InsertEmployee(e *model.Employee) error
	SeedingEmployees(idStart, n int) error
	UpdateEmployee(e *model.Employee) error
	DeleteEmployee(id int) error
	CheckHealth() error
}

type Construct interface {
	CreateIndex() error
}

// search service area
type SearchService struct {
	repository repository.Repository
}

func NewSearchService(repo repository.Repository) *SearchService {
	return &SearchService{repository: repo}
}

func (s *SearchService) SearchEmployees(keywordString string) ([]*model.Employee, error) {
	keyword := model.NewKeyword(keywordString)
	if err := keyword.ValidateBadKeyword(); err != nil {
		log.Printf("found a bad keyword: %v\n", keyword)
		return nil, err
	}
	employees, err := s.repository.SearchData(keyword)
	if err != nil {
		log.Println("failed search employees: ", err)
		return nil, err
	}
	return employees, nil
}

// sync service area
type SyncService struct {
	repository repository.Repository
}

func NewSyncService(repo repository.Repository) *SyncService {
	return &SyncService{repository: repo}
}

func (s *SyncService) InsertEmployee(e *model.Employee) error {
	err := s.repository.InsertData(e)
	if err != nil {
		log.Println("failed insert employee: ", err)
		return err
	}
	return nil
}

func (s *SyncService) SeedingEmployees(idStart, n int) error {
	err := s.repository.SeedingData(idStart, n)
	if err != nil {
		log.Println("failed seeding employees: ", err)
		return err
	}
	return nil
}

func (s *SyncService) UpdateEmployee(e *model.Employee) error {
	err := s.repository.UpdateData(e)
	if err != nil {
		log.Println("failed update employee: ", err)
		return err
	}
	return nil
}

func (s *SyncService) DeleteEmployee(id int) error {
	err := s.repository.DeleteData(id)
	if err != nil {
		log.Println("failed delete employee: ", err)
		return err
	}
	return nil
}

func (s *SyncService) CheckHealth() error {
	err := s.repository.CheckHealth()
	if err != nil {
		log.Println("failed do health check: ", err)
		return err
	}
	return nil
}

// construct service area
type ConstructService struct {
	repository repository.Repository
}

func NewConstructService(repo repository.Repository) *ConstructService {
	return &ConstructService{repository: repo}
}

func (s *ConstructService) CreateIndex() error {
	err := s.repository.CreateIndex()
	if err != nil {
		log.Println("failed create index: ", err)
		return err
	}
	return nil
}
