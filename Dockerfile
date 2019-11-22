FROM golang:alpine AS build
WORKDIR /src
ADD . .
WORKDIR /src
RUN go build -o numfiles_exporter

FROM alpine
WORKDIR /bin
COPY --from=build /src/numfiles_exporter .
COPY --from=build /src/targets.yaml /opt/numfiles_exporter/
CMD ["numfiles_exporter"] # specify the required flags!
