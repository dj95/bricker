package remoteswitch

const (
	function_get_switching_state = uint8(2) // working
	function_set_repeats         = uint8(4)
	function_get_repeats         = uint8(5)
	function_switch_socket_a     = uint8(6) // working
	function_switch_socket_b     = uint8(7) // working
	function_dim_socket_b        = uint8(8)
	function_switch_socket_c     = uint8(9) // working
	function_get_identity        = uint8(255)
)
