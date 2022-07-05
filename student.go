type StudentHandler struct {
	StudentRepo model.StudentRepo
	Validator *validator.Validate
	}
	type Student struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	ClassID int64 `json:"class_id" validate:"required"`
	Score float64 `json:"score" validate:"required,gte=0,lte=20"`
	}
	type StudentCreate struct {
	ClassID int64 `json:"class_id" validate:"required"`
	Students []Student `json:"students" validate:"required,dive"`
	}
	func (s StudentHandler) Create(c echo.Context) error {
	req := new(StudentCreate)
	if err := c.Bind(req); err != nil {
	return echo.ErrBadRequest
	}
	if err := s.Validator.Struct(req); err != nil {
	return echo.ErrBadRequest
	}
	err := s.StudentRepo.SaveForClass(req.ClassID, req.Students)
	if err != nil {
	return echo.ErrInternalServerError
	}
	return c.NoContent(http.StatusCreated)
	}