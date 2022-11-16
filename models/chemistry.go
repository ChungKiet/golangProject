package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// define enum for type material, type spectrum

type Chemistry struct {
	ID           primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	TypeMaterial string             `json:"typeMaterial" bson:"type_material"`
	TypeSpectrum string             `json:"typeSpectrum" bson:"type_spectrum"`
	Chemical     string             `json:"chemical" bson:"chemical"`
	HTMLText     string             `json:"htmlText" bson:"html_text"`
	VideoUrl     string             `json:"videoUrl" bson:"video_url"`
}

// write func post, put, delete to use in backend
