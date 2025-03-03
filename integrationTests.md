Create and run a new PostgreSQL database inside a Docker container
test -db is the container name(we can start/stop it later)
-e POSTGRES_USER=test will set the database username to test
-e POSTGRES_PASSWORD=test Stes password to test
-e POSTGRES_DB=testdb Creates a test Database names testdb
-p 5433.5432 Maps port 5433 on your local machine to PostgreSQL's default port(5432) inside container
-d Runs in detached mode
docker run --name test-db -e POSTGRES_USER=test -e POSTGRES_PASSWORD=test -e POSTGRES_DB=testdb -p 5433:5432 -d postgres

After tests finish, stop the container
docker stop test-db

Can also remove container entirely
docker rm -f test-db

Running tests with Docker PostgreSQL
1. Start test PostgreSQL database if not already running
   docker start test-db
2. Run Go tests:
   go test ./...
3. Stop test database when done:
   docker stop test-db