package github

import (
	"encoding/json"
	"testing"
)

func TestRepoPermissionLevel(t *testing.T) {
	get := func(v RepoPermissionLevel) *RepoPermissionLevel {
		return &v
	}
	cases := []struct {
		input    string
		expected *RepoPermissionLevel
	}{
		{
			"admin",
			get(Admin),
		},
		{
			"write",
			get(Write),
		},
		{
			"read",
			get(Read),
		},
		{
			"none",
			get(None),
		},
		{
			"",
			nil,
		},
		{
			"unknown",
			nil,
		},
	}
	for _, tc := range cases {
		var actual RepoPermissionLevel
		err := json.Unmarshal([]byte("\""+tc.input+"\""), &actual)
		switch {
		case err == nil && tc.expected == nil:
			t.Errorf("%s: failed to receive an error", tc.input)
		case err != nil && tc.expected != nil:
			t.Errorf("%s: unexpected error: %v", tc.input, err)
		case err == nil && *tc.expected != actual:
			t.Errorf("%s: actual %v != expected %v", tc.input, tc.expected, actual)
		}
	}
}

func TestLevelFromPermissions(t *testing.T) {
	var testCases = []struct {
		permissions RepoPermissions
		level       RepoPermissionLevel
	}{
		{
			permissions: RepoPermissions{},
			level:       None,
		},
		{
			permissions: RepoPermissions{Pull: true},
			level:       Read,
		},
		{
			permissions: RepoPermissions{Pull: true, Push: true},
			level:       Write,
		},
		{
			permissions: RepoPermissions{Pull: true, Push: true, Admin: true},
			level:       Admin,
		},
	}

	for _, testCase := range testCases {
		if actual, expected := LevelFromPermissions(testCase.permissions), testCase.level; actual != expected {
			t.Errorf("got incorrect level from permissions, expected %v but got %v", expected, actual)
		}
	}
}

func TestPermissionsFromTeamPermission(t *testing.T) {
	var testCases = []struct {
		level       TeamPermission
		permissions RepoPermissions
	}{
		{
			level:       TeamPermission("foobar"),
			permissions: RepoPermissions{},
		},
		{
			level:       RepoPull,
			permissions: RepoPermissions{Pull: true},
		},
		{
			level:       RepoPush,
			permissions: RepoPermissions{Pull: true, Push: true},
		},
		{
			level:       RepoAdmin,
			permissions: RepoPermissions{Pull: true, Push: true, Admin: true},
		},
	}

	for _, testCase := range testCases {
		if actual, expected := PermissionsFromTeamPermission(testCase.level), testCase.permissions; actual != expected {
			t.Errorf("got incorrect permissions from team permissions, expected %v but got %v", expected, actual)
		}
	}
}
