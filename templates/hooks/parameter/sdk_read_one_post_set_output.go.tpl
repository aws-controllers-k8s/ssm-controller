	// DataType is optional on creation and AWS defaults it to "text".
	// This causes a diff `"diff":[{"Path":{"Parts":["Spec","DataType"]},"A":null,"B":"text"}]}`
	if r.ko.Spec.DataType == nil && ko.Spec.DataType != nil {
	    r.ko.Spec.DataType = ko.Spec.DataType
	}

	// For SecureString parameters, we should not compare the Value field
	// because AWS returns the encrypted value, not the plaintext.
	if r.ko.Spec.Type != nil && *r.ko.Spec.Type == "SecureString" {
		ko.Spec.Value = r.ko.Spec.Value
	}
