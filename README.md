# Simple file upload server in Golang

## How to build?

Run `go build`

## APIs

| API | Description |
|--|--|
| `GET /` | Echos the data and headers back |
| `POST /echo` | Echos only the data back |
| `POST /error` | Replies with a HTTP 400 and a JSON error message |
| `POST /upload` | Saves file under `uploads` folder with the name specified by the header `Odp-File-Name`. The contents of the file are echoed back with `Content-type: application/octet-stream` |
| `POST /upload2` | Saves file under `uploads2` folder with the name specified by the header `Odp-File-Name`. Returns JSON data that has the filename in it with the format `{"fileName": <name of the file>}` |
