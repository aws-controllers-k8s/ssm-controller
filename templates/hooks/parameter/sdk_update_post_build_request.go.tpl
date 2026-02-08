	// For SSM Parameter, Overwrite must be true for update operations
	// PutParameter is used for both create and update, and requires Overwrite=true to modify existing parameters.
	// This follows what the rds controller implements but using hooks to avoid the Overwrite field from being added to the api Spec
	// https://github.com/aws-controllers-k8s/rds-controller/blob/2d5427a8b4a2f6caf0bc889754370a6ab55dc135/generator.yaml#L103
	input.Overwrite = aws.Bool(true)
