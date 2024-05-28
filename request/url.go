package request

import (
	"fmt"
	"os"
	"strings"
)

func createUrl() string {
	name := os.Getenv("FIBERY_WORKSPACE_NAME")
	id := os.Getenv("FIBERY_SYNC_SOURCE_ID")

	name = strings.ToLower(name)

	return fmt.Sprintf("https://%s.fibery.io/api/data-sync/sync-sources/%s", name, id)
}
