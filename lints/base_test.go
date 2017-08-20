package lints

import "testing"

func TestAllLintsHaveUpdateReport(t *testing.T) {
	for name, lint := range Lints {
		if lint.updateReport == nil {
			t.Errorf("lint %s has nil updateReport", name)
		}
	}
}

func TestAllLintsHaveNameDescriptionProvenance(t *testing.T) {
	for name, lint := range Lints {
		if lint.Name == "" {
			t.Errorf("lint %s has empty name", name)
		}
		if lint.Description == "" {
			t.Errorf("lint %s has empty description", name)
		}
		if lint.Provenance == "" {
			t.Errorf("lint %s has empty providence", name)
		}
	}
}
