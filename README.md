# Go CRUD API

This is a basic CRUD API built with Go and Gin framework. It provides endpoints to create, read, update and delete posts.

## Installation and Usage

To run the API, you will need to have Golang and a PostgreSQL database installed on your machine. You will also need to create a `.env` file with your database credentials.

1.  Clone the repository: `git clone https://github.com/shrine2000/go-crud.git`
2.  Install dependencies: `go mod download`
3.  Create a `.env` file with your database credentials
4.  Run the application: `go run main.go`

The API will be available at `http://localhost:8080`.


## Endpoints

-   `POST /posts` - creates a new post with a title and body
-   `GET /posts` - retrieves a list of all posts
-   `GET /posts/:id` - retrieves a single post by its ID
-   `PUT /posts/:id` - updates a post by its ID
-   `DELETE /posts/:id` - deletes a post by its ID


## Screenshots

### Create
 ![](https://github.com/shrine2000/go-crud/blob/master/screenshots/create.png?raw=true)

### Update
 ![](https://github.com/shrine2000/go-crud/blob/master/screenshots/update.png?raw=true)
 
### Fetch
 ![](https://github.com/shrine2000/go-crud/blob/master/screenshots/fetch.png?raw=true)
### Delete
 ![](https://github.com/shrine2000/go-crud/blob/master/screenshots/delete.png?raw=true)


## Testing

The API endpoints can be tested using the included Postman configuration file located at `go-crud/postman/Post.postman_collection.json`. Import this file into your Postman application and use it to test the API endpoints.


## Contributing

If you would like to contribute to this project, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. 
