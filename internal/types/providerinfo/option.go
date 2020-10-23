package providerinfo

// OptionUpdater updates veProviderInfo properties.
type OptionUpdater func(*veProviderInfo)

// WithFullName updates FullName field of veProviderInfo
func WithFullName(n string) OptionUpdater {
	return func(ti *veProviderInfo) {
		ti.FullName = n
	}
}

// WithGivenName updates GivenName field of veProviderInfo
func WithGivenName(n string) OptionUpdater {
	return func(ti *veProviderInfo) {
		ti.GivenName = n
	}
}

// WithFamilyName updates FamilyName field of veProviderInfo
func WithFamilyName(n string) OptionUpdater {
	return func(ti *veProviderInfo) {
		ti.FamilyName = n
	}
}

// WithPicture updates Picture field of veProviderInfo
func WithPicture(p string) OptionUpdater {
	return func(ti *veProviderInfo) {
		ti.Picture = p
	}
}

// WithEmail updates Email field of an veProviderInfo
func WithEmail(e string) OptionUpdater {
	return func(ti *veProviderInfo) {
		ti.Email = e
	}
}
