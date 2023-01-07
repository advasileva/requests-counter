A simple training Go project, 
which consists of TG bot, http server and database

They are used in the expected way to count the number of requests from one user

To run project:
```bash
make update
```

`.env` configuration:
```
TOKEN=*your token*
POSTGRES_DB=postgres
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
SERVER_PORT=8080
```
