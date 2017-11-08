package slug

import (
	"testing"
)

// expectation bool for whether the function will be valid, return a nil error
type stringSliceStruct struct {
	field       string
	expectation []string
}

// Is some value in the slice?
func inSlice(slice []string, val string) bool {
	for _, j := range slice {
		if j == val {
			return true
		}
	}
	return false
}

// Compare two slices and return the difference
func sliceEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for _, j := range slice1 {
		if !inSlice(slice2, j) {
			return false
		}
	}
	return true
}

func TestIsItemTag(t *testing.T) {
	var list = []defaultStruct{
		{"", false},
		{"0123456789", true},
		{"this is a test", true},
		{"Version 2", true},
		{"this-is-a-test", false},
		{"thIS is A TEst", true},
		{"this is a  test", false},
		{"this  is  a test", false},
		{"thIs is  a test", false},
		{" this is a test", false},
		{"this is a test ", false},
		{" this is a test ", false},
		{"this is a test     ", false},
		{"      this is a test", false},
		{"      this is a test       ", false},
		{"   s    this is a test    d    ", false},
		{"this is a test, another", false},
		{"this  is  a  test", false},
		{"THIS IS A TEST", true},
		{"Th1S i3 A te3t", true},
		{"Th1S i3 @ te3t", false},
		{"this_is_a_test", false},
		{"this.is.a.test", false},
		{"this+is+a+test", false},
		{"this/is/a/test", false},
		{"this\\is\\a\\test", false},
		{"~!@#$%^&*()-+={}|:;<>?'`", false},
		{"this-is a test", false},
		{"a", true},
		{"aaaaaaaaaaaaaaaaaa", true},
		{"aaaaaaaaaaaaaaaaaa bbbbbbbbbbbbbbbbbbbbb", true},
		{"aaaaaaaaaaaaaaaaaa bbbbbbbbbbbbbbbbbbbbb ccccccccccccccccccccccccccccccccccc", true},
		{"a b c d e f g h i j k l m n o p q r s t u v w x y z 0 1 2 3 4 5 6 7 8 9", true},
		{"\u0070 here", true},   //UTF-8(ASCII): p
		{"\u0026 there", false}, //UTF-8(ASCII): &
		{"forneça here", false},
		{"n'est pas", false},
		{"gültige test", false},
		{"𐅪", false},
		{"𐅪4", false},
		{"zulässig", false},
		{"1.02", false},
		{" 1234", false},
		{"a=b", false},
		{"How are you?", false},
		{"Test.", false},
		{"中 文 网", false},
		{"소 주", false},
		{"test ¾", false},
	}

	for _, l := range list {
		valid := false
		if IsItemTag(l.field) {
			valid = true
		}

		if l.expectation != valid {
			t.Errorf("IsItemTag(%v): Valid[%t]. Expected: %t", l.field, valid, l.expectation)
		}
	}
}

func TestIsUTF8ItemTag(t *testing.T) {
	var list = []defaultStruct{
		{"", false},
		{"0123456789", true},
		{"this is a test", true},
		{"Version 2", true},
		{"this-is-a-test", false},
		{"thIS is A TEst", true},
		{"this is a  test", false},
		{"this  is  a test", false},
		{"thIs is  a test", false},
		{" this is a test", false},
		{"this is a test ", false},
		{" this is a test ", false},
		{"this is a test     ", false},
		{"      this is a test", false},
		{"      this is a test       ", false},
		{"   s    this is a test    d    ", false},
		{"this is a test, another", false},
		{"this  is  a  test", false},
		{"THIS IS A TEST", true},
		{"Th1S i3 A te3t", true},
		{"Th1S i3 @ te3t", false},
		{"this_is_a_test", false},
		{"this.is.a.test", false},
		{"this+is+a+test", false},
		{"this/is/a/test", false},
		{"this\\is\\a\\test", false},
		{"~!@#$%^&*()-+={}|:;<>?'`", false},
		{"this-is a test", false},
		{"a", true},
		{"aaaaaaaaaaaaaaaaaa", true},
		{"aaaaaaaaaaaaaaaaaa bbbbbbbbbbbbbbbbbbbbb", true},
		{"aaaaaaaaaaaaaaaaaa bbbbbbbbbbbbbbbbbbbbb ccccccccccccccccccccccccccccccccccc", true},
		{"a b c d e f g h i j k l m n o p q r s t u v w x y z 0 1 2 3 4 5 6 7 8 9", true},
		{"\u0070 here", true},   //UTF-8(ASCII): p
		{"\u0026 there", false}, //UTF-8(ASCII): &
		{"forneça here", true},
		{"n'est pas", false},
		{"gültige test", true},
		{"𐅪", false},
		{"𐅪4", false},
		{"zulässig", true},
		{"1.02", false},
		{" 1234", false},
		{"a=b", false},
		{"How are you?", false},
		{"Test.", false},
		{"ẞ 123 test", true},
		{"é here", true},
		{"中 文 网", true},
		{"소 주", true},
		{"test ¾", false},

		{"あ い う え お", true},
		{"あ い う え お か", true},
		{"あ い う え", true},
		{"ひらがな カタカナ 漢字", true},
		{"３ー０　ａ＠ｃｏｍ", false},
		{"Ｆｶﾀｶﾅﾞﾬ", true},
		{"￩left", false},
		{"\x19test\x7F", false},
	}

	for _, l := range list {
		valid := false
		if IsUTF8ItemTag(l.field) {
			valid = true
		}

		if l.expectation != valid {
			t.Errorf("IsUTF8ItemTag(%v): Valid[%t]. Expected: %t", l.field, valid, l.expectation)
		}
	}
}

func TestGetTagSlugListWithBlanks(t *testing.T) {
	var list = []stringSliceStruct{
		{"", []string{"-"}},

		{"one,two 234,three", []string{"one", "two-234", "three"}},
		{"oNE,TwO 234,tHReE", []string{"one", "two-234", "three"}},
		{",one,two,three", []string{"-", "one", "two", "three"}},
		{"one,two,three,", []string{"one", "two", "three", "-"}},
		{",one,two,three,", []string{"-", "one", "two", "three", "-"}},
		{",,,,,,,,one,,,,,,,two,three", []string{"-", "-", "-", "-", "-", "-", "-", "-", "one", "-", "-", "-", "-", "-", "-", "two", "three"}},
		{"one,two,,,,,,,,,,three,,,,,,,", []string{"one", "two", "-", "-", "-", "-", "-", "-", "-", "-", "-", "three", "-", "-", "-", "-", "-", "-", "-"}},
		{",,,,,,,,one,,,,,,,,,two,three,,,,,,,,,four,,,,,,,", []string{"-", "-", "-", "-", "-", "-", "-", "-", "one", "-", "-", "-", "-", "-", "-", "-", "-", "two", "three", "-", "-", "-", "-", "-", "-", "-", "-", "four", "-", "-", "-", "-", "-", "-", "-"}},
		{"one,two,three, ", []string{"one", "two", "three", "-"}},
		{" ,one,two,three", []string{"-", "one", "two", "three"}},
		{" ,one,two,three, ", []string{"-", "one", "two", "three", "-"}},
		{"        ,one,two,three, ", []string{"-", "one", "two", "three", "-"}},
		{" one,two,three, ", []string{"one", "two", "three", "-"}},
		{"one,two,three ", []string{"one", "two", "three"}},
		{"        one,two,three, ", []string{"one", "two", "three", "-"}},
		{"one,two,three         ", []string{"one", "two", "three"}},
		{" ,one,two,three,         ", []string{"-", "one", "two", "three", "-"}},
		{",   ,   ,  one,   two,    three,   ,   , ", []string{"-", "-", "-", "one", "two", "three", "-", "-", "-"}},
		{",   ,   ,  one   123   456,   two  7   8   9,    three,   ,   , ", []string{"-", "-", "-", "one-123-456", "two-7-8-9", "three", "-", "-", "-"}},

		{"THIS IS A TEST", []string{"this-is-a-test"}},
		{"ThiS iS A teSt", []string{"this-is-a-test"}},
		{"Th1S i3 A te3t", []string{"th1s-i3-a-te3t"}},
		{"Th1S i3 @ te3t", []string{"th1s-i3-te3t"}},
		{"this_is_a_test", []string{"this-is-a-test"}},
		{"this.is.a.test", []string{"this-is-a-test"}},
		{"this+is+a+test", []string{"this-is-a-test"}},
		{"this-is-a-test", []string{"this-is-a-test"}},
		{"this is a test-", []string{"this-is-a-test"}},
		{"-this is a test", []string{"this-is-a-test"}},
		{"--this is a test---", []string{"this-is-a-test"}},
		{"this/is/a/test", []string{"this-is-a-test"}},
		{"this\\is\\a\\test", []string{"this-is-a-test"}},
		{"this=is=a=test", []string{"this-is-a-test"}},

		{"a", []string{"a"}},
		{"0123456789", []string{"0123456789"}},
		{"aaaaaaaaaaaaaaaaaa", []string{"aaaaaaaaaaaaaaaaaa"}},
		{"aaaaaaaaaaaaaaaaaa,bbbbbbbbbbbbbbbbbbbbb", []string{"aaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbb"}},
		{"aaaaaaaaaaaaaaaaaa,bbbbbbbbbbbbbbbbbbbbb,ccccccccccccccccccccccccccccccccccc", []string{"aaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbb", "ccccccccccccccccccccccccccccccccccc"}},
		{"aaaaaaaaaa aaaaaaaaaa aaaaaaaaaa aaaaaaaaaa aaaaaaaaaa", []string{"aaaaaaaaaa-aaaaaaaaaa-aaaaaaaaaa-aaaaaaaaaa-aaaaaaaaaa"}},
		{"a  ?;'\"\\  b c d||[]e f    g h i j   %^& k l~!@m n o p,,,,,,q r s t u+++\\/-+v w x<>??{y z 0 1 2 3 4 5 6 7 8 9", []string{"a-b-c-d-e-f-g-h-i-j-k-lm-n-o-p", "-", "-", "-", "-", "-", "q-r-s-t-u-v-w-xy-z-0-1-2-3-4-5-6-7-8-9"}},

		{"~!@#$%^&*()-+={}|:;<>?'`", []string{"-"}},
		{"\u0070 here", []string{"p-here"}}, //UTF-8(ASCII): p
		{"\u0026 there", []string{"there"}}, //UTF-8(ASCII): &

		{"	àåáâäãå\tèéêëę,ìíîïı.òóôõöøőð/ùúûüŭů\\çćčĉ-żźž_śşšŝ", []string{"aaaaaaa-eeeee", "iiiii-oooooood-uuuuuu-cccc-zzz-ssss"}},
		{"	===ñń,,,++++ýÿ  ğĝ,,, ř ł  đ ß Þ ĥ ĵ ", []string{"nn", "-", "-", "yy-gg", "-", "-", "r-l-d-ss-th-h-j"}},
		{" , Hêlło======1234\\, WhaT+++++do.......you///////want  , thîs______is@				a--------TEšT,    ,               ,", []string{"-", "hello-1234", "what-do-you-want", "this-is-a-test", "-", "-", "-"}},
		{"b!(!)#()%^^!,Hel{}];;lo-12소-주34,Wh<><,.l;aT-do-y\u0026ou-want,+++++++___a", []string{"b", "hel-lo-12so-ju34", "wh", "lat-do-you-want", "a"}},

		{"forneça here,gültige zulässig test,ẞ 123 test", []string{"forneca-here", "gultige-zulassig-test", "ss-123-test"}},
		{"é here,é here,é here", []string{"e-here", "e-here", "e-here"}},
		{"中 文 网,中 文 网", []string{"zhong-wen-wang", "zhong-wen-wang"}},
		{"소 주,주,소", []string{"so-ju", "ju", "so"}},
		{"𐅪", []string{"-"}},
		{"𐅪4", []string{"4"}},
		{"n'est pas", []string{"nest-pas"}},
		{"1.02", []string{"1-02"}},
		{"-1234", []string{"1234"}},
		{"a=b", []string{"a-b"}},
		{"How are you?", []string{"how-are-you"}},
		{"Test.", []string{"test"}},
		{"test ¾,test 34", []string{"test-34", "test-34"}},

		{"あ い, う, え, お", []string{"a-i", "u", "e", "o"}},
		{"あ い う, え お, か", []string{"a-i-u", "e-o", "ka"}},
		{"あ い う え", []string{"a-i-u-e"}},
		{"ひら,がな, カタ,    カナ,    漢字", []string{"hira", "gana", "kata", "kana", "han-zi"}},
		{"ひら,~!@#$%^&*(がな, カタ,    カ++++++ナ,    漢:\"{_+字", []string{"hira", "gana", "kata", "ka-na", "han-zi"}},
		{"３ー０　ａ＠ｃｏｍ", []string{"30-acom"}},
		{"Ｆｶﾀｶﾅﾞﾬ", []string{"fkatakanalb"}},
		{"￩left", []string{"left"}},
		{"   ,    , ,,,,   ,,  \x19test\x7F,,,,,,\x19test\x7F", []string{"-", "-", "-", "-", "-", "-", "-", "-", "test", "-", "-", "-", "-", "-", "test"}},
	}

	for _, l := range list {
		result := getTagSlugListWithBlanks(l.field)

		if !sliceEqual(l.expectation, result) {
			t.Errorf("getTagSlugListWithBlanks(%v): Result:%s. Expected: %s", l.field, result, l.expectation)
		}
	}
}

func TestGetPlainTagListWithBlanks(t *testing.T) {
	var list = []stringSliceStruct{
		{"", []string{"-"}},
		{"one,two 234,three", []string{"one", "two 234", "three"}},
		{"oNE,TwO 234,tHReE", []string{"oNE", "TwO 234", "tHReE"}},
		{",one?,two!,three$", []string{"-", "one?", "two!", "three$"}},
		{"one,two,three,", []string{"one", "two", "three", "-"}},
		{",one,two,three,", []string{"-", "one", "two", "three", "-"}},
		{",,,,,,,,O@nE,,,,,,,TWo!#,tHreE$", []string{"-", "-", "-", "-", "-", "-", "-", "-", "O@nE", "-", "-", "-", "-", "-", "-", "TWo!#", "tHreE$"}},
	}

	for _, l := range list {
		result := getPlainTagListWithBlanks(l.field)

		if !sliceEqual(l.expectation, result) {
			t.Errorf("getPlainTagListWithBlanks(%v): Result:%s. Expected: %s", l.field, result, l.expectation)
		}
	}
}

func TestGetTagsAndTagSlugs(t *testing.T) {
	var list = []struct {
		field    string
		tagList  []string
		slugList []string
	}{
		{"", []string{}, []string{}},
		{
			",,,,,,,,O@nE,,,,,,,TWo!#,tHreE$ ",
			[]string{"O@nE", "TWo!#", "tHreE$ "},
			[]string{"one", "two", "three"},
		},
		{
			",,,,,,,,     O@nE    ,,,,,     O@nE    ,,,,  T  W  o  !#,,     O@nE    ,tHr    eE$ ",
			[]string{"     O@nE    ", "  T  W  o  !#", "tHr    eE$ "},
			[]string{"one", "t-w-o", "thr-ee"},
		},
		{
			"@#$%^&*(,abc@#$%^&*(def, One in the middle too., abc@#$%^&*(def, abc def, ABc dEf, Another post here!,What about this?",
			[]string{"abc@#$%^&*(def", " One in the middle too.", " Another post here!", "What about this?"},
			[]string{"abc-def", "one-in-the-middle-too", "another-post-here", "what-about-this"},
		},
	}

	for _, l := range list {
		tagList, slugList := GetTagsAndTagSlugs(l.field)

		if !sliceEqual(tagList, l.tagList) {
			t.Errorf("GetTagsAndTagSlugs(%v): Result:%s. Expected: %s", l.field, tagList, l.tagList)
		}
		if !sliceEqual(slugList, l.slugList) {
			t.Errorf("GetTagsAndTagSlugs(%v): Result:%s. Expected: %s", l.field, slugList, l.slugList)
		}
	}
}

func TestIsItemTagListRegex(t *testing.T) {
	var list = []defaultStruct{
		{"", false},

		{"test123|test123|123hello", true},
		{"test-123|test-123|123-hello", true},
		{"-test", false}, {"test-", false}, {"-test-", false},
		{" test", false}, {"test ", false}, {" test ", false},
		{"~test", false}, {"test~", false}, {"~test~", false},
		{"!test", false}, {"test!", false}, {"!test!", false},
		{"@test", false}, {"test@", false}, {"@test@", false},
		{"#test", false}, {"test#", false}, {"#test#", false},

		{"|this is a test|here is another 123|test this|how about that|", false},
		{"|||||||d||||||||this is a test|here is another 123|test this|how about that", false},
		{"this is a test|here is another 123|test this|how about that|||||e|||||||", false},
		{"|||||d|||||||this is a test|here is another 123|test this|how about that||||||s|||||||", false},
		{"this is a test|here is another 123|test this|how about that| ", false},
		{" |this is a test|here is another 123|test this|how about that", false},
		{" |this is a test|here is another 123|test this|how about that| ", false},
		{"       |this is a test|here is another 123|test this|how about that| ", false},
		{" this is a test|here is another 123|test this|how about that", false},
		{"this is a test|here is another 123|test this|how about that ", false},
		{"       this is a test|here is another 123|test this|how about that", false},
		{"this is a test|here is another 123|test this|how about that          ", false},
		{" |this is a test|here is another 123|test this|how about that|       ", false},
	}

	for _, l := range list {
		valid := IsItemTagListRegex(l.field)

		if l.expectation != valid {
			t.Errorf("IsItemTagListRegex(%v): Valid[%t]. Expected: %t", l.field, valid, l.expectation)
		}
	}
}

/*
uses commas for separator character instead of pipe (|)

func TestIsItemTagListRegex(t *testing.T) {
	var list = []defaultStruct{
		{"", false},

		{"test123,test123,123hello", true},
		{"test-123,test-123,123-hello", true},
		{"-test", false}, {"test-", false}, {"-test-", false},
		{" test", false}, {"test ", false}, {" test ", false},
		{"~test", false}, {"test~", false}, {"~test~", false},
		{"!test", false}, {"test!", false}, {"!test!", false},
		{"@test", false}, {"test@", false}, {"@test@", false},
		{"#test", false}, {"test#", false}, {"#test#", false},

		{",this is a test,here is another 123,test this,how about that,", false},
		{",,,,,,,d,,,,,,,,this is a test,here is another 123,test this,how about that", false},
		{"this is a test,here is another 123,test this,how about that,,,,,e,,,,,,,", false},
		{",,,,,d,,,,,,,this is a test,here is another 123,test this,how about that,,,,,,s,,,,,,,", false},
		{"this is a test,here is another 123,test this,how about that, ", false},
		{" ,this is a test,here is another 123,test this,how about that", false},
		{" ,this is a test,here is another 123,test this,how about that, ", false},
		{"       ,this is a test,here is another 123,test this,how about that, ", false},
		{" this is a test,here is another 123,test this,how about that", false},
		{"this is a test,here is another 123,test this,how about that ", false},
		{"       this is a test,here is another 123,test this,how about that", false},
		{"this is a test,here is another 123,test this,how about that          ", false},
		{" ,this is a test,here is another 123,test this,how about that,       ", false},

	}

	for _, l := range list {
		valid := IsItemTagListRegex(l.field)

		if l.expectation != valid {
			t.Errorf("IsItemTagListRegex(%v): Valid[%t]. Expected: %t", l.field, valid, l.expectation)
		}
	}
}
*/

func TestIsItemTagList(t *testing.T) {
	var list = []defaultStruct{
		{"", false},
		{"this is a test", true},
		{"this is a test,here is another 123,test this,how about that", true},
		{"thIS is A TEst,heRE is aNOther 123,teST this,hOW about THat", true},
		{"this is a test,here is  another 123,test this,how about that", false},
		{",this is a test,here is another 123,test this,how about that", false},
		{"this is a test,here is another 123,test this,how about that,", false},
		{",this is a test,here is another 123,test this,how about that,", false},
		{",,,,,,,d,,,,,,,,this is a test,here is another 123,test this,how about that", false},
		{"this is a test,here is another 123,test this,how about that,,,,,e,,,,,,,", false},
		{",,,,,d,,,,,,,this is a test,here is another 123,test this,how about that,,,,,,s,,,,,,,", false},
		{"this is a test,here is another 123,test this,how about that, ", false},
		{" ,this is a test,here is another 123,test this,how about that", false},
		{" ,this is a test,here is another 123,test this,how about that, ", false},
		{"       ,this is a test,here is another 123,test this,how about that, ", false},
		{" this is a test,here is another 123,test this,how about that", false},
		{"this is a test,here is another 123,test this,how about that ", false},
		{"       this is a test,here is another 123,test this,how about that", false},
		{"this is a test,here is another 123,test this,how about that          ", false},
		{" ,this is a test,here is another 123,test this,how about that,       ", false},
		{"THIS IS A TEST", true},
		{"ThiS iS A teSt", true},
		{"Th1S i3 A te3t", true},
		{"Th1S i3 @ te3t", false},
		{"this_is_a_test", false},
		{"this.is.a.test", false},
		{"this+is+a+test", false},
		{"this-is-a-test", false},
		{"this is a test-", false},
		{"-this is a test", false},
		{"--this is a test---", false},
		{"this/is/a/test", false},
		{"this\\is\\a\\test", false},
		{"this=is=a=test", false},
		{"a", true},
		{"0123456789", true},
		{"aaaaaaaaaaaaaaaaaa", true},
		{"aaaaaaaaaaaaaaaaaa bbbbbbbbbbbbbbbbbbbbb", true},
		{"aaaaaaaaaaaaaaaaaa bbbbbbbbbbbbbbbbbbbbb ccccccccccccccccccccccccccccccccccc", true},
		{"a b c d e f g h i j k l m n o p q r s t u v w x y z 0 1 2 3 4 5 6 7 8 9", true},
		{"this is @ test", false},
		{"this is.a test", false},
		{" this is a test", false},
		{"this is a test ", false},
		{" this is a test ", false},
		{" this is a test            ", false},
		{"          this is a test", false},
		{"          this is a test             ", false},
		{"~!@#$%^&*()-+={}|:;<>?'`", false},
		{"\u0070 here", true},   //UTF-8(ASCII): p
		{"\u0026 there", false}, //UTF-8(ASCII): &
	}

	for _, l := range list {
		valid := IsItemTagList(l.field)

		if l.expectation != valid {
			t.Errorf("IsItemTagList(%v): Valid[%t]. Expected: %t", l.field, valid, l.expectation)
		}
	}
}

func TestIsUTF8ItemTagList(t *testing.T) {
	var list = []defaultStruct{
		{"", false},
		{"this is a test", true},
		{"this is a test,here is another 123,test this,how about that", true},
		{"thIS is A TEst,heRE is aNOther 123,teST this,hOW about THat", true},
		{"this is a test,here is  another 123,test this,how about that", false},
		{",this is a test,here is another 123,test this,how about that", false},
		{"this is a test,here is another 123,test this,how about that,", false},
		{",this is a test,here is another 123,test this,how about that,", false},
		{",,,,,,,d,,,,,,,,this is a test,here is another 123,test this,how about that", false},
		{"this is a test,here is another 123,test this,how about that,,,,,e,,,,,,,", false},
		{",,,,,d,,,,,,,this is a test,here is another 123,test this,how about that,,,,,,s,,,,,,,", false},
		{"this is a test,here is another 123,test this,how about that, ", false},
		{" ,this is a test,here is another 123,test this,how about that", false},
		{" ,this is a test,here is another 123,test this,how about that, ", false},
		{"       ,this is a test,here is another 123,test this,how about that, ", false},
		{" this is a test,here is another 123,test this,how about that", false},
		{"this is a test,here is another 123,test this,how about that ", false},
		{"       this is a test,here is another 123,test this,how about that", false},
		{"this is a test,here is another 123,test this,how about that          ", false},
		{" ,this is a test,here is another 123,test this,how about that,       ", false},
		{"THIS IS A TEST", true},
		{"ThiS iS A teSt", true},
		{"Th1S i3 A te3t", true},
		{"Th1S i3 @ te3t", false},
		{"this_is_a_test", false},
		{"this.is.a.test", false},
		{"this+is+a+test", false},
		{"this-is-a-test", false},
		{"this is a test-", false},
		{"-this is a test", false},
		{"--this is a test---", false},
		{"this/is/a/test", false},
		{"this\\is\\a\\test", false},
		{"this=is=a=test", false},
		{"a", true},
		{"0123456789", true},
		{"aaaaaaaaaaaaaaaaaa", true},
		{"aaaaaaaaaaaaaaaaaa bbbbbbbbbbbbbbbbbbbbb", true},
		{"aaaaaaaaaaaaaaaaaa bbbbbbbbbbbbbbbbbbbbb ccccccccccccccccccccccccccccccccccc", true},
		{"a b c d e f g h i j k l m n o p q r s t u v w x y z 0 1 2 3 4 5 6 7 8 9", true},
		{"this is @ test", false},
		{"this is.a test", false},
		{" this is a test", false},
		{"this is a test ", false},
		{" this is a test ", false},
		{" this is a test            ", false},
		{"          this is a test", false},
		{"          this is a test             ", false},
		{"~!@#$%^&*()-+={}|:;<>?'`", false},
		{"\u0070 here", true},   //UTF-8(ASCII): p
		{"\u0026 there", false}, //UTF-8(ASCII): &

		{"forneça here,gültige test,ẞ 123 test", true},
		{"forneça here", true},
		{"gültige test", true},
		{"zulässig", true},
		{"ẞ 123 test", true},
		{"é here,é here,é here", true},
		{"中 文 网,中 文 网", true},
		{"소 주,주,소", true},
		{"𐅪", false},
		{"𐅪4", false},
		{"n'est pas", false},
		{"1.02", false},
		{"-1234", false},
		{"a=b", false},
		{"How are you?", false},
		{"Test.", false},
		{"test ¾,test 34", false},

		{"あ い う え お", true},
		{"あ い う え お か", true},
		{"あ い う え", true},
		{"ひらがな カタカナ 漢字", true},
		{"３ー０　ａ＠ｃｏｍ", false},
		{"Ｆｶﾀｶﾅﾞﾬ", true},
		{"￩left", false},
		{"\x19test\x7F", false},
	}

	for _, l := range list {
		valid := IsUTF8ItemTagList(l.field)

		if l.expectation != valid {
			t.Errorf("IsUTF8ItemTagList(%v): Valid[%t]. Expected: %t", l.field, valid, l.expectation)
		}
	}
}
