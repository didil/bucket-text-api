package api

import "net/http"

type SaveTextReqData struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// HandleSaveText saves text
func (app *App) HandleSaveText(w http.ResponseWriter, r *http.Request) {
	reqData := &SaveTextReqData{}

	err := readJSON(r, reqData)
	if err != nil {
		app.HandleError(w, r, err)
		return
	}

	err = app.GCPSvc.CreateFile(r.Context(), reqData.Name, reqData.Content)
	if err != nil {
		app.HandleError(w, r, err)
		return
	}

	JSONOk(w, &struct{}{})
}
