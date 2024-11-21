    if ko.Status.BaselineID != nil {
        tags, err := rm.fetchCurrentTags(ko.Status.BaselineID)
        if err != nil {
            return nil, err
        }
        ko.Spec.Tags = FromACKTags(tags)
    }