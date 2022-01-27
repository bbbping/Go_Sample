package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_export(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)
	ExportToFile("/Users/oldlam/Downloads/Go_demo/stock/alfred_pugin/abc","哈哈哈")
}


func Test_read(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)
	result := ReadFromFile("/Users/oldlam/Downloads/Go_demo/stock/alfred_pugin/abc")
	fmt.Println(result)
}

func Test_readAndWtite(t *testing.T){
	result := ReadFromFile("./abc")
	fmt.Println("长度:", len(result))
}

func Test_arg(t *testing.T){
	fmt.Println(os.Args)
}
