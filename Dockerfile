# build frontend
FROM node:20.11.1-alpine3.18 as frontend-build

WORKDIR /web

COPY web/package*.json ./

RUN npm install

COPY /web .

RUN npm run build

# build backend and run
FROM golang:1.22.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./
COPY --from=frontend-build /web/dist ./web/dist

RUN CGO_ENABLED=0 GOOS=linux go build -o /nostalgie_replay_finder

EXPOSE 8080

CMD ["/nostalgie_replay_finder"]