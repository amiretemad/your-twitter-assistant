package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"tweak_twitter/pkg/model"
)

const ErrorFileObjIsEmpty = "file object is empty"

type UserFileRepository struct {
	File *os.File
}

func (u *UserFileRepository) SaveList(user []model.User) (bool, error) {

	if u.File == nil {
		return false, errors.New("file object is empty")
	}

	marshal, err := json.Marshal(user)
	if err != nil {
		return false, err
	}

	_, err = u.File.Write(marshal)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
