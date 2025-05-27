package agristat

import (
	"encoding/json"
	"net/http"
)

type ReportHandler struct {
	service *ReportService
}

func NewReportHandler(s *ReportService) *ReportHandler {
	return &ReportHandler{service: s}
}

func (h *ReportHandler) GetCropStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.service.GetCropStats()
	if err != nil {
		http.Error(w, "Ошибка при получении статистики", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
