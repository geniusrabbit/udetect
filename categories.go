package udetect

import "github.com/rtb-0/rtbdict/r0/lightdict/content"

// R0CategoryIDs resolves OpenRTB category codes to r0 ids. cattax 0 is Content Taxonomy 1.0. Unknown codes are skipped.
func R0CategoryIDs(codes []string, cattax int) []uint {
	if len(codes) == 0 {
		return nil
	}
	if cattax == 0 {
		cattax = 1
	}
	out := make([]uint, 0, len(codes))
	for _, code := range codes {
		n := content.ByCattax(cattax, code)
		if n == nil || n.ID() <= 0 {
			continue
		}
		out = append(out, uint(n.ID()))
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func content10Codes(ids []uint) []string {
	if len(ids) == 0 {
		return nil
	}
	out := make([]string, 0, len(ids))
	for _, id := range ids {
		if id == 0 {
			continue
		}
		node := content.ByID(int(id))
		if node == nil {
			continue
		}
		out = append(out, node.IDs(1)...)
	}
	if len(out) == 0 {
		return nil
	}
	return out
}
