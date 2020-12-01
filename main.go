package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
)

type createCheckoutSessionResponse struct {
	SessionID string `json:"id"`
}

func main() {
	stripe.Key = "sk_test_51Hpq4CIkuecL7QkUnVXbFdWYIskjwYpA8248lJ0sffdPBqslK7uzOScCG6RCXNhW3ZkziF7CJYEx65dVzAO9TZgf001jfrLSwH"

	var dir string
	flag.StringVar(&dir, "dir", ".", "./static")
	flag.Parse()
	r := mux.NewRouter()

	//Serve static html from localhost:8080/home
	r.PathPrefix("/home").Handler(http.StripPrefix("/home", http.FileServer(http.Dir(dir))))

	svr := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}

	log.Fatal(svr.ListenAndServe())
}

func createSession(w http.ResponseWriter, req *http.Request) {
	domain := "http://localhost:8080"
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(string(stripe.CurrencyUSD)),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("NVIDIA RTX 3090"),
					},
					UnitAmount: stripe.Int64(2000),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/success.html"),
		CancelURL:  stripe.String(domain + "/cancel.html"),
	}
	session, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
	}
	data := createCheckoutSessionResponse{
		SessionID: session.ID,
	}
	js, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
