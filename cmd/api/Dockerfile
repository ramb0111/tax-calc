FROM golang

COPY . /go/src/github.com/ramb0111/tax-calc
WORKDIR /go/src/github.com/ramb0111/tax-calc
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/gorilla/handlers"
RUN go get "github.com/lib/pq"
RUN go get "github.com/pkg/errors"
CMD [ "go","run","cmd/api/api.go"]

