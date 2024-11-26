FROM golang:1.23

# sets the working directory inside the container to /app and creates it if it doesn't exist
WORKDIR /app

# copy the application files
COPY . .

# build the Go application with optimizations to reduce binary size
RUN go build -o groupie .

# add labels for metadata
LABEL maintainer="mzakri <mostafazakri3@gmail.com> || helbadao <elbadaouihoussam4@gmail.com>"
LABEL description="Groupie-Tracker Project"
LABEL version="1.0"

# default port for the application
EXPOSE 8080

# default command to run the application
CMD ["./groupie"]
