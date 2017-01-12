package common

import "github.com/olebedev/when/rules"

var All = []rules.Rule{
	SlashDMY(rules.Override),
}
