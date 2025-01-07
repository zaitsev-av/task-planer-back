package task

import (
	"context"
	"testing"
	"time"

	"task-planer-back/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Мок для интерфейса Storage
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) CreateTask(ctx context.Context, task *Task) (*Task, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(*Task), args.Error(1)
}

func (m *MockStorage) GetTask(ctx context.Context, id string) (Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Task), args.Error(1)
}

func (m *MockStorage) DeleteTask(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockStorage) RenameTask(ctx context.Context, id string, name string) (*ChangeNameDTO, error) {
	args := m.Called(ctx, id, name)
	return args.Get(0).(*ChangeNameDTO), args.Error(1)
}

func (m *MockStorage) ChangeDescriptionTask(ctx context.Context, id string, description string) error {
	args := m.Called(ctx, id, description)
	return args.Error(0)
}

func TestService_CreateTask(t *testing.T) {
	mockRepo := new(MockStorage)
	service := NewServices(mockRepo)

	ctx := context.Background()
	dto := &CreateTaskDTO{
		Name:        "Test Task",
		Description: "Description",
		UserID:      uuid.New(),
		IsCompleted: false,
	}

	expectedTask := &Task{
		ID:          uuid.New(),
		Name:        dto.Name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: dto.Description,
		Priority:    models.LowPriority,
		UserID:      dto.UserID,
		IsCompleted: dto.IsCompleted,
	}

	mockRepo.On("CreateTask", ctx, mock.MatchedBy(func(task *Task) bool {
		return task.Name == dto.Name &&
			task.Description == dto.Description &&
			task.UserID == dto.UserID &&
			task.IsCompleted == dto.IsCompleted
	})).Return(expectedTask, nil)

	task, err := service.CreateTask(ctx, dto)

	require.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, expectedTask.Name, task.Name)

	mockRepo.AssertExpectations(t)
}

func TestService_DeleteTask(t *testing.T) {
	mockRepo := new(MockStorage)
	service := NewServices(mockRepo)

	ctx := context.Background()
	id := "task-id"

	mockRepo.On("DeleteTask", ctx, id).Return(nil)

	err := service.DeleteTask(ctx, id)

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "DeleteTask", ctx, id)
}
