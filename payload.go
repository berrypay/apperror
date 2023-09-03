/*
 * Project: Application Error Library
 * Filename: /payload.go
 * Created Date: Thursday April 13th 2023 13:05:20 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Tuesday April 18th 2023 16:30:01 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package apperrors

import "fmt"

const (
	ERR_PAYLOAD_MISSING = "ERR_PAYLOAD_MISSING"
	ERR_PAYLOAD_INVALID = "ERR_PAYLOAD_INVALID"
)

type PayloadError struct {
	ErrorCode   string `json:"errorCode"`
	Message     string `json:"message"`
	SourceError error  `json:"sourceError"`
}

func (e *PayloadError) Error() string {
	return fmt.Sprintf("[%s] %s", e.ErrorCode, e.Message)
}

func NewPayloadMissingError(message string, err error) *PayloadError {
	return &PayloadError{
		ErrorCode:   ERR_PAYLOAD_MISSING,
		Message:     message,
		SourceError: err,
	}
}

func NewPayloadInvalidError(message string, err error) *PayloadError {
	return &PayloadError{
		ErrorCode:   ERR_PAYLOAD_INVALID,
		Message:     message,
		SourceError: err,
	}
}
