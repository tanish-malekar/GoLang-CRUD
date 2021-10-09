# AppointyTask
<b>Run using: 
<i>go run server.go</i><b><br>
<b>Test using:
<i>go test</i></b>

<h2> <img src = "https://media2.giphy.com/media/QssGEmpkyEOhBCb7e1/giphy.gif?cid=ecf05e47a0n3gi1bfqntqmob8g9aid1oyj2wr3ds3mg700bl&rid=giphy.gif" width = 32px> Endpoints   </h2>

* **Create an User: http://localhost:8080/users (POST request)**
* **Get a user using id: http://localhost:8080/users/<id here> (GET request)**
* **Create a Post: http://localhost:8080/posts (POST request)** 
* **Get a post using id: http://localhost:8080/posts/<id here> (GET request)**. 
* **List all posts of a user: http://localhost:8080/posts/users/<id here>?offset=<offset here>&limit=<limit here> (GET request)(Pagination has been implemented)**.

<h2> <img src = "https://media2.giphy.com/media/QssGEmpkyEOhBCb7e1/giphy.gif?cid=ecf05e47a0n3gi1bfqntqmob8g9aid1oyj2wr3ds3mg700bl&rid=giphy.gif" width = 32px> Features Implemented  </h2>

- The API was developed using Go
- MongoDB was used for storage
- Only Golang Standard library and mongo-driver packege was used

- Passwords are securely stored in the database with encryption
- The server has been made thread safe
- Pagination was used for list endpoints
- Unit tests were added
 
 <h2> <img src = "https://media2.giphy.com/media/QssGEmpkyEOhBCb7e1/giphy.gif?cid=ecf05e47a0n3gi1bfqntqmob8g9aid1oyj2wr3ds3mg700bl&rid=giphy.gif" width = 32px>Screenshots</h2>
 Post User
<img src = "https://github.com/tanish-malekar/Appointy-Task/blob/main/Screenshots/postUser.jpg?raw=true" width = 100px>
 Get user by ID
 <img src = "https://github.com/tanish-malekar/Appointy-Task/blob/main/Screenshots/getUserByID.jpg?raw=true" width = 100px>
 Post a post
 <img src = "https://github.com/tanish-malekar/Appointy-Task/blob/main/Screenshots/postPost.jpg?raw=true" width = 100px>
 Get Post by Id
 <img src = "https://github.com/tanish-malekar/Appointy-Task/blob/main/Screenshots/getPostByID.jpg?raw=true" width = 100px>
 Get Posts by ID
  <img src = "https://github.com/tanish-malekar/Appointy-Task/blob/main/Screenshots/getPostsByUserID.jpg?raw=true" width = 100px>
 Get Posts by ID (with pagination)
 <img src = "https://github.com/tanish-malekar/Appointy-Task/blob/main/Screenshots/getPostsByUserID.jpg?raw=true" width = 100px>


