package main
 
import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var target = os.Getenv("TARGET_HOSTNAME")
var referer = os.Getenv("REFERER")
 
func main() {
	log.Printf("Starting with target %s and referer %s", target, referer)
	http.HandleFunc("/", ProxyFunc)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
 
func ProxyFunc(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse("https://" + target)
	log.Printf("Proxy %s to %s", r.URL, u)

	proxy := httputil.NewSingleHostReverseProxy(u)
	r.Header = map[string][]string{}
	r.Header.Add("Referer", referer)  
	r.Host = target
	proxy.ServeHTTP(w, r)
 
}
