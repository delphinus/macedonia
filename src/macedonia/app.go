package macedonia

import "macedonia/controller"

// Build creates handlers
func Build() *gin.Engine {
	r := gin.New()
	r.GET("/update", controller.Update)
	return r
}
