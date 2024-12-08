# ASCII Art Web Export File

## Description
ASCII Art Web Export File is a web server that lets you generate ASCII art from text input using different banners (fonts). Users can input text, select a banner style, ASCII color and see the ASCII art representation on the web. They can also export the produced ASCII art into:
- Text
- XML
- JSON
- html

## Authors
- Marinos Kouvaras
- Ypatios Chaniotakos

## Usage: how to run
To run the server:

```bash
cd app
go run main.go
```

## Implementation Details
The server uses Go's net/http package to handle requests. ASCII art generation is based on ASCII banner fonts. Each font is stored in a text file and read dynamically based on user selection. On the execution algorithm parses the line number for each character of the text and after it builds an eight (8) line ascii art forming the text inserted.

## HTTP status codes examples
```
<curl -i http://localhost:8080/> //Status 200 OK
<curl -i http://localhost:8080/gr> //Status 404 Page not found
<curl -X POST -d "banner=standard" http://localhost:8080/ascii-art -i> //Status 400 Bad Request
<curl -X POST -d "text=Hello&banner=heroe" http://localhost:8080/ascii-art -i> //Internal Server Error, for unhandled errors.
```

## API
- ui
```
<http://localhost:8080/static/api/api.html>
```
- CLI
```
curl -X POST -H "Content-Type: application/json" \
-d '{"text":"asd", "banner":"standard"}' \
http://localhost:8080/api/ascii-art
```

## Docker
The application can be easily deployed to a docker using the Dockerfile