package buildings_test

import (
	"context"
	"os"
	buildings "rentRoom/businesses/buildings"
	buildingMock "rentRoom/businesses/buildings/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	buildingRepository buildingMock.Repository
	buildingUseCase    buildings.Usecase
)

func setup() {
	buildingUseCase = buildings.NewBuildingsUsecase(&buildingRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestStore(t *testing.T) {
	t.Run("test 1: valid test", func(t *testing.T) {
		var domain = buildings.Domain{
			BuildingName:    "building1",
			BuildingAddr: "surabaya",
		}
		buildingRepository.On("Store", mock.Anything, mock.Anything).Return(domain, nil).Once()
		result, err := buildingUseCase.Store(context.Background(), &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain.BuildingName, result.BuildingName)
		assert.Equal(t, domain.BuildingAddr, result.BuildingAddr)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("test 1: valid test", func(t *testing.T) {
		var domain = []buildings.Domain{
			{
				ID:      1,
				BuildingName:    "building1",
				BuildingAddr: "surabaya",
			},
			{
				ID:      2,
				BuildingName:    "building2",
				BuildingAddr: "malang",
			},
		}
		buildingRepository.On("Find", context.Background(),mock.AnythingOfType("string")).Return(domain, nil).Once()
		result, err := buildingUseCase.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, domain[0].ID, result[0].ID)
		assert.Equal(t, domain[0].BuildingName, result[0].BuildingName)
		assert.Equal(t, domain[0].BuildingAddr, result[0].BuildingAddr)
	})
}