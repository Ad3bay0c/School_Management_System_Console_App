package students

import (
	"testing"
)

func TestApplicant_CheckStatus(t *testing.T) {
	applicants := []struct{
		input		Applicant
		output		bool
		word 		string
	}{
		{
			input: Applicant{
				ID: 2,
			},
			output: true,
			word: "Check Status 1",
		},
		{
			input: Applicant{
				ID: 20,
			},
			output: false,
			word: "Check Status 1",
		},
	}

	for _, applicant := range applicants {
		t.Run(applicant.word, func(t *testing.T) {
			result := applicant.input.CheckStatus()
			if result != applicant.output {
				t.Errorf("Expected %v; Got %v", applicant.output, result)
			}
		})
	}
}