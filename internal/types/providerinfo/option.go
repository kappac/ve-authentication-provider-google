package providerinfo

// OptionUpdater updates veProviderInfo properties.
type OptionUpdater func(*veProviderInfo)

// WithFullName updates FullName field of veProviderInfo
func WithFullName(n string) OptionUpdater {
	return func(ti *veProviderInfo) {
		ti.PFullName = n
	}
}

// WithGivenName updates GivenName field of veProviderInfo
func WithGivenName(n string) OptionUpdater {
	return func(ti *veProviderInfo) {
		ti.PGivenName = n
	}
}

// WithFamilyName updates FamilyName field of veProviderInfo
func WithFamilyName(n string) OptionUpdater {
	return func(ti *veProviderInfo) {
		ti.PFamilyName = n
	}
}

// WithPicture updates Picture field of veProviderInfo
func WithPicture(p string) OptionUpdater {
	return func(ti *veProviderInfo) {
		ti.PPicture = p
	}
}

// WithEmail updates Email field of an veProviderInfo
func WithEmail(e string) OptionUpdater {
	return func(ti *veProviderInfo) {
		ti.PEmail = e
	}
}
