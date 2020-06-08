package access

type AccessFlag uint16

const (
	AccessPublic     AccessFlag = 0x0001
	AccessFinal      AccessFlag = 0x0010
	AccessSuper      AccessFlag = 0x0020
	AccessInterface  AccessFlag = 0x0200
	AccessAbstract   AccessFlag = 0x0400
	AccessSynthetic  AccessFlag = 0x1000
	AccessAnnotation AccessFlag = 0x2000
	AccessEnum       AccessFlag = 0x4000
)

var allFlags = []AccessFlag{
	AccessPublic,
	AccessFinal,
	AccessSuper,
	AccessInterface,
	AccessAbstract,
	AccessSynthetic,
	AccessAnnotation,
	AccessEnum,
}

func GetFlags(mask uint16) (flags []AccessFlag) {
	for _, flag := range allFlags {
		if mask & uint16(flag) == uint16(flag) {
			flags = append(flags, flag)
		}
	}

	return
}
