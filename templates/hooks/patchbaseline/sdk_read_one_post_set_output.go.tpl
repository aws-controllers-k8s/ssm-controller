    if ko.Status.BaselineID != nil {
        tags, err := rm.fetchCurrentTags(ctx, ko.Status.BaselineID)
        if err != nil {
            return nil, err
        }
        ko.Spec.Tags = fromACKTags(tags, nil)
    }