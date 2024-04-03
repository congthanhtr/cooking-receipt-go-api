# Simple Cooking API with CRUD and search function

## Structure:

```
cooking-receipt
    |_connector
        |_sqliteConnector
    |_controller
        |_receiptController.go
        |_receiptController_test.go
    |_model
        |_ingredient
        |_receipt
    |_route
        |_handleRequest.go
        |_receiptRoute.go
    |_wrapper
        |_receiptWrapper.go
        |_receiptWrapper_test.go    

```

## Technologies

- Golang
- Sqlite3 Database