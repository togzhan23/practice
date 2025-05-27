package agristat

import (
	"log"
	"practice/models"
)

type ReportService struct {
	repo *ReportRepository
}

func NewReportService(repo *ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetCropStats() ([]models.CropStat, error) {
	stats, err := s.repo.GetCropStatistics()
	if err != nil {
		log.Printf("Error fetching crop statistics: %v", err)
		return nil, err
	}
	return stats, nil
}
