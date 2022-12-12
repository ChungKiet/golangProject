package controllers

import (
	"github.com/gin-gonic/gin"
	"kietchung/models"
	"kietchung/request"
	"kietchung/services"
	"net/http"
)

type ChemistryController struct {
	ChemistryService services.ChemistryService
}

func New(chemistryService services.ChemistryService) ChemistryController {
	return ChemistryController{
		ChemistryService: chemistryService,
	}
}

func (uc *ChemistryController) ImportMaterial(ctx *gin.Context) {
	var input request.InsertChemistryReq

	// Get name from link
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Không thể parse input"})
		return
	}

	// Check already register
	chemistry, err := uc.ChemistryService.GetMaterialUrl(&request.GetChemistryReq{
		TypeChemical: input.TypeChemical,
		GroupName:    input.GroupName,
		TypeSpectrum: input.TypeSpectrum,
		Chemical:     input.Chemical,
	})

	if chemistry != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Tài liệu trùng!"})
		return
	}

	// Init new user
	var chemistryReq = models.Chemistry{
		GroupName:    input.GroupName,
		TypeChemical: input.TypeChemical,
		TypeSpectrum: input.TypeSpectrum,
		Chemical:     input.Chemical,
		VideoUrl:     input.VideoUrl,
	}

	// Create user
	_, err = uc.ChemistryService.ImportMaterial(&chemistryReq)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	// Response success
	ctx.JSON(http.StatusOK, gin.H{"message": "Import tai lieu thanh cong"})
}

func (uc *ChemistryController) GetMaterialUrl(ctx *gin.Context) {
	var getChemistryReq request.GetChemistryReq
	err := ctx.ShouldBindQuery(&getChemistryReq)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	chemistries, err := uc.ChemistryService.GetMaterialUrl(&getChemistryReq)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, chemistries)
}

func (uc *ChemistryController) UpdateMaterial(ctx *gin.Context) {
	var updateChemistry models.Chemistry
	err := ctx.ShouldBindJSON(&updateChemistry)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Khong the parse input"})
		return
	}

	chemistry, err := uc.ChemistryService.UpdateMaterial(&updateChemistry)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Tai lieu khong ton tai"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "chemistry": chemistry})
}

func (uc *ChemistryController) DeleteMaterial(ctx *gin.Context) {
	var updateChemistry request.DeleteChemistryReq
	err := ctx.ShouldBindQuery(&updateChemistry)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Khong the parse input"})
		return
	}

	err = uc.ChemistryService.DeleteMaterial(&updateChemistry)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Tai lieu khong ton tai"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Xoa thanh cong!"})
}

func (uc *ChemistryController) RegisterUserRoutes(rg *gin.RouterGroup) {
	chemistryRoute := rg.Group("/chemistry")
	chemistryRoute.POST("/import-material", uc.ImportMaterial)
	chemistryRoute.GET("/get-material", uc.GetMaterialUrl)
	chemistryRoute.PUT("/update-material", uc.UpdateMaterial)
	chemistryRoute.DELETE("/delete-material", uc.DeleteMaterial)
}

func (uc *ChemistryController) GetReferenceDocument(ctx *gin.Context) {
	var getRefDocument request.GetRefDocument
	err := ctx.ShouldBindQuery(&getRefDocument)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	refDocument, err := uc.ChemistryService.GetReferenceDocument(&getRefDocument)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, refDocument)
}

func (uc *ChemistryController) ImportReferenceDocument(ctx *gin.Context) {
	var getRefDocument request.GetRefDocument
	err := ctx.ShouldBindQuery(&getRefDocument)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	refDocument, err := uc.ChemistryService.GetReferenceDocument(&getRefDocument)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, refDocument)
}
