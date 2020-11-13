# Golang REST API with Gin and Gorm

## List Person
Request :
- Method: GET
- Endpoint: `/persons`
- Header :
  - Accept : application/json

Response :
```json
{
  "count": "integer",
  "result": [
    {
      "id": "integer",
      "FirstName": "string",
      "LastName": "string",
      "CreatedAt": "string",
      "UpdatedAt": "string",
      "DeletedAt": "string"
    }
  ]
}
```

## Get Person
Request :
- Method : GET
- Endpoint : `/person/{id_person}`
- Header :
  - Accept : application/json

Response :
```json
{
  "count": "integer",
  "result": {
    "id": "integer",
    "FirstName": "string",
    "LastName": "string",
    "CreatedAt": "string",
    "UpdatedAt": "string",
    "DeletedAt": "string"
  }
}
```

## Create Person
Request :
- Method : POST
- Endpoint : `/person`
- Header :
    - Accept : application/json
    - Content-Type : application/json
- Body :
    - first_name : string
    - last_name : string

Response :

```json
{
  "message": "string",
  "result": {
    "id": "integer",
    "FirstName": "string",
    "LastName": "string",
    "CreatedAt": "string",
    "UpdatedAt": "string",
    "DeletedAt": "string"
  }
}
```

## Update Person
Request :
- Method : PUT
- Endpoint : `/person/{id_person}`
- Header :
    - Accept : application/json
    - Content-Type : application/json
- Body :
    - first_name : string
    - last_name : string

Response :

```json
{
  "result": "string",
}
```

## Delete Person
Request :
- Method : DELETE
- Endpoint : `/person/{id_person}`
- Header : 
    - Accept : application/json

Response :
```json
{
  "result": "string",
}
```