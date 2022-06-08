package business

import (
	"errors"
	"github.com/labstack/gommon/log"
	"strconv"
)

var (
	ErrInvalidBody   = errors.New("given body cannot be parsed to struct")
	ErrGetDataFromDB = errors.New("error get data from db")
	ErrDataNotFound  = errors.New("error data not found")
	ErrInsertData    = errors.New("error insert data")
)

func GenerateErrDataNotFound(funcName string, business string, tableName string, data int) error {
	log.Info(business + "; " + funcName + ": Pencarian id: " + strconv.Itoa(data) + " pada tabel " + tableName + " tidak ditemukan")
	return ErrDataNotFound
}
