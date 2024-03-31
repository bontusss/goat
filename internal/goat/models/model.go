package models

// CustomFields defines an interface for user-specific data
type CustomFields interface {
	// Get and set the user fields (optional)
	GetName() string
	SetName(name string)
	GetBio() string
	SetBio(bio string)
	// define methods for accessing or manipulating custom data here (optional)
}

type UserCustomFields struct {
	Name *string `json:"name,omitempty"`
	Bio  *string `json:"bio,omitempty"`
}

type User struct {
	ID           uint         `json:"id"`
	Email        string       `json:"email"`
	Password     string       `json:"password"`
	CustomFields CustomFields `json:"custom_fields,omitempty"`
}

// GetName implements the CustomFields interface
func (f *UserCustomFields) GetName() string {
	if f.Name == nil {
	  return ""
	}
	return *f.Name
  }

  // SetName implements the CustomFields interface
  func (f *UserCustomFields) SetName(name string) {
	f.Name = &name
  }

  // GetBio implements the CustomFields interface
  func (f *UserCustomFields) GetBio() string {
	if f.Bio == nil {
	  return ""
	}
	return *f.Bio
  }

  // SetBio implements the CustomFields interface
  func (f *UserCustomFields) SetBio(bio string) {
	f.Bio = &bio
  }