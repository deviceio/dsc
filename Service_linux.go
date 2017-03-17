package dsc

import "fmt"

func (t *Service) Apply() (bool, error) {
	return false, fmt.Errorf("Not Yet Supported")
}

func (t *Service) Refresh() error {
	return fmt.Errorf("Not Yet Supported")
}

func (t *Service) Requires() []string {
	if t.Relation == nil || t.Relation.Require == nil {
		return []string{}
	}

	return t.Relation.Require
}

func (t *Service) Notifies() []string {
	if t.Relation == nil || t.Relation.Notify == nil {
		return []string{}
	}

	return t.Relation.Notify
}

func (t *Service) Refreshes() []string {
	if t.Relation == nil || t.Relation.Refresh == nil {
		return []string{}
	}

	return t.Relation.Refresh
}
