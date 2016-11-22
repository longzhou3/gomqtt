package gate

/* We need to transfer external user to our internal acc and user*/

func accTrans(ci *connInfo) {
	ci.acc = ci.cp.Username()
	ci.user = ci.cp.ClientId()
}
