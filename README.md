# # Maximum Path Sum of a Binary Tree

This project provides a service to calculate the sum of numbers provided in a JSON payload.

## Structure:

- `types/`: Contains data structures.
- `services/`: Contains business logic.
- `handlers/`: Contains HTTP handlers.

## Running the Service:

1. `go run main.go -port 8080`
2. Make a POST request to `http://localhost:8080/calculate`.

## Testing:

Run `go test ./...`.

## Testing with `curl`

You can test the service using `curl` to send the provided payload:

```bash
curl -X POST http://localhost:8080/calculate \
     -H "Content-Type: application/json" \
     -d '{
         "tree": {
             "nodes": [
                 {"id": "1", "left": "2", "right": "3", "value": 1},
                 {"id": "3", "left": null, "right": null, "value": 3},
                 {"id": "2", "left": null, "right": null, "value": 2}
             ],
             "root": "1" 
         }
     }'
