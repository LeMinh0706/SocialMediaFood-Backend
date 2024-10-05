package response

import "github.com/LeMinh0706/SocialMediaFood-Backend/db"

func UserRes(user db.User) UserResponse {
	return UserResponse{ID: user.ID, Fullname: user.Fullname, Gender: user.Gender, UrlAvatar: user.UrlAvatar, UrlBackground: user.UrlBackgroundProfile, RoleID: user.RoleID, DateCreateAccount: user.DateCreateAccount}
}
