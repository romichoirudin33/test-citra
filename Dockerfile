FROM golang:1.14.3

COPY . /app/

CMD ["go", "run", "/app/main.go"]