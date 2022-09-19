package envt

import "godevil/src/utils"

func GetKernelVersion() (string, error) {
	return utils.RunCommand(cmd_kernel_version[0], cmd_kernel_version[1:]...)
}
