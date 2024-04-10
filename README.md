# Simple REST API 

### Система состоит из:
1) Publisher
2) Reader
3) Retrieval
4) Proxy Adapter
5) Cron
6) JetStream Nats
7) PostgreSQL

### Чтобы запустить систему необходимо воспользоваться командой `make`

   
### Примеры запросов: 
POST
```localhost:8000/employees/```

```
{
    "id": 1936,
    "firstName": "Nikita",
    "lastName": "Chereshnev",
    "position": "developer",
    "department": "dev",
    "hireDate": "20.04.2024",
    "salary": 199999,
    "email": "chereshnev.n.s@gmail.com",
    "phoneNumber": "89191113882",
    "address": "Pushkina 1, 20"
}
``` 

GET
```localhost:8000/employees/1936/```
