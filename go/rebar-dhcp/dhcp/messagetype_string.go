// generated by stringer -type=MessageType; DO NOT EDIT

package dhcp

import "fmt"

const _MessageType_name = "DiscoverOfferRequestDeclineACKNAKReleaseInform"

var _MessageType_index = [...]uint8{0, 8, 13, 20, 27, 30, 33, 40, 46}

func (i MessageType) String() string {
	i -= 1
	if i+1 >= MessageType(len(_MessageType_index)) {
		return fmt.Sprintf("MessageType(%d)", i+1)
	}
	return _MessageType_name[_MessageType_index[i]:_MessageType_index[i+1]]
}
