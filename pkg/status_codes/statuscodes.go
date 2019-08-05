package status_codes

import (
    "errors"
    "fmt"
)

var statusErrorMap map[uint16] error

const CONTINUE = uint16(100)
const SWITCHING_PROTOCOLS = uint16(101)
const OK = uint16(200)
const CREATED = uint16(201)
const ACCEPTED = uint16(202)
const NONAUTHORITATIVE_INFORMATION = uint16(203)
const NO_CONTENT = uint16(204)
const RESET_CONTENT = uint16(205)
const PARTIAL_CONTENT = uint16(206)
const MULTIPLE_CHOICES = uint16(300)
const MOVED_PERMANENTLY = uint16(301)
const FOUND = uint16(302)
const SEE_OTHER = uint16(303)
const NOT_MODIFIED = uint16(304)
const USE_PROXY = uint16(305)
const TEMPORARY_REDIRECT = uint16(307)
const BAD_REQUEST = uint16(400)
const UNAUTHORIZED = uint16(401)
const PAYMENT_REQUIRED = uint16(402)
const FORBIDDEN = uint16(403)
const NOT_FOUND = uint16(404)
const METHOD_NOT_ALLOWED = uint16(405)
const NOT_ACCEPTABLE = uint16(406)
const PROXY_AUTHENTICATION_REQUIRED = uint16(407)
const REQUEST_TIMEOUT = uint16(408)
const CONFLICT = uint16(409)
const GONE = uint16(410)
const LENGTH_REQUIRED = uint16(411)
const PRECONDITION_FAILED = uint16(412)
const REQUEST_ENTITY_TOO_LARGE = uint16(413)
const REQUESTURI_TOO_LONG = uint16(414)
const UNSUPPORTED_MEDIA_TYPE = uint16(415)
const REQUESTED_RANGE_NOT_SATISFIABLE = uint16(416)
const EXPECTATION_FAILED = uint16(417)
const INTERNAL_SERVER_ERROR = uint16(500)
const NOT_IMPLEMENTED = uint16(501)
const BAD_GATEWAY = uint16(502)
const SERVICE_UNAVAILABLE = uint16(503)
const GATEWAY_TIMEOUT = uint16(504)
const HTTP_VERSION_NOT_SUPPORTED = uint16(505)

func init() {
    statusErrorMap = make(map[uint16] error)
    statusErrorMap[100] = errors.New("100: Continue")
    statusErrorMap[101] = errors.New("101: Switching Protocols")
    statusErrorMap[200] = errors.New("200: OK")
    statusErrorMap[201] = errors.New("201: Created")
    statusErrorMap[202] = errors.New("202: Accepted")
    statusErrorMap[203] = errors.New("203: Non-Authoritative Information")
    statusErrorMap[204] = errors.New("204: No Content")
    statusErrorMap[205] = errors.New("205: Reset Content")
    statusErrorMap[206] = errors.New("206: Partial Content")
    statusErrorMap[300] = errors.New("300: Multiple Choices")
    statusErrorMap[301] = errors.New("301: Moved Permanently")
    statusErrorMap[302] = errors.New("302: Found")
    statusErrorMap[303] = errors.New("303: See Other")
    statusErrorMap[304] = errors.New("304: Not Modified")
    statusErrorMap[305] = errors.New("305: Use Proxy")
    statusErrorMap[306] = errors.New("306: (Unused)")
    statusErrorMap[307] = errors.New("307: Temporary Redirect")
    statusErrorMap[400] = errors.New("400: Bad Request")
    statusErrorMap[401] = errors.New("401: Unauthorized")
    statusErrorMap[402] = errors.New("402: Payment Required")
    statusErrorMap[403] = errors.New("403: Forbidden")
    statusErrorMap[404] = errors.New("404: Not Found")
    statusErrorMap[405] = errors.New("405: Method Not Allowed")
    statusErrorMap[406] = errors.New("406: Not Acceptable")
    statusErrorMap[407] = errors.New("407: Proxy Authentication Required")
    statusErrorMap[408] = errors.New("408: Request Timeout")
    statusErrorMap[409] = errors.New("409: Conflict")
    statusErrorMap[410] = errors.New("410: Gone")
    statusErrorMap[411] = errors.New("411: Length Required")
    statusErrorMap[412] = errors.New("412: Precondition Failed")
    statusErrorMap[413] = errors.New("413: Request Entity Too Large")
    statusErrorMap[414] = errors.New("414: Request-URI Too Long")
    statusErrorMap[415] = errors.New("415: Unsupported Media Type")
    statusErrorMap[416] = errors.New("416: Requested Range Not Satisfiable")
    statusErrorMap[417] = errors.New("417: Expectation Failed")
    statusErrorMap[500] = errors.New("500: Internal Server Error")
    statusErrorMap[501] = errors.New("501: Not Implemented")
    statusErrorMap[502] = errors.New("502: Bad Gateway")
    statusErrorMap[503] = errors.New("503: Service Unavailable")
    statusErrorMap[504] = errors.New("504: Gateway Timeout")
    statusErrorMap[505] = errors.New("505: HTTP Version Not Supported")
}

func StatusToError(status uint16) (error) {
    if val, ok := statusErrorMap[status]; ok {
        return val
    } else {
        return errors.New(fmt.Sprintf("%d: Unknown status code"))
    }
}

func StatusToErrorFilter(status uint16) (error) {
    if status < 400 {
        return nil
    } else {
        return StatusToError(status)
    }
}

func IsInformational(status uint16) (bool) {
    return status/100 == 1
}

func IsSuccessful(status uint16) (bool) {
    return status/100 == 2
}

func IsMultipleChoices(status uint16) (bool) {
    return status/100 == 3
}

func IsClientError(status uint16) (bool) {
    return status/100 == 2
}

func IsServerError(status uint16) (bool) {
    return status/100 == 2
}