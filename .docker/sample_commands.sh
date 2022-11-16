export HOST_ADDR="localhost"
export HOST_PORT="81"
# connect
curl http://$HOST_ADDR:$HOST_PORT/connect -d '{"connectionId": "91123456", "username": "Aishatutu"}' --header "Content-Type:application/json"
# disconnect
curl http://$HOST_ADDR:$HOST_PORT/disconnect -d '{"connectionId": "91123456", "username": "Aishatutu"}' --header "Content-Type:application/json"
# list users
curl http://$HOST_ADDR:$HOST_PORT/online -d '{"connectionId": "91123456", "username": "Aishatutu"}' --header "Content-Type:application/json"
# message
curl http://localhost:$HOST_PORT/message -d '{"connectionId": "91123456", "from_username": "Aishatu"}' --header "Content-Type:application/json"
