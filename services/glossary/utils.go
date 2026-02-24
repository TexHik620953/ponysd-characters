package glossary

import (
	"math/rand"
)

func PickRandomGlossary(data []GlossaryRecord, rng *rand.Rand) GlossaryRecord {
	if len(data) == 0 {
		return GlossaryRecord{}
	}
	return data[rng.Int31n(int32(len(data)))]
}
func SetIfInvalid(field *string, glossary []GlossaryRecord, random GlossaryRecord) {
	if len(*field) == 0 || !InGlossary(glossary, *field) {
		*field = random.Value
	}
}
func InGlossary(data []GlossaryRecord, val string) bool {
	for _, v := range data {
		if v.Value == val {
			return true
		}
	}
	return false
}
