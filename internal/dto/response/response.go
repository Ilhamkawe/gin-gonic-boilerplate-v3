package dto

type Pagination struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}

type Meta struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Response struct {
	Meta  Meta        `json:"meta"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error,omitempty"`
}

func NewResponse(code int, message string, data interface{}) Response {
	return Response{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
		Data: data,
	}
}

func NewPaginatedResponse(code int, message string, data interface{}, page, limit int, total int64) Response {
	return Response{
		Meta: Meta{
			Code:    code,
			Message: message,
			Pagination: &Pagination{
				Page:  page,
				Limit: limit,
				Total: total,
			},
		},
		Data: data,
	}
}

func NewErrorResponse(code int, message string, err error) Response {
	return Response{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
		Error: err.Error(),
	}
}
