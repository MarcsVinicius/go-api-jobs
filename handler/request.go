package handler

import "fmt"

func errParamIsRequired(name string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func errIdNotFound(id string) error {
	return fmt.Errorf("ID: %s not found", id)
}

//CreateOpening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or invalid")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("Company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("Location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("Link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("Remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("Salary", "int64")
	}
	return nil
}

//UpdateOpening

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
