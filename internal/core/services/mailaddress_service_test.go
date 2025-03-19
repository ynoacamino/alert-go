package services_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ynoacamino/alert-go/internal/adapters/db"
	"github.com/ynoacamino/alert-go/internal/core/domain"
	"github.com/ynoacamino/alert-go/internal/core/services"
)

type MockMailAddressRepository struct {
	mock.Mock
}

func (m *MockMailAddressRepository) Save(ctx context.Context, mailAddress *domain.MailAddress) (*domain.MailAddress, error) {
	args := m.Called(ctx, mailAddress)
	return args.Get(0).(*domain.MailAddress), args.Error(1)
}

func (m *MockMailAddressRepository) FindByID(ctx context.Context, id string) (*domain.MailAddress, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.MailAddress), args.Error(1)
}

func (m *MockMailAddressRepository) List(ctx context.Context, limit, offset int64) ([]*domain.MailAddress, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]*domain.MailAddress), args.Error(1)
}

func (m *MockMailAddressRepository) Update(ctx context.Context, result *domain.MailAddress) error {
	args := m.Called(ctx, result)
	return args.Error(0)
}

func (m *MockMailAddressRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestMailAddressService_CreateMailAddress(t *testing.T) {
	mockRepo := new(MockMailAddressRepository)
	service := services.NewMailAddressService(mockRepo)

	ctx := context.Background()

	t.Run("Successful Creation", func(t *testing.T) {
		createdMailAddress := &domain.MailAddress{
			ID:      "1",
			Address: "test@test.com",
			Active:  true,
		}

		mockRepo.On("Save", ctx, mock.AnythingOfType("*domain.MailAddress")).Return(createdMailAddress, nil)

		mailAddress, err := service.CreateMailAddress(ctx, "1", "test@test.com", true)

		assert.NoError(t, err)
		assert.NotNil(t, mailAddress)
		assert.Equal(t, "1", mailAddress.ID)
		assert.Equal(t, "test@test.com", mailAddress.Address)
		assert.Equal(t, true, mailAddress.Active)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Empty ID", func(t *testing.T) {
		mailAddress, err := service.CreateMailAddress(ctx, "", "test@test.com", true)

		assert.Nil(t, mailAddress)
		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrDomain)

		mockRepo.AssertNotCalled(t, "Save")
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Invalid Email", func(t *testing.T) {
		mailAddress, err := service.CreateMailAddress(ctx, "1", "test", true)

		assert.Nil(t, mailAddress)
		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrDomain)

		mockRepo.AssertNotCalled(t, "Save")
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Database Connection", func(t *testing.T) {
		mockRepo.On("Save", ctx, mock.AnythingOfType("*domain.MailAddress")).Return((*domain.MailAddress)(nil), services.ErrRepository)

		mailAddress, err := service.CreateMailAddress(ctx, "1", "test@test.com", true)

		assert.Nil(t, mailAddress)
		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})
}

func TestMailAddressService_GetMailAddress(t *testing.T) {
	mockRepo := new(MockMailAddressRepository)
	service := services.NewMailAddressService(mockRepo)

	ctx := context.Background()

	t.Run("Successful Get", func(t *testing.T) {
		mailAddress := &domain.MailAddress{
			ID:      "1",
			Address: "test@test.com",
			Active:  true,
		}

		mockRepo.On("FindByID", ctx, "1").Return(mailAddress, nil)

		result, err := service.GetMailAddress(ctx, "1")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "1", result.ID)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Empty ID", func(t *testing.T) {
		_, err := service.GetMailAddress(ctx, "")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrDomain)

		mockRepo.AssertNotCalled(t, "FindByID")
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Not Found", func(t *testing.T) {
		mockRepo.On("FindByID", ctx, "2").Return((*domain.MailAddress)(nil), db.ErrNotFound)

		result, err := service.GetMailAddress(ctx, "2")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Database Connection", func(t *testing.T) {
		mockRepo.On("FindByID", ctx, "1").Return((*domain.MailAddress)(nil), db.ErrConnection)

		result, err := service.GetMailAddress(ctx, "1")

		assert.Nil(t, result)
		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})
}

func TestMailAddressService_ListMailAddresses(t *testing.T) {
	mockRepo := new(MockMailAddressRepository)
	service := services.NewMailAddressService(mockRepo)

	ctx := context.Background()

	t.Run("Successful List", func(t *testing.T) {
		expectedMailAddresses := []*domain.MailAddress{
			{ID: "1", Address: "test1@test.com", Active: true},
			{ID: "2", Address: "test2@test.com", Active: false},
		}

		mockRepo.On("List", ctx, int64(10), int64(0)).Return(expectedMailAddresses, nil)

		mailAddresses, err := service.ListMailAddresses(ctx, 10, 0)

		assert.NoError(t, err)
		assert.NotNil(t, mailAddresses)
		assert.Len(t, mailAddresses, 2)
		assert.Equal(t, "1", mailAddresses[0].ID)
		assert.Equal(t, "test1@test.com", mailAddresses[0].Address)
		assert.Equal(t, true, mailAddresses[0].Active)
		assert.Equal(t, "2", mailAddresses[1].ID)
		assert.Equal(t, "test2@test.com", mailAddresses[1].Address)
		assert.Equal(t, false, mailAddresses[1].Active)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Empty List", func(t *testing.T) {
		mockRepo.On("List", ctx, int64(10), int64(0)).Return([]*domain.MailAddress{}, nil)

		mailAddresses, err := service.ListMailAddresses(ctx, 10, 0)

		assert.NoError(t, err)
		assert.NotNil(t, mailAddresses)
		assert.Empty(t, mailAddresses)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Database Connection Error", func(t *testing.T) {
		mockRepo.On("List", ctx, int64(10), int64(0)).Return(([]*domain.MailAddress{}), db.ErrConnection)

		mailAddresses, err := service.ListMailAddresses(ctx, 10, 0)

		assert.Error(t, err)
		assert.Nil(t, mailAddresses)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})
}

func TestMailAddressService_UpdateMailAddress(t *testing.T) {
	mockRepo := new(MockMailAddressRepository)
	service := services.NewMailAddressService(mockRepo)

	ctx := context.Background()

	t.Run("Successful Update", func(t *testing.T) {
		existingMailAddress := &domain.MailAddress{
			ID:      "1",
			Address: "old@test.com",
			Active:  false,
		}

		mockRepo.On("FindByID", ctx, "1").Return(existingMailAddress, nil)
		mockRepo.On("Update", ctx, mock.AnythingOfType("*domain.MailAddress")).Return(nil)

		err := service.UpdateMailAddress(ctx, "1", "new@test.com", true)

		assert.NoError(t, err)
		assert.Equal(t, "new@test.com", existingMailAddress.Address)
		assert.Equal(t, true, existingMailAddress.Active)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Empty ID", func(t *testing.T) {
		err := service.UpdateMailAddress(ctx, "", "new@test.com", true)

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrDomain)

		mockRepo.AssertNotCalled(t, "FindByID")
		mockRepo.AssertNotCalled(t, "Update")
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Invalid Email", func(t *testing.T) {
		err := service.UpdateMailAddress(ctx, "1", "new", true)

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrDomain)

		mockRepo.AssertNotCalled(t, "FindByID")
		mockRepo.AssertNotCalled(t, "Update")
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Not Found", func(t *testing.T) {
		mockRepo.On("FindByID", ctx, "1").Return((*domain.MailAddress)(nil), db.ErrNotFound)

		err := service.UpdateMailAddress(ctx, "1", "test@test.com", true)

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Database Connection Error on Find", func(t *testing.T) {
		mockRepo.On("FindByID", ctx, "1").Return((*domain.MailAddress)(nil), db.ErrConnection)

		err := service.UpdateMailAddress(ctx, "1", "new@test.com", true)

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertNotCalled(t, "Update")
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Database Connection Error on Update", func(t *testing.T) {
		existingMailAddress := &domain.MailAddress{
			ID:      "1",
			Address: "old@test.com",
			Active:  false,
		}

		mockRepo.On("FindByID", ctx, "1").Return(existingMailAddress, nil)
		mockRepo.On("Update", ctx, mock.AnythingOfType("*domain.MailAddress")).Return(db.ErrConnection)

		err := service.UpdateMailAddress(ctx, "1", "new@test.com", true)

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})
}

func TestMailAddressService_DeleteMailAddress(t *testing.T) {
	mockRepo := new(MockMailAddressRepository)
	service := services.NewMailAddressService(mockRepo)

	ctx := context.Background()

	t.Run("Successful Deletion", func(t *testing.T) {
		mockRepo.On("Delete", ctx, "1").Return(nil)

		err := service.DeleteMailAddress(ctx, "1")

		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty ID", func(t *testing.T) {
		err := service.DeleteMailAddress(ctx, "")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrDomain)

		mockRepo.AssertNotCalled(t, "Delete")
	})

	t.Run("Not Found", func(t *testing.T) {
		mockRepo.On("Delete", ctx, "2").Return(db.ErrNotFound)

		err := service.DeleteMailAddress(ctx, "2")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Database Connection", func(t *testing.T) {
		mockRepo.On("Delete", ctx, "3").Return(db.ErrConnection)

		err := service.DeleteMailAddress(ctx, "3")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
	})
}
