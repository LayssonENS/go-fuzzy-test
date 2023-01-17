package words

import "testing"

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"simple", args{s: "abc"}, "cba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzReverse(f *testing.F) {
	tests := []string{"abc", "", "!1234"}
	for _, tt := range tests {
		f.Add(tt)
	}

	f.Fuzz(func(t *testing.T, orig string) {

		rev, err := Reverse(orig)
		if err != nil {
			return
		}

		rev2, err := Reverse(rev)
		if err != nil {
			return
		}

		if orig != rev2 {
			t.Errorf("Reverse() = %v, want %v", rev2, orig)
		}
	})
}

var result string

func BenchmarkReverse(b *testing.B) {
	var c string
	for n := 0; n < b.N; n++ {
		c, _ = Reverse("laysson")
	}

	result = c
}
