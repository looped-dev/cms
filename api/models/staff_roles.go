package models

type StaffRole string

const (
	// Trusted staff user who should be able to manage all content and users, as well
	// as site settings and options.
	StaffRoleAdministrator StaffRole = "ADMINISTRATOR"
	// Can invite and manage other Authors and Contributors, as well as edit and
	// publish any posts on the site.
	StaffRoleEditor StaffRole = "EDITOR"
	// A trusted user who can create, edit and publish their own posts, but can’t
	// modify others.
	StaffRoleAuthor StaffRole = "AUTHOR"
)

func (s StaffRole) String() string {
	return string(s)
}

func (e StaffRole) IsValid() bool {
	switch e {
	case StaffRoleAdministrator, StaffRoleEditor, StaffRoleAuthor:
		return true
	}
	return false
}
