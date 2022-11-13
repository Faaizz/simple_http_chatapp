export MONGODB_USER="root"
export MONGODB_PASSWORD="password"
export MONGODB_HOST="localhost"
export MONGODB_PORT="27017"
export MONGODB_DATABASE="chat"

export DB_TYPE="MONGODB"

export HTTP_PORT="81"
export TABLE_NAME="my_ddb_table_27184"

docker-compose up -d && cd ../src
go run .
cd ../.docker && docker-compose down
