package staff

import (
	"fmt"
	"time"

	"github.com/looped-dev/cms/api/models"
)

func validateInviteCode(code string, dbCode models.InviteCode) error {
	if code != dbCode.Code {
		return fmt.Errorf("Invalid invite code")
	}
	now := time.Now()
	expiry := time.Unix(int64(dbCode.Expiry.T), 0)
	if now.After(expiry) {
		return fmt.Errorf("Invite code has expired")
	}
	return nil
}
