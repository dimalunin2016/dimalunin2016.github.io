package main

import (
    "encoding/json"
    "net/http"
)

var urls map[string]string

func getAns(url string) string {
    ans := ""
    i := len(urls) + 1
    for ; i > 0; i /= 26 {
        ans += string('a' + i % 26);
    }
    urls[ans] = url
    return ans
}

type Req struct {
    url string `json:"url"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        var req Req
        dec := json.NewDecoder(r.Body)
        er := dec.Decode(&req)
        if er != nil || req.url == "" {
            http.Error(w, "some error", 400)
            return
        }
        rMap := make(map[string]string)
        rMap["key"] = getAns(req.url)
        res, _ := json.Marshal(rMap)
        w.Write(res)
    } else if r.Method == "GET" {
        cUrl := r.RequestURI[1:]
        dUrl, ex := urls[cUrl]
        if !ex {
            http.Error(w, "not found", 404)
            return
        }
        http.Redirect(w, r, dUrl, 301)
    }
}

func main() {
    urls = make(map[string] string)
    http.HandleFunc("/",  handler)
    http.ListenAndServe(":8082", nil)
}
