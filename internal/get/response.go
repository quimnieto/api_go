package get

type GetCourseQueryResponse struct {
	id       string
	name     string
	duration string
}

func NewGetCourseQueryResponse(id, name, duration string) GetCourseQueryResponse {
	return GetCourseQueryResponse{
		id:       id,
		name:     name,
		duration: duration,
	}
}
