package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// define enum for type material, type spectrum

type Chemistry struct {
	ID           primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	TypeChemical string             `json:"typeChemical,omitempty" bson:"type_chemical,omitempty"`
	GroupName    string             `json:"groupName,omitempty" bson:"group_name,omitempty"`
	TypeSpectrum string             `json:"typeSpectrum,omitempty" bson:"type_spectrum,omitempty"`
	Chemical     string             `json:"chemical,omitempty" bson:"chemical,omitempty"`
	VideoUrl     string             `json:"videoUrl,omitempty" bson:"video_url,omitempty"`
}

type ReferenceDocument struct {
	ID   primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
	Type string             `json:"type,omitempty" bson:"type,omitempty"`
	Url  string             `json:"url,omitempty,omitempty" bson:"url,omitempty"`
}

// write func post, put, delete to use in backend
