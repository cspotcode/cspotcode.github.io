package main

import (
    "net/http"
    "google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
)

type Redirect struct {
    sourceDomain string
    targetDomain string
}

var redirect Redirect

func init() {
    http.HandleFunc("/", redirect)
}

func getRedirect(r *http.Request) {
    if(redirect != nil) {
        return
    }
    ctx := appengine.NewContext(r)
    redirect = new(Redirect)
    k := datastore.NewKey(ctx, "redirect", "singleton", 0, nil)
    datastore.Get(ctx, k, redirect)
}

func redirect(w http.ResponseWriter, r *http.Request) {
    getRedirect(r)
    if(r.Host == redirect.sourceDomain) {
        redirectTo := "https://" + redirect.targetDomain + r.URL.Path
        if(r.URL.RawQuery != "") {
            redirectTo += "?" + r.URL.RawQuery
        }
        http.Redirect(w, r, redirectTo, http.StatusMovedPermanently)
        return
    }
    http.Error(w, "", http.StatusNotFound)
}

