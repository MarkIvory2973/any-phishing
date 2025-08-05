FROM golang:1.25rc2-alpine AS build

WORKDIR /src/

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build .

FROM alpine

WORKDIR /dist/

COPY --from=build /src/dbs/ ./dbs
COPY --from=build /src/backend .

CMD [ "./backend" ]