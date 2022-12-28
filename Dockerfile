FROM golang:1.19.0-alpine as build
WORKDIR /src
COPY . .
RUN go build -o /out/example/ .
FROM golang:1.19.0-alpine as sample
WORKDIR /usr/local/bin/app
COPY --from=build /out/example/simple-go-project ./
ENV PATH "$PATH:/usr/local/bin/app"
CMD [ "simple-go-project" ]