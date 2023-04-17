package extra25519

// IsLowOrder checks if the passed group element is of low order.
// Algorithm translated from the same source as the blacklist (see above).
func IsEdLowOrder(ge []byte) bool {
	var (
		c    [7]byte
		k    int
		i, j int
	)

	// cases j = 0..30
	for j = 0; j < 31; j++ {
		for i = 0; i < len(edBlacklist); i++ {
			c[i] |= ge[j] ^ edBlacklist[i][j]
		}
	}

	// case j = 31, ignore highest bit
	for i = 0; i < len(edBlacklist); i++ {
		c[i] |= (ge[j] & 0x7f) ^ edBlacklist[i][j]
	}

	k = 0
	for i = 0; i < len(edBlacklist); i++ {
		k |= int(c[i]) - 1
	}

	return ((k >> 8) & 1) == 1
}

// edBlacklist is a list of elements of the ed25519 curve that have low order.
// The list was copied from https://github.com/jedisct1/libsodium/blob/141288535127c22162944e12fcadb8bc269671cc/src/libsodium/crypto_core/ed25519/ref10/ed25519_ref10.c
var edBlacklist = [7][32]byte{
	/* 0 (order 4) */
	{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	/* 1 (order 1) */
	{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	/* 2707385501144840649318225287225658788936804267575313519463743609750303402022
	   (order 8) */
	{0x26, 0xe8, 0x95, 0x8f, 0xc2, 0xb2, 0x27, 0xb0, 0x45, 0xc3, 0xf4,
		0x89, 0xf2, 0xef, 0x98, 0xf0, 0xd5, 0xdf, 0xac, 0x05, 0xd3, 0xc6,
		0x33, 0x39, 0xb1, 0x38, 0x02, 0x88, 0x6d, 0x53, 0xfc, 0x05},
	/* 55188659117513257062467267217118295137698188065244968500265048394206261417927
	   (order 8) */
	{0xc7, 0x17, 0x6a, 0x70, 0x3d, 0x4d, 0xd8, 0x4f, 0xba, 0x3c, 0x0b,
		0x76, 0x0d, 0x10, 0x67, 0x0f, 0x2a, 0x20, 0x53, 0xfa, 0x2c, 0x39,
		0xcc, 0xc6, 0x4e, 0xc7, 0xfd, 0x77, 0x92, 0xac, 0x03, 0x7a},
	/* p-1 (order 2) */
	{0xec, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
	/* p (=0, order 4) */
	{0xed, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
	/* p+1 (=1, order 1) */
	{0xee, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
}
