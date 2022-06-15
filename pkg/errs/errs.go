package errs

var ErrFmt = "err: [%v]"

// common errors
var (
	ErrMarshalingJson        = _new("CMN001", "error marshaling json")
	ErrUnmarshalingJson      = _new("CMN002", "error unmarshaling json")
	ErrParsingTime           = _new("CMN003", "error parsing time")
	ErrNoEntityIdProvided    = _new("CMN004", "entity ID is required but none was provided")
	ErrNoDateProvided        = _new("CMN005", "error no date provided")
	ErrNoPayloadData         = _new("CMN006", "error event contains no payload data")
	ErrRepoMockAction        = _new("CMN007", "error repo mock action")
	ErrUnknownErrorType      = _new("CMN008", "error unknown error type")
	ErrInvalidDate           = _new("CMN009", "error invalid date format")
	ErrConvertingStringToInt = _new("CMN010", "error converting string to int")
	ErrGettingParam          = _new("CMN011", "error getting param")
	ErrUnmarshalingXML       = _new("CMN012", "error unmarshaling XML")
	ErrReadingBytes          = _new("CMN013", "error reading bytes from request")
)

// pkg/api
var (
	ErrResponseWriter = _new("API000", "error writing to response writer")
	ErrRequest        = _new("API001", "error request")
	ErrNewRequest     = _new("API002", "error creating new request")
)
