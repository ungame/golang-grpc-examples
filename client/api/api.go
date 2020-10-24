package api

import (
	"encoding/json"
	"fmt"
	"golang-grpc-examples/messages/messenger"
	"io/ioutil"
	"net/http"
)

type messengerAPI struct {
	client messenger.MessengerServiceClient
}

func NewMessengerAPI(client messenger.MessengerServiceClient) *messengerAPI {
	return &messengerAPI{client: client}
}

func (a *messengerAPI) Handlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.getMessage(w, r)
	case http.MethodPost:
		a.postMessage(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method not implemented: %v", r.Method)
	}
}

func (a *messengerAPI) postMessage(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("body can't be empty"))
		return
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	message := new(messenger.Message)
	err = json.Unmarshal(b, message)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	resp, err := a.client.WriteMessage(r.Context(), message)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	json.NewEncoder(w).Encode(resp)

}

func (a *messengerAPI) getMessage(w http.ResponseWriter, r *http.Request) {
	resp, err := a.client.ReadMessage(r.Context(), &messenger.Empty{})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	json.NewEncoder(w).Encode(resp)
}
