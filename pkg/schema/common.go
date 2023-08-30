package schema

const (
	TraceKey = "X-Request-Id"
	RespKey  = "Resp"
)

type (
	RespStruct struct {
		Resp      any
		ErrorCode int
	}
)
