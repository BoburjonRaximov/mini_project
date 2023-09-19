package test

import (
	"fmt"
	"net/http"
	"new_project/models"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

func TestCreateBranch(t *testing.T) {
	response := &models.Branch{}

	request := &models.CreateBranch{
		Name:    faker.FirstName(),
		Address: faker.LastName(),
	}

	resp, err := makeRequest(http.MethodPost, "/branch", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)
}
