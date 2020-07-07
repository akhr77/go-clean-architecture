package mock_repository

import (
	"reflect"
	"testing"

	"github.com/akhr77/go-clean-architecture/domain/model"
	"github.com/akhr77/go-clean-architecture/usecase"
	"github.com/golang/mock/gomock"
)

func TestView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected []*model.Todo
	var err error

	mockSample := NewMockTodoRepository(ctrl)
	mockSample.EXPECT().FindAll().Return(expected, err)

	todoUsecase := usecase.NewTodoUsecase(mockSample)
	result, err := todoUsecase.View()

	if err != nil {
		t.Errorf("Actual FindAll() is nos same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual FindAll() is not same as expected")
	}
}
