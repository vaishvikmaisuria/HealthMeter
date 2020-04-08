package db

import (
	"log"

	"github.com/vaishvikmaisuria/CoronaVision/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dataCollection = "CoronaData"
	countryTable   = "CoronaTable"
)

// DataInsert <function>
// is used to insert an array of coronaVirus data to DB
func (s Service) DataInsert(data []models.CoronaStates) error {
	c := s.Collection(dataCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	if err := c.Drop(ctx); err != nil {
		log.Fatal(err)
	}

	c = s.Collection(dataCollection)
	ctx, cncl = s.CTX(5)
	defer cncl()

	// we cast it to array of interfaces because mongo driver asks so
	iData := []interface{}{} // []models.CoronaStates
	for _, n := range data {
		iData = append(iData, n)
	}

	_, err := c.InsertMany(
		ctx,
		iData,
		s.InsertManyOptionsOrdered(),
	)

	return err
}

// TableInsert <function>
// is used to insert an array of coronaVirus data to DB
func (s Service) TableInsert(data []models.CountryTable) error {
	c := s.Collection(countryTable)
	ctx, cncl := s.CTX(5)
	defer cncl()

	if err := c.Drop(ctx); err != nil {
		log.Fatal(err)
	}

	c = s.Collection(countryTable)
	ctx, cncl = s.CTX(5)
	defer cncl()

	// we cast it to array of interfaces because mongo driver asks so
	iData := []interface{}{} // []models.CountryTable
	for _, n := range data {
		iData = append(iData, n)
	}

	_, err := c.InsertMany(
		ctx,
		iData,
		s.InsertManyOptionsOrdered(),
	)

	return err
}

// DeleteOldData <function>
// is used to delete all array of coronaVirus data in DB
func (s Service) DeleteOldData() {
	c := s.Collection(dataCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	err := c.Drop(ctx)

	if err != nil {
		log.Fatal(err)
	}
}

// DataGet <function>
// is used to get data with given count
func (s Service) DataGet(count int64) ([]models.CoronaStates, error) {
	c := s.Collection(dataCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	var data []models.CoronaStates
	options := options.Find()
	cur, err := c.Find(
		ctx,
		bson.D{},
		options.SetLimit(count), // limit to <count>
	)
	if err != nil {
		return []models.CoronaStates{}, err
	}

	for cur.Next(ctx) {
		var n models.CoronaStates
		err = cur.Decode(&n)
		if err != nil {
			return []models.CoronaStates{}, err
		}

		data = append(data, n)
	}

	if err = cur.Err(); err != nil {
		return []models.CoronaStates{}, err
	}

	cur.Close(ctx)

	return data, nil
}

// CTableGet <function>
// is used to get data with given count
func (s Service) CTableGet() ([]models.CountryTable, error) {
	c := s.Collection(countryTable)
	ctx, cncl := s.CTX(5)
	defer cncl()

	var table []models.CountryTable
	options := options.Find()
	cur, err := c.Find(
		ctx,
		bson.D{},
		options.SetLimit(2), // limit to <count>
	)
	if err != nil {
		return []models.CountryTable{}, err
	}

	for cur.Next(ctx) {
		var n models.CountryTable
		err = cur.Decode(&n)
		if err != nil {
			return []models.CountryTable{}, err
		}

		table = append(table, n)
	}

	if err = cur.Err(); err != nil {
		return []models.CountryTable{}, err
	}

	cur.Close(ctx)

	return table, nil
}
