func WritingCompanyCreate(ctx context.Context, params WritingCompany) uuids.UUID {
	query := queries.WritingCompany.Query()
	query.SetValue(query.NaicCode, params.NaicCode)
	query.SetValue(query.ClaimPhone, params.ClaimPhone)
	query.SetValue(query.FullName, params.Name)
	query.SetValue(query.Phone, params.Phone)
	query.SetValue(query.WebPageUrl, params.ClaimUrl)
	query.SetValue(query.SfId, params.SfId)
	query.SetValue(query.EmailAddress, params.Email)
	query.SetValue(query.AddressLines, params.AddressLines)
	query.SetValue(query.City, params.City)
	query.SetValue(query.State, params.State)
	query.SetValue(query.ZipCode, params.ZipCode)
	Id, err := query.Insert(ctx)
	if err != errors.NoErr {
		panic(err)
	}

	return Id
}