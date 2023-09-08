package chttp

type PolicyName string

const (
	PolicyNameCookie       PolicyName = "cookie"
	PolicyNameFirst        PolicyName = "first"
	PolicyNameHeader       PolicyName = "header"
	PolicyNameIPHash       PolicyName = "ip_hash"
	PolicyNameLeastConn    PolicyName = "least_conn"
	PolicyNameRandomChoose PolicyName = "random_choose"
	PolicyNameRandom       PolicyName = "random"
	PolicyNameRoundRobin   PolicyName = "round_robin"
	PolicyNameURIHash      PolicyName = "uri_hash"
)

type SelectionPolicy interface {
	Policy() PolicyName
}

type BaseSelectionPolicy struct {
	PolicyName PolicyName `json:"policy"`
}

func (b *BaseSelectionPolicy) Policy(name PolicyName) PolicyName {
	b.PolicyName = name
	return b.PolicyName
}

func NewPolicy[S SelectionPolicy](policy S) S {
	policy.Policy()
	return policy
}

// policy definitions
type (
	PolicyCookie struct {
		BaseSelectionPolicy
		Name   string `json:"name"`
		Secret string `json:"secret"`
	}

	PolicyFirst struct {
		PolicyCookie
	}

	PolicyHeader struct {
		BaseSelectionPolicy
		Field string `json:"field"`
	}

	PolicyIPHash struct {
		BaseSelectionPolicy
	}

	PolicyLeastConn struct {
		BaseSelectionPolicy
	}

	PolicyRandomChoose struct {
		BaseSelectionPolicy
		Choose int `json:"choose"`
	}

	PolicyRandom struct {
		BaseSelectionPolicy
	}

	PolicyRoundRobin struct {
		BaseSelectionPolicy
	}

	PolicyURIHash struct {
		BaseSelectionPolicy
	}
)

// policy methods
func (p *PolicyCookie) Policy() PolicyName { return p.BaseSelectionPolicy.Policy(PolicyNameCookie) }

func (p *PolicyFirst) Policy() PolicyName { return p.BaseSelectionPolicy.Policy(PolicyNameFirst) }

func (p *PolicyHeader) Policy() PolicyName { return p.BaseSelectionPolicy.Policy(PolicyNameHeader) }

func (p *PolicyIPHash) Policy() PolicyName { return p.BaseSelectionPolicy.Policy(PolicyNameIPHash) }

func (p *PolicyLeastConn) Policy() PolicyName {
	return p.BaseSelectionPolicy.Policy(PolicyNameLeastConn)
}

func (p *PolicyRandomChoose) Policy() PolicyName {
	return p.BaseSelectionPolicy.Policy(PolicyNameRandomChoose)
}

func (p *PolicyRandom) Policy() PolicyName { return p.BaseSelectionPolicy.Policy(PolicyNameRandom) }

func (p *PolicyRoundRobin) Policy() PolicyName {
	return p.BaseSelectionPolicy.Policy(PolicyNameRoundRobin)
}

func (p *PolicyURIHash) Policy() PolicyName { return p.BaseSelectionPolicy.Policy(PolicyNameURIHash) }

// Interface guards
var (
	_ SelectionPolicy = (*PolicyCookie)(nil)
	_ SelectionPolicy = (*PolicyFirst)(nil)
	_ SelectionPolicy = (*PolicyHeader)(nil)
	_ SelectionPolicy = (*PolicyIPHash)(nil)
	_ SelectionPolicy = (*PolicyLeastConn)(nil)
	_ SelectionPolicy = (*PolicyRandomChoose)(nil)
	_ SelectionPolicy = (*PolicyRandom)(nil)
	_ SelectionPolicy = (*PolicyRoundRobin)(nil)
	_ SelectionPolicy = (*PolicyURIHash)(nil)
)
