package db_errors

import "errors"

var (
	ErrEmptyValue    = errors.New("empty value")
	ErrGetRecord     = errors.New("error getting student")
	ErrScanRow       = errors.New("error scanning row")
	ErrIteration     = errors.New("error during row iteration")
	ErrAddNewRow     = errors.New("error adding student into DB")
	ErrUpdateRow     = errors.New("error updating student")
	ErrDeleteRow     = errors.New("error deleting student")
	ErrTruncateTable = errors.New("error truncating students table")
	ErrCreateTable   = errors.New("error creating students table")
	ErrDropTable     = errors.New("error dropping students table")
	ErrFindUser      = errors.New("login or password incorrect")
)
