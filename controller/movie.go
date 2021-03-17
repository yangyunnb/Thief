package controller

import (
	"time"

	"github.com/Thief.git/common"
	"github.com/Thief.git/models"
	"github.com/Thief.git/protocol"
	"github.com/kataras/iris"
)

type Movie struct {
}

func (m *Movie) PostAdd(ctx iris.Context) {
	addMovieReq := protocol.AddMovieReq{}
	if err := ctx.ReadJSON(&addMovieReq); err != nil {
		ctx.JSON(map[string]interface{}{"err": err.Error()})
	}
	dbModel := models.MovieModel{
		DB: common.GetDB(),
	}

	newMovie := models.Movie{
		Name:       addMovieReq.Name,
		Year:       addMovieReq.Year,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := dbModel.Insert(&newMovie); err != nil {
		ctx.JSON(map[string]interface{}{"err": err.Error()})
	}
	ctx.JSON(newMovie)
}
