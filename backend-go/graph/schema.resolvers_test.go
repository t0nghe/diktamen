package graph

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
)

func TestTodo(t *testing.T) {
	c := client.New(handler.NewDefaultServer(NewExecutableSchema(New())))

	var resp struct {
		CreateTodo struct{ ID string }
	}

	c.MustPost(`mutation { createTodo(todo:{text:"Fery importante"}){id}`, &resp) // response?

	require.Equal(t, "5", resp.CreateTodo.ID)

	t.Run("update the todo text", func(t *testing.T) {
		var resp struct {
			UpdateTodo struct{ Text string }
		}
		c.MustPost(
			`mutation($id: ID!, $text: String!) { updateTodo(id: $id, changes:{text:$text}) { text } }`,
			&resp,
			client.Var("id", 5),
			client.Var("text", "Very important"),
		)

		require.Equal(t, "Very important", resp.UpdateTodo.Text)
	})
}
