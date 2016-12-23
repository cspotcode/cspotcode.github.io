package main

import {
    "net/http"
}

func init() {
    http.HandleFunc("/", redirect)
}

func redirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, r.url.RawPath + r.url.RawQuery, http.StatusMovedPermanently)
}
