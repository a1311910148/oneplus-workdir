package main

import "net/http"

func getQuery(req *http.Request, query string) string {
	q := req.URL.Query().Get("query")
	return q
}

func GetUserQuery(req *http.Request) (string, string, string) {
	return getQuery(req, "name"), getQuery(req, "age"), getQuery(req, "email")
}
