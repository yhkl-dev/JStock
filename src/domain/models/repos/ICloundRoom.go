package repos

type ICloudRoomModel interface {
	New(IModel) error
	FindByID(IModel) error
	Load(IModel) error
	ListCloudRoom(m IModel, page, pageSize int) (interface{}, error)
}
