package db

import "github.com/liangdas/mqant/log"

func (m *Module) rpcFindPlayerByUsername(username string) (bool, string) {
	exists, err := m.ExistPlayer(username)
	if err != nil {
		return exists, err.Error()
	}
	return exists, ""
}

func (m *Module) rpcCreatePlayerData(username string, nickname string) (string, string) {
	player, err := m.CreatePlayer(username, nickname)
	if err != nil {
		return player, err.Error()
	}
	return player, ""
}

func (m *Module) rpcLoadPlayerDataByUsername(username string) (string, string) {
	player, err := m.LoadPlayer(username)
	if err != nil {
		return player, err.Error()
	}
	return player, ""
}

func (m *Module) rpcSavePlayerData(username string, player string) (bool, string) {
	result, err := m.SavePlayer(username, player)
	if err != nil {
		return result, err.Error()
	}

	log.Info("rpcSavePlayerData username=%v success", username)
	return result, ""
}
