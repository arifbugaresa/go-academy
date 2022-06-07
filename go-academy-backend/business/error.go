package business

import "errors"

var (
	// ErrInvalidBody Error when data given is not valid on update or insert
	ErrInvalidBody = errors.New("given body cannot be parsed to struct")
	ErrGetDataFromDB = errors.New("error get data from db")
	ErrDataNotFound = errors.New("error data not found")
	ErrInsertData = errors.New("error insert data")
)