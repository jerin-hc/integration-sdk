package schema

import "context"

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

type HandleFuncRequest struct {
	Event     Event
	Resources []Resource
}

type Ctx context.Context

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

type IntegrationServer struct {
	HandleFunc func(event Event, resources []Resource, ctx Ctx) *ResourceResponse
}
