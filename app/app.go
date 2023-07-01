package app

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// setRouters sets the all required routers
func (a *App) setRouters() {

}
