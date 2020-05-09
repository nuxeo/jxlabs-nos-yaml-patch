module github.com/nxmatic/jxlabs-nos-helmfile-patch

go 1.14

require (
	github.com/mattbaird/jsonpatch v0.0.0-20171005235357-81af80346b1a
	github.com/spf13/cobra v1.0.0
)

replace (
	github.com/jenkins-x/jx => github.com/jenkins-x/jx v1.3.981-0.20200508204445-1b2ed275a50e
)
