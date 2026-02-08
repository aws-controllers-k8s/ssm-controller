	// Set WithDecryption to true for SecureString parameters so we can retrieve
	// the plaintext value. This allows proper comparison of the desired vs actual
	// value without getting encrypted values.
	if r.ko.Spec.Type != nil && *r.ko.Spec.Type == "SecureString" {
		input.WithDecryption = aws.Bool(true)
	}
