package repositories

import "errors"

const (
	ErrPrepareStatement = "failed to prepare SQL statement"
	ErrExecuteStatement = "failed to execute statement"
	ErrExecuteQuery     = "failed to execute query"
	ErrScanData         = "failed to scan data"
	ErrBeginTransaction = "failed to begin transaction"
	ErrRollback         = "failed to rollback transaction"
	ErrCommit           = "failed to commit transaction"
	ErrRetrieveRows     = "failed to retrieve rows affected"
)

// Repository Errors
var (
	ErrUserProfileNotFound = errors.New("user profile not found")
)
