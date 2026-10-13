	if ko.Status.AssociationID != nil {
		tags, err := rm.fetchCurrentTags(ctx, ko.Status.AssociationID)
		if err != nil {
			return nil, err
		}
		ko.Spec.Tags = fromACKTags(tags, nil)
	}
