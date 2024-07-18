##
## Build
##
FROM golang:1.17-alpine AS build

WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./
COPY . .
RUN go mod download

COPY *.go ./

ENV CGO_ENABLED=0
RUN go build -o /Go-Instagram

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /instagram /instagram

EXPOSE 8080
EXPOSE 5432

USER nonroot:nonroot

ENTRYPOINT ["/instagram"]
# unzip /home/mahdi/Downloads/build-tools_r33-linux.zip -d build-tools/33.0.0 && cd build-tools/33.0.0 && mv android-*/* . && rm -rf android-*

km6585534
55967
sudo killall openconnect

#sudo openconnect -script /etc/vpnc/vpnc-script  c1.kmak.us:443

sudo openconnect  –script=/etc/vpnc/vpnc-script <c1.kmak.us:443>
sudo openconnect  –s=/etc/vpnc/vpnc-scrip c1.kmak.us:443

#docker build -t mahdiabdolmaleki/Go-Instagram:latest -f Dockerfile .
#docker push mahdiabdolmaleki/Go-Instagram
#docker run -d --restart=always  --name Go-Instagram -p 8080:8080 --network postgres_network mahdiabdolmaleki/Go-Instagram:latest

#docker run  -d --name PostgreSQL  --restart=always  -p 5432:5432  --network postgres_network  -e POSTGRES_USER=mahdi  -e POSTGRES_PASSWORD=aliali  -v "$PWD/postgresdb":/var/lib/postgresql/data  postgres


