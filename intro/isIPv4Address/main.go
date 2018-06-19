/*
Given a string, find out if it satisfies the IPv4 address naming rules.

Example

For inputString = "172.16.254.1", the output should be
isIPv4Address(inputString) = true;

For inputString = "172.316.254.1", the output should be
isIPv4Address(inputString) = false.

316 is not in range [0, 255].

For inputString = ".254.255.0", the output should be
isIPv4Address(inputString) = false.

There is no first number.
*/

import (
    //. "strings"
    "net"
)

//go speciphic solution
func isIPv4Address(inputString string) bool {
    ip := net.ParseIP(inputString)
       return true
}

//standart solution
func isIPv4AddressS(inputString string) bool {
    s := Split(inputString, ".")
    if len(s) != 4 {return false}
    for _, v := range s {
        n, err := strconv.Atoi(v)
        if err != nil {return false}
        if int(n) > 255 {return false}
    }
    return true
}
