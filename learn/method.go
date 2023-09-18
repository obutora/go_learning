package learn

import "log"

type Example struct {
	Value int
}

func (v Example) Increment1(n Example) int {
	return v.Value + n.Value
}

func (v *Example) Increment2(n Example) {
	v.Value += n.Value
}

func Method() {
	v1 := Example{100}
	v2 := Example{200}

	res := v1.Increment1(v2)
	log.Println(res)

	v1.Increment2(v2)
	log.Println(v1.Value)

	v1.Increment2(v2)
	log.Println(v1.Value)

}
