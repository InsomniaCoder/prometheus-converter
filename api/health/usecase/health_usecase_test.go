package usecase_test

import (
	"context"
	"encoding/json"
	"errors"
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
	mockHealthInfo *domain.HealthInfo
)

type MockHealthRepository struct {
	mock.Mock
}

func (m *MockHealthRepository) GetHealthInfo(c context.Context) (info *domain.HealthInfo, err error) {
	ret := m.Called(c)
	return ret.Get(0).(*domain.HealthInfo), ret.Error(1)
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

	code := m.Run()
	os.Exit(code)
}

func TestGetPrometheusHealthInfoAllUpShouldProduceAllUp(t *testing.T) {

	mr := new(MockHealthRepository)
	mr.On("GetHealthInfo", mock.Anything).Return(mockHealthInfo, nil).Once()

	mockPromethesText := "gateway_up 1\nface_comparison_up 1\nthai_id_up 1\nantispoofing_up 1\n"

	uc := usecase.NewHealthUsecase(mr)
	returnPrometheusText, err := uc.GetPrometheusHealthInfo(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, mockPromethesText, returnPrometheusText)
}

func TestGetPrometheusHealthInfoDownShouldProduceDown(t *testing.T) {

	mr := new(MockHealthRepository)
	mr.On("GetHealthInfo", mock.Anything).Return(&domain.HealthInfo{}, errors.New("some API error occured")).Once()

	mockPromethesText := "gateway_up 0\nface_comparison_up 0\nthai_id_up 0\nantispoofing_up 0\n"

	uc := usecase.NewHealthUsecase(mr)
	returnPrometheusText, err := uc.GetPrometheusHealthInfo(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, mockPromethesText, returnPrometheusText)
}

func TestGetPrometheusHealthInfoFaceCompareDownShouldProduceOnlyFaceCompareDown(t *testing.T) {

	mockFaceCompareDown := mockHealthInfo
	mockFaceCompareDown.FaceComparison.Instances[1].Status = ""

	mr := new(MockHealthRepository)
	mr.On("GetHealthInfo", mock.Anything).Return(mockFaceCompareDown, nil).Once()

	mockPromethesText := "gateway_up 1\nface_comparison_up 0\nthai_id_up 1\nantispoofing_up 1\n"

	uc := usecase.NewHealthUsecase(mr)
	returnPrometheusText, err := uc.GetPrometheusHealthInfo(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, mockPromethesText, returnPrometheusText)
}

func TestGetPrometheusHealthInfoThaiIDDownShouldProduceOnlyThaiIDDown(t *testing.T) {

	mockThaiIDDown := mockHealthInfo
	mockThaiIDDown.ThaiID.Instances[0].Status = ""

	mr := new(MockHealthRepository)
	mr.On("GetHealthInfo", mock.Anything).Return(mockThaiIDDown, nil).Once()

	mockPromethesText := "gateway_up 1\nface_comparison_up 1\nthai_id_up 0\nantispoofing_up 1\n"

	uc := usecase.NewHealthUsecase(mr)
	returnPrometheusText, err := uc.GetPrometheusHealthInfo(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, mockPromethesText, returnPrometheusText)
}

func TestGetPrometheusHealthInfoSpoofDownShouldProduceOnlySpoofDown(t *testing.T) {

	mockSpoofDown := mockHealthInfo
	mockSpoofDown.Antispoofing.Instances[1].Status = ""

	mr := new(MockHealthRepository)
	mr.On("GetHealthInfo", mock.Anything).Return(mockSpoofDown, nil).Once()

	mockPromethesText := "gateway_up 1\nface_comparison_up 1\nthai_id_up 1\nantispoofing_up 0\n"

	uc := usecase.NewHealthUsecase(mr)
	returnPrometheusText, err := uc.GetPrometheusHealthInfo(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, mockPromethesText, returnPrometheusText)
}
