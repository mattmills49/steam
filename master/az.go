package master

import (
	"fmt"
	"log"
	"net/http"

	auth "github.com/abbot/go-http-auth"
	"github.com/h2oai/steamY/master/az"
	"github.com/h2oai/steamY/master/data"
)

type DefaultAz struct {
	ds *data.Datastore
}

func newDefaultAz(ds *data.Datastore) *DefaultAz {
	return &DefaultAz{ds}
}

func (a *DefaultAz) Authenticate(username string) string {
	pz, err := a.ds.NewPrincipal(username)
	if err != nil {
		log.Printf("User %s read failed: %s\n", username, err)
		return ""
	}

	if pz == nil {
		log.Printf("User %s does not exist\n", username)
		return ""
	}
	log.Println("User logged in:", username)
	return pz.Password()
}

func (a *DefaultAz) Identify(r *http.Request) (az.Principal, error) {
	username := r.Header.Get(auth.AuthUsernameHeader)
	log.Println("User identified:", username)
	pz, err := a.ds.NewPrincipal(username)
	if err != nil {
		return nil, err
	}

	if pz == nil {
		return nil, fmt.Errorf("User %s does not exist\n", username)
	}

	return pz, nil
}

func serveNoop(w http.ResponseWriter, r *http.Request) {}
func authNoop(user, realm string) string               { return "" }
