### EXAMPLE

```
package main

import (
	integratonsdk "github.com/jerin-hc/integration-sdk/internal"
)

func main() {
	s := integratonsdk.New("example-cost-analyzer-v1")

	// Handle multiple event types, e.g., post-plan or a custom assessment trigger
	s.Handle([]integratonsdk.Event{integratonsdk.PostPlan, integratonsdk.PostResourceAssessment}, func(event integratonsdk.Event, resources []integratonsdk.Resource, ctx integratonsdk.Ctx) *integratonsdk.ResourceResponse 	{
    		// add custom logic here...
	})

	s.Run()
}
```
