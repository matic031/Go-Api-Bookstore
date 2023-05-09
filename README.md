# Go-Api-Bookstore
 
Commands:

run GO server: go run main.go
check list of books: curl http://localhost:8080/books
create new book from JSON payload (body.json): curl -X POST -H "Content-Type: application/json" -d @body.json http://localhost:8080/books
remove book: curl -X DELETE http://localhost:8080/books/{id}
Update a book by id: curl -X PATCH -H "Content-Type: application/json" -d "{\"title\": \"New Title\", \"author\": \"New Author\", \"quantity\": 10}" http://localhost:8080/books/{id}
Checkout (quantity -1) book by id: curl -X PATCH http://localhost:8080/checkout?id={id}
Return (quantity +1) a book by id: curl -X PATCH http://localhost:8080/return?id={id}
