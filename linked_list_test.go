package main

import "testing"

func TestAdd(t *testing.T) {
	want := []string{"t1", "t2", "t3"}
	list := LinkedList{}
	list.Add(want[0])
	list.Add(want[1])
	list.Add(want[2])

	got := []string{}

	got = append(got, list.head.value)
	got = append(got, list.head.next.value)
	got = append(got, list.head.next.next.value)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got[i], want[i])
		}
	}

}

func TestFind(t *testing.T) {
	want := []string{"t1", "t2", "t3"}
	list := LinkedList{}
	list.Add(want[0])
	list.Add(want[1])
	list.Add(want[2])

	got := []string{}

	for _, x := range want {
		val, err := list.Find(x)
		if err != nil {
			t.Errorf("error returned: %q", err)
			return
		}
		got = append(got, val.value)
	}
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got[i], want[i])
		}
	}
}

func TestFindByIndex(t *testing.T) {
	want := []string{"t1", "t2", "t3"}
	list := LinkedList{}
	list.Add(want[0])
	list.Add(want[1])
	list.Add(want[2])

	got := []string{}

	for x := range want {
		val, err := list.FindByIndex(uint(x))
		if err != nil {
			t.Errorf("error returned: %q", err)
			return
		}
		got = append(got, val.value)
	}
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got[i], want[i])
		}
	}
}
