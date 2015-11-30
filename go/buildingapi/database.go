package main

//Almost every API will interact with some kind of datastore
//We need to:
// 1. Connect when progrm is first run(expensive)
// 2. Disconnect when the program is terminated
// 3. Create a session per request(checp)
// 4. Clean up after each request

func withDB(s *mgo.Session, h http.Handler) http.Handler {
	return &dbwrapper{dbSession: s, h: h}
}

type dbwrapper struct {
	dbSession *mgo.Session
	h         http.Handler
}

func (dbwrapper *dbwrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// copy the session
	dbcopy := dbwrapper.dbSession.Copy()
	defer dbcopy.Close()
	context.Set(r, "db", dbcopy)
	dbwrapper.h.ServeHTTP(w, r)
}

func handleThingRead(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*mgo.Session)

	var results []interface{}
	if err := db.DB("myapp").C("things").Find(nil).All(&results); err != nil {
		respond.With(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.With(w, r, http.StatusOK, results)
}

func main() {
	db, _ := mgo.Dial("localhost")
	defer db.Close()

	router := mux.NewRouter()

	router.Handle("/things", withDB(db, http.HandlerFunc(handleThingsRead)))
	router.Handle("/status", http.HandlerFunc(handleStatus))

	http.ListenAndServe(":8080", context.ClearHandler(router))
}
