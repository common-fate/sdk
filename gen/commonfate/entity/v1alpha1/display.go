package entityv1alpha1

import (
	"fmt"
	"strings"
)

func (u *EID) Display() string {
	if strings.Contains(u.Id, " ") {
		return fmt.Sprintf(`%s::"%s"`, u.Type, u.Id)
	}

	return fmt.Sprintf(`%s::%s`, u.Type, u.Id)
}
