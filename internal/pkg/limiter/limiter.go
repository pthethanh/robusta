package limiter

import (
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

type (
	// Visitor Create a custom Visitor struct which holds the rate limiter for each
	// Visitor and the last time that the Visitor was seen.
	Visitor struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	Store interface {
		Set(string, *Visitor) error
		Get(string) (*Visitor, error)
		Del(string) error
		Clean(t time.Duration) error
	}

	Config struct {
		Rate               rate.Limit    `envconfig:"LIMITER_RATE" default:"200"`
		Burst              int           `evnconfig:"LIMITER_BURST" default:"1000"`
		TrustForwardHeader bool          `envconfig:"LIMITER_TRUST_FORWARD_HEADER" default:"false"`
		Interval           time.Duration `envconfig:"LIMITER_INTERVAL" default:"5m"`
		MaxAge             time.Duration `envconfig:"LIMITER_MAX_AGE" default:"15m"`
	}

	Limiter struct {
		store Store
		conf  Config
		mu    *sync.Mutex
	}
)

func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

func New(conf Config) *Limiter {
	return &Limiter{
		conf:  conf,
		store: NewMemStore(),
		mu:    &sync.Mutex{},
	}
}

func (l *Limiter) UseStore(store Store) {
	l.store = store
}

func (l *Limiter) addVisitor(ip string) *rate.Limiter {
	limiter := rate.NewLimiter(l.conf.Rate, l.conf.Burst)
	l.mu.Lock()
	l.store.Set(ip, &Visitor{limiter, time.Now()})
	l.mu.Unlock()
	return limiter
}

func (l *Limiter) getVisitor(ip string) *rate.Limiter {
	l.mu.Lock()
	v, err := l.store.Get(ip)
	if err != nil || v == nil {
		l.mu.Unlock()
		return l.addVisitor(ip)
	}
	// Update the last seen time for the Visitor.
	v.lastSeen = time.Now()
	l.store.Set(ip, v)
	l.mu.Unlock()
	return v.limiter
}

// Every minute check the map for Visitors that haven't been seen for
// more than 3 minutes and delete the entries.
func (l *Limiter) cleanupVisitors() {
	for {
		time.Sleep(l.conf.Interval)
		l.mu.Lock()
		if err := l.store.Clean(l.conf.MaxAge); err != nil {
			log.Errorf("failed to clean up, err: %v", err)
		}
		l.mu.Unlock()
	}
}

func (l *Limiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := string(getIP(r, l.conf.TrustForwardHeader))
		limiter := l.getVisitor(ip)
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// getIP returns IP address from request.
// If options is defined and TrustForwardHeader is true, it will lookup IP in
// X-Forwarded-For and X-Real-IP headers.
func getIP(r *http.Request, trustForwardHeader bool) net.IP {
	if trustForwardHeader {
		ip := r.Header.Get("X-Forwarded-For")
		if ip != "" {
			parts := strings.SplitN(ip, ",", 2)
			part := strings.TrimSpace(parts[0])
			return net.ParseIP(part)
		}
		ip = strings.TrimSpace(r.Header.Get("X-Real-IP"))
		if ip != "" {
			return net.ParseIP(ip)
		}
	}

	remoteAddr := strings.TrimSpace(r.RemoteAddr)
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return net.ParseIP(remoteAddr)
	}

	return net.ParseIP(host)
}
