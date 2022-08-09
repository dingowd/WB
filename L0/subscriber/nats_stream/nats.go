package natsstream

import (
	"fmt"
	"github.com/dingowd/WB/L0/app"
	"github.com/dingowd/WB/L0/internal/utils"
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
	//var o *model.Order
	o, err := utils.Validate(msg.Data)
	if err != nil {
		n.App.Log.Error(err.Error())
		return
	}
	s := fmt.Sprintln("Получено: ", o)
	n.App.Log.Info(s)
	errC := n.App.Store.CreateOrder(*o)
	if errC != nil {
		n.App.Log.Error(errC.Error())
		return
	}
	n.App.Cache.WriteToCache(*o)
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
	// Connect to STAN
	n.SC, err = stan.Connect(n.ClusterID, n.ClientID, stan.NatsConn(n.NC),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			s := fmt.Sprintf("Соединение потеряно, причина: %v", reason)
			n.App.Log.Error(s)
			os.Exit(1)
		}))
	if err != nil {
		s := fmt.Sprintf("Невозможно установить соединение: %v.\nПроверьте, что NATS Streaming Server запущен по адресу: %s", err, n.URL)
		n.App.Log.Error(s)
		os.Exit(1)
	}
	s := fmt.Sprintf("Установлено соединение с %s clusterID: [%s] clientID: [%s]", n.URL, n.ClusterID, n.ClientID)
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

	n.Sub, err = n.SC.QueueSubscribe(n.Subj, n.Qgroup, n.MsgHandler, startOpt, stan.DurableName(n.Durable)) //TODO MsgHandler
	if err != nil {
		n.SC.Close()
		n.App.Log.Error(err.Error())
		os.Exit(1)
	}

	s = fmt.Sprintf("Ждем сообщений [%s], clientID=[%s], qgroup=[%s] durable=[%s]\n", n.Subj, n.ClientID, n.Qgroup, n.Durable)
	n.App.Log.Info(s)

	if n.ShowTime {
		log.SetFlags(log.LstdFlags)
	}
}

func (n *NatsStream) Stop() {
	n.App.Log.Info("Остановка подписчика, отписка и закрытие соединения...")
	if n.Durable == "" || n.Unsubscribe {
		n.Sub.Unsubscribe()
	}
	n.SC.Close()
	n.NC.Close()
	n.App.Log.Info("Подписчик остановлен")
}
