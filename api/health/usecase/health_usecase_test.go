package usecase_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/insomniacoder/prometheus-converter/api/domain"
	"github.com/insomniacoder/prometheus-converter/api/health/usecase"

	"testing"
)

var (
	mockHealthInfo    *domain.HealthInfo
	mockPromethesText string
)

type MockHealthRepository struct {
	mock.Mock
}

func (m *MockHealthRepository) GetHealthInfo(c context.Context) (info *domain.HealthInfo, err error) {
	ret := m.Called()
	return ret.Get(0).(*domain.HealthInfo), nil
}

func TestMain(m *testing.M) {

	jsonFile, err := os.Open("../../../static/example-response.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &mockHealthInfo)

	mockPromethesText = "gateway_up 1\nface_comparison_up 1\nthai_id_up 1\nantispoofing_up 1\n"

	code := m.Run()
	os.Exit(code)
}

func TestGetPrometheusHealthInfoAllUpShouldProduceAllUp(t *testing.T) {

	mr := new(MockHealthRepository)
	mr.On("GetHealthInfo", mock.Anything).Return(mockHealthInfo, nil).Once()
	uc := usecase.NewHealthUsecase(mr)

	returnPrometheusText, err := uc.GetPrometheusHealthInfo(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, mockPromethesText, returnPrometheusText)
}
