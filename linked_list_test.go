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



func TestInsertAfter(t *testing.T) {
    afterT1 := "after t1"
    afterT2 := "after t2"
    afterT3 := "after t3"

	list := LinkedList{}
	list.Add("t1")
	list.Add("t2")
	list.Add("t3")


    list.InsertAfter("t1", afterT1)
    wantT1 := afterT1
    gotT1result, err := list.Find("t1")
    if err != nil {
			t.Errorf("error returned: %q", err)
        return
    }

    gotT1 := gotT1result.next.value

    if gotT1 != wantT1 {
        t.Errorf("got %q, wanted %q", gotT1, wantT1)
    }

    list.InsertAfter("t2", afterT2)
    wantT2 := afterT2
    gotT2result, err := list.Find("t2")
    if err != nil {
			t.Errorf("error returned: %q", err)
        return
    }

    gotT2 := gotT2result.next.value

    if gotT2 != wantT2 {
        t.Errorf("got %q, wanted %q", gotT2, wantT2)
    }


    list.InsertAfter("t3", afterT3)
    wantT3 := afterT3
    gotT3result, err := list.Find("t3")
    if err != nil {
			t.Errorf("error returned: %q", err)
        return
    }

    gotT3 := gotT3result.next.value

    if gotT3 != wantT3 {
        t.Errorf("got %q, wanted %q", gotT3, wantT3)
    }

}
