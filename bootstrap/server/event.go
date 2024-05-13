package server

import (
	"obyoy-backend/container"
	"obyoy-backend/event"
	"obyoy-backend/ws"

	"github.com/codeginga/locevt"
)

// Event registers event related providers
func Event(c container.Container) {
	c.Register(func() locevt.Event {
		return locevt.NewEvent()
	})

	c.Resolve(func(e locevt.Event, hub ws.Hub) {
		e.Register(
			event.NameWSNotification,
			event.WSNotificationWorker(hub),
		)
	})
}
