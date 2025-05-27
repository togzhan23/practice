package agristat

import (
	"database/sql"
	"practice/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

// Пример метода получения статистики по урожаю
func (r *ReportRepository) GetCropStatistics() ([]models.CropStat, error) {
	rows, err := r.db.Query("SELECT crop_name, SUM(quantity) as total_quantity FROM crop_production GROUP BY crop_name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []models.CropStat
	for rows.Next() {
		var stat models.CropStat
		if err := rows.Scan(&stat.CropName, &stat.TotalQuantity); err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}

	return stats, nil
}
