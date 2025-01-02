package task

import (
	"context"
	"errors"
	"testing"
	"time"

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

func TestCreateTask(t *testing.T) {
	mockRepo := new(MockStorage)
	service := &Service{Repo: mockRepo}

	ctx := context.Background()
	dto := &CreateTaskDTO{
		Name:        "Test Task",
		Description: "Test Description",
		UserID:      uuid.New(),
		IsCompleted: false,
	}

	expectedTask := &Task{
		ID:          uuid.New(),
		Name:        dto.Name,
		Description: dto.Description,
		UserID:      dto.UserID,
		IsCompleted: dto.IsCompleted,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("CreateTask", ctx, mock.AnythingOfType("*Task")).Return(expectedTask, nil)

	task, err := service.CreateTask(ctx, dto)

	require.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, expectedTask, task)

	mockRepo.AssertCalled(t, "CreateTask", ctx, mock.AnythingOfType("*Task"))

	mockRepo.On("CreateTask", ctx, mock.AnythingOfType("*Task")).Return(nil, errors.New("repository error"))

	task, err = service.CreateTask(ctx, dto)

	require.Error(t, err)
	assert.Nil(t, task)
	assert.EqualError(t, err, "repository error")

	mockRepo.AssertCalled(t, "CreateTask", ctx, mock.AnythingOfType("*Task"))
}
