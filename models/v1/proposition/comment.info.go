package proposition

import "github.com/e-politica/api/models/v1/user"

type CommentInfo struct {
	Comment     string          `json:"comment"`
	UserPubInfo user.PublicInfo `json:"user_public_info"`
}
