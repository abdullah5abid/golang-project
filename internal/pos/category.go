package pos

type Category struct {
	Identifier rune
	Fields     []Field
}

func (c Category) IsEmpty() bool {
	return c.Identifier == 0
}

var (
	Adjective    = Category{'A', []Field{Type, Degree, Gender, Nmber, PossessorPerson, PossessorNumber}}
	AdPosition   = Category{'S', []Field{Type}}
	Adverb       = Category{'R', []Field{Type}}
	Conjunction  = Category{'C', []Field{Type}}
	Date         = Category{'W', []Field{}}
	Determiner   = Category{'D', []Field{Type, Person, Gender, Nmber, PossessorNumber}}
	Interjection = Category{'I', []Field{}}
	Noun         = Category{'N', []Field{Type, Gender, Nmber, NounClass, NounSubClass, Degree}}
	Number       = Category{'Z', []Field{Type}}
	Pronoun      = Category{'P', []Field{Type, Person, Gender, Nmber, Case, Polite}}
	Punctuation  = Category{'F', []Field{Type, PunctuationEnclose}}
	Verb         = Category{'V', []Field{Type, Mood, Tense, Person, Nmber, Gender}}
)

func GetCategory(c rune) Category {
	switch c {
	case 'A':
		return Adjective
	case 'C':
		return Conjunction
	case 'D':
		return Determiner
	case 'N':
		return Noun
	case 'P':
		return Pronoun
	case 'R':
		return Adverb
	case 'S':
		return AdPosition
	case 'V':
		return Verb
	case 'Z':
		return Number
	case 'W':
		return Date
	case 'I':
		return Interjection
	case 'F':
		return Punctuation
	}

	return Category{}
}
