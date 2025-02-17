package usecase

type Policy interface {
	IsSysadmin(userID string) bool
	CanManageChannel(userID string, channelID string) bool
}
