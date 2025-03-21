package resp

import "backend/models/sqlmodels"

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	sqlmodels.MenuModel
	Banners []Banner `json:"banners"`
}

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}
