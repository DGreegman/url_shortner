package utils

import(
	"strings"
)

// DetectDeviceType parses User_Agent to determine device type

func DetectDeviceType(userAgent string) string {
	ua := strings.ToLower(userAgent)

	mobileKeywords := []string {"mobile", "android", "iphone", "ipad", "ipod", "blackberry", "windows phone"}

	for _, keyword := range mobileKeywords {
		if strings.Contains(ua, keyword){
			return "Mobile"
		}
	}
	return "Desktop"
}

// GetCountry performs a simple country lookup (you can enhance this)
// For now, returns a placeholder. You can integrate MaxMind GeoIP or IP2Location later

func GetCountry(ip string) string {
	// TODO: implemment GeoIP lookup here
	// For now, return empty string or "Unknown"
	// You can later integrate: github.com/maxmind/geoip2-golang

	return "Unknown"
}

// GetClientIP extracts real client IP from request headers
func GetClientIP(forwarded, realIP, remoteAddr string) string {

	// Check X-Fowarded-For first (proxy/load balancer)

	if forwarded != "" {
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}

	// Check X-Real-IP
	if realIP != "" {
		return realIP
	}

	// Fall back to remote address 

	if remoteAddr != "" {
		// Remove Port if present

		if idx := strings.LastIndex(remoteAddr, ":"); idx != -1 {
			return remoteAddr[:idx]
		}
		return remoteAddr
	}
	return "unknown"
}