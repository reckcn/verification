package slide_verification

import (
	"sync"
	"fmt"
	"io"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"
)

// 定义一个全局的session管理器
type Manager struct {
	cookieName string
	lock sync.Mutex
	provider Provider
	maxLifeTime int64

}

var provides = make(map[string]Provider)

func NewManager(provideName,cookieName string,maxLifeTime int64) (*Manager,error) {
	provider,ok := provides[provideName]
	if !ok{
		return nil,fmt.Errorf("session:unKnown provide %q (forgotten import?)",provideName)
	}
	return &Manager{provider:provider,cookieName:cookieName,maxLifeTime:maxLifeTime},nil
}

func Register(name string,provider Provider)  {
	if provider == nil{
		panic("session: Register provide is nil")
	}
	if _,dup :=provides[name];dup{
		panic("session: Register called twice for provide "+name)
	}
	provides[name] = provider

}

// 全局的唯一的sessionID
func (manager *Manager) sessionId() string  {
	b := make([]byte, 32)
	if _,err := io.ReadFull(rand.Reader,b);err != nil{
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// session创建
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge:0}
		http.SetCookie(w, &cookie)
		} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
		}
		return
	}

// 表征session管理器底层存储结构
type Provider interface {
	SessionInit(sid string) (Session,error)
	SessionRead(sid string) (Session,error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

// session 接口
type Session interface {
	Set(key,value interface{}) error // set session value
	Get(key interface{}) interface{} // get session value
	Delete(key interface{}) error //delete session value
	SessionID() string // back current sessionID
}
