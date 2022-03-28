package interview

import (
	"errors"
	"strings"
	uuid "github.com/satori/go.uuid"
)

// Account represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty" valid:"uuid"`
	OrganisationID string             `json:"organisation_id,omitempty" valid:"uuid"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty" validate:"required,oneof='Personal' 'Business'"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

//Validade is responsible to validate if the struct has the all mandatory fields to create a new record.
func (account *AccountData) Validate() error {
	if len(account.Attributes.Name) == 0 {
		return errors.New("Name is required")
	}

	if strings.ToLower(*account.Attributes.AccountClassification) != "personal" && strings.ToLower(*account.Attributes.AccountClassification) != "business" {
		return errors.New("Account classification should be one of [Personal Business]")
	}

	if account.Type == "" {
		return errors.New("Type is required")
	}

	if account.ID == "" {
		return errors.New("ID is required")
	}

    if _, err := uuid.FromString(account.ID); err != nil {
		return errors.New("ID isn't a valid uuid")
	}

	if account.OrganisationID == "" {
		return errors.New("OrganisationID is required")
	}

    if _, err := uuid.FromString(account.OrganisationID); err != nil {
		return errors.New("OrganisationID isn't a valid uuid")
	}

	if *account.Attributes.Country == "" {
		return errors.New("Country is required")
	}

	return nil
}