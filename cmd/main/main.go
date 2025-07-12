package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	integratonsdk "github.com/jerin-hc/integration-sdk/internal"
	"github.com/jerin-hc/integration-sdk/jsoncodec"
	"github.com/jerin-hc/integration-sdk/schema"
)

func main() {

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		fmt.Fprintln(os.Stderr, "plugin received SIGINT/SIGTERM, exiting...")
		os.Exit(0)
	}()

	jsoncodec.Init()

	// setting for developmet purpose
	os.Setenv("TF_PLUGIN_MAGIC_COOKIE", "d602bf8f470bc67ca7faa0386276bbdd4330efaf76d1a219cb4d6991ca9872b2")

	s := integratonsdk.New("example-cost-analyzer-v1")

	// Handle multiple event types, e.g., post-plan or a custom assessment trigger
	s.Handle([]schema.Event{schema.PostPlan, schema.PostResourceAssessment}, func(event schema.Event, resources []schema.Resource, ctx schema.Ctx) *schema.ResourceResponse {
		updates := make([]schema.ResourceChange, 0)
		totalCost := 0.0
		instanceCount := 0

		for _, resource := range resources {
			// Example: Focus on EC2 instances or equivalent
			if resource.Type == "aws_instance" { // Simplified type check
				instanceCount++
				// Placeholder: Imagine fetching actual cost data via an API or internal logic
				currentCost := 50.0                  // Simulated current cost
				suggestedInstanceType := "t4.medium" // Simulated suggestion
				// costAfterChange := 45.67 // Simulated cost after change

				// Annotate with cost
				annotation := map[string]interface{}{
					"current_monthly_cost": fmt.Sprintf("%.2f", currentCost),
					"currency":             "USD",
				}

				// Assess and suggest changes
				comment := schema.Comment{
					Pass:    currentCost <= 40.0, // Example budget check
					Message: fmt.Sprintf("Instance cost: $%.2f/mo.", currentCost),
				}
				if currentCost > 40.0 {
					comment.Message += " Exceeds budget of $40/mo. Consider changing to " + suggestedInstanceType + "."
				}

				change := schema.ResourceChange{
					Identity: resource.Identity,
					Mutate: map[string]interface{}{ // Suggest attribute changes
						"instance_type": suggestedInstanceType,
						"tags":          []string{"budget-review-needed"},
					},
					Annotation: annotation,
					Comment:    comment,
				}
				updates = append(updates, change)
				totalCost += currentCost // Or costAfterChange if applying mutations
			}
		}

		overallAssessmentComment := fmt.Sprintf("Analyzed %d instances. Total current monthly cost: $%.2f", instanceCount, totalCost)
		if totalCost > 100.0 { // Example overall budget
			overallAssessmentComment += ". Overall budget exceeded."
		}

		return &schema.ResourceResponse{
			Resources: updates,
			Comment: schema.Comment{ // Overall assessment for the task run
				Pass:    totalCost <= 100.0,
				Message: overallAssessmentComment,
			},
		}
	})

	s.Run()
}
