package slack_notice

import (
	"encoding/json"
	"testing"
)

func TestSetDividerBlock(t *testing.T) {
	t.Run("without block id", func(t *testing.T) {
		want := "[{\"type\":\"divider\"}]"

		dividerBlock := SetDividerBlock(nil, "")
		got, err := json.Marshal(dividerBlock)
		if err != nil {
			t.Fatal(err)
		}

		if want != string(got) {
			t.Fatalf("want %s, got %s", want, got)
		}
	})
	t.Run("with block id", func(t *testing.T) {
		want := "[{\"type\":\"divider\",\"block_id\":\"divider1\"}]"

		dividerBlock := SetDividerBlock(nil, "divider1")
		got, err := json.Marshal(dividerBlock)
		if err != nil {
			t.Fatal(err)
		}

		if want != string(got) {
			t.Fatalf("want %s, got %s", want, got)
		}
	})
}

func TestNewDividerBlock(t *testing.T) {
	t.Run("without block id", func(t *testing.T) {
		want := "{\"type\":\"divider\"}"

		divider := NewDividerBlock("")
		got, err := json.Marshal(divider)
		if err != nil {
			t.Fatal(err)
		}

		if want != string(got) {
			t.Fatalf("want %s, got %s", want, got)
		}
	})
	t.Run("with block id", func(t *testing.T) {
		want := "{\"type\":\"divider\",\"block_id\":\"divider1\"}"

		divider := NewDividerBlock("divider1")
		got, err := json.Marshal(divider)
		if err != nil {
			t.Fatal(err)
		}

		if want != string(got) {
			t.Fatalf("want %s, got %s", want, got)
		}
	})

}
