package api

type Message struct {
	messageReceived   string
	messageToTransmit string
}
type Transmitter struct {
	channel  string
	user     string
	password string
	vhost    string
}
type Listener struct {
	channel     string
	user        string
	password    string
	vhost       string
	messages    map[Message]Message
	transmitter Transmitter
}
