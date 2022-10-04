package envt

var (
	// cmd_list_files     string = "ls"
	cmd_kernel_version    = []string{"uname", "-r"}
	cmd_list_kern_modules = "lsmod | awk '{ print $1 }'"
)
