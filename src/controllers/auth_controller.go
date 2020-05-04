package controllers

import "simple-auth/src/requests"

//AuthController godoc
type AuthController struct {
}

//Login godoc
// @Summary Update a player item
// @Tags Teams
// @Accept  json
// @Produce json
// @Param model body requestmodels.PlayerUpdateModel true "Update Team"
// @Param id path string true "Team ID" string
// @Param playerid path string true "Player ID" string
// @Success 204
// @Failure 400 {object} responsemodels.ErrorModel
// @Router /teams/:id/players/:playerid [put]
func (authController *AuthController) Login(model requests.Login) {

}
