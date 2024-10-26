package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"security-go/entity"
	"security-go/service"
	"security-go/util"
)

type HospitalController struct {
	HospitalService service.HospitalService
}

func NewHospitalController(HospitalService service.HospitalService) *HospitalController {
	return &HospitalController{
		HospitalService: HospitalService,
	}
}

func (ctrl *HospitalController) CreateHospital(c *gin.Context) {
	var Hospital entity.Hospital
	if err := c.ShouldBindJSON(&Hospital); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.HospitalService.CreateHospital(&Hospital); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Hospital)
}

func (ctrl *HospitalController) GetHospitalById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	Hospital, err := ctrl.HospitalService.GetHospitalById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hospital not found"})
		return
	}
	c.JSON(http.StatusOK, Hospital)
}

func (ctrl *HospitalController) UpdateHospital(c *gin.Context) {
	var Hospital entity.Hospital
	if err := c.ShouldBindJSON(&Hospital); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.HospitalService.UpdateHospital(&Hospital); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Hospital)
}

func (ctrl *HospitalController) DeleteHospital(c *gin.Context) {
	id, _ := util.StringToUint(c.Param("id"))
	if err := ctrl.HospitalService.DeleteHospital(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
