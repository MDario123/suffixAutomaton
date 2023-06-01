package suffy

import "testing"

func TestSuffy_CornerCases(t *testing.T) {
	suffy := New()
	{
		err := suffy.InsertString("abacabca")
		if err != nil {
			t.Errorf("unexpected error inserting string %s", "abacabca")
		}
	}
	testCases := []struct {
		name      string
		subString string
		wantErr   bool
		ans       bool
	}{
		{"Full string", "abacabca", false, true},
		{"Empty string", "", false, true},
		{"Invalid UTF-8 string", string([]byte{0xff}), true, false /*does not really matter*/},
		{"Full string plus a letter", "abacabcaa", false, false},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			b, err := suffy.IsSubstring(tt.subString)
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr is %t and err is %e", tt.wantErr, err)
			}
			if b != tt.ans {
				var shouldFind string
				if !b {
					shouldFind = "not"
				}
				t.Errorf("substring %s %s found in %s", tt.subString, shouldFind, "abacabca")
			}
		})
	}
}
