package basic

type ReaPackRoom struct {
	id              int
	maxPlayerNumber int
	Banker          *CommonUser
	Players         []*CommonUser
}

func (r *ReaPackRoom) JoinRoom(player *CommonUser, roomId int) int {

	if len(r.Players) >= r.maxPlayerNumber {
		return -1
	} else {
		r.Players = append(r.Players, player)
		return 1
	}

}
