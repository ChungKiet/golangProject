package services

import (
	"kietchung/models"
	"kietchung/request"
)

type ChemistryService interface {
	GetMaterialUrl(chemistry *request.GetChemistryReq) ([]*models.Chemistry, error)
	ImportMaterial(chemistry *models.Chemistry) (*models.Chemistry, error)
	UpdateMaterial(chemistry *models.Chemistry) (*models.Chemistry, error)
	DeleteMaterial(chemistry *request.DeleteChemistryReq) error
}
