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

type MockResultRepository struct {
	mock.Mock
}

func (m *MockResultRepository) Save(ctx context.Context, result *domain.Result) (*domain.Result, error) {
	args := m.Called(ctx, result)
	return args.Get(0).(*domain.Result), args.Error(1)
}

func (m *MockResultRepository) FindByID(ctx context.Context, id string) (*domain.Result, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Result), args.Error(1)
}

func (m *MockResultRepository) List(ctx context.Context, limit, offset int64) ([]*domain.Result, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]*domain.Result), args.Error(1)
}

func (m *MockResultRepository) Update(ctx context.Context, result *domain.Result) error {
	args := m.Called(ctx, result)
	return args.Error(0)
}

func (m *MockResultRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestResultService_CreateResult(t *testing.T) {
	mockRepo := new(MockResultRepository)
	service := services.NewResultService(mockRepo)

	ctx := context.Background()

	t.Run("Successful Creation", func(t *testing.T) {
		expectedResult := &domain.Result{
			ID:     "1",
			Status: "AVAILABLE",
		}
		mockRepo.On("Save", ctx, mock.AnythingOfType("*domain.Result")).Return(expectedResult, nil)

		createdResult, err := service.CreateResult(ctx, "1", "AVAILABLE")

		assert.NoError(t, err)
		assert.NotNil(t, createdResult)
		assert.Equal(t, "1", createdResult.ID)
		assert.Equal(t, "AVAILABLE", createdResult.Status)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Empty ID", func(t *testing.T) {
		_, err := service.CreateResult(ctx, "", "AVAILABLE")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrDomain)

		mockRepo.AssertNotCalled(t, "Save")
	})

	t.Run("Invalid Status", func(t *testing.T) {
		_, err := service.CreateResult(ctx, "1", "INVALID")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrDomain)

		mockRepo.AssertNotCalled(t, "Save")
	})

	t.Run("Duplicate mail", func(t *testing.T) {
		mockRepo.On("Save", ctx, mock.AnythingOfType("*domain.Result")).Return((*domain.Result)(nil), db.ErrUnique)

		_, err := service.CreateResult(ctx, "1", "AVAILABLE")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Database Connection", func(t *testing.T) {
		mockRepo.On("Save", ctx, mock.AnythingOfType("*domain.Result")).Return((*domain.Result)(nil), db.ErrConnection)

		_, err := service.CreateResult(ctx, "1", "AVAILABLE")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})
}

func TestResultService_GetResult(t *testing.T) {
	mockRepo := new(MockResultRepository)
	service := services.NewResultService(mockRepo)

	ctx := context.Background()

	t.Run("Successful Get", func(t *testing.T) {
		expectedResult := &domain.Result{
			ID:     "1",
			Status: "AVAILABLE",
		}

		mockRepo.On("FindByID", ctx, "1").Return(expectedResult, nil)

		result, err := service.GetResult(ctx, "1")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "1", result.ID)
		assert.Equal(t, "AVAILABLE", result.Status)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Result Not Found", func(t *testing.T) {
		mockRepo.On("FindByID", ctx, "2").Return((*domain.Result)(nil), db.ErrNotFound)

		result, err := service.GetResult(ctx, "2")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Empty ID", func(t *testing.T) {
		_, err := service.GetResult(ctx, "")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrDomain)

		mockRepo.AssertNotCalled(t, "FindByID")
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Not Found", func(t *testing.T) {
		mockRepo.On("FindByID", ctx, "1").Return((*domain.Result)(nil), db.ErrNotFound)

		result, err := service.GetResult(ctx, "1")

		assert.Nil(t, result)
		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Database Connection", func(t *testing.T) {
		mockRepo.On("FindByID", ctx, "2").Return((*domain.Result)(nil), db.ErrConnection)

		result, err := service.GetResult(ctx, "2")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})
}

func TestResultService_ListResult(t *testing.T) {
	mockRepo := new(MockResultRepository)
	service := services.NewResultService(mockRepo)

	ctx := context.Background()

	t.Run("Successful List", func(t *testing.T) {
		expectedResults := []*domain.Result{
			{ID: "1", Status: "AVAILABLE"},
			{ID: "2", Status: "NOT_AVAILABLE"},
		}

		mockRepo.On("List", ctx, int64(10), int64(0)).Return(expectedResults, nil)

		results, err := service.ListResult(ctx, 10, 0)

		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.Len(t, results, 2)
		assert.Equal(t, "1", results[0].ID)
		assert.Equal(t, "AVAILABLE", results[0].Status)
		assert.Equal(t, "2", results[1].ID)
		assert.Equal(t, "NOT_AVAILABLE", results[1].Status)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Empty List", func(t *testing.T) {
		mockRepo.On("List", ctx, int64(10), int64(0)).Return([]*domain.Result{}, nil)

		results, err := service.ListResult(ctx, 10, 0)

		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.Empty(t, results)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Database Connection", func(t *testing.T) {
		mockRepo.On("List", ctx, int64(10), int64(0)).Return(([]*domain.Result{}), db.ErrConnection)

		results, err := service.ListResult(ctx, 10, 0)

		assert.Error(t, err)
		assert.Nil(t, results)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})
}

func TestResultService_UpdateResult(t *testing.T) {
	mockRepo := new(MockResultRepository)
	service := services.NewResultService(mockRepo)

	ctx := context.Background()

	t.Run("Successful Update", func(t *testing.T) {
		existingResult := &domain.Result{ID: "1", Status: "AVAILABLE"}

		mockRepo.On("FindByID", ctx, "1").Return(existingResult, nil)
		mockRepo.On("Update", ctx, mock.AnythingOfType("*domain.Result")).Return(nil)

		err := service.UpdateResult(ctx, "1", "NOT_AVAILABLE")

		assert.NoError(t, err)
		assert.Equal(t, "NOT_AVAILABLE", existingResult.Status)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Empty ID", func(t *testing.T) {
		err := service.UpdateResult(ctx, "", "AVAILABLE")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrDomain)

		mockRepo.AssertNotCalled(t, "FindByID")
		mockRepo.AssertNotCalled(t, "Update")
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Invalid Status", func(t *testing.T) {
		existingResult := &domain.Result{ID: "1", Status: "AVAILABLE"}

		mockRepo.On("FindByID", ctx, "1").Return(existingResult, nil)

		err := service.UpdateResult(ctx, "1", "INVALID")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrDomain)

		mockRepo.AssertNotCalled(t, "Update")
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Result Not Found", func(t *testing.T) {
		mockRepo.On("FindByID", ctx, "1").Return((*domain.Result)(nil), db.ErrNotFound)

		err := service.UpdateResult(ctx, "1", "NOT_AVAILABLE")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertNotCalled(t, "Update")
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Database Connection", func(t *testing.T) {
		existingResult := &domain.Result{ID: "1", Status: "AVAILABLE"}

		mockRepo.On("FindByID", ctx, "1").Return(existingResult, nil)
		mockRepo.On("Update", ctx, mock.AnythingOfType("*domain.Result")).Return(db.ErrConnection)

		err := service.UpdateResult(ctx, "1", "NOT_AVAILABLE")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})
}

func TestResultService_DeleteResult(t *testing.T) {
	mockRepo := new(MockResultRepository)
	service := services.NewResultService(mockRepo)

	ctx := context.Background()

	t.Run("Successful Deletion", func(t *testing.T) {
		mockRepo.On("Delete", ctx, "1").Return(nil)

		err := service.DeleteResult(ctx, "1")

		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Empty ID", func(t *testing.T) {
		err := service.DeleteResult(ctx, "")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrDomain)

		mockRepo.AssertNotCalled(t, "Delete")
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Result Not Found", func(t *testing.T) {
		mockRepo.On("Delete", ctx, "2").Return(db.ErrNotFound)

		err := service.DeleteResult(ctx, "2")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})

	t.Run("Database Connection", func(t *testing.T) {
		mockRepo.On("Delete", ctx, "2").Return(db.ErrConnection)

		err := service.DeleteResult(ctx, "2")

		assert.Error(t, err)
		assert.ErrorIs(t, err, services.ErrRepository)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil
	})
}
