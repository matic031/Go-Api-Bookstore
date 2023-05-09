# Go-Api-Bookstore
 
Commands:

run GO server: go run main.go<br>
check list of books: curl http://localhost:8080/books<br>
create new book from JSON payload (body.json): curl -X POST -H "Content-Type: application/json" -d @body.json http://localhost:8080/books<br>
remove book: curl -X DELETE http://localhost:8080/books/{id}<br>
Update a book by id: curl -X PATCH -H "Content-Type: application/json" -d "{\"title\": \"New Title\", \"author\": \"New Author\", \"quantity\": 10}" http://localhost:8080/books/{id}<br>
Checkout (quantity -1) book by id: curl -X PATCH http://localhost:8080/checkout?id={id}<br>
Return (quantity +1) a book by id: curl -X PATCH http://localhost:8080/return?id={id}<br>
