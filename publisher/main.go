// Copyright 2016-2019 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var usageStr = `
Usage: stan-pub [options] <subject> <message>

Options:
	-s,  --server   <url>            NATS Streaming server URL(s)
	-c,  --cluster  <cluster name>   NATS Streaming cluster name
	-id, --clientid <client ID>      NATS Streaming client ID
	-a,  --async                     Asynchronous publish mode
	-cr, --creds    <credentials>    NATS 2.0 Credentials
`

func send(m string) {
	var (
		clusterID string
		clientID  string
		URL       string
		async     bool
		userCreds string
	)

	clusterID = "WB_Cluster"
	clientID = "stan-pub"
	args := make([]string, 2)
	args[0] = "order"
	args[1] = m
	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Streaming Example Publisher")}
	// Use UserCredentials
	if userCreds != "" {
		opts = append(opts, nats.UserCredentials(userCreds))
	}

	// Connect to NATS
	nc, err := nats.Connect(URL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, URL)
	}
	defer sc.Close()

	subj, msg := args[0], []byte(args[1])

	ch := make(chan bool)
	var glock sync.Mutex
	var guid string
	acb := func(lguid string, err error) {
		glock.Lock()
		log.Printf("Received ACK for guid %s\n", lguid)
		defer glock.Unlock()
		if err != nil {
			log.Fatalf("Error in server ack for guid %s: %v\n", lguid, err)
		}
		if lguid != guid {
			log.Fatalf("Expected a matching guid in ack callback, got %s vs %s\n", lguid, guid)
		}
		ch <- true
	}

	if !async {
		err = sc.Publish(subj, msg)
		if err != nil {
			log.Fatalf("Error during publish: %v\n", err)
		}
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	} else {
		glock.Lock()
		guid, err = sc.PublishAsync(subj, msg, acb)
		if err != nil {
			log.Fatalf("Error during async publish: %v\n", err)
		}
		glock.Unlock()
		if guid == "" {
			log.Fatal("Expected non-empty guid to be returned.")
		}
		log.Printf("Published [%s] : '%s' [guid: %s]\n", subj, msg, guid)

		select {
		case <-ch:
			break
		case <-time.After(5 * time.Second):
			log.Fatal("timeout")
		}

	}
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	send(string(b))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/send", sendHandler).Methods("POST")
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
