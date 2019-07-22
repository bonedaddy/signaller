package memory

import (
	"sync"

	"github.com/nxshock/signaller/internal"
	"github.com/nxshock/signaller/internal/models/createroom"
	"github.com/nxshock/signaller/internal/models/rooms"
)

type Room struct {
	id         string
	Visibility createroom.VisibilityType
	aliasName  string
	name       string
	topic      string

	creator internal.User
	joined  []internal.User

	events []RoomEvent

	mutex sync.RWMutex
}

func (room *Room) ID() string {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.id
}

func (room *Room) Name() string {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.name
}

func (room *Room) AliasName() string {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.aliasName
}

func (room *Room) Topic() string {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.topic
}

func (room *Room) Users() []internal.User {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.joined
}

func (room *Room) Events() []rooms.Event {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	result := make([]rooms.Event, 0)
	for _, v := range room.events {
		result = append(result, v.ToEvent())
	}

	return result
}

func (room *Room) Creator() internal.User {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.creator
}
