package interview

import (
	"testing"
	"github.com/stretchr/testify/assert"
	uuid "github.com/satori/go.uuid"
)

func TestCreate(t *testing.T)  {
	c := NewClient()
	account := generateAccount()

	res, err := c.Create(account)

	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.NotEqual(t, res.ID, "")
}

func TestCreateFail(t *testing.T)  {
	c := NewClient()
	account := generateAccount()

	account.ID = ""

	res, err := c.Create(account)

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T)  {
	c := NewClient()

	ID := insertAccountMock()
	assert.NotEqual(t, ID, "")

	res, err := c.Delete(ID, 0)

	assert.Nil(t, err)
	assert.Equal(t, res, true)
}

func TestDeleteFail(t *testing.T)  {
	c := NewClient()

	res, err := c.Delete("44554-545-s", 0)

	assert.NotNil(t, err)
	assert.Equal(t, res, false)
}

func TestFetch(t *testing.T)  {
	c := NewClient()
	
	ID := insertAccountMock()
	assert.NotEqual(t, ID, "")
	
	res, err := c.Fetch(ID)

	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, res.ID, ID)
}

func TestFetchFail(t *testing.T)  {
	c := NewClient()

	_, err := c.Fetch("44554-545-s")

	assert.NotNil(t, err)
}

func insertAccountMock() (string) {
	c := NewClient()
	account := generateAccount()

	res, err := c.Create(account)

	if err != nil {
		return ""
	}

	return res.ID
}

func generateAccount() (*AccountData) {
	Country := "GB"
	AccountClassification := "Personal"

	return &AccountData {
		ID: uuid.NewV4().String(),
		Type: "accounts",
		OrganisationID: uuid.NewV4().String(),
		Attributes: &AccountAttributes {
			Country: &Country,
			BaseCurrency: "GBP",
			BankID: "400302",
			BankIDCode: "GBDSC",
			AccountNumber: "10000004",
			Iban: "GB28NWBK40030212764204",
			Bic: "NWBKGB42",
			AccountClassification: &AccountClassification,
			Name: []string{"Ronaldo Lemos Junior"},
		},
	}
}