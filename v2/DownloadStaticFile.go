package toolkit

import (
	"fmt"
	"net/http"
)

// DownloadStaticFile a file and tries to force the browser to avoid displaying it in the browser window by setting content disposition.
// It also allows specification of the display name
func (t *Tools) DownloadStaticFile(w http.ResponseWriter, r *http.Request, pathName, displayName string) {
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", displayName))

	http.ServeFile(w, r, pathName)
}
