# Mannual execution steps:
0. Clone master branch
1. Grab the latest database dump form paper-review-go/SQLSchema/dumps/
2. Import the dump to mysql
3. Get dependencies as mentioned in the Dockerfile:
RUN go get github.com/rcrowley/go-tigertonic
RUN go get github.com/alexedwards/scs
RUN go get github.com/Jeffail/gabs
RUN go get github.com/stretchr/testify/assert
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/mr-ma/paper-review-go

4. Run taxonomy backend: go run taxonomyserver.go -mysqluser "USER" -mysqlpass "PASS"
5. Run frontend server: cd github.com/mr-ma/paper-review-go/frontend go run server.go
6. Browse on localhost:8001

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
