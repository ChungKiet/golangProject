package services

import (
	"context"
	"errors"
	"kietchung/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"kietchung/models"
)

/*
@Author: DevProblems(Sarang Kumar)
@YTChannel: https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ
*/
type ChemistryServiceImpl struct {
	refDocumentCollection *mongo.Collection
	chemistryCollection   *mongo.Collection
	ctx                   context.Context
}

func NewUserService(chemistryCollection *mongo.Collection, refDocument *mongo.Collection, ctx context.Context) ChemistryService {
	return &ChemistryServiceImpl{
		refDocumentCollection: refDocument,
		chemistryCollection:   chemistryCollection,
		ctx:                   ctx,
	}
}

func (c *ChemistryServiceImpl) ImportMaterial(chemistry *models.Chemistry) (*models.Chemistry, error) {
	_, err := c.chemistryCollection.InsertOne(c.ctx, chemistry)
	return chemistry, err
}

func (c *ChemistryServiceImpl) GetMaterialUrl(chemistry *request.GetChemistryReq) ([]*models.Chemistry, error) {
	filter := bson.M{}
	var res []*models.Chemistry
	if chemistry.TypeChemical != "" {
		filter["type_chemical"] = chemistry.TypeChemical
	}

	if chemistry.GroupName != "" {
		filter["group_name"] = chemistry.GroupName
	}

	if chemistry.TypeSpectrum != "" {
		filter["type_spectrum"] = chemistry.TypeSpectrum
	}

	if chemistry.Chemical != "" {
		filter["chemical"] = chemistry.Chemical
	}

	cursor, err := c.chemistryCollection.Find(c.ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(c.ctx) {
		var chemistryRes models.Chemistry
		err := cursor.Decode(&chemistryRes)
		if err != nil {
			return nil, err
		}
		res = append(res, &chemistryRes)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(c.ctx)

	if len(res) == 0 {
		return nil, errors.New("documents not found")
	}
	return res, err
}

// them truong id

func (c *ChemistryServiceImpl) UpdateMaterial(chemistry *models.Chemistry) (*models.Chemistry, error) {
	filter := bson.M{}
	if chemistry.TypeChemical == "" || chemistry.TypeSpectrum == "" || chemistry.Chemical == "" || chemistry.GroupName == "" {
		return nil, errors.New("Request invalid!")
	}

	filter["group_name"] = chemistry.GroupName
	filter["type_chemical"] = chemistry.TypeChemical
	filter["type_spectrum"] = chemistry.TypeSpectrum
	filter["chemical"] = chemistry.Chemical

	update := bson.M{}

	if chemistry.VideoUrl != "" {
		update["video_url"] = chemistry.VideoUrl
	}

	result, _ := c.chemistryCollection.UpdateOne(c.ctx, filter, bson.M{
		"$set": update,
	})

	if result.MatchedCount != 1 {
		return nil, errors.New("no matched document found for update")
	}
	return chemistry, nil
}

func (c *ChemistryServiceImpl) DeleteMaterial(chemistry *request.DeleteChemistryReq) error {
	filter := bson.M{}
	if chemistry.TypeChemical == "" || chemistry.TypeSpectrum == "" || chemistry.Chemical == "" || chemistry.GroupName == "" {
		return errors.New("Request invalid!")
	}

	filter["group_name"] = chemistry.GroupName
	filter["type_chemical"] = chemistry.TypeChemical
	filter["type_spectrum"] = chemistry.TypeSpectrum
	filter["chemical"] = chemistry.Chemical

	result, _ := c.chemistryCollection.DeleteOne(c.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

func (c *ChemistryServiceImpl) GetReferenceDocument(refDoc *request.GetRefDocument) ([]*models.ReferenceDocument, error) {
	filter := bson.M{}
	var res []*models.ReferenceDocument
	if refDoc.Type != "" {
		filter["type"] = refDoc.Type
	}

	cursor, err := c.refDocumentCollection.Find(c.ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(c.ctx) {
		var chemistryRes models.ReferenceDocument
		err := cursor.Decode(&chemistryRes)
		if err != nil {
			return nil, err
		}
		res = append(res, &chemistryRes)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(c.ctx)

	if len(res) == 0 {
		return nil, errors.New("documents not found")
	}
	return res, err
}
