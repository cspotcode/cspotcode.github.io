package main

import (
    "net/http"
    "time"
    "google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
)

type RedirectConfig struct {
    SourceDomain string
    TargetDomain string
}

type Log struct {
    Time int32
    SourceDomain string
    TargetDomain string
    Host string
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

func writeLog(r *http.Request, l *Log) {
    ctx := appengine.NewContext(r)
    k := datastore.NewKey(ctx, "log", "mostrecent", 0, nil)
    datastore.Put(ctx, k, l)
}

func onRequest(w http.ResponseWriter, r *http.Request) {
    getRedirectConfig(r)
    writeLog(r, &Log{
        Time: int32(time.Now().Unix()),
        SourceDomain: redirect.SourceDomain,
        TargetDomain: redirect.TargetDomain,
        Host: r.Host,
    })
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

