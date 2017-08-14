package boom

/*
Boom :
*/
type Boom struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Detail  interface{} `json:"detail,omitempty"`
}

func (b *Boom) Error() string {
	return b.Message
}

/*
BadRequest :
*/
func BadRequest(msg string) *Boom {
	return &Boom{
		Code:    400,
		Message: msg,
	}
}

/*
Unauthorized :
*/
func Unauthorized(msg string) *Boom {
	return &Boom{
		Code:    401,
		Message: msg,
	}
}

/*
PaymentRequired :
*/
func PaymentRequired(msg string) *Boom {
	return &Boom{
		Code:    402,
		Message: msg,
	}
}

/*
Forbidden :
*/
func Forbidden(msg string) *Boom {
	return &Boom{
		Code:    403,
		Message: msg,
	}
}

/*
NotFound :
*/
func NotFound(msg string) *Boom {
	return &Boom{
		Code:    404,
		Message: msg,
	}
}

/*
MethodNotAllowed :
*/
func MethodNotAllowed(msg string) *Boom {
	return &Boom{
		Code:    405,
		Message: msg,
	}
}

/*
NotAcceptable :
*/
func NotAcceptable(msg string) *Boom {
	return &Boom{
		Code:    406,
		Message: msg,
	}
}

/*
ProxyAuthRequired :
*/
func ProxyAuthRequired(msg string) *Boom {
	return &Boom{
		Code:    407,
		Message: msg,
	}
}

/*
ClientTimeout :
*/
func ClientTimeout(msg string) *Boom {
	return &Boom{
		Code:    408,
		Message: msg,
	}
}

/*
Conflict :
*/
func Conflict(msg string) *Boom {
	return &Boom{
		Code:    409,
		Message: msg,
	}
}

/*
ResourceGone :
*/
func ResourceGone(msg string) *Boom {
	return &Boom{
		Code:    410,
		Message: msg,
	}
}

/*
LengthRequired :
*/
func LengthRequired(msg string) *Boom {
	return &Boom{
		Code:    411,
		Message: msg,
	}
}

/*
PreconditionFailed :
*/
func PreconditionFailed(msg string) *Boom {
	return &Boom{
		Code:    412,
		Message: msg,
	}
}

/*
EntityTooLarge :
*/
func EntityTooLarge(msg string) *Boom {
	return &Boom{
		Code:    413,
		Message: msg,
	}
}

/*
URITooLong :
*/
func URITooLong(msg string) *Boom {
	return &Boom{
		Code:    414,
		Message: msg,
	}
}

/*
UnsupportedMediaType :
*/
func UnsupportedMediaType(msg string) *Boom {
	return &Boom{
		Code:    415,
		Message: msg,
	}
}

/*
RangeNotSatisfiable :
*/
func RangeNotSatisfiable(msg string) *Boom {
	return &Boom{
		Code:    416,
		Message: msg,
	}
}

/*
BadData :
*/
func BadData(msg string) *Boom {
	return &Boom{
		Code:    422,
		Message: msg,
	}
}

/*
PreconditionRequired :
*/
func PreconditionRequired(msg string) *Boom {
	return &Boom{
		Code:    428,
		Message: msg,
	}
}

/*
TooManyRequests :
*/
func TooManyRequests(msg string) *Boom {
	return &Boom{
		Code:    429,
		Message: msg,
	}
}

/*
Illegal :
*/
func Illegal(msg string) *Boom {
	return &Boom{
		Code:    451,
		Message: msg,
	}
}

/*
BadImplementation :
*/
func BadImplementation(msg string) *Boom {
	return &Boom{
		Code:    500,
		Message: msg,
	}
}

/*
NotImplemented :
*/
func NotImplemented(msg string) *Boom {
	return &Boom{
		Code:    501,
		Message: msg,
	}
}

/*
BadGateway :
*/
func BadGateway(msg string) *Boom {
	return &Boom{
		Code:    502,
		Message: msg,
	}
}

/*
ServerUnavailable :
*/
func ServerUnavailable(msg string) *Boom {
	return &Boom{
		Code:    503,
		Message: msg,
	}
}

/*
GatewayTimeout :
*/
func GatewayTimeout(msg string) *Boom {
	return &Boom{
		Code:    504,
		Message: msg,
	}
}
