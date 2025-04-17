package query

import (
	"net"
	"strings"
	"unicode/utf8"
)

type IPType uint

const (
	Single  = iota + 1 // 单IP 192.168.0.1
	Segment            // IP段 192.168.0.1-192.168.0.100
	Mask               // IP掩码 192.168.0.1/24
)

const (
	IPAddressError = "ip address error"
)

type IPSegment struct {
	Raw   string // 原始 IP
	Begin string // IP 起 (10 进制)
	End   string // IP 止 (10 进制)
	Type  IPType
}

func NewIPSegment(raw string) (*IPSegment, error) {
	ipSegment := &IPSegment{Raw: raw}
	switch {
	case raw == "":
		return nil, errors.New(IPAddressError)
	case strings.Contains(raw, "-") && strings.Contains(raw, "/"):
		return nil, errors.New(IPAddressError)
	case strings.Contains(raw, "-"):
		ipSegment.Type = Segment
	default:
		_, network, err := net.ParseCIDR(raw)
		if err != nil {
			return nil, err
		}
		if network == nil {
			ipSegment.Type = Single
		} else {
			ipSegment.Type = Mask
		}
	}
	return ipSegment, nil
}

func (s *IPSegment) Build() error {
	switch s.Type {
	case Single:
		if err := CheckIPAddr(s.Raw); err != nil {
			return err
		}
		s.Begin = strings.TrimSpace(s.Raw)
		s.End = strings.TrimSpace(s.Raw)
	case Segment:
		segments := strings.Split(s.Raw, "-")
		if len(segments) != 2 {
			return errors.New("ip raw error")
		}
		for _, seg := range segments {
			if err := CheckIPAddr(strings.TrimSpace(seg)); err != nil {
				return err
			}
		}
		if strings.TrimSpace(segments[1]) < strings.TrimSpace(segments[0]) {
			return errors.New("ip segment error")
		}
		s.Begin = strings.TrimSpace(segments[0])
		s.End = strings.TrimSpace(segments[1])
	case Mask:
		start, end, err := GetIpRangeFromNetworkSegment(s.Raw)
		if err != nil {
			return errors.New("network segment parse failed")
		}
		s.Begin = start
		s.End = end
	default:
		return errors.New("bad ip type")
	}
	return nil
}

func CheckIPAddr(ip string) error {
	address := net.ParseIP(ip)
	if address == nil {
		return errors.New(IPAddressError)
	}
	return nil
}

func GetIpRangeFromNetworkSegment(networkSegment string) (string, string, error) {
	ip, network, err := net.ParseCIDR(networkSegment)
	if err != nil {
		return "", "", errors.Wrap(err, "parse cidr failed")
	}
	firstIP := ip.Mask(network.Mask)
	var lastIP net.IP
	if ip.To4() != nil {
		lastIP = net.IPv4(0, 0, 0, 0).To4()
		for i := range network.Mask {
			lastIP[i] = firstIP[i] | ^network.Mask[i]
		}
		return firstIP.String(), lastIP.String(), nil
	} else {
		lastIP = net.ParseIP("0:0:0:0:0:0:0:0")
		for i := range network.Mask {
			lastIP[i] = firstIP[i] | ^network.Mask[i]
		}
		return Ipv6Recover(firstIP.String()), Ipv6Recover(lastIP.String()), nil
	}
}

// Ipv6Recover 将压缩后的 ipv6 地址恢复为标准格式
func Ipv6Recover(ip string) string {
	var recoveredIPSplits []string
	ipSplits := strings.Split(ip, ":")
	for _, val := range ipSplits {
		switch len(val) {
		case 4:
			recoveredIPSplits = append(recoveredIPSplits, val)
		case 3:
			recoveredIPSplits = append(recoveredIPSplits, "0"+val)
		case 2:
			recoveredIPSplits = append(recoveredIPSplits, "00"+val)
		case 1:
			recoveredIPSplits = append(recoveredIPSplits, "000"+val)
		case 0:
			continue
		}
	}
	t := 8 - len(recoveredIPSplits)
	if len(recoveredIPSplits) < 8 {
		for i := 0; i < t; i++ {
			recoveredIPSplits = append(recoveredIPSplits, "0000")
		}
	}
	return strings.Join(recoveredIPSplits, ":")
}

func EscapeStringValue(value []byte) string {
	var builder strings.Builder
	for len(value) > 0 {
		r, size := utf8.DecodeRune(value)

		switch r {
		case '\\':
			builder.WriteRune(92)
			builder.WriteRune(92)
		case '\r':
			// write "\\r"
			builder.WriteRune(rune(92))
			builder.WriteRune(rune(92))
			builder.WriteRune(114)
		case '\n':
			// write "\\n"
			builder.WriteRune(rune(92))
			builder.WriteRune(rune(92))
			builder.WriteRune(rune(110))
		case '\t':
			// write "\\t"
			builder.WriteRune(rune(92))
			builder.WriteRune(rune(92))
			builder.WriteRune(rune(116))
		case '\'':
			builder.WriteRune(39)
			builder.WriteRune(39)
			// builder.WriteRune(rune(92))
			// builder.WriteRune(rune(92))
			// builder.WriteRune(39)
		default:
			builder.WriteRune(r)
		}
		value = value[size:]
	}
	return builder.String()
}

// unescape
// 针对高级查询中的字段值进行反转义
// 由于传输JSON字符串需要转义，所以前端需要先将 " -> \", ' -> ”
// sql解析，然后再经过反转义返回之前的字符串
// "ABCD_2132_language=sc'.print(md5(817046143)).'" -> "ABCD_2132_language=sc\\'.print(md5(817046143)).\\'" -> "ABCD_2132_language=sc'.print(md5(817046143)).'"
// "JSESSIONID=EA5468BAD9025928650C0E4D810304D9; loginPageURL=""" -> "JSESSIONID=EA5468BAD9025928650C0E4D810304D9; loginPageURL=\\\"\\\"" -> "ABCD_2132_language=sc'.print(md5(817046143)).'"
func unescape(s string) string {
	s1 := strings.ReplaceAll(s, "\\'", "'")
	s2 := strings.ReplaceAll(s1, "\\\"", `"`)
	s3 := strings.ReplaceAll(s2, "\\r", "\r")
	s4 := strings.ReplaceAll(s3, "\\n", "\n")
	s5 := strings.ReplaceAll(s4, "\\t", "\t")
	return s5
}
