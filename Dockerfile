version: "3"

services:
  app:
    build: .
    container_name: appproduct
    ports:
      - "9000:9000"
    volumes:
      - .:/go/src/