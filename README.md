# Go Todo App

## Prerequites

### Storage

[PostgreSQL](https://www.postgresql.org/) is used as the peristence layer, in order to run the application on your machine ensure you have [PostgreSQL](https://www.postgresql.org/) installed and running.

### Environment Variables

The application requires a `.env` with the following variables.

```bash
API_PORT=8080

DB_HOST=
DB_PORT=
DB_USER=
DB_NAME=
DB_PASSWORD=
```

## Development with Podman

### Building the app using a Containerfile

```bash
podman build \
--build-arg=CGO_ENABLED=0 \
--build-arg=GOOS=linux \
--build-arg=GOARCH=amd64 \
-t go-todo:<version /> .
```

### Running the app using Podman

```bash
podman run --name=go-todo --env-file=.env -p 8080:8080 localhost/go-todo:<version />
```

```bash
podman run --name=go-todo \
--pod=go-todo-pod \
-e API_PORT=8080 \
-e DB_HOST=localhost \
-e DB_PORT=5432 \
-e DB_USER=username \
-e DB_PASSWORD=password \
-e DB_NAME=tasks \
-p 8080:8080 localhost/go-todo:<version />
```

```bash
oc new-app --name todo-api --image=quay.io/cbennett/go-todo:v1.0.3 \
-e API_PORT=8080 \
-e DB_HOST= \
-e DB_PORT=5432 \
-e DB_USER=user \
-e DB_PASSWORD=password \
-e DB_NAME=tasks
```
### Running Postgres PGAdmin with Podman

```bash
oc new-app \
  --name todo-db-postgresql \
  --image=registry.redhat.io/rhel8/postgresql-12:1-177 \
  -e POSTGRESQL_USER=user \
  -e POSTGRESQL_DATABASE=db \
  -e POSTGRESQL_PASSWORD=password
```

```bash
podman run \
--pod=go-todo-pod \
--name postgresdb \
-e POSTGRES_USER=username \
-e POSTGRES_PASSWORD=password \
-p 5432:5432 \
-v ~/dev/postgres/data \
-d postgres
```

```bash
podman run --pod=go-todo-pod \
--name=go-todo \
--env-file=.env localhost/go-todo:v1.0.0
```

```bash
podman run --pod=go-todo-pod \
--name postgresdb \
-e POSTGRES_USER=username \
-e POSTGRES_PASSWORD=password \
-v ~/dev/postgres/data \
-d postgres
```

```bash
podman run --pod=go-todo-pod \
-e 'PGADMIN_DEFAULT_EMAIL=user@example.com' \
-e 'PGADMIN_DEFAULT_PASSWORD=topsecret' \
--name pgadmin \
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
