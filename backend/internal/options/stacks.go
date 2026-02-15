package options

type Stack struct {
	ID   int    `json:"id"`
	Data string `json:"data"`
}

func GetStacks() []Stack {
	return []Stack{
		// Go Ecosystem
		{ID: 1, Data: "Go - Built in (net/http)"},
		{ID: 2, Data: "Go - Fiber"},
		{ID: 3, Data: "Go - Gin"},
		{ID: 4, Data: "Go - Echo"},
		{ID: 5, Data: "Go - gRPC Services"},

		// Java Ecosystem
		{ID: 6, Data: "Java - Spring Boot"},
		{ID: 7, Data: "Java - Spring WebFlux"},
		{ID: 8, Data: "Java - Quarkus"},
		{ID: 9, Data: "Java - Micronaut"},

		// .NET
		{ID: 10, Data: ".NET - ASP.NET Core"},
		{ID: 11, Data: ".NET - Minimal APIs"},

		// Node / TypeScript Backend
		{ID: 12, Data: "TypeScript - Node (Express)"},
		{ID: 13, Data: "TypeScript - Fastify"},
		{ID: 14, Data: "TypeScript - NestJS"},
		{ID: 15, Data: "TypeScript - tRPC"},

		// Frontend Frameworks
		{ID: 16, Data: "TypeScript - React (Vite)"},
		{ID: 17, Data: "TypeScript - Next.js"},
		{ID: 18, Data: "TypeScript - Angular"},
		{ID: 19, Data: "TypeScript - Vue"},
		{ID: 20, Data: "TypeScript - Svelte"},

		// Python
		{ID: 21, Data: "Python - FastAPI"},
		{ID: 22, Data: "Python - Django"},
		{ID: 23, Data: "Python - Flask"},

		// Databases
		{ID: 24, Data: "PostgreSQL"},
		{ID: 25, Data: "MySQL"},
		{ID: 26, Data: "MongoDB"},
		{ID: 27, Data: "Redis"},
		{ID: 28, Data: "SQLite"},

		// Messaging / Async
		{ID: 29, Data: "RabbitMQ"},
		{ID: 30, Data: "Kafka"},
		{ID: 31, Data: "NATS"},
		{ID: 32, Data: "AWS SQS/SNS"},

		// Realtime / APIs
		{ID: 33, Data: "WebSockets"},
		{ID: 34, Data: "GraphQL"},
		{ID: 35, Data: "REST APIs"},
		{ID: 36, Data: "gRPC"},

		// DevOps / Infra
		{ID: 37, Data: "Docker"},
		{ID: 38, Data: "Kubernetes"},
		{ID: 39, Data: "NGINX"},
		{ID: 40, Data: "Terraform"},

		// Observability
		{ID: 41, Data: "Prometheus"},
		{ID: 42, Data: "Grafana"},
		{ID: 43, Data: "OpenTelemetry"},

		// Cloud
		{ID: 44, Data: "AWS"},
		{ID: 45, Data: "Azure"},
		{ID: 46, Data: "GCP"},

		// Architecture Styles
		{ID: 47, Data: "Monolith"},
		{ID: 48, Data: "Microservices"},
		{ID: 49, Data: "Event-Driven"},
		{ID: 50, Data: "Serverless"},
	}
}
