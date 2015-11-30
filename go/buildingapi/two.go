package test

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServerHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(int)
}

func main() {
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/goodbye", handleGoodbye)
	http.HandleFunc("/", handleIndex)
}
