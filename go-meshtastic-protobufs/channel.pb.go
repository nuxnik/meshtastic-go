// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: channel.proto

package meshtastic_go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//
// Standard predefined channel settings
// Note: these mappings must match ModemConfigChoice in the device code.
type ChannelSettings_ModemConfig int32

const (
	//
	// < Bw = 125 kHz, Cr = 4/5, Sf(7) = 128chips/symbol, CRC
	// < on. Default medium range (5.469 kbps)
	ChannelSettings_Bw125Cr45Sf128 ChannelSettings_ModemConfig = 0
	//
	// < Bw = 500 kHz, Cr = 4/5, Sf(7) = 128chips/symbol, CRC
	// < on. Fast+short range (21.875 kbps)
	ChannelSettings_Bw500Cr45Sf128 ChannelSettings_ModemConfig = 1
	//
	// < Bw = 31.25 kHz, Cr = 4/8, Sf(9) = 512chips/symbol,
	// < CRC on. Slow+long range (275 bps)
	ChannelSettings_Bw31_25Cr48Sf512 ChannelSettings_ModemConfig = 2
	//
	// < Bw = 125 kHz, Cr = 4/8, Sf(12) = 4096chips/symbol, CRC
	// < on. Slow+long range (183 bps)
	ChannelSettings_Bw125Cr48Sf4096 ChannelSettings_ModemConfig = 3
)

// Enum value maps for ChannelSettings_ModemConfig.
var (
	ChannelSettings_ModemConfig_name = map[int32]string{
		0: "Bw125Cr45Sf128",
		1: "Bw500Cr45Sf128",
		2: "Bw31_25Cr48Sf512",
		3: "Bw125Cr48Sf4096",
	}
	ChannelSettings_ModemConfig_value = map[string]int32{
		"Bw125Cr45Sf128":   0,
		"Bw500Cr45Sf128":   1,
		"Bw31_25Cr48Sf512": 2,
		"Bw125Cr48Sf4096":  3,
	}
)

func (x ChannelSettings_ModemConfig) Enum() *ChannelSettings_ModemConfig {
	p := new(ChannelSettings_ModemConfig)
	*p = x
	return p
}

func (x ChannelSettings_ModemConfig) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChannelSettings_ModemConfig) Descriptor() protoreflect.EnumDescriptor {
	return file_channel_proto_enumTypes[0].Descriptor()
}

func (ChannelSettings_ModemConfig) Type() protoreflect.EnumType {
	return &file_channel_proto_enumTypes[0]
}

func (x ChannelSettings_ModemConfig) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChannelSettings_ModemConfig.Descriptor instead.
func (ChannelSettings_ModemConfig) EnumDescriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{0, 0}
}

//
// How this channel is being used (or not).
//
// Note: this field is an enum to give us options for the future.  In particular, someday
// we might make a 'SCANNING' option.  SCANNING channels could have different frequencies and the radio would
// occasionally check that freq to see if anything is being transmitted.
//
// For devices that have multiple physical radios attached, we could keep multiple PRIMARY/SCANNING channels active at once to allow
// cross band routing as needed.  If a device has only a single radio (the common case) only one channel can be PRIMARY at a time
// (but any number of SECONDARY channels can't be sent received on that common frequency)
type Channel_Role int32

const (
	//
	// This channel is not in use right now
	Channel_DISABLED Channel_Role = 0
	//
	// This channel is used to set the frequency for the radio - all other enabled channels must be SECONDARY
	Channel_PRIMARY Channel_Role = 1
	//
	// Secondary channels are only used for encryption/decryption/authentication purposes.  Their radio settings (freq etc)
	// are ignored, only psk is used.
	Channel_SECONDARY Channel_Role = 2
)

// Enum value maps for Channel_Role.
var (
	Channel_Role_name = map[int32]string{
		0: "DISABLED",
		1: "PRIMARY",
		2: "SECONDARY",
	}
	Channel_Role_value = map[string]int32{
		"DISABLED":  0,
		"PRIMARY":   1,
		"SECONDARY": 2,
	}
)

func (x Channel_Role) Enum() *Channel_Role {
	p := new(Channel_Role)
	*p = x
	return p
}

func (x Channel_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Channel_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_channel_proto_enumTypes[1].Descriptor()
}

func (Channel_Role) Type() protoreflect.EnumType {
	return &file_channel_proto_enumTypes[1]
}

func (x Channel_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Channel_Role.Descriptor instead.
func (Channel_Role) EnumDescriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{1, 0}
}

//
// Full settings (center freq, spread factor, pre-shared secret key etc...)
// needed to configure a radio for speaking on a particular channel This
// information can be encoded as a QRcode/url so that other users can configure
// their radio to join the same channel.
// A note about how channel names are shown to users: channelname-Xy
// poundsymbol is a prefix used to indicate this is a channel name (idea from @professr).
// Where X is a letter from A-Z (base 26) representing a hash of the PSK for this
// channel - so that if the user changes anything about the channel (which does
// force a new PSK) this letter will also change. Thus preventing user confusion if
// two friends try to type in a channel name of "BobsChan" and then can't talk
// because their PSKs will be different.  The PSK is hashed into this letter by
// "0x41 + [xor all bytes of the psk ] modulo 26"
// This also allows the option of someday if people have the PSK off (zero), the
// users COULD type in a channel name and be able to talk.
// Y is a lower case letter from a-z that represents the channel 'speed' settings
// (for some future definition of speed)
//
// FIXME: Add description of multi-channel support and how primary vs secondary channels are used.
// FIXME: explain how apps use channels for security.  explain how remote settings and
// remote gpio are managed as an example
type ChannelSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	// If zero then, use default max legal continuous power (ie. something that won't
	// burn out the radio hardware)
	// In most cases you should use zero here.  Units are in dBm.
	TxPower int32 `protobuf:"varint,1,opt,name=tx_power,json=txPower,proto3" json:"tx_power,omitempty"`
	//
	// Note: This is the 'old' mechanism for specifying channel parameters.
	// Either modem_config or bandwidth/spreading/coding will be specified - NOT BOTH.
	// As a heuristic: If bandwidth is specified, do not use modem_config.
	// Because protobufs take ZERO space when the value is zero this works out nicely.
	// This value is replaced by bandwidth/spread_factor/coding_rate.
	// If you'd like to experiment with other options add them to MeshRadio.cpp in the device code.
	ModemConfig ChannelSettings_ModemConfig `protobuf:"varint,3,opt,name=modem_config,json=modemConfig,proto3,enum=ChannelSettings_ModemConfig" json:"modem_config,omitempty"`
	//
	// Bandwidth in MHz
	// Certain bandwidth numbers are 'special' and will be converted to the
	// appropriate floating point value: 31 -> 31.25MHz
	Bandwidth uint32 `protobuf:"varint,6,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	//
	// A number from 7 to 12.  Indicates number of chirps per symbol as
	// 1<<spread_factor.
	SpreadFactor uint32 `protobuf:"varint,7,opt,name=spread_factor,json=spreadFactor,proto3" json:"spread_factor,omitempty"`
	//
	// The denominator of the coding rate.  ie for 4/8, the value is 8. 5/8 the value is 5.
	CodingRate uint32 `protobuf:"varint,8,opt,name=coding_rate,json=codingRate,proto3" json:"coding_rate,omitempty"`
	// NOTE: this field is _independent_ and unrelated to the concepts in channel.proto.
	// this is controlling the actual hardware frequency the radio is transmitting on.  In a perfect world
	// we would have called it something else (band?) but I forgot to make this change during the big 1.2 renaming.
	// Most users should never need to be exposed to this field/concept.
	// A channel number between 1 and 13 (or whatever the max is in the current
	// region). If ZERO then the rule is "use the old channel name hash based
	// algorithm to derive the channel number")
	// If using the hash algorithm the channel number will be: hash(channel_name) %
	// NUM_CHANNELS (Where num channels depends on the regulatory region).
	// NUM_CHANNELS_US is 13, for other values see MeshRadio.h in the device code.
	// hash a string into an integer - djb2 by Dan Bernstein. -
	// http://www.cse.yorku.ca/~oz/hash.html
	// unsigned long hash(char *str) {
	//   unsigned long hash = 5381; int c;
	//   while ((c = *str++) != 0)
	//     hash = ((hash << 5) + hash) + (unsigned char) c;
	//   return hash;
	// }
	ChannelNum uint32 `protobuf:"varint,9,opt,name=channel_num,json=channelNum,proto3" json:"channel_num,omitempty"`
	//
	// A simple pre-shared key for now for crypto.  Must be either 0 bytes (no
	// crypto), 16 bytes (AES128), or 32 bytes (AES256)
	// A special shorthand is used for 1 byte long psks.
	// These psks should be treated as only minimally secure,
	// because they are listed in this source code.  Those bytes are mapped using the following scheme:
	// `0` = No crypto
	// `1` = The special "default" channel key: {0xd4, 0xf1, 0xbb, 0x3a, 0x20, 0x29, 0x07, 0x59, 0xf0, 0xbc, 0xff, 0xab, 0xcf, 0x4e, 0x69, 0xbf}
	// `2` through 10 = The default channel key, except with 1 through 9 added to the last byte.  Shown to user as simple1 through 10
	Psk []byte `protobuf:"bytes,4,opt,name=psk,proto3" json:"psk,omitempty"`
	//
	// A SHORT name that will be packed into the URL.  Less than 12 bytes.
	// Something for end users to call the channel
	// If this is the empty string it is assumed that this channel
	// is the special (minimally secure) "Default"channel.
	// In user interfaces it should be rendered as a local language translation of "X".  For channel_num
	// hashing empty string will be treated as "X".
	// Where "X" is selected based on the English words listed above for ModemConfig
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	//
	// Used to construct a globally unique channel ID.  The full globally unique ID will be: "name.id"
	// where ID is shown as base36.  Assuming that the number of meshtastic users is below 20K (true for a long time)
	// the chance of this 64 bit random number colliding with anyone else is super low.  And the penalty for
	// collision is low as well, it just means that anyone trying to decrypt channel messages might need to
	// try multiple candidate channels.
	// Any time a non wire compatible change is made to a channel, this field should be regenerated.
	// There are a small number of 'special' globally known (and fairly) insecure standard channels.
	// Those channels do not have a numeric id included in the settings, but instead it is pulled from
	// a table of well known IDs.  (see Well Known Channels FIXME)
	Id uint32 `protobuf:"fixed32,10,opt,name=id,proto3" json:"id,omitempty"`
	//
	// If true, messages on the mesh will be sent to the *public* internet by any gateway ndoe
	UplinkEnabled bool `protobuf:"varint,16,opt,name=uplink_enabled,json=uplinkEnabled,proto3" json:"uplink_enabled,omitempty"`
	//
	// If true, messages seen on the internet will be forwarded to the local mesh.
	DownlinkEnabled bool `protobuf:"varint,17,opt,name=downlink_enabled,json=downlinkEnabled,proto3" json:"downlink_enabled,omitempty"`
}

func (x *ChannelSettings) Reset() {
	*x = ChannelSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelSettings) ProtoMessage() {}

func (x *ChannelSettings) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelSettings.ProtoReflect.Descriptor instead.
func (*ChannelSettings) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{0}
}

func (x *ChannelSettings) GetTxPower() int32 {
	if x != nil {
		return x.TxPower
	}
	return 0
}

func (x *ChannelSettings) GetModemConfig() ChannelSettings_ModemConfig {
	if x != nil {
		return x.ModemConfig
	}
	return ChannelSettings_Bw125Cr45Sf128
}

func (x *ChannelSettings) GetBandwidth() uint32 {
	if x != nil {
		return x.Bandwidth
	}
	return 0
}

func (x *ChannelSettings) GetSpreadFactor() uint32 {
	if x != nil {
		return x.SpreadFactor
	}
	return 0
}

func (x *ChannelSettings) GetCodingRate() uint32 {
	if x != nil {
		return x.CodingRate
	}
	return 0
}

func (x *ChannelSettings) GetChannelNum() uint32 {
	if x != nil {
		return x.ChannelNum
	}
	return 0
}

func (x *ChannelSettings) GetPsk() []byte {
	if x != nil {
		return x.Psk
	}
	return nil
}

func (x *ChannelSettings) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChannelSettings) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChannelSettings) GetUplinkEnabled() bool {
	if x != nil {
		return x.UplinkEnabled
	}
	return false
}

func (x *ChannelSettings) GetDownlinkEnabled() bool {
	if x != nil {
		return x.DownlinkEnabled
	}
	return false
}

//
// A pair of a channel number, mode and the (sharable) settings for that channel
type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	// The index of this channel in the channel table (from 0 to MAX_NUM_CHANNELS-1)
	// (Someday - not currently implemented) An index of -1 could be used to mean "set by name",
	// in which case the target node will find and set the channel by settings.name.
	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	//
	// The new settings, or NULL to disable that channel
	Settings *ChannelSettings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
	Role     Channel_Role     `protobuf:"varint,3,opt,name=role,proto3,enum=Channel_Role" json:"role,omitempty"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{1}
}

func (x *Channel) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Channel) GetSettings() *ChannelSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *Channel) GetRole() Channel_Role {
	if x != nil {
		return x.Role
	}
	return Channel_DISABLED
}

var File_channel_proto protoreflect.FileDescriptor

var file_channel_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xdc, 0x03, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x78, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x3f,
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1c, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x46, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x70, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x07, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x70,
	0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x75, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x60, 0x0a, 0x0b,
	0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x0e, 0x42,
	0x77, 0x31, 0x32, 0x35, 0x43, 0x72, 0x34, 0x35, 0x53, 0x66, 0x31, 0x32, 0x38, 0x10, 0x00, 0x12,
	0x12, 0x0a, 0x0e, 0x42, 0x77, 0x35, 0x30, 0x30, 0x43, 0x72, 0x34, 0x35, 0x53, 0x66, 0x31, 0x32,
	0x38, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x77, 0x33, 0x31, 0x5f, 0x32, 0x35, 0x43, 0x72,
	0x34, 0x38, 0x53, 0x66, 0x35, 0x31, 0x32, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x77, 0x31,
	0x32, 0x35, 0x43, 0x72, 0x34, 0x38, 0x53, 0x66, 0x34, 0x30, 0x39, 0x36, 0x10, 0x03, 0x22, 0xa2,
	0x01, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x2c, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x21,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x22, 0x30, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53,
	0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x4d, 0x41,
	0x52, 0x59, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x41, 0x52,
	0x59, 0x10, 0x02, 0x42, 0x48, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x73,
	0x76, 0x69, 0x6c, 0x6c, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x42, 0x0d, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x03, 0x5a, 0x20, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x37, 0x2f,
	0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2d, 0x67, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_channel_proto_rawDescOnce sync.Once
	file_channel_proto_rawDescData = file_channel_proto_rawDesc
)

func file_channel_proto_rawDescGZIP() []byte {
	file_channel_proto_rawDescOnce.Do(func() {
		file_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_channel_proto_rawDescData)
	})
	return file_channel_proto_rawDescData
}

var file_channel_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_channel_proto_goTypes = []interface{}{
	(ChannelSettings_ModemConfig)(0), // 0: ChannelSettings.ModemConfig
	(Channel_Role)(0),                // 1: Channel.Role
	(*ChannelSettings)(nil),          // 2: ChannelSettings
	(*Channel)(nil),                  // 3: Channel
}
var file_channel_proto_depIdxs = []int32{
	0, // 0: ChannelSettings.modem_config:type_name -> ChannelSettings.ModemConfig
	2, // 1: Channel.settings:type_name -> ChannelSettings
	1, // 2: Channel.role:type_name -> Channel.Role
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_channel_proto_init() }
func file_channel_proto_init() {
	if File_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelSettings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_channel_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_channel_proto_goTypes,
		DependencyIndexes: file_channel_proto_depIdxs,
		EnumInfos:         file_channel_proto_enumTypes,
		MessageInfos:      file_channel_proto_msgTypes,
	}.Build()
	File_channel_proto = out.File
	file_channel_proto_rawDesc = nil
	file_channel_proto_goTypes = nil
	file_channel_proto_depIdxs = nil
}
