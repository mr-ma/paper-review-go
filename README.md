# Screencast (demo) of the tool
[![Demo](https://img.youtube.com/vi/na_DjN1tdc4/0.jpg)](https://www.youtube.com/watch?v=na_DjN1tdc4)

# Dockerized approach
- create a `.env` file in the root folder

- set my sql root password variable in .env file

  `MYSQL_DATABASE=classification`

  `MYSQL_ROOT_PASSWORD=SECURE_PASSWORD`

  `MYSQL_ROOT_HOST=%`

- run `docker-compose up`

  This sets up nginx reverse proxy, runs an instance of go server, and a mysql instance (initialized with the db_schema.sql file) all in separated containers.
- navigate to localhost:80, everything should work out of the box!
- to modify the reverse proxy edit nginx config file in `nginx/config.d/nginx.config`

- Use the default root user credentials to approve users
`username: root@root.com`
`password: root`

# Important note
Make sure that you change the root password, or deleting the user, before deploying the platform to any production environment!


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

# Citation
Please cite this work as:
```
@article{DBLP:journals/corr/abs-1906-11217,
  author    = {Mohsen Ahmadvand and
               Amjad Ibrahim and
               Felix Huber},
  title     = {Taxonomy-as-a-Service: How To Structure Your Related Work},
  journal   = {CoRR},
  volume    = {abs/1906.11217},
  year      = {2019},
  url       = {http://arxiv.org/abs/1906.11217},
  archivePrefix = {arXiv},
  eprint    = {1906.11217},
  timestamp = {Thu, 27 Jun 2019 18:54:51 +0200},
  biburl    = {https://dblp.org/rec/bib/journals/corr/abs-1906-11217},
  bibsource = {dblp computer science bibliography, https://dblp.org}
}
```
