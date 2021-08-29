package domain

var (
	_ ChatMember = ChatMemberOwner{}
	_ ChatMember = ChatMemberAdministrator{}
	_ ChatMember = ChatMemberMember{}
	_ ChatMember = ChatMemberRestricted{}
	_ ChatMember = ChatMemberLeft{}
	_ ChatMember = ChatMemberBanned{}
)

type ChatMember interface{ isChatMember() }
