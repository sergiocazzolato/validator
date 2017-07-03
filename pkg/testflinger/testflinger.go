package testflinger

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/fgimenez/validator/pkg/types"
)

const (
	FromTargetFmt = `job_queue: dragonboard
provision_data:
    channel: %s
test_data:
    test_cmds:
        - sudo apt update && sudo apt install -y git curl
        - git clone https://github.com/snapcore/snapd
        - curl -s -O https://niemeyer.s3.amazonaws.com/spread-amd64.tar.gz && tar xzvf spread-amd64.tar.gz
        - snapd/tests/lib/external/prepare-ssh.sh {device_ip} 22 ubuntu
        - cd snapd && export SPREAD_EXTERNAL_ADDRESS={device_ip}:22 && git checkout %s && ../spread -v %s
`
	FromStableFmt = `job_queue: dragonboard
provision_data:
    channel: stable
test_data:
    test_cmds:
        - sudo apt update && sudo apt install -y git curl
        - git clone https://github.com/snapcore/snapd
        - curl -s -O https://niemeyer.s3.amazonaws.com/spread-amd64.tar.gz && tar xzvf spread-amd64.tar.gz
        - snapd/tests/lib/external/prepare-ssh.sh {device_ip} 22 ubuntu
        - ssh -q -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no ubuntu@{device_ip} sudo snap refresh --%s core
        - cd snapd && export SPREAD_EXTERNAL_ADDRESS={device_ip}:22 && git checkout %s && ../spread -v %s
`
)

type Testflinger struct{}

func (t *Testflinger) GenerateCfg(options *types.Options, input [][]string) []string {
	var result []string

	var tpl string
	if options.From == "stable" {
		tpl = FromStableFmt
	} else {
		tpl = FromTargetFmt
	}

	for _, item := range input {
		mergedLines := strings.Join(item, " ")
		content := []byte(fmt.Sprintf(tpl, options.Channel, options.Release, mergedLines))

		tmpfile, _ := ioutil.TempFile("", "")
		if _, err := tmpfile.Write(content); err != nil {
			log.Fatal(err)
		}
		if err := tmpfile.Close(); err != nil {
			log.Fatal(err)
		}
		result = append(result, tmpfile.Name())
	}
	return result
}
