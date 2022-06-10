package mongodb

import "go.mongodb.org/mongo-driver/bson/primitive"

type addressDocument struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Address []string           `bson:"address"`
	Postal  string             `bson:"postal"`
	City    string             `bson:"city"`
	State   string             `bson:"state"`
	Country string             `bson:"country"`
}

type companyDocument struct {
	Id          primitive.ObjectID   `bson:"_id,omitempty"`
	Name        string               `bson:"name"`
	Description string               `bson:"description,omitempty"`
	Labels      []string             `bson:"labels,omitempty"`
	Addresses   []addressDocument    `bson:"addresses,omitempty"`
	ContactIds  []primitive.ObjectID `bson:"contact_ids,omitempty"`
}

type contactDocument struct {
	Id         primitive.ObjectID   `bson:"_id,omitempty"`
	FamilyName string               `bson:"family_name"`
	GivenName  string               `bson:"give_name"`
	Labels     []string             `bson:"labels,omitempty"`
	Addresses  []addressDocument    `bson:"addresses,omitempty"`
	CompanyIds []primitive.ObjectID `bson:"company_ids,omitempty"`
}
