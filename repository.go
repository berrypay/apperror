/*
 * Project: Application Error Library
 * Filename: /repository.go
 * Created Date: Thursday April 13th 2023 15:36:22 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Tuesday April 18th 2023 16:29:04 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package apperrors

import "fmt"

const (
	ERR_REPOSITORY_FAILED_CONNECTION = "ERR_REPOSITORY_FAILED_CONNECTION"
	ERR_REPOSITORY_FAILED_QUERY      = "ERR_REPOSITORY_FAILED_QUERY"
	ERR_REPOSITORY_FAILED_UPDATE     = "ERR_REPOSITORY_FAILED_UPDATE"
	ERR_REPOSITORY_FAILED_DELETE     = "ERR_REPOSITORY_FAILED_DELETE"
)

type RepositoryError struct {
	ErrorCode   string `json:"errorCode"`
	Message     string `json:"message"`
	SourceError error  `json:"sourceError"`
}

func (e *RepositoryError) Error() string {
	return fmt.Sprintf("[%s] %s", e.ErrorCode, e.Message)
}

func NewRepositoryFailedConnectionError(message string, err error) *RepositoryError {
	return &RepositoryError{
		ErrorCode:   ERR_REPOSITORY_FAILED_CONNECTION,
		Message:     message,
		SourceError: err,
	}
}

func NewRepositoryFailedQueryError(message string, err error) *RepositoryError {
	return &RepositoryError{
		ErrorCode:   ERR_REPOSITORY_FAILED_QUERY,
		Message:     message,
		SourceError: err,
	}
}

func NewRepositoryFailedUpdateError(message string, err error) *RepositoryError {
	return &RepositoryError{
		ErrorCode:   ERR_REPOSITORY_FAILED_UPDATE,
		Message:     message,
		SourceError: err,
	}
}

func NewRepositoryFailedDeleteError(message string, err error) *RepositoryError {
	return &RepositoryError{
		ErrorCode:   ERR_REPOSITORY_FAILED_DELETE,
		Message:     message,
		SourceError: err,
	}
}
