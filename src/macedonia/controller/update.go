package controller

import (
	"macedonia/lib/updater"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// Update means `/update`
func Update(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)
	if err := updater.Updater(ctx); err != nil {
		log.Warningf(ctx, "error in Updater: %v", err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}
