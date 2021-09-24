package user

import (
	"github.com/gin-gonic/gin"
	. "vtmtea.com/f.cli/handler"
	"vtmtea.com/f.cli/model"
	"vtmtea.com/f.cli/pkg/errno"
)

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user [post]
func Create(c *gin.Context) {

	var r CreateRequest

	var err error

	if err = c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	if err = r.checkParam(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// Encrypt the user password.
	if err = u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Insert the user to the database.
	if err = u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	resp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, err, resp)
}

func (r *CreateRequest) checkParam() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}

	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}

	return nil
}
