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

func TestSuffy_InsertString(t *testing.T) {
	testCases := []struct {
		name    string
		string  string
		wantErr bool
	}{
		{
			"Common string #1",
			"The quick brown fox jumps over the lazy dog",
			false,
		},
		{
			"Common string #2",
			"A rough-coated, dough-faced, thoughtful ploughman strode through the streets of Scarborough; after falling into a slough, he coughed and hiccoughed",
			false,
		},
		{
			"Common string #3",
			"This exceeding trifling witling, considering ranting criticizing concerning adopting fitting wording being exhibiting transcending learning," +
				" was displaying, notwithstanding ridiculing, surpassing boasting swelling reasoning," +
				" respecting correcting erring writing, and touching detecting deceiving arguing during debating",
			false,
		},
		{
			"Common string #4",
			"I do not know where family doctors acquired illegibly perplexing handwriting; nevertheless, extraordinary pharmaceutical intellectuality," +
				" counterbalancing indecipherability, transcendentalizes intercommunications’ incomprehensibleness",
			false,
		},
		{
			"Common string #5",
			"Buffalo buffalo Buffalo buffalo buffalo buffalo Buffalo buffalo",
			false,
		},
		{
			"Invalid UTF-8 string",
			string([]byte{0xff}),
			true,
		},
		{
			"Empty string",
			"",
			false,
		},
		{
			"Uncommon characters",
			"☺☻♥♦♣♠•◘○◙♂♀♪-.◄↕‼¶§▬↨↑↓→←∟↔▲▼!#$%&'()*+,-./日大年中会人本月長国\n",
			false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			suffy := New()
			err := suffy.InsertString(tt.string)
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr is %t and err is %e (string=%s)", tt.wantErr, err, tt.string)
			}
		})
	}
}
