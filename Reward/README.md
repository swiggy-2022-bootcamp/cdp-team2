# Reward Microservice

This microservice handles all the operations related to reward associated with a customer.

1. **Rest API Endpoints** available: <br>
   i. GET `/reward/v1/healthcheck`

   - Checks the health of the given microservice.
   - Returns the json stating the status of the service.
   - In case of internal error, http error is sent with `internal error message` and HTTP Response code 500 is set. <br>

   ii. POST `/reward/v1/reward`

   - Allows user with role **admin** to give the reward for a given a customer.
   - Request Body - CustomerId of type string, Description of type String and reward points of type int are required.
   - Returns - HTTP StatusCode `200 OK`
   - This API adds the reward given by admin to the given customer reward total balance. <br>

2. **gRPC services** available: <br>
   i. `GetReward`

   - Gives the reward details of the customer to the microservice calling this gRPC function.
   - It accepts `RewardRequest` proto message with field `CustomerId`
   - Returns
     - `RewardResponse` proto message with fields `CustomerId`, `Description` and `Points`.
     - `Error` of type error is returned in case of any error. It returns the gRPC error status and associated error message from Reward MS.<br>

   ii. `RedeemReward`

   - Allows the customer to use the reward at checkout if redeem reward option is selected.
   - It accepts `RedeemRewardRequest` proto message with fields `orderId` and `points` reward points customer wants to apply.
   - Returns
     - `RedeemRewardResponse` proto message with field `finalPrice` which is the final price after deducting the reward points from the total pricing.
     - `Error` of type error is returned in case of any error. It returns the gRPC error status and associated error message from Reward MS.<br>

## Associated Dependencies

This section contains the list of all the upstream and downstream dependencies of the Reward Microservice.

1. DynamoDB
2. Order Microservice
3. Checkout Microservice
4. Customer-Frontstore Microservice

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

## Flow Diagram of Reward MS

## Local Setup for Development

**Prerequisties** 1. Golang ~ 1.17 or above should be installed on the system. 2. DynamoDb - should be set locally. <br>

- To install all the dependent packages, run `go mod tidy`
- To bring up the service, run the following command `cd cmd/server`
  and `go run main.go`

## References

- [Meeting Notes of the Project](https://docs.google.com/document/d/1VR5kihnHYApgUbRzAXx9rW-7rzDE1Mtgy3_hjCs-70g/edit?usp=sharing)
