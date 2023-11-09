package get

import mooc "api_go/internal"

type GetCourseService struct {
	repository mooc.CourseRepository
}

func NewGetCourseService(courseRepository mooc.CourseRepository) GetCourseService {
	return GetCourseService{
		repository: courseRepository,
	}
}

func (service GetCourseService) ById(id string) (GetCourseQueryResponse, error) {
	courseId, err := mooc.NewCourseId(id)

	if err != nil {
		return GetCourseQueryResponse{}, err
	}

	course, err := service.repository.ById(courseId)

	if err != nil {
		return GetCourseQueryResponse{}, err
	}

	response := NewGetCourseQueryResponse(course.Id(), course.Name(), course.Duration())

	return response, nil
}
