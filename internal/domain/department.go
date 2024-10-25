package domain

type Department struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id,omitempty"`
	Flags    int8   `json:"flags"`
}

func (d *Department) IsActive() bool {
	return d.Flags&1 != 0
}

func (d *Department) IsDeleted() bool {
	return d.Flags&2 != 0
}

func (d *Department) IsApproved() bool {
	return d.Flags&4 != 0
}
