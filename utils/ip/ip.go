package ip

import (
	"net/http"
	"strings"
)

const (
	UnKnown = "unknown"
)

func GetUserIP(head http.Header) (ip string) {
	if xForWardedFor := head.Get("X-Forwarded-For"); xForWardedFor != "" && !strings.Contains(xForWardedFor, UnKnown) {
		xs := strings.Split(xForWardedFor, ",")
		ip = xs[0]
	} else if clientIP := head.Get("CLIENT-IP"); clientIP != "" && !strings.Contains(clientIP, UnKnown) {
		ip = clientIP
	} else if remoteAddr := head.Get("REMOTE-ADDR"); remoteAddr != "" && !strings.Contains(remoteAddr, UnKnown) {
		ip = remoteAddr
	} else if realIP := head.Get("X-Real-Ip"); realIP != "" && !strings.Contains(realIP, UnKnown) {
		ip = realIP
	}
	return
}
