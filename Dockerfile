FROM golang
WORKDIR /robin
COPY . .
CMD [ "make","run/api"]