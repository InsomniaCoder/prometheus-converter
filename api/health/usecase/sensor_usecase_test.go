package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/insomniacoder/iot-api/api/domain"
	"github.com/insomniacoder/iot-api/api/sensor/usecase"

	"testing"
)

var (
	mockSlice = []domain.Sensor{{ID: 1, SoilMoisture: 10.5}, {ID: 2, SoilMoisture: 15}}
)

type MockSensorRepository struct {
	mock.Mock
}

func (m *MockSensorRepository) FetchAll() (sensorSlice *[]domain.Sensor, err error) {
	ret := m.Called()
	return ret.Get(0).(*[]domain.Sensor), nil
}

func (m *MockSensorRepository) Store(sensorData *domain.Sensor) (createdSensor *domain.Sensor, err error) {
	ret := m.Called(sensorData)
	return ret.Get(0).(*domain.Sensor), nil
}

func TestFetchAll(t *testing.T) {

	ms := new(MockSensorRepository)
	ms.On("FetchAll", mock.Anything).Return(&mockSlice, nil).Once()
	uc := usecase.NewSensorUsecase(ms)

	returnedSlice, err := uc.FetchAll()

	assert.Nil(t, err)
	assert.Equal(t, &mockSlice, returnedSlice)
}

func TestStore(t *testing.T) {

	creatingSensor := domain.Sensor{SoilMoisture: 10.5}
	expectedSensor := domain.Sensor{ID: 1, SoilMoisture: 10.5}

	ms := new(MockSensorRepository)

	ms.On("Store", mock.Anything).Return(&expectedSensor, nil).Once()

	uc := usecase.NewSensorUsecase(ms)

	returnedSensor, err := uc.Store(&creatingSensor)

	assert.Nil(t, err)
	assert.Equal(t, &expectedSensor, returnedSensor)
}
