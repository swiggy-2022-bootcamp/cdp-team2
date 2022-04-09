# Mock Server

## Setup Steps

1. Run the following command `npm install -g json-server`
2. Run `json-server --watch db.json` to start the server in the folder `mock-server`

## Cart Microservice

1. **GET `localhost:3000/cart`** <br>
   Returns all items in the cart. <br>

- Headers - Authorization: Bearer JWT_TOKEN
- Response
  - Status Code - 200 OK
  - Response Body
    > [{
        "id": 1,
        "product_id": 56,
        "quantity": 2
    },<br>
    {"id": 2,
    "product_id": 45,
    "quantity": 5
    }]

2. **POST `localhost:3000/cart`** <br>
   Adds an item to cart with the given data.

- Headers - Authorization: Bearer JWT_TOKEN
- Request Body
  > {
  > "product_id": 100,
  > "quantity": 5
  > }
- Response
  - Status Code - 201 Created
  - Response Body
    > {
        "product_id": 100,
        "quantity": 5,
        "id": 3
        }

3. **PUT `localhost:3000/cart/{id}`** <br>
   Updates the quantity of cart item given as id in request URL.

- Headers - Authorization: Bearer JWT_TOKEN
- Request Body
  > {
  > "product_id": 100,
  > "quantity": 20
  > }
- Headers - Authorization: Bearer JWT_TOKEN
- Response
  - Status Code - 200 OK
  - Response Body
    > {
        "product_id": 100,
        "quantity": 20,
        "id": 1
        }

5. **DELETE `localhost:3000/cart/{id}`** <br>
   Delete cart item by id given in path parameter in request.

- Headers - Authorization: Bearer JWT_TOKEN
- Response:
  - Status Code - 204 No Content

7. **DELETE `localhost:3000/cart/empty`** <br>
   This endpoint is not supported by mock-server.

## Reward Microservice

1. **POST `localhost:3000/reward/{user_id}`** <br>
   Add reward points to the customer using user id by Admin.

- Headers - Authorization: Bearer JWT_TOKEN
- Request Body
  > {
      "user_id": 100,
      "description": "Reward points for shopping over Silver limit.",
      "points": 5,
  }
- Response
  - Status Code - 200 OK
  - Response Body
    > {
    > "user_id": 100,
    > "description": "Reward points for shopping over Silver limit.",
    > "points": 5,
    > "id": 6
    > }
