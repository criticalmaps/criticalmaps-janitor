package controllers

import (
	"errors"
	"net/http"

	"github.com/criticalmaps/criticalmaps-janitor/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	gallery := []models.Gallery{}

	result := DB.Select("id").Find(&gallery)
	if result.Error != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}

	return c.Render(gallery)

}

func (c App) Thumbnail(id int) revel.Result {
	var gallery models.Gallery

	result := DB.Select("thumbnail").First(&gallery, id)
	if result.Error != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}

	c.Response.WriteHeader(http.StatusOK, "image/jpeg")
	c.Response.GetWriter().Write(gallery.Thumbnail)

	return c.Result
}

func (c App) Image(id int) revel.Result {
	var gallery models.Gallery

	result := DB.Select("image").First(&gallery, id)
	if result.Error != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}

	c.Response.WriteHeader(http.StatusOK, "image/jpeg")
	c.Response.GetWriter().Write(gallery.Image)

	return c.Result
}

func (c App) Delete(id int) revel.Result {
	var imageToDelete models.Gallery

	result := DB.First(&imageToDelete, id)
	if result.Error != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}

	DB.Delete(&imageToDelete)

	return c.RenderText("success")
}
