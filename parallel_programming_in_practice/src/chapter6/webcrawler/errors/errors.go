package errors

import (
	"bytes"
	"fmt"
	"strings"
)

type ErrorType string

const (
	ERROR_TYPE_DOWNLOADER ErrorType = "downloader error"
	ERROR_TYPE_ANALYZER   ErrorType = "analyzer error"
	ERROR_TYPE_PIPELINE   ErrorType = "pipeline error"
	ERROR_TYPE_SCHEDULER  ErrorType = "scheduler error"
)

type CrawlerError interface {
	Type() ErrorType
	Error() string
}

type myCrawlerError struct {
	errType    ErrorType
	errMsg     string
	fullErrMsg string
}

func NewCrawlerError(errType ErrorType, errMsg string) CrawlerError {
	return &myCrawlerError{
		errType: errType,
		errMsg:  strings.TrimSpace(errMsg),
	}
}

func NewCrawlerErrorBy(errType ErrorType, err error) CrawlerError {
	return NewCrawlerError(errType, err.Error())
}

func (m *myCrawlerError) Type() ErrorType {
	return m.errType
}

func (m *myCrawlerError) Error() string {
	if m.fullErrMsg == "" {
		m.genFullErrMsg()
	}
	return m.fullErrMsg
}

func (m *myCrawlerError) genFullErrMsg() {
	var buf bytes.Buffer
	buf.WriteString("crawler error:")
	if m.errType != "" {
		buf.WriteString(string(m.errType))
		buf.WriteString(": ")
	}
	buf.WriteString(m.errMsg)
	m.fullErrMsg = fmt.Sprintf("%s", buf.String())
	return
}

type IllegalParameterError struct {
	msg string
}

func NewIllegalParameterError(errMsg string) IllegalParameterError {
	return IllegalParameterError{
		msg: fmt.Sprintf("illegal parameter: %s", strings.TrimSpace(errMsg)),
	}
}

func (ipe IllegalParameterError) Error() string {
	return ipe.msg
}
