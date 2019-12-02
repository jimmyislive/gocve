package cli

import (
	"os/exec"
	"testing"
)

func TestConfigSet(t *testing.T) {

	// valid case
	out, err := exec.Command("../../../_output/bin/linux/amd64/gocve-linux-amd64", "config", "set-db", "--dbType", "postgres", "--dbHost", "testing", "--dbPort", "5432", "--dbUser", "postgres", "--tableName", "cve", "--config", "/tmp/gocve-test.yaml").Output()
	if err != nil {
		t.Errorf("Error while setting config %s %s", err, out)
	}

	// invalid db
	out, err = exec.Command("../../../_output/bin/linux/amd64/gocve-linux-amd64", "config", "set-db", "--dbType", "xxx").Output()
	if err == nil {
		t.Errorf("Error while setting config %s %s", err, out)
	}
}
