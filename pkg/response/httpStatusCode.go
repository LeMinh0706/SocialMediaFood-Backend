package response

const (
	StatusOk              = 200
	CodeSuccess           = 201
	UpdateFriend          = 20101
	ChangePassword        = 20102
	AddEmail              = 20103
	SeenNoti              = 20104
	DeleteSuccess         = 204
	DeleteNoti            = 20401
	ErrBadRequest         = 40000
	ErrBadRequestPage     = 40001
	ErrBadRequestPageSize = 40002
	ErrBadRequestMime     = 40003
	ErrBadRequestId       = 40004
	ErrImageLen           = 40005
	ErrImageSize          = 40006
	ErrGender             = 40007
	ErrDistance           = 40017
	ErrUnauthorize        = 40101
	ErrInvalid            = 40102
	ErrYourSelf           = 40103
	ErrUserExist          = 40900
	ErrLogin              = 40104
	ErrNotFoundUser       = 40401
	ErrFindPost           = 40402
	ErrUnlike             = 40403
	ErrLike               = 40404
	ErrFileTooLarge       = 41300
	ErrWrongPassword      = 40105
	ErrEmailExists        = 40901
	ErrUsernameChar       = 40008
	ErrMinPassword        = 40009
	ErrMinFullname        = 40010
	ContentNull           = 40011
	ErrAccountID          = 40012
	ErrPositionField      = 40013
	ErrImageWasDelete     = 40408
	HaveFollow            = 40111
	AcceptForbidden       = 40303
	TheirFriend           = 40411
	WaitngAccept          = 40412
	ErrBadLngLat          = 40020
	ErrSaveImage          = 40021
	ErrEmptyContent       = 40022
	ErrAccountExists      = 40414
	ErrDeleteComment      = 40119
	ErrInputFollow        = 40415
	ErrFullNameNull       = 40014
	ErrTokenInvalid       = 40304
	ErrInputSearch        = 40015
	ErrEmailInvalid       = 40016
	ErrReport             = 40116
	ErrEmailNotExists     = 40420
	YouHaveRequest        = 40307
	PasswordHaveChange    = 40308
	ResetPasswordTimeOut  = 40309
	CantDelete            = 40425
)

var msg = map[int]string{
	StatusOk:              "Ok",
	CodeSuccess:           "Success",
	DeleteSuccess:         "Delete no error",
	UpdateFriend:          "Success update to friend",
	ChangePassword:        "Success to reset password",
	AddEmail:              "Add email complete",
	SeenNoti:              "Seen notification success",
	DeleteNoti:            "Delete notification success",
	ErrBadRequest:         "Bad request",
	ErrBadRequestPage:     "Page should be number and greater 1",
	ErrBadRequestPageSize: "Page size should be number and greater 1",
	ErrBadLngLat:          "Lat, Lng should be number",
	ErrBadRequestMime:     "Can only use file .png, .jpg, .jpeg, .gif",
	ErrBadRequestId:       "Id must be number, can't convert from this request",
	ErrImageLen:           "Images shoud less than 10",
	ErrGender:             "Gender should be 0 (for female) or 1 (for male)",
	ErrUnauthorize:        "Unauthorized",
	ErrInvalid:            "Invalid Token",
	ErrYourSelf:           "Not your self, cant create/update/delete anything for another user",
	ErrUserExist:          "User exist",
	ErrLogin:              "User doesn't exists",
	ErrNotFoundUser:       "User not found",
	ErrFindPost:           "Can not found post or post was deleted",
	ErrUnlike:             "You didn't like this post yet",
	ErrLike:               "You have liked post yet",
	ErrFileTooLarge:       "File too large, only allowed 4MB",
	ErrWrongPassword:      "Wrong password",
	ErrEmailExists:        "Email exists",
	ErrEmailNotExists:     "Email doesn't exist",
	ErrUsernameChar:       "Username need to be between 6 to 16 character",
	ErrMinPassword:        "Password need at least 8 character",
	ErrMinFullname:        "Fullname need at least 6 character",
	ContentNull:           "Description for comment can't null",
	ErrAccountID:          "Account id must be number",
	ErrPositionField:      "LNG or Lat must be both fill or both empty",
	ErrImageWasDelete:     "Image not found or was deleted",
	HaveFollow:            "You have followed this person or they waiting for your acceptance",
	AcceptForbidden:       "Waiting for their reply",
	TheirFriend:           "You're their friend",
	WaitngAccept:          "They're waiting for your acceptance",
	ErrSaveImage:          "Failed to save image",
	ErrEmptyContent:       "Description or images can't be empty",
	ErrAccountExists:      "This account doesn't exist",
	ErrDeleteComment:      "Comment not found",
	ErrInputFollow:        "Error follow status input",
	ErrFullNameNull:       "Fullname can't be empty",
	ErrTokenInvalid:       "Invalid token",
	ErrInputSearch:        "Seaching bar can't be empty",
	ErrEmailInvalid:       "This email is invalid",
	ErrReport:             "You have report this post with this issue",
	YouHaveRequest:        "You requested a few minutes ago, please wait",
	PasswordHaveChange:    "You have changed password before",
	ResetPasswordTimeOut:  "You're to late, try request forgot-password again",
	ErrDistance:           "Distance must be number",
	CantDelete:            "Notification not found",
}

// ErrOutOfDate:   "Token is out of date",

// real err
var (
	EmailExists   = "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)"
	UserExists    = "ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)"
	WrongUsername = "wrong username"
	WrongPassword = "wrong password"
	FollowGhost   = "ERROR: insert or update on table \"follower\" violates foreign key constraint \"follower_to_follow_fkey\" (SQLSTATE 23503)"
)
