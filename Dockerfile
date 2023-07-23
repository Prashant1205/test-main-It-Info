FROM golang:1.19-alpine


RUN mkdir -p /ltinfo

WORKDIR /ltinfo
COPY . .

RUN apk update && apk add --no-cache git
RUN go mod tidy

RUN go build /ltinfo

EXPOSE 8080
ENTRYPOINT [ "./ltinfo" ]
