package composition

import (
	"io"
	"os"
	"testing"
)

func TestHuman_Greet(t *testing.T) {
	type fields struct {
		Name string
		Age  uint
		Want string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "first",
			fields: fields{
				Name: "Bob",
				Age:  8,
				Want: "Hello! My name is Bob and I'm 8 years old\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rescStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			h := Human{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if h.Age != tt.fields.Age || h.Name != tt.fields.Name {
				t.Errorf(
					"got Age = %d, Name = %s\nwant Age = %d, Name = %s",
					h.Age,
					h.Name,
					tt.fields.Age,
					tt.fields.Name,
				)
			}
			h.Greet()
			w.Close()
			got, _ := io.ReadAll(r)
			os.Stdout = rescStdout
			if string(got) != tt.fields.Want {
				t.Errorf("got %q\nwant %q", string(got), tt.fields.Want)
			}
		})
	}
}

func TestAction_Greet(t *testing.T) {
	type fields struct {
		Name string
		Age  uint
		Want string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "first",
			fields: fields{
				Name: "Bob",
				Age:  8,
				Want: "Hello! My name is Bob and I'm 8 years old\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rescStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			a := Action{
				Human{
					Name: tt.fields.Name,
					Age:  tt.fields.Age,
				},
			}
			if a.Age != tt.fields.Age || a.Name != tt.fields.Name {
				t.Errorf(
					"got Age = %d, Name = %s\nwant Age = %d, Name = %s",
					a.Age,
					a.Name,
					tt.fields.Age,
					tt.fields.Name,
				)
			}
			a.Greet()
			w.Close()
			got, _ := io.ReadAll(r)
			os.Stdout = rescStdout
			if string(got) != tt.fields.Want {
				t.Errorf("got %q\nwant %q", string(got), tt.fields.Want)
			}
		})
	}
}

func TestAction2_Greet(t *testing.T) {
	type fields struct {
		Name string
		Age  uint
		Want string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "first",
			fields: fields{
				Name: "Bob",
				Age:  8,
				Want: "Hello! My name is Bob and I'm 8 years old\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rescStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			a := Action2{
				Human{
					Name: tt.fields.Name,
					Age:  tt.fields.Age,
				},
			}
			if a.Person.Age != tt.fields.Age || a.Person.Name != tt.fields.Name {
				t.Errorf(
					"got Age = %d, Name = %s\nwant Age = %d, Name = %s",
					a.Person.Age,
					a.Person.Name,
					tt.fields.Age,
					tt.fields.Name,
				)
			}
			a.Person.Greet()
			w.Close()
			got, _ := io.ReadAll(r)
			os.Stdout = rescStdout
			if string(got) != tt.fields.Want {
				t.Errorf("got %q\nwant %q", string(got), tt.fields.Want)
			}
		})
	}
}
