package mqtt

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Topic struct {
	Name  string
	Topic string 
	Gauge prometheus.Gauge
}

var (
	totalBytesRecieved = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "total_bytes_recieved",
			Help:      "The total number of bytes received since the broker started.",
		})
	totalBytesSent = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "total_bytes_sent",
			Help:      "The total number of bytes sent since the broker started.",
		})
	connectedClients = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "connected_clients",
			Help:      " The number of currently connected clients.",
		})
	expiredClients = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "expired_clients",
			Help:      "The number of disconnected persistent clients that have been expired and removed through the persistent_client_expiration option.",
		})
	disconnectedClients = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "disconnected_clients",
			Help:      "The total number of persistent clients (with clean session disabled) that are registered at the broker but are currently disconnected.",
		})
	maxConnClients = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "max_conn_clients",
			Help:      "The maximum number of clients that have been connected to the broker at the same time.",
		})
	totalClients = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "total_clients",
			Help:      "The total number of active and inactive clients currently connected and registered on the broker.",
		})
	currentHeap = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "current_heap",
			Help:      "The current size of the heap memory in use by mosquitto. Note that this Topic may be unavailable depending on compile time options.",
		})
	largestHeapUsed = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "largest_heap_used",
			Help:      "The largest amount of heap memory used by mosquitto. Note that this Topic may be unavailable depending on compile time options.",
		})
	noOfConnectPackets = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "no_of_connect_packets",
			Help:      "The moving average of the number of CONNECT packets received by the broker over different time intervals. The final " + " of the hierarchy can be 1min, 5min or 15min. The value returned represents the number of connections received in 1 minute, averaged over 1, 5 or 15 minutes.",
		})
	bytesRecieved = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "bytes_recieved",
			Help:      "The moving average of the number of bytes received by the broker over 1min time intervals.",
		})
	bytesSent = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "bytes_sent",
			Help:      "The moving average of the number of bytes sent by the broker over 1min time intervals.",
		})
	messagesRecieved = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "messages_recieved",
			Help:      "The moving average of the number of all types of MQTT messages received by the broker over 1min time intervals.",
		})
	messagesSent = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "messages_sent",
			Help:      "The moving average of the number of all types of MQTT messages sent by the broker over 1min time intervals.",
		})
	publishDropMsg = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "publish_drop_msg",
			Help:      "The moving average of the number of publish messages dropped by the broker over 1min time intervals. This shows the rate at which durable clients that are disconnected are losing messages.",
		})
	publishRecieveMsg = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "publish_recieve_msg",
			Help:      "The moving average of the number of publish messages received by the broker over 1min time intervals.",
		})
	publishSentMsg = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "publish_sent_msg",
			Help:      "The moving average of the number of publish messages sent by the broker over 1min time intervals.",
		})
	openSocketConnection = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "open_socket_connection",
			Help:      "The moving average of the number of socket connections opened to the broker over different time intervals.",
		})
	inflight = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "inflight",
			Help:      "The number of messages with QoS>0 that are awaiting acknowledgments.",
		})
	totalMessagesRecieved = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "total_messages_recieved",
			Help:      "The total number of messages of any type received since the broker started.",
		})
	totalMessagesSent = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "total_messages_sent",
			Help:      "The total number of messages of any type sent since the broker started.",
		})
	stored = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "stored",
			Help:      "The number of messages currently held in the message store. This includes retained messages and messages queued for durable clients",
		})
	retainedMsgCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "retained_msg_count",
			Help:      "The total number of retained messages active on the broker.",
		})
	totalActiveSubscriptions = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "total_active_subscriptions",
			Help:      "The total number of subscriptions active on the broker.",
		})
	uptime = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "uptime",
			Help:      "The amount of time in seconds the broker has been online.",
		})
	online = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "online",
			Help:      "1 if broker is Online and 0 if broker is offline",
		})
)

var Topics = []Topic{
	{Name: "totalBytesRecieved", Topic: "$SYS/broker/bytes/received", Gauge: totalBytesRecieved},
	{Name: "totalBytesSent", Topic: "$SYS/broker/bytes/sent", Gauge: totalBytesSent},
	{Name: "connectedClients", Topic: "$SYS/broker/clients/connected", Gauge: connectedClients},
	{Name: "expiredClients", Topic: "$SYS/broker/clients/expired", Gauge: expiredClients},
	{Name: "disconnectedClients", Topic: "$SYS/broker/clients/disconnected", Gauge: disconnectedClients},
	{Name: "maxConnClients", Topic: "$SYS/broker/clients/maximum", Gauge: maxConnClients},
	{Name: "totalClients", Topic: "$SYS/broker/clients/total", Gauge: totalClients},
	{Name: "currentHeap", Topic: "$SYS/broker/heap/current size,", Gauge: currentHeap},
	{Name: "largestHeapUsed", Topic: "$SYS/broker/heap/maximum size", Gauge: largestHeapUsed},
	{Name: "noOfConnectPackets", Topic: "$SYS/broker/load/connections/+", Gauge: noOfConnectPackets},
	{Name: "bytesRecieved", Topic: "$SYS/broker/load/bytes/received/1min", Gauge: bytesRecieved},
	{Name: "bytesSent", Topic: "$SYS/broker/load/bytes/sent/1min", Gauge: bytesSent},
	{Name: "messagesRecieved", Topic: "$SYS/broker/load/messages/received/1min", Gauge: messagesRecieved},
	{Name: "messagesSent", Topic: "$SYS/broker/load/messages/sent/1min", Gauge: messagesSent},
	{Name: "publishDropMsg", Topic: "$SYS/broker/load/publish/dropped/1min", Gauge: publishDropMsg},
	{Name: "publishRecieveMsg", Topic: "$SYS/broker/load/publish/received/1min", Gauge: publishRecieveMsg},
	{Name: "publishSentMsg", Topic: "$SYS/broker/load/publish/sent/1min", Gauge: publishSentMsg},
	{Name: "openSocketConnection", Topic: "$SYS/broker/load/sockets/1min", Gauge: openSocketConnection},
	{Name: "inflight", Topic: "$SYS/broker/messages/inflight", Gauge: inflight},
	{Name: "totalMessagesRecieved", Topic: "$SYS/broker/messages/received", Gauge: totalMessagesRecieved},
	{Name: "totalMessagesSent", Topic: "$SYS/broker/messages/sent", Gauge: totalMessagesSent},
	{Name: "stored", Topic: "$SYS/broker/messages/stored", Gauge: stored},
	{Name: "retainedMsgCount", Topic: "$SYS/broker/retained messages/count", Gauge: retainedMsgCount},
	{Name: "totalActiveSubscriptions", Topic: "$SYS/broker/subscriptions/count", Gauge: totalActiveSubscriptions},
	{Name: "uptime", Topic: "$SYS/broker/uptime", Gauge: uptime},
	{Name: "version", Topic: "$SYS/broker/version", Gauge: nil},
	{Name: "online", Topic: "", Gauge: online},
}
