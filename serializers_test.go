package prefer

import "testing"

const (
	MOCK_NAME  = "Mock Name"
	MOCK_VALUE = 12
)

type MockSubject struct {
	Name  string
	Value int
}

func checkTest(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func TestJSONSerializer(t *testing.T) {
	serializer := JSONSerializer{}

	subject := MockSubject{
		Name:  MOCK_NAME,
		Value: MOCK_VALUE,
	}

	serialized, err := serializer.Serialize(subject)
	checkTest(t, err)

	result := MockSubject{}
	checkTest(t, serializer.Deserialize(serialized, &result))

	if result != subject {
		t.Error("Result does not match original serialized object.")
	}
}

func TestXMLSerializer(t *testing.T) {
	serializer := XMLSerializer{}

	subject := MockSubject{
		Name:  MOCK_NAME,
		Value: MOCK_VALUE,
	}

	serialized, err := serializer.Serialize(subject)
	checkTest(t, err)

	result := MockSubject{}
	checkTest(t, serializer.Deserialize(serialized, &result))

	if result != subject {
		t.Error("Result does not match original serialized object.")
	}
}
