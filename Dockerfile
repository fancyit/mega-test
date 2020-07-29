# We specify the base image we need for our
# go application
FROM golang:alpine
# We create an /app directory within our
# image that will hold our application source
# files
RUN apk add --no-cache git

RUN mkdir /app
# We copy everything in the root directory
# into our /app directory
ADD . /app
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /app

# we run go build to compile the binary
# executable of our Go program
RUN go get -d -v github.com/gorilla/mux
RUN go get -d -v github.com/joho/godotenv

#RUN go install -v source/helpers
#RUN go install -v source/controllers

RUN go build -o mega-test .
# Our start command which kicks off
# our newly created binary executable
CMD ["/app/mega-test"]
