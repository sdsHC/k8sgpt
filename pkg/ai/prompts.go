package ai

const (
	//default_prompt = `Simplify the following Kubernetes error message delimited by triple dashes written in --- %s --- language; --- %s ---.
	//Provide the most possible solution in a step by step style in no more than 280 characters. Write the output in the following format:
	//Error: {Explain error here}
	//Solution: {Step by step solution here}
	//`
	default_prompt = `Simplify the following Kubernetes error message delimited by triple dashes written in --- %s --- language; --- %s ---.
        Provide the most possible solution in a step by step style in no more than 2000 characters. Write the output in the following format:
        Error: {Explain error here}
        Solution: {Step by step solution, kubectl command here}
        Object: {Related Kubernetes object}
        `

	trivy_vuln_prompt = "Explain the following trivy scan result and the detail risk or root cause of the CVE ID, then provide a solution. Response in %s: %s"
	trivy_conf_prompt = "Explain the following trivy scan result and the detail risk or root cause of the security check, then provide a solution."
)

var PromptMap = map[string]string{
	"default":             default_prompt,
	"VulnerabilityReport": trivy_vuln_prompt, // for Trivy integration, the key should match `Result.Kind` in pkg/common/types.go
	"ConfigAuditReport":   trivy_conf_prompt,
}
