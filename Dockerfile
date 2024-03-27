FROM golang:latest

COPY . .

RUN go install github.com/go-task/task/v3/cmd/task@latest
RUN task install
RUN task build

EXPOSE 3000

CMD [ "./dist/main" ]