package github

import (
	"testing"
)

func TestIssueCapsLogin(t *testing.T) {
	// some valid logins that should all normalize to match the first
	validLoginVariants := []string{
		"BenTheElder",
		"BENTHEELDER",
		"bentheelder",
		"BenTHEElder",
		"BeNtHeElDeR",
		"bEnThEeLdEr",
	}
	// add an explicitly normalized version for sanity
	validLoginVariants = append(validLoginVariants, NormLogin(validLoginVariants[0]))

	issue := Issue{
		User: User{
			Login: validLoginVariants[0],
		},
		Assignees: []User{
			{
				Login: validLoginVariants[0],
			},
		},
	}
	for _, login := range validLoginVariants {
		if !issue.IsAuthor(login) {
			t.Errorf("expected issue.IsAuthor(%s) to be true", login)
		}
		if !issue.IsAssignee(login) {
			t.Errorf("expected issue.IsAssignee(%s) to be true", login)
		}
	}
}

func TestUnmarshalClientError(t *testing.T) {
	var testcases = []struct {
		name string
		body string
	}{
		{
			name: "invalid JSON",
			body: `{"message":"Problems parsing JSON"}`,
		},
		{
			name: "wrong type of JSON values",
			body: `{"message":"Body should be a JSON object"}`,
		},
		{
			name: "invalid fields",
			body: `{
				"message": "Validation Failed",
				"errors": [
				  {
					"resource": "Issue",
					"field": "title",
					"code": "missing_field"
				  }
				]
			  }`,
		},
		{
			name: "requires authentication",
			body: `{
				"message": "Requires authentication",
				"documentation_url": "https://developer.github.com/v3"
			  }`,
		},
		{
			name: "validation failed, position is invalid",
			body: `{
				"message": "Validation Failed",
				"errors": [
				  "Position is invalid"
				],
				"documentation_url": "https://developer.github.com/v3/pulls/reviews/#create-a-pull-request-review"
			  }`,
		},
	}
	for _, tc := range testcases {
		b := []byte(tc.body)
		err := unmarshalClientError(b)
		_, isClientError := err.(ClientError)
		_, isAlternativeClientError := err.(AlternativeClientError)
		if !(isClientError || isAlternativeClientError) {
			t.Errorf("For case %s, json.Unmarshal error: %v", tc.name, err)
		}
	}
}
