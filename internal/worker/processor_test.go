package worker

import (
	"encoding/csv"
	"errors"
	"excel-processor/internal/entity"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStudentRepo struct {
	mock.Mock
}

func (m *MockStudentRepo) Create(student entity.Student) error  {
	args := m.Called(student)

	return args.Error(0)
}

func TestProcessor_Success(t *testing.T)  {
	filename := "test_data.csv"
	file, _ := os.Create(filename)
	writer := csv.NewWriter(file)

	writer.Write([]string{"Test User 1", "test1@mail.com"})
	writer.Write([]string{"Test User 2", "test2@mail.com"})
	writer.Flush()
	file.Close()
	defer os.Remove(filename)

	mockRepo := new(MockStudentRepo)

	mockRepo.On("Create", mock.Anything).Return(nil).Times(2)

	processor := NewProcessor(mockRepo)

	processor.Start(filename)

	assert.Equal(t, 2, processor.SuccessCount)
	assert.Equal(t, 0, processor.FailCount)

	mockRepo.AssertExpectations(t)
}

func TestProcessor_PartialFailure(t *testing.T)  {
	filename := "test_partial_error.csv"
	file, _ := os.Create(filename)
	writer := csv.NewWriter(file)

	writer.Write([]string{"User Aman 1", "aman1@mail.com"})
	writer.Write([]string{"User Error", "error@mail.com"})
	writer.Write([]string{"User Aman 2", "aman2@mail.com"})

	writer.Flush()
	file.Close()
	defer os.Remove(filename)

	mockRepo := new(MockStudentRepo)

	// Because Workers run concurrently (randomly), we don't know the order of the calls
	// So we match based on the DATA, not the order
	mockRepo.On("Create", mock.MatchedBy(func (s entity.Student) bool  {
		return s.Email == "error@mail.com"
	})).Return(errors.New("DB connection failed")).Times(1)

	mockRepo.On("Create", mock.MatchedBy(func (s entity.Student) bool  {
		return s.Email != "error@mail.com"
	})).Return(nil).Times(2)

	// run processor
	processor := NewProcessor(mockRepo)
	processor.Start(filename)

	assert.Equal(t, 1, processor.FailCount, "Harusnya ada 1 data gagal")
	assert.Equal(t, 2, processor.SuccessCount, "Harusnya ada 2 data sukses")

	mockRepo.AssertExpectations(t)
}