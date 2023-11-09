package get

import (
	"api_go/kit/query"
	"context"
	"errors"
)

const GetCourseQueryType query.Type = "query.get.course"

type GetCourseQuery struct {
	id string
}

func NewGetCourseQuery(id string) GetCourseQuery {
	return GetCourseQuery{
		id: id,
	}
}

func (query GetCourseQuery) Type() query.Type {
	return GetCourseQueryType
}

type GetCourseQueryHandler struct {
	service GetCourseService
}

func NewGetCourseQueryHandler(service GetCourseService) GetCourseQueryHandler {
	return GetCourseQueryHandler{
		service: service,
	}
}

func (handler GetCourseQueryHandler) handle(ctx context.Context, query query.Query) (query.Response, error) {
	getCourseQuery, ok := query.(GetCourseQuery)

	if !ok {
		return errors.New("unexpected query")
	}
	
	response, err := handler.service.ById(query.GetCourseQueryType)
}
