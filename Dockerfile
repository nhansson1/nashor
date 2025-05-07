FROM node:24-alpine AS client

WORKDIR app/client

COPY client/package*.json ./

RUN npm install

COPY client .

RUN npm run build

FROM golang:1.24.0-alpine

WORKDIR app/server 

COPY server/go.mod server/go.sum ./
RUN go mod download

COPY server .

COPY --from=client /app/client/dist ./internal/dist

RUN go build -o ./bin/nashor ./cmd/nashor

EXPOSE 8080

CMD ["./bin/nashor"]
