package main

import (
	"context"

	"nox/v1/common/discover"
)

func discoverAndPingServers(port int, ts uint32, data []byte) {
	ctx := context.Background()
	err := discover.EachServer(ctx, func(s discover.Server) error {
		_, err := sendToServer(s.Addr, port, data)
		return err
	})
	if err != nil {
		discover.Log.Println(err)
	}
	for _, l := range discover.XWISRooms() {
		if l.Game == nil {
			continue
		}
		g := l.Game
		// TODO: more fields
		onLobbyServer(&LobbyServerInfo{
			Addr:       g.Addr,
			Port:       port, // TODO: this should come from the server
			Name:       g.Name,
			Map:        g.Map,
			Players:    byte(g.Players),
			MaxPlayers: byte(g.MaxPlayers),
			Flags:      uint16(g.Flags) | uint16(g.MapType),
			Ticks:      uint64(ts),
		})
	}
}
