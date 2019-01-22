package basic

type Room interface {

	//加入房间
	JoinRoom(uerId int, roomId int) int

	//离开房间
	ExitRoom(uerId int, roomId int) error

	//发包
	Dispense(uerId int, pack interface{}) error

	//抢包
	RobPack(uerId int, packId int) error

	////发送抢庄包
	//DispenseBanker(roomId int) error

	//抢庄包
	RobBanker(uerId int, packId int) error
}
