package slug

import (
	"net/url"
	"testing"
	fv "github.com/dholtzmann/formvalidator"
)

// expectation bool for whether the function will be valid, return a nil error
type defaultStruct struct {
	field       string
	expectation bool
}

type stringStruct struct {
	field       string
	expectation string
}

func TestGetSlug(t *testing.T) {
	var list = []stringStruct{
		{"", ""},
		{"Hello world!", "hello-world"},
		{"Hello\t\tworld!", "hello-world"},
		{"Hello\nworld!", "hello-world"},
		{"Hello      world!", "hello-world"},
		{"     Hello world!", "hello-world"},
		{"Hello world         !", "hello-world"},
		{"     Hello      world     ", "hello-world"},
		{"HeLlO WOrlD!", "hello-world"},
		{"H3ll8 W8R11d!", "h3ll8-w8r11d"},
		{"~!@#$%^&*()_+{}|:'\"<>?/\\|[]", ""},
		{"a~!@#$%^&*()_+{}b|:'\"<>?/\\|[]c", "a-b-c"},
		{"E=mc2", "e-mc2"},
		{"1+1=2", "1-1-2"},
		{"One,two,three", "one-two-three"},
		{"This.is.a.test", "this-is-a-test"},
		{"Here//is//another", "here-is-another"},
		{"One\\\\more\\\\time", "one-more-time"},
		{"----This---is---a---test----", "this-is-a-test"},
		{"____This___is___a___test____", "this-is-a-test"},
		{"This+is+a+test+", "this-is-a-test"},
		{"=====This====is=====a=====test=======", "this-is-a-test"},
		{"This+++is++++a+++++test++++", "this-is-a-test"},
		{"\u0070 here", "p-here"}, //UTF-8(ASCII): p
		{"\u0026 there", "there"}, //UTF-8(ASCII): &
		{"𐅪", ""},
		{"𐅪4", "4"},
		{"Forneça here", "forneca-here"},
		{"n'est pas", "nest-pas"},
		{"Gültige Test", "gultige-test"},
		{"Zulässig", "zulassig"},
		{"gültige@Heiẞe.de", "gultigeheisse-de"},
		{"中-文-网", "zhong-wen-wang"},
		{"소-주", "so-ju"},
		{"test-¾", "test-3-4"},
		{"あいうえお", "aiueo"},
		{"あいうえおか", "aiueoka"},
		{"あいうえ", "aiue"},
		{"ひらがな・カタカナ、．漢字", "hiraganakatakana-han-zi"},
		{"３ー０　ａ＠ｃｏｍ", "30-acom"},
		{"Ｆｶﾀｶﾅﾞﾬ", "fkatakanalb"},
		{"￩left", "left"},
		{"\x19test\x7F", "test"},
		{"	àåáâäãåą	èéêëę,ìíîïı.òóôõöøőð/ùúûüŭů\\çćčĉ-żźž_śşšŝ=ñń++++++ýÿ     ğĝ ř ł     đ ß Þ ĥ ĵ   ", "aaaaaaaa-eeeee-iiiii-oooooood-uuuuuu-cccc-zzz-ssss-nn-yy-gg-r-l-d-ss-th-h-j"},
	}

	for _, l := range list {
		result := ""
		result = GetAsciiSlug(l.field)

		if l.expectation != result {
			t.Errorf("GetAsciiSlug(%v): Result[%s]. Expected: %s", l.field, result, l.expectation)
		}
	}
}

func TestGetSlugAndIsSlug(t *testing.T) {
	var list = []defaultStruct{
		{"", false},
		{"Hello world!", true},
		{"Hello\t\tworld!", true},
		{"Hello\nworld!", true},
		{"Hello      world!", true},
		{"     Hello world!", true},
		{"Hello world         !", true},
		{"     Hello      world     ", true},
		{"HeLlO WOrlD!", true},
		{"H3ll8 W8R11d!", true},
		{"~!@#$%^&*()_+{}|:'\"<>?/\\|[]", false},
		{"a~!@#$%^&*()_+{}b|:'\"<>?/\\|[]c", true},
		{"E=mc2", true},
		{"1+1=2", true},
		{"One,two,three", true},
		{"This.is.a.test", true},
		{"Here//is//another", true},
		{"One\\\\more\\\\time", true},
		{"----This---is---a---test----", true},
		{"____This___is___a___test____", true},
		{"This+is+a+test+", true},
		{"=====This====is=====a=====test=======", true},
		{"This+++is++++a+++++test++++", true},
		{"\u0070 here", true},  //UTF-8(ASCII): p
		{"\u0026 there", true}, //UTF-8(ASCII): &
		{"𐅪", false},
		{"𐅪4", true},
		{"Forneça here", true},
		{"n'est pas", true},
		{"Gültige Test", true},
		{"Zulässig", true},
		{"gültige@Heiẞe.de", true},
		{"中-文-网", true},
		{"소-주", true},
		{"test-¾", true},
		{"あいうえお", true},
		{"あいうえおか", true},
		{"あいうえ", true},
		{"ひらがな・カタカナ、．漢字", true},
		{"３ー０　ａ＠ｃｏｍ", true},
		{"Ｆｶﾀｶﾅﾞﾬ", true},
		{"￩left", true},
		{"\x19test\x7F", true},
		{"	àåáâäãåą	èéêëę,ìíîïı.òóôõöøőð/ùúûüŭů\\çćčĉ-żźž_śşšŝ=ñń++++++ýÿ     ğĝ ř ł     đ ß Þ ĥ ĵ   ", true},
	}

	for _, l := range list {
		result := GetAsciiSlug(l.field)
		expectation := IsSlug(result)

		if expectation != l.expectation {
			t.Errorf("GetAsciiSlug(%v) -> IsSlug(%s): Result[%t]", l.field, result, expectation)
		}
	}
}

func TestIsSlug(t *testing.T) {
	var list = []defaultStruct{
		{"", false},
		{"this-is-a-test", true},
		{"THIS-IS-A-TEST", true},
		{"ThiS-iS-A-teSt", true},
		{"Th1S-i3-A-te3t", true},
		{"Th1S-i3-@-te3t", false},
		{"this_is_a_test", false},
		{"this.is.a.test", false},
		{"this+is+a+test", false},
		{"this is a test", false},
		{"this/is/a/test", false},
		{"this\\is\\a\\test", false},
		{"~!@#$%^&*()-+={}|:;<>?'`", false},
		{"this-is-a test", false},
		{"a", true},
		{"0123456789", true},
		{"aaaaaaaaaaaaaaaaaa", true},
		{"aaaaaaaaaaaaaaaaaa-bbbbbbbbbbbbbbbbbbbbb", true},
		{"aaaaaaaaaaaaaaaaaa-bbbbbbbbbbbbbbbbbbbbb-ccccccccccccccccccccccccccccccccccc", true},
		{"a-b-c-d-e-f-g-h-i-j-k-l-m-n-o-p-q-r-s-t-u-v-w-x-y-z-0-1-2-3-4-5-6-7-8-9", true},
		{"this-is-@-test", false},
		{"this-is.a-test", false},
		{"-this-is-a-test", false},
		{"this-is-a-test-", false},
		{"-this-is-a-test-", false},
		{"-this-is-a-test------------", false},
		{"----------this-is-a-test", false},
		{"----------this-is-a-test-------------", false},
		{"\u0070-here", true},   //UTF-8(ASCII): p
		{"\u0026-there", false}, //UTF-8(ASCII): &
		{"𐅪", false},
		{"𐅪4", false},
		{"forneça-here", false},
		{"n'est pas", false},
		{"gültige-test", false},
		{"zulässig", false},
		{"1.02", false},
		{"-1234", false},
		{"a=b", false},
		{"How-are-you?", false},
		{"Test.", false},
		{"中-文-网", false},
		{"소-주", false},
		{"test-¾", false},
	}

	for _, l := range list {
		valid := false
		if IsSlug(l.field) {
			valid = true
		}

		if l.expectation != valid {
			t.Errorf("IsSlug(%v): Valid[%t]. Expected: %t", l.field, valid, l.expectation)
		}
	}
}

func TestIsUTF8Slug(t *testing.T) {
	var list = []defaultStruct{
		{"", false},
		{"this-is-a-test", true},
		{"THIS-IS-A-TEST", true},
		{"ThiS-iS-A-teSt", true},
		{"Th1S-i3-A-te3t", true},
		{"Th1S-i3-@-te3t", false},
		{"this_is_a_test", false},
		{"this.is.a.test", false},
		{"this+is+a+test", false},
		{"this is a test", false},
		{"this/is/a/test", false},
		{"this\\is\\a\\test", false},
		{"~!@#$%^&*()-+={}|:;<>?'`", false},
		{"this-is-a test", false},
		{"a", true},
		{"0123456789", true},
		{"aaaaaaaaaaaaaaaaaa", true},
		{"aaaaaaaaaaaaaaaaaa-bbbbbbbbbbbbbbbbbbbbb", true},
		{"aaaaaaaaaaaaaaaaaa-bbbbbbbbbbbbbbbbbbbbb-ccccccccccccccccccccccccccccccccccc", true},
		{"a-b-c-d-e-f-g-h-i-j-k-l-m-n-o-p-q-r-s-t-u-v-w-x-y-z-0-1-2-3-4-5-6-7-8-9", true},
		{"this-is-@-test", false},
		{"this-is.a-test", false},
		{"-this-is-a-test", false},
		{"this-is-a-test-", false},
		{"-this-is-a-test-", false},
		{"-this-is-a-test------------", false},
		{"----------this-is-a-test", false},
		{"----------this-is-a-test-------------", false},
		{"\u0070-here", true},   //UTF-8(ASCII): p
		{"\u0026-there", false}, //UTF-8(ASCII): &
		{"forneça-here", true},
		{"gültige-test", true},
		{"zulässig", true},
		{"ẞ-123-test", true},
		{"é-here", true},
		{"中-文-网", true},
		{"소-주", true},
		{"𐅪", false},
		{"𐅪4", false},
		{"n'est pas", false},
		{"1.02", false},
		{"-1234", false},
		{"a=b", false},
		{"How-are-you?", false},
		{"Test.", false},
		{"test-¾", false},
	}

	for _, l := range list {
		valid := false
		if IsUTF8Slug(l.field) {
			valid = true
		}

		if l.expectation != valid {
			t.Errorf("IsUTF8Slug(%v): Valid[%t]. Expected: %t", l.field, valid, l.expectation)
		}
	}
}

func Test_IsSlugField(t *testing.T) {
	var list = []struct {
		field       string
		expectation bool
	}{
		{"true", true},
		{"10", true},
		{"-5.20301", true},
		{"a", true},
		{"[]string{}", true},
		{`[]string{"qwerty"}`, true},
		{"", false},
		{"%", false},
		{"~!@#$%^&*()_+", false},
		{"~@$^&t&*())", true},
		{"<>>.P<>:", true},
		{"{\":><\"}[", false},
		{"0123456789", true},
		{"E = mc^2", true},
		{"How do you do?", true},
	}

	for _, l := range list {
		valid := false
		var rule fv.Rule = IsSlugField()

		if e, _ := rule.Validate([]string{l.field}, make(map[string]string)); e == nil {
			valid = true
		}

		if l.expectation != valid {
			t.Errorf("isSlug(%s): Valid[%t]. Expected: %t", l.field, valid, l.expectation)
		}
	}
}

func Test_formValidate(t *testing.T) {
	form := url.Values{}
	form.Set("Slug", "Test123%")
	form.Set("Slug2", "<>)(*^&")

	// this should validate without a problem
	rules := map[string][]fv.Rule{
		"Slug": fv.RuleChain(IsSlugField()),
		//		"Slug2":	fv.RuleChain(IsSlugField()),
	}

	err, validator := fv.New(rules)
	if err != nil {
		t.Errorf("Error making a new form validator type! %s", err.Error())
	}

	_, errors := validator.Validate(form)

	if len(errors) > 0 {
		for field, val := range errors {
			for _, e := range val {
				t.Errorf("Test_formValidate(): %s: %s", field, e.Error())
			}
		}
	}
}

func Benchmark_GetSlug(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetAsciiSlug("ひらがな カタカナ 漢字")
	}
}

func Benchmark_IsSlug(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsSlug(",,,,,d,,,,,,,this is a test,here is another 123,test this,how about that,,,,,,s,,,,,,,")
	}
}

func Benchmark_IsUTF8Slug(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsUTF8Slug("ひらがな カタカナ 漢字")
	}
}
