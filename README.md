# API Server

Run

```console
./setupdb.sh
./runserver.sh
```

## Example queries
### GET USERS BY ROLE
those require field Authorization:"IMADMIN" in the header
```
GET http://localhost:8080/admin/users           //Default role_id=1, page_num=1, page_size=5
GET http://localhost:8080/admin/users?role_id=1
```
Response 
```json
{
    "links": {
        "next": {
            "href": "/admin/users?page_num=2"
        }
    },
    "users": [
        {
            "id": 4,
            "username": "user1",
            "role": "пользователь"
        },
        {
            "id": 5,
            "username": "user2",
            "role": "пользователь"
        },
        {
            "id": 6,
            "username": "user3",
            "role": "пользователь"
        },
        {
            "id": 7,
            "username": "user4",
            "role": "пользователь"
        },
        {
            "id": 8,
            "username": "user5",
            "role": "пользователь"
        }
    ]
}
```


### DELETE BY ID
those require field Authorization:"IMADMIN" in the header
admin cannot be deleted, only users
```
DELETE http://localhost:8080/admin/users/4
```
Response 
```json
{
    "deleted": 1
}
```


### GET BY ID
```
GET http://localhost:8080/users/4
```
Response 
```json
{
    "user": {
        "id": 6,
        "username": "user3"
    }
}
```

### CREATE USER
```
POST http://localhost:8080/users
{
    "username": "user10",
    "password": "password"
}
```
Response 
```json
{
    "user": {
        "id": 13,
        "username": "user121132"
    }
}
```




