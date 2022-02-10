package trace

import (
	"encoding/hex"
	"regexp"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/hashcode"
)

func DefSampler(dktrace DatakitTrace) (DatakitTrace, bool) {
	for i := range dktrace {
		if IsRootSpan(dktrace[i]) {
			switch dktrace[i].Priority {
			case PriorityAuto:
				hex.EncodeToString([]byte(dktrace[i].TraceID))
			case PriorityReject:
				return nil, true
			case PriorityKeep:
				return dktrace, true
			default:
				log.Debug("unrecognized trace proority")
			}
		}
	}

	return dktrace, false
}

func CloseResourceWrapper(ignoredResources map[string]*regexp.Regexp) FilterFunc {
	if len(ignoredResources) == 0 {
		return func(dktrace DatakitTrace) (DatakitTrace, bool) {
			return dktrace, false
		}
	} else {
		return func(dktrace DatakitTrace) (DatakitTrace, bool) {
			for i := range dktrace {
				if IsRootSpan(dktrace[i]) {
					for k := range ignoredResources {
						if ignoredResources[k].MatchString(dktrace[i].Resource) {
							return nil, true
						}
					}
				}
			}

			return dktrace, false
		}
	}
}

func KeepRareResourceWrapper(presentMap map[string]time.Time) FilterFunc {
	if presentMap == nil {
		presentMap = make(map[string]time.Time)
	}

	return func(dktrace DatakitTrace) (DatakitTrace, bool) {
		var skip bool
		for i := range dktrace {
			if IsRootSpan(dktrace[i]) {
				var (
					checksum = hashcode.GenMapHash(map[string]string{
						"service":  dktrace[i].Service,
						"resource": dktrace[i].Resource,
						"env":      dktrace[i].Env,
					})
					lastCheck time.Time
					ok        bool
				)
				if lastCheck, ok = presentMap[checksum]; !ok || time.Since(lastCheck) >= time.Hour {
					skip = true
				}
				presentMap[checksum] = time.Now()
				break
			}
		}

		return dktrace, skip
	}
}
