package accessv1alpha1

import "fmt"

func (u *User) Display() string {
	if u.Name != "" && u.Email != "" {
		return fmt.Sprintf("%s <%s>", u.Name, u.Email)
	}

	if u.Name != "" {
		return u.Name
	}

	if u.Email != "" {
		return u.Email
	}

	if u.Eid == nil {
		return ""
	}

	return u.Eid.Display()
}

func (u *NamedEID) Display() string {
	if u.Name != "" {
		return u.Name
	}

	if u.Eid == nil {
		return ""
	}

	return u.Eid.Display()
}
