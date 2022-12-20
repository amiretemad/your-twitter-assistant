package repository

import (
	"testing"
	"tweak_twitter/pkg/model"
	"tweak_twitter/pkg/repository"
)

func TestSaveListReturnsProperFileError(t *testing.T) {
	fileRepository := repository.UserFileRepository{}

	var users []model.User
	users = append(users, model.User{})

	_, err := fileRepository.SaveList(users)
	if err.Error() != repository.ErrorFileObjIsEmpty {
		t.Errorf("Error actual = %v , got %v", err.Error(), repository.ErrorFileObjIsEmpty)
	}
}
