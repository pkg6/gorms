package gorms

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

const (
	ErrDBMapAddExist int = iota + 1
	ErrDBMapGetNotFind
	ErrDBMapStruct
)

type ErrMapDB struct {
	Errors  int
	Name    string
	Message string
}

func NewErrMapDB(name string, Errors int) *ErrMapDB {
	e := &ErrMapDB{
		Name:   name,
		Errors: Errors,
	}
	switch Errors {
	case ErrDBMapAddExist:
		e.Message = fmt.Sprintf("%s already exists", name)
	case ErrDBMapGetNotFind:
		e.Message = fmt.Sprintf("%s does not exist", name)
	case ErrDBMapStruct:
		e.Message = "struct parsing failed"
	default:
		e.Message = "Inexplicable error"
	}
	return e
}

func (e *ErrMapDB) Error() string {
	return e.Message
}

func ErrRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
func ErrInvalidTransaction(err error) bool {
	return errors.Is(err, gorm.ErrInvalidTransaction)
}
func ErrNotImplemented(err error) bool {
	return errors.Is(err, gorm.ErrNotImplemented)
}
func ErrMissingWhereClause(err error) bool {
	return errors.Is(err, gorm.ErrMissingWhereClause)
}
func ErrUnsupportedRelation(err error) bool {
	return errors.Is(err, gorm.ErrUnsupportedRelation)
}
func ErrPrimaryKeyRequired(err error) bool {
	return errors.Is(err, gorm.ErrPrimaryKeyRequired)
}
func ErrModelValueRequired(err error) bool {
	return errors.Is(err, gorm.ErrModelValueRequired)
}
func ErrModelAccessibleFieldsRequired(err error) bool {
	return errors.Is(err, gorm.ErrModelAccessibleFieldsRequired)
}
func ErrSubQueryRequired(err error) bool {
	return errors.Is(err, gorm.ErrSubQueryRequired)
}
func ErrInvalidData(err error) bool {
	return errors.Is(err, gorm.ErrInvalidData)
}
func ErrUnsupportedDriver(err error) bool {
	return errors.Is(err, gorm.ErrUnsupportedDriver)
}
func ErrRegistered(err error) bool {
	return errors.Is(err, gorm.ErrRegistered)
}
func ErrInvalidField(err error) bool {
	return errors.Is(err, gorm.ErrInvalidField)
}
func ErrEmptySlice(err error) bool {
	return errors.Is(err, gorm.ErrEmptySlice)
}
func ErrDryRunModeUnsupported(err error) bool {
	return errors.Is(err, gorm.ErrDryRunModeUnsupported)
}
func ErrInvalidDB(err error) bool {
	return errors.Is(err, gorm.ErrInvalidDB)
}
func ErrInvalidValue(err error) bool {
	return errors.Is(err, gorm.ErrInvalidValue)
}
func ErrInvalidValueOfLength(err error) bool {
	return errors.Is(err, gorm.ErrInvalidValueOfLength)
}
func ErrPreloadNotAllowed(err error) bool {
	return errors.Is(err, gorm.ErrPreloadNotAllowed)
}
func ErrDuplicatedKey(err error) bool {
	return errors.Is(err, gorm.ErrDuplicatedKey)
}
func ErrForeignKeyViolated(err error) bool {
	return errors.Is(err, gorm.ErrForeignKeyViolated)
}
