# Shopping Manager

## Basic info

A web aplication supporting concious shopping. The aplication take part in the competition "Get Noticed!“.

***

## Technology

### Database
- [MongoDB](https://www.mongodb.com/)

### Backend
- [Go Programming Language](https://golang.org/)

### Frontend
- [Angular2](https://angular.io/)
- [Angular CLI](https://cli.angular.io/)

***

## Instalation and running

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
- Before first runtime use `npm install`
- Run web aplication `ng serve` or simple `ng s`

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
