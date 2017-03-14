# Shopping Manager

## Basic info

A web aplication supporting concious shopping. The aplication take part in the competition "Get Noticed!“.

***

## Technology

### Backend
- [Go Programming Language](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)

### Frontend
- [Angular2](https://angular.io/)
- [Angular CLI](https://cli.angular.io/)

***

## Running

### Database

TODO

### Backend
- Open backend directory `cd api`
- Run go server `go run main.go`

#### The server is running:
```
http://localhost:8001
```

### Frontend
- Open frontend directory `cd web`
- Run web aplication `ng serve`

#### The application is running:
```
http://localhost:4200
```

Running unit tests: `ng test` via [Karma]

Running end-to-end tests: `ng e2e` via [Protractor]

***

## Endpoints list

#### Product

- `POST` `/product` Create product
- `PUT` `/product/{id}` Update product
- `GET` `/product` Get products list
- `GET` `/product/{id}` Get product by id
- `DELETE` `/product/{id}` Delete product
