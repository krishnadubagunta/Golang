package twitch

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func requests(w http.ResponseWriter, r *http.Request, v interface{}, url string, method string) (http.ResponseWriter, interface{}) {
	req, _ := http.NewRequest(method, url, nil)
	req.Header = r.Header
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(v)
	w.Header().Set("Content-Type", "application/json")
	return w, v
}
