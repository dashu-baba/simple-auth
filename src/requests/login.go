package requests

//Login godoc
//@Summary defines the login request model
type Login struct {
	// {"method": "email|phone"}
	Method    string `json:"method" form:"method" xml:"method" binding:"required"`
	Userid    string `json:"userId" form:"userid" xml:"userId" binding:"required"`
	Password  string `json:"password" form:"password" xml:"password" binding:"required"`
	Appid     string `json:"appid" form:"appid" xml:"appid" binding:"required"`
	Tennantid string `json:"tennantId" form:"tennantid" xml:"tennantId"`
	State     string `json:"state" form:"state" xml:"state" binding:"required"`
}
