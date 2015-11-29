package funcwrapper

/* An htt.HandlerFunc wrapper is a function that has one input
argument and one output argument, both of type http.HandlerFunc
*/
func log(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before")
		fn(w, r)
		log.Println("After")
	}
}

/* http.HandleFucn("/path", handleThing)
--> http.HandlerFunc("/path", log(handleThing))
*/

// When would you use wrappers ?
/* This approach can be used to address lots of different situations, including but not limited to:
1. Logging and tracing
2. Connecting and disconnection to databases
3. Validation the request, such as checking authentication credentials
4. Writing common response headers
*/

/* Sharing state
If our 'http.HandlerFunc' were to create some useful object that our original handler might want to
use(such as d database connection), we need to make that object available to handlers. A common and
forgivable solution is to create your own alternative to the 'http.HandlerFunc' where you take the
object as an additional argument. This is not recommanded, because suddenly your code stops working
with normal 'http.HandlerFunc' functions. Instead consider a solution like Gorilla's context package.
*/

func MustParams(fn http.HandlerFunc, params ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, param := range params {
			if len(r.URL.Query().Get(param)) == 0 {
				http.Error(w, "missing "+param, http.StatusBadRequest)
				return
			}
		}
		fn(w, r)
	}
}

func MustAuth(fn http.HandlerFunc) http.HandlerFunc {
	return MustParams(func(w http.ResponseWriter, r *http.Request) {
		//TODO: use auth arguments to validate request
		fun(w, r)
	}, "auth")
}

func log(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := NewResponseLogger(w)
		fn(logger, r)
	}
}

func (r *ResponseLogger) Write(b []byte) (int, error) {
	log.Print(string(b))
	return r.w.Write(b)
}

func log(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := NewResponseLogger(w)
		fn(logger, r)
	}
}

func log(fn http.HandlerFunc) http.HandlerFunc {
	name := runtime.FuncForPC(reflect.ValuseOf(fn).Pointer()).Name()
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before ", name)
		defer log.Println("After ", name)
		fn(w, r)
	}
}

func log(fn http.HandlerFunc) http.HandlerFunc {
	name := runtime.FuncForPC(reflect.ValuseOf(fn).Pointer()).Name()
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before ", name)
		defer log.Println("After ", name)
		fn(w, r)
	}
}

//3
const appVersionStr = "1.0"

func commonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ReponseWriter, r *http.Request) {
		w.Header().Set("X-App-Version", appVersionStr)
		fn(w, r)
	}
}

func main() {
	http.HandlerFunc("/user", MustParams(handleUser, "key", "auth"))
	http.HandlerFunc("/group", MustParams(handleGroup, "key"))
	http.HandlerFunc("/items", MustParams(handleItems, "key", "q"))
}

//4
type Server struct {
	dbsession *mgo.Session
}

func NewServer() (*Server, error) {
	dbsession, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return &Server{dbsession: dbsession}, nil
}

func (s *Server) Close() {
	s.dbsession.Close()
}

//version 1
func (s *Server) WithData(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbcopy := s.dbsession.Copy()
		defer dbcopy.Close()
		context.Set(r, "db", dbcopy)
		fn(w, r)
	}
}

//version 2
func (s *Server) WithData(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, present := context.GetOk(r, "db"); !present {
			dbcopy := s.dbsession.Copy()
			defer dbcopy.Close()
			context.Set(r, "db", dbcopy)
		}
		fn(w, r)
	}
}

func handleThing(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*mgo.Session)
	err := db.DB("myapp").C("things").Insert(bson.M{"val": 1})
	if err != nil {
		//handle error
	}
	io.WriteString(w, "Inserted")
}

func db(r *http.Request) *mgo.Session {
	db, ok := context.GetOk(r, "db")
	if !ok {
		panic("db missing: wrap with WithData")
	}
	return db
}

func main() {
	srv, err := NewServer()
	if err != nil {
		// handle error
	}

	defer srv.Close()
	//setup handlers
	http.HandlerFunc("/things", srv.WithData(handleThing))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		//handle error
	}
}

//5
/*
{
	_id:  "some ID",
	name: "user's name",
	auth: {
		"apiky" : "token",
		"apike2" : " token"
	}
}
*/

type User struct {
	ID   bson.ObjectId     `bson:"_id"`
	Name string            `bson:"name"`
	Auth map[string]string `bson:"auth"`
}

/*
Our "WithAuth" wapper will:
	1. Validate the auth token
	2. Use our 'db' helper function to get a valid database connection
	3. Query our 'users' collection in MongoDB for a user that has the specifed auth token for the specified API key
	4. If no user was found, assume the keys are wrong and return with an http.StatusUnauthorized
	5. If there are no errors, save the User object in the context for this request, and call the original handler
*/

func WithAuth(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get and check the auth paramter
		authken := r.URL.Query().Get("auth")
		if len(authtoken) == 0 {
			http.Error(w, "missing auth param", http.StatusUnauthorized)
			return
		}

		//assume key has already been validated
		key := r.URL.Query().Get("key")
		var user User
		if err := db(r).DB("myapp").C("users").Find(bson.M{
			"auth." + key: authtoken}).One(&user); err != nil {
			if err == mgo.ErrNotFound { //no user found
				http.Error(w, "Bad auth param", http.StatusUnauthorized)
				return
			}

			//some other error
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//save the user in the context
		context.Set(r, "user", &user)

		//call the original handler
		fn(w, r)
	}
}

//7
/* Conclusion
"http.HandlerFunc" Wrappers let you do things befoer and/or after your original handler functions are called,
providing pretty powerful middleware like capabilities using only the Go Standard Library.
*/

func MyWrapper(fn http.HandlerFunc) http.HandlerFunc {
	// called once per wrapping
	return func(w http.ResponseWriter, r *http.Request) {
		//called for each request
		//defer clean-up
		fn(w, r) //call original
	}
}
