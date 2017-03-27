package local

import (
	"bytes"
	"strings"
	"testing"

	"github.com/docker/infrakit/pkg/types"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestZero(t *testing.T) {
	require.True(t, zero("string", ""))
	require.True(t, zero("int", 0))
	require.True(t, zero("bool", false))
	require.True(t, zero("float", 0.))
}

func TestContext(t *testing.T) {

	// A template file containing flags and prompts will be parsed and used to configure
	// the cobra command

	script := `
{{/* The directive here tells infrakit to run this script with sh:  =% sh %=  */}}

{{/* The function 'flag' will create a flag in the CLI; the function 'prompt' will ask user for input */}}

{{ $doCommit := flag "commit" "bool" "true to commit" false }}
{{ $clusterName := flag "cluster-name" "string" "the name of the cluster" "swarm" }}
{{ $clusterSize := flag "size" "int" "the size of the cluster" 20 }}
{{ $floatValue := flag "param" "float" "some float param" 25.5 }}

{{ $user := prompt "Please enter your user name" "string" }}

{{/* An example here where we expose a flag and if not set, ask the user */}}
{{ $instanceType := flag "instance-type" "string" "VM instance type" | prompt "Please specify vm instance type:" "string"}}

{{/* generate a json so we can parse back and check result */}}
{
  "clusterName" : "{{$clusterName}}",
  "clusterSize" : {{$clusterSize}},
  "username" : "{{$user}}",
  "doCommit" : {{$doCommit}},
  "instanceType" : "{{$instanceType}}",
  "param" : {{$floatValue}}
}
`

	c := &Context{
		cmd: &cobra.Command{
			Use:   "test",
			Short: "test",
		},
		src:   "str://" + script,
		input: bytes.NewBufferString("username\n"),
	}

	c.exec = false
	err := c.buildFlags()
	require.NoError(t, err)

	for _, n := range []string{"commit", "cluster-name", "size", "instance-type", "param"} {
		require.NotNil(t, c.cmd.Flag(n))
	}

	err = c.cmd.Flags().Parse(strings.Split("--param 75.0 --cluster-name swarm1 --commit --size 20 --instance-type large", " "))
	require.NoError(t, err)

	err = c.loadBackend()
	require.NoError(t, err)
	require.NotNil(t, c.run)

	err = c.execute()
	require.NoError(t, err)

	m := map[string]interface{}{}
	err = types.AnyString(c.script).Decode(&m)
	require.NoError(t, err)

	// compare by the encoded json value
	require.Equal(t, types.AnyValueMust(map[string]interface{}{
		"clusterName":  "swarm1",
		"clusterSize":  20,
		"param":        75.0,
		"username":     "username",
		"doCommit":     true,
		"instanceType": "large",
	}).String(), types.AnyValueMust(m).String())
}
