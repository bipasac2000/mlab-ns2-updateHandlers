package handlers 

import (
    "appengine"
    "appengine/datastore"
    "net/http"
    "fmt"
    "strings"

    "code.google.com/p/mlab-ns2/gae/ns/data"
)

func init(){
    http.HandleFunc("/testSite", testSite)    
    http.HandleFunc("/testSliverTool", testSliverTool)
}

func testSite(w http.ResponseWriter, r *http.Request){
    
    c := appengine.NewContext(r)
    q := datastore.NewQuery("Site")
    var sites []*data.Site
    _, err := q.GetAll(c, &sites)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) 
        return
    }

    q = datastore.NewQuery("Sites")
    var sitesNew []*data.Site
    _, err = q.GetAll(c, &sitesNew)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) 
        return
    }

    if len(sites) != len(sitesNew) {
        fmt.Fprintf(w, "sites datastore length does not match")
        return
    }

    for i,_ := range sites {
        a := sites[i]
        b := sitesNew[i]
        if !strings.EqualFold(a.SiteID, b.SiteID) {
            fmt.Fprintf(w, "sites data does not match %s-%s", a.SiteID, b.SiteID)
        }
    }
    fmt.Fprintf(w, "OK")
}

func testSliverTool(w http.ResponseWriter, r *http.Request){
    
    c := appengine.NewContext(r)
    q := datastore.NewQuery("SliverTool")
    var sls []*data.SliverTool
    _, err := q.GetAll(c, &sls)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) 
        return
    }

    for i,_ := range sls {
        a := sls[i]
        q = datastore.NewQuery("SliverTools").Filter("fqdn=", a.FQDN)
        count, _ := q.Count(c)
        if count < 1 {
            fmt.Fprintf(w, "%s ", a.FQDN)    
        }
    }
    fmt.Fprintf(w, "OK")
}
