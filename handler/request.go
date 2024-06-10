package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create Opportunity

type CreateOpportunityRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (cr *CreateOpportunityRequest) Validate() error {
	if cr.Role == "" &&
		cr.Company == "" &&
		cr.Location == "" &&
		cr.Link == "" &&
		cr.Remote == nil &&
		cr.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if cr.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if cr.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if cr.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if cr.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if cr.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if cr.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}

// UpdateOpportunityRequest UpdateOpportunity
type UpdateOpportunityRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpportunityRequest) Validate() error {
	// If any field is provided, validation is truthy
	if r.Role != "" ||
		r.Company != "" ||
		r.Location != "" ||
		r.Link != "" ||
		r.Remote != nil ||
		r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("request body is empty or malformed")
}
