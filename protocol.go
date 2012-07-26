package kartoffel

// Common definitions for MQTT protocol

// CONNECT command message - see section 3.1 of MQTT 3.1 specification
type Connect struct {
}

// CONNACK command message - see section 3.2 of MQTT 3.1 specification
type ConnAck struct {
}

// PUBLISH command message - see section 3.3 of MQTT 3.1 specification
type Publish struct {
}

// PUBACK command message - see section 3.4 of MQTT 3.1 specification
type PublishAck struct {
}

// PUBREC command message - see section 3.5 of MQTT 3.1 specification
type PublishRec struct {
}

// PUBREL command message - see section 3.6 of MQTT 3.1 specification
type PublishRel struct {
}

// PUBCOMP command message - see section 3.7 of MQTT 3.1 specification
type PublishComp struct {
}

// SUBSCRIBE command message - see section 3.8 of MQTT 3.1 specification
type Subscribe struct {
}

// SUBACK command message - see section 3.9 of MQTT 3.1 specification
type SubscribeAck struct {
}

// UNSUBSCRIBE command message - see section 3.10 of MQTT 3.1 specification
type Unsubscribe struct {
}

// UNSUBACK command message - see section 3.11 of MQTT 3.1 specification
type UnsubscribeAck struct {
}

// PINGREQ command message - see section 3.12 of MQTT 3.1 specification
type PingReq struct {
}

// PINGRESP command message - see section 3.13 of MQTT 3.1 specification
type PingResp struct {
}

// DISCONNECT command message - see section 3.14 of MQTT 3.1 specification
type Disconnect struct {
}

type ProtocolError struct {
	Message string
}

func (e ProtocolError) Error() string {
	return e.Message
}

var notImplemented = ProtocolError{"Not implemented"}

func connect(c Connect) (ack *ConnAck, err error) {
	return nil, notImplemented
}

func publishQosZero(p Publish) error {
	return notImplemented
}

func publishQosOne(p Publish) (ack *PublishAck, err error) {
	return nil, notImplemented
}

func publishQosTwo(p Publish) (ack *PublishRec, err error) {
	return nil, notImplemented
}

func publishRelease(p PublishRel) (comp *PublishComp, err error)	{
	return nil, notImplemented
}

func subscribe(s Subscribe) (ack *SubscribeAck, err error)	{
	return nil, notImplemented
}

func unsubscribe(s Unsubscribe) (ack *UnsubscribeAck, err error)	{
	return nil, notImplemented
}

func ping(p PingReq) (resp *PingResp, err error)	{
	return nil, notImplemented
}

func disconnect(d Disconnect) error	{
	return notImplemented
}
