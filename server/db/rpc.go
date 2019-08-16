package db

func (m *Module) rpcFindPlayerByUsername(username string) (bool, string) {
	exists, err := m.ExistPlayer(username)
	if err != nil {
		return exists, err.Error()
	}
	return exists, ""
}

func (m *Module) rpcCreatePlayerData(username string, nickname string) (string, string){
	player, err := m.CreatePlayer(username, nickname)
	if err != nil {
		return player, err.Error()
	}
	return player, ""
}

func (m *Module) rpcLoadPlayerDataByUsername() {

}

func (m *Module) rpcSavePlayerData() {

}