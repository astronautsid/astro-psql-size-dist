# go-psql-size
A program to get the size of tables in PostgreSQL tables' based on configured label

# Setup
## Install dependencies
```
$ go mod tidy
```

## Compiling
```
$ go build cmd/*.go
```

# Configuration
Create a file in `config/config.yaml`
Fill the file with the following content:
```
username: $YOUR_PSQL_USERNAME
password: $YOUR_PSQL_PASSWORD
host: $YOUR_PSQL_HOST
port: $YOUR_PSQL_PORT
dbname: $YOUR_PSQL_DBNAME

```

# Running
To run the binary, execute this following line:
```
$ ./main config/config.yaml config/rule.json ouput.json
```
