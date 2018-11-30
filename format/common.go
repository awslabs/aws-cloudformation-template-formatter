package format

func ordering(path []string) []string {
	if len(path) == 0 {
		return []string{
			"AWSTemplateFormatVersion",
			"Description",
			"Metadata",
			"Parameters",
			"Mappings",
			"Conditions",
			"Transform",
			"Resources",
			"Outputs",
		}
	} else if path[0] == "Parameters" && len(path) == 2 {
		return []string{
			"Type",
			"Default",
		}
	} else if path[0] == "Transform" || path[len(path)-1] == "Fn::Transform" {
		return []string{
			"Name",
			"Parameters",
		}
	} else if path[0] == "Resources" && len(path) == 2 {
		return []string{
			"Type",
		}
	} else if path[0] == "Outputs" && len(path) == 2 {
		return []string{
			"Description",
			"Value",
			"Export",
		}
	}

	return []string{}
}
