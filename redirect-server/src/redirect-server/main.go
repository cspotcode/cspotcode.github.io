package main

import (
    "net/http"
    "google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
)

type RedirectConfig struct {
    SourceDomain string
    TargetDomain string
}

var redirect *RedirectConfig

func init() {
    http.HandleFunc("/", onRequest)
}

func getRedirectConfig(r *http.Request) {
    if(redirect != nil) {
        return
    }
    ctx := appengine.NewContext(r)
    redirect = new(RedirectConfig)
    k := datastore.NewKey(ctx, "redirect", "singleton", 0, nil)
    datastore.Get(ctx, k, redirect)
}

func onRequest(w http.ResponseWriter, r *http.Request) {
    getRedirectConfig(r)
    if(r.Host == redirect.SourceDomain) {
        redirectTo := "https://" + redirect.TargetDomain + r.URL.Path
        if(r.URL.RawQuery != "") {
            redirectTo += "?" + r.URL.RawQuery
        }
        http.Redirect(w, r, redirectTo, http.StatusMovedPermanently)
        return
    }
    http.Error(w, "", http.StatusNotFound)
}

