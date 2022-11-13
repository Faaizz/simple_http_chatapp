export DB_TYPE="DYNAMODB"

export HTTP_PORT="81"
export TABLE_NAME="my_ddb_table_$RANDOM"

cd ../src
go run .
