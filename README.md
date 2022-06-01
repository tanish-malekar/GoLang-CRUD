# AppointyTask
<b>Run using: 
<i>go run server.go</i><b><br>
<b>Test using:
<i>go test</i></b>

<h2>Endpoints   </h2>

* **Create an User: http://localhost:8080/users (POST request)**
* **Get a user using id: http://localhost:8080/users/'id here' (GET request)**
* **Create a Post: http://localhost:8080/posts (POST request)** 
* **Get a post using id: http://localhost:8080/posts/'id here' (GET request)**. 
* **List all posts of a user: http://localhost:8080/posts/users/'id here'?offset='offset here'&limit='limit here; (GET request)<b>(Pagination has been implemented)<b>**.

<h2>Features Implemented  </h2>

- The API was developed using Go
- MongoDB was used for storage
- Only Golang Standard library and mongo-driver packege was used

- Passwords are securely stored in the database with encryption
- The server has been made thread safe
- Pagination was used for list endpoints
- Unit tests were added
 
 <h2>Screenshots</h2>
 Post User<br>
<img src = "https://github.com/tanish-malekar/Appointy-Task/blob/main/Screenshots/postUser.jpg?raw=true" width = 1000px>
 Get user by ID<br>
 <img src = "https://github.com/tanish-malekar/Appointy-Task/blob/main/Screenshots/getUserByID.jpg?raw=true" width = 1000px>
 Post a post<br>
 <img src = "https://github.com/tanish-malekar/Appointy-Task/blob/main/Screenshots/postPost.jpg?raw=true" width = 1000px>
 Get Post by Id<br>
 <img src = "https://github.com/tanish-malekar/Appointy-Task/blob/main/Screenshots/getPostByID.jpg?raw=true" width = 1000px>
 Get Posts by ID<br>
  <img src = "https://github.com/tanish-malekar/Appointy-Task/blob/main/Screenshots/getPostsByUserID.jpg?raw=true" width = 1000px>
 Get Posts by ID (with pagination)<br>
 <img src = "https://github.com/tanish-malekar/Appointy-Task/blob/main/Screenshots/getPostsByUserIDPagination.jpg?raw=true" width = 1000px>


