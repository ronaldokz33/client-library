package interview

import (
	"net/http"
	"fmt"
	"bytes"
	"encoding/json"
	"errors"
	uuid "github.com/satori/go.uuid"
)

//BaseRequest is a helper to handle with POST request, to put all the data inside a "data" object
type baseRequest struct {
	Data *AccountData `json:"data"`
}

//Create is a method used to create a new account. Is expected a struct AccountData as a parameter
// and the return will be the created object or an error.
func (c *Client) Create(account *AccountData) (*AccountData, error)  {

	err := account.Validate()

	if err != nil {
		return nil, err
	}

	request := baseRequest{Data: account}

	plainText, _ := json.Marshal(&request)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/organisation/accounts", c.baseURL), bytes.NewBuffer(plainText))

	if err != nil {
		return nil, err
	}

	res := AccountData{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

//Fetch is a method used to find an account by id. Is expected a  valid uuid as a parameter
// and the return will be the account found object or an error.
func (c *Client) Fetch(id string) (*AccountData, error)  {
	_, err := uuid.FromString(id)

    if err != nil {
		return nil, errors.New("ID isn't a valid uuid")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/organisation/accounts/%s", c.baseURL, id), nil)

	if err != nil {
		return nil, err
	}

	res := AccountData{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

//Delete is a method used to delete an account by id. Is expected a valid uuid and the version of the account as a parameter
// and the return will be a boolean.
func (c *Client) Delete(id string, version int) (bool, error)  {
	_, err := uuid.FromString(id)

    if err != nil {
		return false, errors.New("Id isn't a valid uuid")
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/organisation/accounts/%s?version=%d", c.baseURL, id, version), nil)

	if err != nil {
		return false, err
	}

	res := AccountData{}

	if err := c.sendRequest(req, &res); err != nil {
		return false, err
	}

	return true, nil
}