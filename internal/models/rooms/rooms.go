package rooms

import (
	"encoding/json"

	"github.com/nxshock/signaller/internal/models/events"
)

// https://matrix.org/docs/spec/client_server/latest#room-event-fields
type Event struct {
	Content        json.RawMessage `json:"content"`            // Required. The fields in this object will vary depending on the type of event. When interacting with the REST API, this is the HTTP body.
	Type           string          `json:"type"`               // Required. The type of event. This SHOULD be namespaced similar to Java package naming conventions e.g. 'com.example.subdomain.event.type'
	EventID        string          `json:"event_id"`           // Required. The globally unique event identifier.
	Sender         string          `json:"sender"`             // Required. Contains the fully-qualified ID of the user who sent this event.
	OriginServerTS int             `json:"origin_server_ts"`   // Required. Timestamp in milliseconds on originating homeserver when this event was sent.
	Unsigned       UnsignedData    `json:"unsigned,omitempty"` // Contains optional extra information about the event.
	RoomID         string          `json:"room_id"`            // Required. The ID of the room associated with this event. Will not be present on events that arrive through /sync, despite being required everywhere else.
}

type UnsignedData struct {
	Age             int          `json:"age"`                        // The time in milliseconds that has elapsed since the event was sent. This field is generated by the local homeserver, and may be incorrect if the local time on at least one of the two servers is out of sync, which can cause the age to either be negative or greater than it actually is.
	RedactedBecause events.Event `json:"redacted_because,omitempty"` // Optional. The event that redacted this event, if any.
	TransactionID   string       `json:"transaction_id"`             // The client-supplied transaction ID, if the client being given the event is the same one which sent it.
}