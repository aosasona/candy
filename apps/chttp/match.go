package chttp

type MatchName string

const (
	MatchNameClientIP MatchName = "client_ip"
	MatchNameHeader   MatchName = "header"
	MatchNameHeaderRE MatchName = "header_regexp"
	MatchNameHost     MatchName = "host"
	MatchNameMethod   MatchName = "method"
	MatchNameNot      MatchName = "not"
	MatchNamePath     MatchName = "path"
	MatchNamePathRE   MatchName = "path_regexp"
	MatchNameQuery    MatchName = "query"
	MatchNameProtocol MatchName = "protocol"
	MatchNameRemoteIP MatchName = "remote_ip"
	MatchNameVars     MatchName = "vars"
	MatchNameVarsRE   MatchName = "vars_regexp"
)

type Match interface {
	Match() MatchName
}

type PatternMatch struct {
	Name    string `json:"name"`
	Pattern string `json:"pattern"`
}

type RouteMatcher map[MatchName]Match

// Match types definitions
type MatchClientIP struct {
	Ranges []string `json:"ranges,omitempty"`
}

type MatchHeader map[string][]string

type MatchHeaderREField PatternMatch

type MatchHeaderRE map[string]MatchHeaderREField

type MatchHost []string

type MatchMethod []string

type MatchNot []Match

type MatchPath []string

type MatchPathRE PatternMatch

type MatchQuery map[string][]string

type MatchProtocol string

type MatchRemoteIP struct {
	Ranges    []string `json:"ranges,omitempty"`
	Forwarded bool     `json:"forwarded"`
}

type MatchVarsRE map[string]PatternMatch

type MatchVars map[string][]string

// Match methods definitions
func (MatchClientIP) Match() MatchName { return MatchNameClientIP }

func (MatchHeader) Match() MatchName { return MatchNameHeader }

func (MatchHeaderRE) Match() MatchName { return MatchNameHeaderRE }

func (MatchHost) Match() MatchName { return MatchNameHost }

func (MatchMethod) Match() MatchName { return MatchNameMethod }

func (MatchNot) Match() MatchName { return MatchNameNot }

func (MatchPath) Match() MatchName { return MatchNamePath }

func (MatchPathRE) Match() MatchName { return MatchNamePathRE }

func (MatchProtocol) Match() MatchName { return MatchNameProtocol }

func (MatchQuery) Match() MatchName { return MatchNameQuery }

func (MatchRemoteIP) Match() MatchName { return MatchNameRemoteIP }

func (MatchVarsRE) Match() MatchName { return MatchNameVarsRE }

func (MatchVars) Match() MatchName { return MatchNameVars }
