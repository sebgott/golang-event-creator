package handlers

import (
	"fmt"
	"net/http"
)

func AddConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
	<div class="config-item">
	  <input name="config_key[]" placeholder="key">
	  <input name="config_val[]" placeholder="value">
	</div>
	`)
}
