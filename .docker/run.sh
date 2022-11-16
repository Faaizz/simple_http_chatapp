if [ -f .env ]
then
  export $(cat .env | xargs)
fi

# export DB_TYPE="MONGODB"
export DB_TYPE="DYNAMODB"

export DYNAMODB_ENDPOINT_URL="http://localhost:8000"


export HTTP_PORT="81"

cd ../src

docker-compose -f ../.docker/docker-compose.yml up -d
# wait 5 seconds to create DB table
sleep 5
go run .
docker-compose -f ../.docker/docker-compose.yml down
