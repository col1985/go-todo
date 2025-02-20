# Go Todo App

## Prerequites

### Storage

[PostgreSQL](https://www.postgresql.org/) is used as the peristence layer, in order to run the application on your machine ensure you have [PostgreSQL](https://www.postgresql.org/) installed and running.

### Environment Variables

The application requires a `.env` with the following variables.

```bash
API_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=user
DB_NAME=todos
DB_PASSWORD=password
```

## Development with Podman

### Running app in development environment

Run the app using the command below pass thr `IS_DEV` env var so that the `.env` contents are loaded as expected.

```bash
IS_DEV=true go run .
```

### Building the app using a Containerfile

```bash
podman build \
--build-arg=CGO_ENABLED=0 \
--build-arg=GOOS=linux \
--build-arg=GOARCH=amd64 \
-t go-todo:<version /> .
```

### Running the app using Podman

#### Using a `.env` file

```bash
podman run --name=go-todo --env-file=.env -p 8080:8080 localhost/go-todo:<version />
```

#### Passing Env var as parameters

```bash
  podman run --name=go-todo \
  -e API_PORT=8080 \
  -e DB_HOST=localhost \
  -e DB_PORT=5432 \
  -e DB_USER=user \
  -e DB_PASSWORD=password \
  -e DB_NAME=todos \
  -p 8080:8080 localhost/go-todo:v1.0.3
```

### Create app and database using the OC Client in Openshift

You should have created the targetted namespace before running the following commands.

#### Create database

```bash
oc new-app \
  --name todo-db \
  --image=registry.redhat.io/rhel8/postgresql-12:1-177 \
  -e POSTGRESQL_USER=user \
  -e POSTGRESQL_DATABASE=todos \
  -e POSTGRESQL_PASSWORD=password
```

#### Create Todo App

```bash
oc new-app --name todo-api --image=quay.io/cbennett/go-todo:v1.0.3 \
-e API_PORT=8080 \
-e DB_HOST=<address of postgres database /> \
-e DB_PORT=5432 \
-e DB_USER=user \
-e DB_PASSWORD=password \
-e DB_NAME=todos
```

### Running Postgres PGAdmin with Podman

```bash
podman run \
--name postgresdb \
-e POSTGRES_USER=user \
-e POSTGRES_PASSWORD=password \
-p 5432:5432 \
-v ~/dev/postgres/data \
-d postgres
```

### Running PGAdmin UI with Podman

Open the userr interface in your browser at `http:localhost:5050` and use the credentials you have set.

```bash
podman run \
--name pgadmin \
-e 'PGADMIN_DEFAULT_EMAIL=user@example.com' \
-e 'PGADMIN_DEFAULT_PASSWORD=topsecret' \
-p 5050:80 \
-d docker.io/dpage/pgadmin4
```

## Testing the API

### Get List

```bash
curl http://localhost:8080/todos | jq
```

### Get Todo using ID

```bash
curl http://localhost:8080/todos/{id} | jq
```

### Create Todo

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"task":"My todo task","author":"Colum B"}' \
  http://localhost:3000/todos
```

### Update Todo

```bash
curl -X PUT \
  --data '{"task":"My todo task that is updated","author":"Colum B"}' \
  http://localhost:8080/todos/{id} | jq
```

### Delete Todo

```bash
curl -X DELETE http://localhost:8080/todos/{id}
```
