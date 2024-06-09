package functions

import (
	"github.com/piyush7833/Chat-Api/types"
)

func CreateFriendRequest(senderId string, receiverId string) (interface{}, types.ErrorType) {
	// friend, _ :=
	// if friend != nil {
	// 	return nil, types.ErrorType{
	// 		Message:    "You are already friends",
	// 		StatusCode: 400,
	// 	}
	// }
	// 	friendReq, _ :=
	// if friendReq != nil {
	// 	return nil, types.ErrorType{
	// 		Message:    "Friend request already exists between sender and receiver",
	// 		StatusCode: 400,
	// 	}
	// }

	// res, err :=

	// if err != nil {
	// 	if strings.Contains(err.Error(), "Unique constraint failed") {
	// 		return nil, types.ErrorType{
	// 			Message:    "Friend request already exists between sender and receiver",
	// 			StatusCode: 400,
	// 		}
	// 	}

	// Handle other errors
	// return nil, types.ErrorType{
	// 	Message:    err.Error(),
	// 	StatusCode: 500,
	// }
	// }

	// return res, error
	return nil, error
}

func UpdateFriendRequestStatus(id string, status string, userId string) (interface{}, types.ErrorType) {
	// FriendRequest, err :=
	// if err != nil {
	// 	error.Message = err.Error()
	// 	error.StatusCode = 500
	// 	return nil, error
	// }
	// if userId != FriendRequest.ReceiverID {
	// 	error.Message = "You can update status of friend request which is recieved to you"
	// 	error.StatusCode = 500
	// 	return nil, error
	// }

	if status == "reject" {
		// res, err :=
		// if err != nil {
		// 	error.Message = err.Error()
		// 	error.StatusCode = 500
		// 	return nil, error
		// }
		// return res, error
		return nil, error
	} else if status == "accept" {
		// res1, err := helpers.Prisma.Friend.CreateOne(db.Friend.User.Link(db.User.ID.Equals(FriendRequest.SenderID)), db.Friend.Friend.Link(db.User.ID.Equals(FriendRequest.ReceiverID))).Exec(ctx)
		// if err != nil {
		// 	error.Message = err.Error()
		// 	error.StatusCode = 500
		// 	return nil, error
		// }

		// _, err2 := helpers.Prisma.FriendRequest.FindUnique(db.FriendRequest.ID.Equals(id)).Delete().Exec(ctx)
		// if err2 != nil {
		// 	error.Message = err2.Error()
		// 	error.StatusCode = 500
		// 	return nil, error
		// }

		return nil, error
		// return res1, error
	}
	error.Message = "Friend request status can only be reject or accept"
	error.StatusCode = 403
	return nil, error
}

func DeleteFriendRequest(id string, userId string) (interface{}, types.ErrorType) {
	// FriendRequest, err :=
	// if err != nil {
	// 	error.Message = err.Error()
	// 	error.StatusCode = 500
	// 	return nil, error
	// }
	// if userId != FriendRequest.SenderID && userId != FriendRequest.ReceiverID {
	// 	error.Message = "You can only delete friend request which is sent to you or recieved to you"
	// 	error.StatusCode = 500
	// 	return nil, error
	// }

	// res, err := helpers.Prisma.FriendRequest.FindUnique(db.FriendRequest.ID.Equals(id)).Delete().Exec(ctx)
	// if err != nil {
	// 	error.Message = err.Error()
	// 	error.StatusCode = 500
	// 	return nil, error
	// }
	// return res, error
	return nil, error
}

func GetParticularFriendRequest(id string, userId string) (interface{}, types.ErrorType) {
	// res, err := helpers.Prisma.FriendRequest.FindUnique(db.FriendRequest.ID.Equals(id)).Exec(ctx)
	// if err != nil {
	// 	error.Message = err.Error()
	// 	error.StatusCode = 500
	// 	return nil, error
	// }
	// if userId != res.SenderID && userId != res.ReceiverID {
	// 	error.Message = "You can view friend request which is sent to you or recieved to you"
	// 	error.StatusCode = 500
	// 	return nil, error
	// }
	return nil, error
	// return res, error
}

func GetAllFriendRequestRecievedByUser(id string, page int) (interface{}, types.ErrorType) {
	// res, err := helpers.Prisma.FriendRequest.FindMany(db.FriendRequest.ReceiverID.Equals(id)).OrderBy(
	// 	db.FriendRequest.CreatedAt.Order(db.SortOrderDesc),
	// ).Take(20).Skip(page * config.RowsPerPageGenral).Exec(ctx)
	// if err != nil {
	// 	error.Message = err.Error()
	// 	error.StatusCode = 500
	// 	return nil, error
	// }
	return nil, error
	// return res, error
}

func GetAllFriendRequestSentByUser(id string, page int) (interface{}, types.ErrorType) {
	// res, err := helpers.Prisma.FriendRequest.FindMany(db.FriendRequest.SenderID.Equals(id)).OrderBy(
	// 	db.FriendRequest.CreatedAt.Order(db.SortOrderDesc),
	// ).Take(config.RowsPerPageGenral).Skip(page * config.RowsPerPageGenral).Exec(ctx)
	// if err != nil {
	// 	error.Message = err.Error()
	// 	error.StatusCode = 500
	// 	return nil, error
	// }
	// return res, error
	return nil, error
}
