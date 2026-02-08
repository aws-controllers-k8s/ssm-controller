	// DataType is optional on creation and AWS defaults it to "text".
	// This causes a diff `"diff":[{"Path":{"Parts":["Spec","DataType"]},"A":null,"B":"text"}]}`
	if r.ko.Spec.DataType == nil && ko.Spec.DataType != nil {
	    r.ko.Spec.DataType = ko.Spec.DataType
	}
