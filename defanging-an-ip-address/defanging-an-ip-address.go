package defanging_an_ip_address

import "strings"

func defangIPaddr(address string) string {
	d := strings.Split(address,".")

	return d[0]+"[.]"+d[1]+"[.]"+d[2]+"[.]"+d[3]
}

//https://leetcode.com/problems/defanging-an-ip-address/

func main() {
	
}
