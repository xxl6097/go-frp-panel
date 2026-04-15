package sub

import (
	"fmt"
)

func Run(cfgFile string, cfgDir string) error {
	if cfgDir != "" {
		go func() {
			_ = runMultipleClients(cfgDir)
		}()
	}
	err := runClient(cfgFile)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
