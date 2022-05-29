# Demo on `docker-compose`
Using simple ReactJs web application and Golang rest api

## Running docker images and containers seperately
```
# Clone repository
git clone git@github.com:viraj-lakshitha/demo-docker-compose
cd docker-compose

# Create docker images for FE and BE
cd demo-be
docker build -t demo-be .

cd ../

cd demo-fe
docker build -t demo-fe .

# Run docker images in daemon mode and expose ports
docker run -it -d -p 3000:3000 demo-fe
docker run -it -d -p 80:80 demo-be
```

## Run with docker-compose
```
cd /demo-docker-compose
docker-compose up --build
```
