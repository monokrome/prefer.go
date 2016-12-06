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

func checkTestError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func getMockSubject() *MockSubject {
	return &MockSubject{
		Name:  MOCK_NAME,
		Value: MOCK_VALUE,
	}
}

func getMockSubjectSerialize(t *testing.T, serializer Serializer) []byte {
	subject := getMockSubject()

	serialized, err := serializer.Serialize(subject)
	checkTestError(t, err)

	return serialized
}

func TestNewSerializerReturnsJSONSerializer(t *testing.T) {
	content := getMockSubjectSerialize(t, JSONSerializer{})
	serializer, _ := NewSerializer("example.json", content)

	if _, ok := serializer.(JSONSerializer); ok != true {
		t.Error("Expected NewSerializer to return a JSONSerializer")
	}
}

func JSONSerializerTestCase(t *testing.T) {
	serializer := JSONSerializer{}
	serialized := getMockSubjectSerialize(t, serializer)

	result := &MockSubject{}
	checkTestError(t, serializer.Deserialize(serialized, result))

	if result != getMockSubject() {
		t.Error("Result does not match original serialized object.")
	}
}

func XMLSerializerTestCase(t *testing.T) {
	serializer := XMLSerializer{}
	serialized := getMockSubjectSerialize(t, serializer)

	result := &MockSubject{}
	checkTestError(t, serializer.Deserialize(serialized, result))

	if result != getMockSubject() {
		t.Error("Result does not match original serialized object.")
	}
}
