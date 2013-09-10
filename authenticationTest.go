package handlers

import (
	"appengine"
    "fmt"
	"io/ioutil"
	"net/http"
    "code.google.com/p/mlab-ns2/gae/ns/digest"
)

func init() {
	http.HandleFunc("admin/TestNagiosAuthentication", TestNagiosAuthentication) 
}

func TestNagiosAuthentication(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
    url := "XXXXXXXXXX?show_state=1&service_name=npad"
    t := digest.GAETransport(c, "XXXXXXXX", "XXXXXXXXXX")
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    } 
    res, err := t.RoundTrip(req) 
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    } 
    txtBlob, err := ioutil.ReadAll(res.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer res.Body.Close()
    fmt.Fprintf(w, "Completed %s", txtBlob)
}

