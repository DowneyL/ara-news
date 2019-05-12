package helper

import "strings"

type Platform int

const (
	iOS     Platform = 1
	android Platform = 3
)

const (
	iOSString     string = "ios"
	androidString string = "android"
)

func GetAttrSetId(platform string) Platform {
	platform = strings.ToLower(platform)
	switch platform {
	case "ios":
		return iOS
	case "android":
		return android
	default:
		return iOS
	}
}

func (p *Platform) String() string {
	switch *p {
	case iOS:
		return iOSString
	case android:
		return androidString
	default:
		return iOSString
	}
}
