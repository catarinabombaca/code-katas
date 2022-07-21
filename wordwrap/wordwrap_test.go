package wordwrap_test

import (
	"github.com/catarinabombaca/code-katas/wordwrap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordWrap(t *testing.T) {
	var (
		assert = assert.New(t)
		tests  = []test{
			{"The prize was delivered to Tom with as much effusion as the superintendent could pump up under the circumstances; but it lacked somewhat of the true gush, for the poor fellow's instinct taught him that there was a mystery here that could not well bear the light, perhaps; it was simply preposterous that this boy had warehoused two thousand sheaves of Scriptural wisdom on his premises—a dozen would strain his capacity, without a doubt.\n\nAmy Lawrence was proud and glad, and she tried to make Tom see it in her face—but he wouldn't look. She wondered; then she was just a grain troubled; next a dim suspicion came and went—came again; she watched; a furtive glance told her worlds—and then her heart broke, and she was jealous, and angry, and the tears came and she hated everybody. Tom most of all (she thought).",
				58,
				"The prize was delivered to Tom with as much effusion as\nthe superintendent could pump up under the circumstances;\nbut it lacked somewhat of the true gush, for the poor\nfellow's instinct taught him that there was a mystery here\nthat could not well bear the light, perhaps; it was simply\npreposterous that this boy had warehoused two thousand\nsheaves of Scriptural wisdom on his premises—a dozen\nwould strain his capacity, without a doubt.\n\nAmy Lawrence was proud and glad, and she tried to make Tom\nsee it in her face—but he wouldn't look. She wondered;\nthen she was just a grain troubled; next a dim suspicion\ncame and went—came again; she watched; a furtive glance\ntold her worlds—and then her heart broke, and she was\njealous, and angry, and the tears came and she hated\neverybody. Tom most of all (she thought)."},
		}
	)

	for _, test := range tests {
		actualOutput := wordwrap.Execute(test.textInput, test.maxChars)
		assert.Equal(test.expectedOutput, actualOutput)
	}
}

type test struct {
	textInput      string
	maxChars       int
	expectedOutput string
}
