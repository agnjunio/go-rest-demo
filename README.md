# go-rest-demo
A demo repository for the go langague containing a simple REST Api for transactions.

## Running Via Docker
Requirements:
- [docker](https://www.docker.com/docker-community)
- [docker-compose](https://docs.docker.com/compose/)

The `docker-compose.yml` file contains a docker-compose script to run the entire stack locally using the latest image published into DockerHub. To start it, just copy the file into your file system and run.

```
docker-compose up
```

This will pull the images and run the stack. After that, the REST API will be available at `http://localhost/`.

## Running Via Source
Requirements:
- [git](https://git-scm.com/)
- [go](https://golang.org/) ^v1.15.6
- [mongodb](https://www.mongodb.com/)

Steps:

1. Setup your mongodb (not covered in this guide) and set the environment variable MONGI_URI to the database using the mongo schema
> You create a .env file in the root of the project to define your environment variables locally
2. Clone the git repository:
```
git clone git@github.com:agnjunio/go-rest-demo.git
```
3. Cd into repository:
```
cd go-rest-demo
```
4. Run the API: 
```
go run .
```

The API should be available in the default port: `http://localhost:8080/`. To change the port, set the PORT environment variable.

## API Documentation

Once the API is up, you can go to the `/docs` to view the API documentation in Swagger UI.
