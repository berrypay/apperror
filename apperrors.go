/*
 * Project: Application Error Library
 * Filename: /apperrors.go
 * Created Date: Thursday April 13th 2023 13:13:53 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Thursday April 13th 2023 15:30:45 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package apperrors

type AppError interface {
	Error() string
}
