package natsstream

import (
	"encoding/json"
	"fmt"
	"github.com/dingowd/WB/L0/app"
	"github.com/dingowd/WB/L0/model"
	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
	"log"
	"os"
	"time"
)

type NatsStream struct {
	ClusterID   string
	ClientID    string
	URL         string
	UserCreds   string
	ShowTime    bool
	Qgroup      string
	Unsubscribe bool
	StartSeq    uint64
	StartDelta  string
	DeliverAll  bool
	NewOnly     bool
	DeliverLast bool
	Durable     string
	Subj        string
	NC          *nats.Conn
	SC          stan.Conn
	Sub         stan.Subscription
	App         *app.App
}

func NewSub(s NatsStream) *NatsStream {
	return &s
}

func (n *NatsStream) MsgHandler(msg *stan.Msg) {
	var o model.Order
	json.Unmarshal(msg.Data, &o)
	s := fmt.Sprintln("Received: ", o)
	n.App.Log.Info(s)
	err := n.App.Store.CreateOrder(o)
	if err != nil {
		n.App.Log.Info(err.Error())
	}
}

func (n *NatsStream) Start() {
	var err error

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Streaming WB Subscriber")}
	// Use UserCredentials
	if n.UserCreds != "" {
		opts = append(opts, nats.UserCredentials(n.UserCreds))
	}

	// Connect to NATS
	n.NC, err = nats.Connect(n.URL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	//defer nc.Close()
	n.SC, err = stan.Connect(n.ClusterID, n.ClientID, stan.NatsConn(n.NC),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			s := fmt.Sprintf("Connection lost, reason: %v", reason)
			n.App.Log.Error(s)
			os.Exit(1)
		}))
	if err != nil {
		s := fmt.Sprintf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, n.URL)
		n.App.Log.Error(s)
		os.Exit(1)
	}
	s := fmt.Sprintf("Connected to %s clusterID: [%s] clientID: [%s]\n", n.URL, n.ClusterID, n.ClientID)
	n.App.Log.Info(s)

	// Process Subscriber Options.
	startOpt := stan.StartAt(pb.StartPosition_NewOnly)
	if n.StartSeq != 0 {
		startOpt = stan.StartAtSequence(n.StartSeq)
	} else if n.DeliverLast {
		startOpt = stan.StartWithLastReceived()
	} else if n.DeliverAll && !n.NewOnly {
		startOpt = stan.DeliverAllAvailable()
	} else if n.StartDelta != "" {
		ago, err := time.ParseDuration(n.StartDelta)
		if err != nil {
			n.SC.Close()
			n.App.Log.Error(err.Error())
			os.Exit(1)
		}
		startOpt = stan.StartAtTimeDelta(ago)
	}

	/*	subj, i := n.Subj, 0
		mcb := func(msg *stan.Msg) {
		i++
		printMsg(msg, i) //TODO
	}*/

	n.Sub, err = n.SC.QueueSubscribe(n.Subj, n.Qgroup, n.MsgHandler, startOpt, stan.DurableName(n.Durable)) //TODO MsgHandler
	if err != nil {
		n.SC.Close()
		n.App.Log.Error(err.Error())
		os.Exit(1)
	}

	s = fmt.Sprintf("Listening on [%s], clientID=[%s], qgroup=[%s] durable=[%s]\n", n.Subj, n.ClientID, n.Qgroup, n.Durable)
	n.App.Log.Info(s)

	if n.ShowTime {
		log.SetFlags(log.LstdFlags) //TODO
	}
	//cleanupDone := make(chan bool)

	/*	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
		// Run cleanup when signal is received
		signalChan := make(chan os.Signal, 1)
		cleanupDone := make(chan bool)
		signal.Notify(signalChan, os.Interrupt)
		go func() {
			for range signalChan {
				fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
				// Do not unsubscribe a durable on exit, except if asked to.
				if n.Durable == "" || n.Unsubscribe {
					sub.Unsubscribe()
				}
				n.SC.Close()
				cleanupDone <- true
			}
		}()
		<-cleanupDone*/
}

func (n *NatsStream) Stop() {
	if n.Durable == "" || n.Unsubscribe {
		n.Sub.Unsubscribe()
	}
	n.SC.Close()
	n.NC.Close()
}
