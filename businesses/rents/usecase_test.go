package rents_test

import (
	"context"
	"os"
	rents "rentRoom/businesses/rents"
	rentMock "rentRoom/businesses/rents/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	rentRepository rentMock.Repository
	rentUseCase    rents.Usecase
)

func setup() {
	rentUseCase = rents.NewRentsUsecase(&rentRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestStore(t *testing.T) {
	t.Run("test 1: valid test", func(t *testing.T) {
		var domain = rents.Domain{
			UserId:    1,
			RoomId: 1,
			JumlahBayar: 100000,
			TanggalPinjam: "2021-08-17",
			StatusPinjam: false,
		}
		rentRepository.On("GetById",mock.Anything,mock.AnythingOfType("int")).Return(domain, nil).Once()
		rentRepository.On("GetRoomById",mock.Anything,mock.AnythingOfType("int")).Return(domain, nil).Once()
		rentRepository.On("Store",mock.Anything,mock.Anything).Return(domain, nil).Once()
		result, err := rentUseCase.Store(context.Background(), &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain.UserId, result.UserId)
		assert.Equal(t, domain.RoomId, result.RoomId)
		assert.Equal(t, domain.JumlahBayar, result.JumlahBayar)
		assert.Equal(t, domain.StatusPinjam, result.StatusPinjam)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("test 1: valid test", func(t *testing.T) {
		var domain = []rents.Domain{
			{
				ID:      1,
				UserId:    1,
				RoomId: 1,
				JumlahBayar: 100000,
				StatusPinjam: false,
			},
			{
				ID:      2,
				UserId:    1,
				RoomId: 2,
				JumlahBayar: 200000,
				StatusPinjam: true,
			},
		}
		rentRepository.On("Find", context.Background(),mock.AnythingOfType("string")).Return(domain, nil).Once()
		result, err := rentUseCase.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, domain[0].ID, result[0].ID)
		assert.Equal(t, domain[0].UserId, result[0].UserId)
		assert.Equal(t, domain[0].RoomId, result[0].RoomId)
		assert.Equal(t, domain[0].JumlahBayar, result[0].JumlahBayar)
		assert.Equal(t, domain[0].StatusPinjam, result[0].StatusPinjam)
	})
}

func TestUpdate(t *testing.T) {
	var domain = rents.Domain{
		ID:       1,
		StatusPinjam: true,
	}
	t.Run("test 1: valid test", func(t *testing.T) {
		rentRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("boolean")).Return(domain, nil).Once()
		result, err := rentUseCase.Update(context.Background(), &domain)
		// result, err = rentMongoRepository.LoginLog(context.Background(), 1)

		assert.Equal(t, domain.ID, result.ID)
		assert.Nil(t, err)
	})

	t.Run("test 2: invalid Input", func(t *testing.T) {
		rentRepository.On("UpdateTrainer",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		result, _ := rentUseCase.Update(context.Background(), &domain)

		assert.Equal(t, rents.Domain{}, result)
	})

	t.Run("test 2: invalid ID", func(t *testing.T) {
		rentRepository.On("UpdateTrainer",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		result, _ := rentUseCase.Update(context.Background(), &domain)

		assert.Equal(t, rents.Domain{}, result)
	})
}