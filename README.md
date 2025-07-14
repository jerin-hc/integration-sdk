# integration-sdk

A framework for building gRPC-based plugins that can handle infrastructure events like post-plan, assessments, and more.

---

## 📦 Project Structure

```
integrations-sdk/
├── cmd/
│   └── main/
│       └── main.go           # Entry point for building your plugin
├── go.mod                    # Go module definition
├── go.sum                    # Go dependencies lock file
├── grpc-plugin/
│   └── plugin.go             # gRPC plugin definitions
├── internal/
│   └── serve.go              # Internal serving logic
├── jsoncodec/
│   └── jsoncodec.go          # JSON codec handling
├── myplugin                  # Built plugin output (binary)
├── README.md                 # Project documentation (this file)
├── schema/
│   └── schema.go             # Event and schema definitions
└── test/
    └── main.go               # Client code for testing your plugin
```

---

## 🔧 How to Build a Plugin

```bash
go build -o myplugin cmd/main/main.go
```

This will compile your plugin and generate a binary named `myplugin`.

---

## 🧪 How to Test

```bash
go run test/main.go
```

This runs the test client which connects to the plugin and sends test event data to validate your plugin logic.  
The file `test/main.go` contains gRPC client code for simulating plugin execution and observing the response.

---

## 🐞 How to Debug

```bash
dlv debug myplugin -- -gcflags=all="-N -l"
```

This command launches your plugin binary in the Delve debugger with optimizations disabled for easier debugging.

---

## 🚀 Example: Creating a Plugin

```go
s := integratonsdk.New("example-cost-analyzer-v1")

// Handle multiple event types
s.Handle([]integratonsdk.Event{
    integratonsdk.PostPlan,
    integratonsdk.PostResourceAssessment,
}, func(event integratonsdk.Event, resources []integratonsdk.Resource, ctx integratonsdk.Ctx) *integratonsdk.ResourceResponse {
    // Add custom plugin logic here...
})

s.Run()
```

---