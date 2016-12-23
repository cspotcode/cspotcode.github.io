package main

import (
    "net/http"
)

var sourceDomain string
var targetDomain string

func init() {
    http.HandleFunc("/", redirect)
    sourceDomain = "www.cspotcode.com"
    targetDomain = "cspotcode.com"
}

func redirect(w http.ResponseWriter, r *http.Request) {
    if(r.Host == sourceDomain) {
        redirectTo := "https://" + targetDomain + r.URL.Path
        if(r.URL.RawQuery != "") {
            redirectTo += "?" + r.URL.RawQuery
        }
        http.Redirect(w, r, redirectTo, http.StatusMovedPermanently)
        return
    }
    http.Error(w, "", http.StatusNotFound)
}

