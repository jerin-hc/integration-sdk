package integratonsdk

import (
	"context"
	"fmt"
)

type Event string

const (
	PrePlan                Event = "pre_plan"
	PostPlan               Event = "post_plan"
	PreApply               Event = "pre_apply"
	PostApply              Event = "post_apply"
	PostResourceAssessment Event = "post_resource_assessment"
)

type Resource struct {
	Type     string
	Identity string
}

type ResourceResponse struct {
	Resources []ResourceChange
	Comment   Comment
}

type Ctx context.Context

type Serve struct{}

type Comment struct {
	Pass    bool
	Message string
}

type ResourceChange struct {
	Identity   string
	Mutate     map[string]interface{}
	Annotation map[string]interface{}
	Comment    Comment
}

func New(id string) *Serve {
	return &Serve{}
}

func (s *Serve) Handle(event []Event, handleFunc func(event Event, resources []Resource, ctx Ctx) *ResourceResponse) {
	fmt.Println("TODO")
}

func (s *Serve) Run() {
	fmt.Println("TODO")
}
