# Cart Microservice

This microservice handles all the operations related to cart associated with a customer.

1. Rest API Endpoints available:
   i. GET `/cart/v1/healthcheck`

   - Checks the health of the given microservice.
   - Returns the json stating the status of the service.
   - In case of internal error, http error is sent with `internal error message` and HTTP Response code 500 is set. <br>

   ii. GET `/cart/v1/cart`

   - Allows user to get all the cart details as list of products present in the cart.
   - Returns - HTTP StatusCode `200 OK` and cart details.<br>

   iii. POST `/cart/v1/cart`

   - Allows the customer to add new product as cart item in the cart.
   - Request Body - productId as string and quantity as number.
   - Returns - `200 ok`

   iv. PUT `cart/v1/cart`

   - Allows customer to update the product quntanity in the cart.
   - Request Body - productId as string and quantity as number.
   - Returns - 200 OK <br>

   v. DELETE `cart/v1/cart/{productId}`

   - Allows the customer to remove the particular product from the cart.
   - Returns - 200 OK on successful deletion

2. gRPC services available:
   i. `GetCart`

   - Gives the cart details of the customer to the microservice calling this gRPC function.
   - It accepts `CartRequest` proto message with field `CustomerId`
   - Returns
     - `CartResponse` proto message with fields `CustomerId`, and list of cart items/products.
     - `Error` of type error is returned in case of any error. It returns the gRPC error status and associated error message from Cart MS.<br>

   ii. `EmptyCart`

   - Allows Checkout MS to empty the cart for the user.
   - It accepts `CartRequest` proto message with field customerId.
   - Returns
     - `EmptyCartResponse` proto message
     - `Error` of type error is returned in case of any error. It returns the gRPC error status and associated error message from Reward MS.<br>

### Tasks

- [x] Completed Writing Microservice Assigned
- [x] Written Test Cases
- [x] Dockerizing an application
- [x] Swagger Documentation
- [x] Effective Use of GRPC
- [x] Container Orchestration
- [x] Have maintained in code repository
- [x] Data base model properly show cased

## Associated Dependencies

This section contains the list of all the upstream and downstream dependencies of the Reward Microservice.

1. DynamoDB
2. Checkout Microservice
3. Customer-Frontstore Microservice

## Technology/Tools Used

This section contains the list of tools and technology stack used while developing the microservice.

1. Golang
2. DynamoDB
3. Docker
4. Swagger
5. Unit Testing
6. gRPC
7. Protocol Buffers
8. Jenkins
9. Ngrok
10. BloomRPC
11. Postman
12. SonarQube

## Flow Diagram of Cart MS

## Local Setup for Development

**Prerequisties** <br>

1. Golang ~ 1.17 or above should be installed on the system.
2. DynamoDb - should be set locally. <br>

**Steps**

- To install all the dependent packages, run `go mod tidy`
- To bring up the service, run the following command `cd cmd/server`
  and `go run main.go`

## References

- [Meeting Notes of the Project](https://docs.google.com/document/d/1VR5kihnHYApgUbRzAXx9rW-7rzDE1Mtgy3_hjCs-70g/edit?usp=sharing)
