FROM golang:1.15

# Meta data:
LABEL maintainer="email@mattglei.ch"
LABEL description="🌲 golang port of Delgan's python loguru"

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

CMD ["make", "local-test"]
