package reuse

func (res ResponseType) Response(code int, body interface{}, message string) ResponseType {
	res.Code = code
	res.Body = body
	res.Message = message
	return res
}
