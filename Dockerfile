FROM golang:alpine AS go-builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY app . 
COPY .config . 

RUN CGO_ENABLED=0 GOOS=linux go build -a -o bin/tv-tracker app/*.go

FROM node:alpine

WORKDIR /app

COPY vue/package*.json ./

RUN npm install

COPY vue .

COPY --from=go-builder /app/bin/tv-tracker .

RUN npm run build

EXPOSE 8080

CMD ["npm", "start", "--", "./tv-tracker"]