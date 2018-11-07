package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "invalid params",
	ERROR_EXIST_TAG:                "tag is exist",
	ERROR_NOT_EXIST_TAG:            "tag is not exist",
	ERROR_NOT_EXIST_ARTICLE:        "article is not exist",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "token check is fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token is expired",
	ERROR_AUTH_TOKEN:               "can not create token",
	ERROR_AUTH:                     "Token error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
