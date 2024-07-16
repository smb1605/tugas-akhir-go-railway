package pengumuman

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/golang_kantin/internal/api"
	"github.com/onainadapdap1/golang_kantin/internal/service/pengumuman"
	"github.com/onainadapdap1/golang_kantin/models"
)

type PengumumanHandler interface {
	CreatePengumuman(c *gin.Context)
	GetAllPengumuman(c *gin.Context)
	UpdatedPengumuman(c *gin.Context)
	DeletePengumumanByID(c *gin.Context)
}

type pengumumanHandler struct {
	service pengumuman.PengumumanService
}

func NewPengumumanHandler(service pengumuman.PengumumanService) PengumumanHandler {
	return &pengumumanHandler{service}
}

func (h *pengumumanHandler) CreatePengumuman(c *gin.Context) {
	var inputPengumuman api.CreatePengumumanInput
	if err := c.ShouldBindJSON(&inputPengumuman); err != nil {
		response := gin.H{
			"success": false,
			"message": "failed to create pengumuman",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// Konversi string tanggal menjadi tipe data time.Time
	// time.Parse("01-2006", inputPengalaman.TahunMulai)

	parsedTime, err := time.Parse("2006-01-02", inputPengumuman.TanggalBerakhir)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	createPengumuman := models.Pengumuman{
		TanggalBerakhir: parsedTime,
		Deskripsi:       inputPengumuman.Deskripsi,
	}

	newPengumuman, err := h.service.CreatePengumuman(createPengumuman)
	if err != nil {
		response := gin.H{
			"success": false,
			"message": "failed to create pengumuman",
			"error":   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	successResp := gin.H{
		"success": true,
		"message": "successfully created new announcement",
		"data":    newPengumuman,
	}
	c.JSON(http.StatusCreated, successResp)
}

func (h *pengumumanHandler) GetAllPengumuman(c *gin.Context) {
	// start := time.Now()

	pengumumans, err := h.service.GetAllPengumuman()
	// elapsed := time.Since(start)
	// log.Printf("GetAllPengumuman took %s", elapsed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no data is found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": pengumumans})
}

func (h *pengumumanHandler) UpdatedPengumuman(c *gin.Context) {
	pengumumanID, _ := strconv.Atoi(c.Param("id")) // 183
	log.Println("error 1 handler")
	var inputPengumuman api.UpdatePengumumanInput
	if err := c.ShouldBindJSON(&inputPengumuman); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	log.Println("error 2 handler")
	updatedPengumuman, err := h.service.UpdatePengumuman(uint(pengumumanID), inputPengumuman)
	if err != nil {
		response := gin.H{
			"success": false,
			"message": "failed to update pengumuman",
			"error":   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	log.Println("error 3 handler")
	successResp := gin.H{
		"success": true,
		"message": "successfully updated new announcement",
		"data":    updatedPengumuman,
	}
	c.JSON(http.StatusOK, successResp)
}

func (h *pengumumanHandler) DeletePengumumanByID(c *gin.Context) {
	pengumumanID, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeletePengumumanByID(uint(pengumumanID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete pengumuman"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "pengumuman successfully deleted",
	})
}