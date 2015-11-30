package test

//Writing middleware

//Don't break the interface

//some people do this:
func xxx(w http.ResponseWriter, r *http.Request, db *mgo.Session, logger *log.Logger) {
}

//But I prefer this:
func xxx(w http.ResponseWriter, r *http.Request) {

}

type Server struct {
	logger *log.Logger
	mailer MailSender
	slack  Notifier
}

func (s *Server) handleSomething(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
