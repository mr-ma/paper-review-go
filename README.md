# paper-review-go
paper review enables researchers to review papers collaboratively and vote for papers 

Install:
go get github.com/mr-ma/paper-review-go

Run backend service:
cd github.com/mr-ma/paper-review-go
go run hello.go

Run frontend service:
cd github.com/mr-ma/paper-review-go/frontend
go run server.go

# taxonomyserver.go

Enables researchers to build their own taxonomy of concepts and offers three visualization methods to analyze and modify them

Available visualizations are:

- 2D Correlation Matrix
- UML-like Visualization
- Circle Packing Visualization