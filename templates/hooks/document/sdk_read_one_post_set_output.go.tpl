	content, err := rm.getDocumentContent(
		ctx,
		ko.Spec.Name,
		resp.Document.DocumentVersion,
		resp.Document.DocumentFormat,
	)
	if err != nil {
		return nil, err
	}
	ko.Spec.Content = content
