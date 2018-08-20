package example

import "github.com/just1689/fun-with-chan/state"

var Topic = state.NewTopic(state.TopicConfig{Name: "Le queue", TimeoutSeconds: 1})

