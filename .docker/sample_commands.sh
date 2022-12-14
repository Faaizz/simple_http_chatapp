export HOST_ADDR="localhost"
export HOST_PORT="81"
# connect
curl http://$HOST_ADDR:$HOST_PORT/connect -d '{"connectionId": "91123456"}' --header "Content-Type:application/json"
# set username
curl http://$HOST_ADDR:$HOST_PORT/username -d '{"connectionId": "91123456", "username": "Asiat"}' --header "Content-Type:application/json"
# disconnect
curl http://$HOST_ADDR:$HOST_PORT/disconnect -d '{"connectionId": "91123456"}' --header "Content-Type:application/json"
# list users
curl http://$HOST_ADDR:$HOST_PORT/online -d '{"connectionId": "91123456"}' --header "Content-Type:application/json"
# message
curl http://$HOST_ADDR:$HOST_PORT/message -d '{"connectionId": "911234567", "username": "Asiat33", "message": "Hello WebSocket", "url": ""}' --header "Content-Type:application/json"
# health
curl http://$HOST_ADDR:$HOST_PORT/healthz --header "Content-Type:application/json"
