package uploadmodel

import (
	"crud-go/common"
	"errors"
)

const EntityName = "Upload"

type Upload struct {
	common.SQLModel `json:",inline"`
	common.Image    `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

var (
	ErrFileTooLarge = common.NewCustomErrorResponse(
		errors.New("file too large"),
		"file too large",
		"Error_File_Too_Large",
	)
)

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomErrorResponse(
		err,
		"file is not an image",
		"Error_File_Is_Not_Image",
	)
}
func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomErrorResponse(
		err,
		"cannot save file",
		"Error_Cannot_Save_File",
	)

}
