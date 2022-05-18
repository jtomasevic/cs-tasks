func CarrierCreate(ctx context.Context, params WritingCompany) (uuids.UUID, errors.CodeErr) {
	query := queries.WritingCompany.Query()
	query.SetValue(query.NaicCode, params.NaicCode)
	query.SetValue(query.ClaimPhone, params.ClaimPhone)
	query.SetValue(query.FullName, params.Name)
	query.SetValue(query.Phone, params.Phone)
	query.SetValue(query.WebPageUrl, params.WebPageContact)
	query.SetValue(query.ClaimPhoneExt, params.ClaimPhoneExt)
	query.SetValue(query.PhoneExt, params.PhoneExt)
	query.SetValue(query.CarrierEmailAddress, params.ClaimContactEmail)
	query.SetValue(query.LegacyType, true)
	query.SetValue(query.AddressLines, params.AddressLines)
	query.SetValue(query.City, params.City)
	query.SetValue(query.State, params.State)
	query.SetValue(query.ZipCode, params.ZipCode)
	Id, err := query.Insert(ctx)
	if err != errors.NoErr {
		panic(err)
	}

	return Id, errors.NoErr	
}