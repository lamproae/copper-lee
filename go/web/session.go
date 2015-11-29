package main

import (
	"fmt"
	"sync"
)

type Manager struct {
	cookieName string
	lock sync.Mutex
	provider Provider
	maxlifetime int64
}

func (manager *Manager) SessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager)SessionStart(w http.ResponseWriter, r *http.Request)(session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value:url.QueryEscape(sid), Path:"/", HttpOnly:true, MaxAge:int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ = url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)n
	}

	return
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxlifetime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key, interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

var providers = make(map[string]Provider)

func Register(name string, provider, Provider) {
	if provider == nil {
		panic("Session: Register provide is nil")
	}

	if _, dup := providers[name]; dup {
		panic("Session: Register called twice for provide " + name)
	}

	providers[name] = provider
}

func NewManager(providerName, cookieName string, maxlifetime int64)(*Manager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provider %q(forgotten import?)", providerName)
	}
	
	return &Manager{provider:provider, cookieName:cookieName, maxlifetime:maxlifetime}, nil
}

var globalSessons *session.Manager

func init() {
	globalSessons = NewManager("memory", "gosessionid", 3600)
}
