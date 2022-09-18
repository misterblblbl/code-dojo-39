package main

import (
	"reflect"
	"testing"
)

func Test_getSong(t *testing.T) {
	expectedSong := `There was an old lady who swallowed a fly.
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a spider;
That wriggled and wiggled and tickled inside her.
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a bird;
How absurd to swallow a bird.
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a cat;
Fancy that to swallow a cat!
She swallowed the cat to catch the bird,
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a dog;
What a hog, to swallow a dog!
She swallowed the dog to catch the cat,
She swallowed the cat to catch the bird,
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a cow;
I don't know how she swallowed a cow!
She swallowed the cow to catch the dog,
She swallowed the dog to catch the cat,
She swallowed the cat to catch the bird,
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a horse...
...She's dead, of course!`

	tests := []struct {
		name  string
		want  string
		parts []part
	}{
		{
			"generates song",
			expectedSong,
			[]part{
				{
					"fly",
					".",
					"I don't know why she swallowed a fly - perhaps she'll die!",
					false,
				},
				{
					"spider",
					";",
					"That wriggled and wiggled and tickled inside her.",
					true,
				},
				{
					"bird",
					";",
					"How absurd to swallow a bird.",
					true,
				},
				{
					"cat",
					";",
					"Fancy that to swallow a cat!",
					true,
				},
				{
					"dog",
					";",
					"What a hog, to swallow a dog!",
					true,
				},
				{
					"cow",
					";",
					"I don't know how she swallowed a cow!",
					true,
				},
				{
					"horse",
					"...",
					"...She's dead, of course!",
					false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateSong(tt.parts)
			equal := compareResults(got, tt.want)

			t.Log(len(got), len(tt.want))

			if !equal {
				t.Errorf("generateSong() = %v \n want %v", got, tt.want)
			}
		})
	}
}

func compareResults(got string, want string) bool {
	return got == want
}

func Test_generateChorus(t *testing.T) {
	tests := []struct {
		name    string
		animals []string
		want    []string
	}{
		{
			name:    "should generate empty chorus when there are 0 animals",
			animals: []string{},
			want:    nil,
		},
		{
			name:    "should generate empty chorus when there is 1 animal",
			animals: []string{"fly"},
			want:    []string{"I don't know why she swallowed a fly - perhaps she'll die!"},
		},
		{
			name:    "should generate minimal chorus when there are 2 animals",
			animals: []string{"fly", "spider"},
			want: []string{
				"She swallowed the spider to catch the fly;",
				"I don't know why she swallowed a fly - perhaps she'll die!",
			},
		},
		{
			name:    "base test",
			animals: []string{"fly", "spider", "bird", "cat"},
			want: []string{
				"She swallowed the cat to catch the bird,",
				"She swallowed the bird to catch the spider,",
				"She swallowed the spider to catch the fly;",
				"I don't know why she swallowed a fly - perhaps she'll die!",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateChorus(tt.animals)

			reflect.DeepEqual(got, tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateChorus() = %v, want %v", got, tt.want)
			}
		})
	}
}
