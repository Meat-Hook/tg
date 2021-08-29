package domain

// ChatMember this object contains information about one member of a chat.
// Currently, the following 6 types of chat members are supported:
// ChatMemberOwner
// ChatMemberAdministrator
// ChatMemberMember
// ChatMemberRestricted
// ChatMemberLeft
// ChatMemberBanned
type ChatMember interface{ isChatMember() }

var (
	_ ChatMember = ChatMemberOwner{}
	_ ChatMember = ChatMemberAdministrator{}
	_ ChatMember = ChatMemberMember{}
	_ ChatMember = ChatMemberRestricted{}
	_ ChatMember = ChatMemberLeft{}
	_ ChatMember = ChatMemberBanned{}
)
