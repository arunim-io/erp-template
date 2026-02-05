package session

import (
	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"

	"github.com/arunim-io/erp-template/internal/config"
	"github.com/arunim-io/erp-template/internal/database"
)

func New(db *database.DB, secure bool, cfg *config.SessionCookieConfig) *scs.SessionManager {
	sm := scs.New()

	sm.Cookie.Secure = secure
	sm.Cookie.Partitioned = secure
	sm.Lifetime = cfg.Lifetime
	sm.Store = pgxstore.New(db.Pool)

	return sm
}
