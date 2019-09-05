package handlers

import (
	"anonbot/domain"
	"anonbot/repositories"
	"anonbot/routing"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
)

func getValues(s string) url.Values {
	//m := map[string]string {}

	values, err := url.ParseQuery(s)
	if err != nil {
		log.Fatalln(err)
	}

	return values
}


func parseSlackQuery(values url.Values) *domain.SlackWebhookRequest {
	slackRequest := domain.SlackWebhookRequest{}
	t := reflect.TypeOf(&slackRequest)

	i := 0
	for i < t.Elem().NumField() {
		field := t.Elem().Field(i)

		tag := field.Tag.Get("json")

		v := reflect.ValueOf(&slackRequest)

		value := values.Get(tag)

		v.Elem().Field(i).Set(reflect.ValueOf(value))
		i++
	}

	fmt.Println(slackRequest)

	return &slackRequest
}

func postSlackCommandHandler(repository repositories.MessageRepository, parent context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, _ := GetDefaultTimeoutContext(parent)

		contentType := r.Header.Get("Content-Type")
		if contentType == "" {
			log.Printf("Content-Type was not supplied")
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := w.Write([]byte("Expected Content-Type to be application/x-www-form-urlencoded but was not set"))
			if err != nil {
				log.Println("ERROR - Writing to response in \"postSlackCommandHandler\": contentType == \"\"")
			}

			return
		}

		if contentType != "application/x-www-form-urlencoded" {
			log.Printf("Content-Type was not application/x-www-form-urlencoded")
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := w.Write([]byte(fmt.Sprintf("Expected Content-Type to be application/x-www-form-urlencoded but was %v", contentType)))
			if err != nil {
				log.Println("ERROR - Writing to response in \"postSlackCommandHandler\": contentType != application/x-www-form-urlencoded")
			}

			return
		}

		if _, ok := ctx.Deadline(); !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}


		//q := r.Form

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}

		bodyValues := getValues(string(body))

		_ = parseSlackQuery(bodyValues)
		//fmt.Println(slackQuery)

		_, err = w.Write([]byte(""))
		if err != nil {
			log.Printf("ERROR - Writing response. %v\n", err)
		}
	}
}

func RegisterSlackWebhookHandler(repository repositories.MessageRepository, router routing.Router, parent context.Context) {
	router.RegisterRoute(http.MethodPost, "/webhooks/slack", postSlackCommandHandler(repository, parent))
}