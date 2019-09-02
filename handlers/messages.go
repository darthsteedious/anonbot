package handlers

import (
	"anonbot/domain"
	"anonbot/repositories"
	"anonbot/routing"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

func getMessage(repository repositories.MessageRepository, parent context.Context) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		ctx, _ := GetDefaultTimeoutContext(parent)

		id, err := strconv.ParseInt(vars["id"], 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte("Path token \"id\" must be an integer"))
			if err != nil {
				log.Println("ERROR - Writing to response in \"getMessage\": strconv.ParseInt")
			}

			return
		}

		message, err := repository.GetMessage(int(id), ctx)
		if err != nil {
			log.Printf("ERROR - Getting message from database. %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte("Internal server error\n"))
			if err != nil {
				log.Println("ERROR - Writing to response in \"getMessage\": repository.GetMessage")
			}

			return
		}

		encoder := json.NewEncoder(w)
		w.Header().Add("Content-Type", "application/json")
		err = encoder.Encode(message)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("ERROR - Writing message to response. %v\n", err)
			_, err := w.Write([]byte("Internal Server Error"))
			if err != nil {
				log.Println("ERROR - Writing to response in \"getMessage\": encoder.Encode")
			}
		}
	}
}

func postMessage(repository repositories.MessageRepository, parent context.Context) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		ctx, _ := GetDefaultTimeoutContext(parent)

		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := w.Write([]byte("Expected JSON content type"))
			if err != nil {
				log.Println("ERROR - Writing to response in \"postMessage\": contentType != \"application/json\"")
			}

			return
		}

		message := &domain.Message{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(message)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			_, err := w.Write([]byte("Unable to deserialize message body\n"))
			if err != nil {
				log.Println("ERROR - Writing to response in \"postMessage\"")
			}

			return
		}

		now := time.Now().UTC()
		message.ReceivedAt = &now

		log.Printf("DEBUG - Message: %v", message)

		insertedId, err := repository.InsertMessage(message, ctx)
		if err != nil {
			log.Printf("Error inserting message. %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("Internal server error\n"))
			if err != nil {
				log.Println("ERROR - Writing to response in \"postMessage\": repository.InsertMessage")
			}

			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Location", fmt.Sprintf("/api/v1/messages/%v", insertedId))
	}
}

func putMessageDelivered(repository repositories.MessageRepository, parent context.Context) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx, _ := GetDefaultTimeoutContext(parent)

		parsedId, err := strconv.ParseInt(vars["id"],  10, 32)
		if err != nil {
			log.Println("ERROR - Error parsing route token \"id\"")
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte("Token \"id\" must be an integer\n"))
			if err != nil {
				log.Println("ERROR - Writing to response in \"putMessageDelivered\": strconv.ParseInt")
			}

			return
		}

		id := int(parsedId)
		_, err = repository.GetMessage(id, ctx)
		if err != nil {
			log.Printf("ERROR - Error getting message from repository. %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Internal server error"))
			if err != nil {
				log.Printf("ERROR - Writing to response in \"putMessageDelivered\": repository.GetMessage")
			}
		}

		err = repository.SetDelivered(id, ctx)
		if err != nil {
			log.Printf("ERROR - Setting message delivered. %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("Internal server error\n"))
			if err != nil {
				log.Println("ERROR - Writing to response in \"putMessageDelivered\": repository.SetDelivered")
			}

			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func RegisterMessageHandlers(repository repositories.MessageRepository, router routing.Router, parent context.Context) {
	router.RegisterRoute(http.MethodGet, "/messages/{id}", getMessage(repository, parent))
	router.RegisterRoute(http.MethodPost, "/messages", postMessage(repository, parent))
	router.RegisterRoute(http.MethodPut, "/messages/{id}/delivered", putMessageDelivered(repository, parent))
}