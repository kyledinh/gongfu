package sys

const (
	ACCESS_OK = "ACCESS_OK"
	ACCESS_NO = "ACCESS_NO"
	USERID    = "USERID"
	BLANK     = ""

	COOKIE_ACCOUNTID  = "COOKIE_ACCOUNTID"
	COOKIE_USERID     = "COOKIE_USERID"
	COOKIE_APP_ACCESS = "COOKIE_APP_ACCESS"
	COOKIE_EMAIL      = "COOKIE_EMAIL"

	ERR_READ_BODY      = "Error reading the data from the Request body"
	ERR_DB_DELETE      = "Error deleting record in the Database"
	ERR_DB_INSERT      = "Error inserting record into the Database"
	ERR_DB_NO_MATCH    = "Error no record match in the Database"
	ERR_DB_UPDATE      = "Error updating record in the Database"
	ERR_SCAN_ROW       = "Error scanning DB record"
	ERR_UNMARSHAL_BODY = "Error while unmarshaling the data from the Request body"

	SUCCESS = "success"
	FAIL    = "fail"

	// HTTP STATUS CODES USED
	CODE_200 = 200 // OK
	CODE_201 = 201 // Created
	CODE_202 = 202 // Accepted
	CODE_204 = 204 // Success delete
	CODE_301 = 301 // Moved Permanently
	CODE_304 = 304 // Not Modified - Cached
	CODE_307 = 307 // Temporary Redirect
	CODE_400 = 400 // Bad Request
	CODE_401 = 401 // Unauthorized
	CODE_403 = 403 // Forbidden
	CODE_404 = 404 // Not Found
	CODE_408 = 408 // Request Timeout
	CODE_410 = 410 // Gone
	CODE_422 = 422 // Unprocessable Entity - Validation Error
	CODE_429 = 429 // Too Many Requests - request rejected due to rate limiting
	CODE_500 = 500 // Internal Server Error
)

var JSON_NOACCESS = map[string]interface{}{"status": "fail", "message": "You do not have access. You need to log in."}
