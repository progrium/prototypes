package foundation

func NSAllHashTableObjects(a uintptr) uintptr                                                 {}
func NSAllMapTableKeys(a uintptr) uintptr                                                     {}
func NSAllMapTableValues(a uintptr) uintptr                                                   {}
func NSAllocateCollectable(a uint64, b uint64) uintptr                                        {}
func NSAllocateMemoryPages(a uint64) uintptr                                                  {}
func NSAllocateObject(a uintptr, b uint64, c uintptr) uintptr                                 {}
func NSClassFromString(a uintptr) uintptr                                                     {}
func NSCompareHashTables(a uintptr, b uintptr) bool                                           {}
func NSCompareMapTables(a uintptr, b uintptr) bool                                            {}
func NSContainsRect(a CGRect, b CGRect) bool                                                  {}
func NSCopyHashTableWithZone(a uintptr, b uintptr) uintptr                                    {}
func NSCopyMapTableWithZone(a uintptr, b uintptr) uintptr                                     {}
func NSCopyMemoryPages(a uintptr, b uintptr, c uint64)                                        {}
func NSCopyObject(a uintptr, b uint64, c uintptr) uintptr                                     {}
func NSCountFrames() uint64                                                                   {}
func NSCountHashTable(a uintptr) uint64                                                       {}
func NSCountMapTable(a uintptr) uint64                                                        {}
func NSCreateHashTable(a NSHashTableCallBacks, b uint64) uintptr                              {}
func NSCreateHashTableWithZone(a NSHashTableCallBacks, b uint64, c uintptr) uintptr           {}
func NSCreateMapTable(a NSMapTableKeyCallBacks, b NSMapTableValueCallBacks, c uint64) uintptr {}
func NSCreateMapTableWithZone(a NSMapTableKeyCallBacks, b NSMapTableValueCallBacks, c uint64, d uintptr) uintptr {
}
func NSCreateZone(a uint64, b uint64, c bool) uintptr                             {}
func NSDeallocateMemoryPages(a uintptr, b uint64)                                 {}
func NSDeallocateObject(a uintptr)                                                {}
func NSDecimalAdd(a uintptr, b uintptr, c uintptr, d uint64) uint64               {}
func NSDecimalCompact(a uintptr)                                                  {}
func NSDecimalCompare(a uintptr, b uintptr) int64                                 {}
func NSDecimalCopy(a uintptr, b uintptr)                                          {}
func NSDecimalDivide(a uintptr, b uintptr, c uintptr, d uint64) uint64            {}
func NSDecimalMultiply(a uintptr, b uintptr, c uintptr, d uint64) uint64          {}
func NSDecimalMultiplyByPowerOf10(a uintptr, b uintptr, c int16, d uint64) uint64 {}
func NSDecimalNormalize(a uintptr, b uintptr, c uint64) uint64                    {}
func NSDecimalPower(a uintptr, b uintptr, c uint64, d uint64) uint64              {}
func NSDecimalRound(a uintptr, b uintptr, c int64, d uint64)                      {}
func NSDecimalString(a uintptr, b uintptr) uintptr                                {}
func NSDecimalSubtract(a uintptr, b uintptr, c uintptr, d uint64) uint64          {}
func NSDecrementExtraRefCountWasZero(a uintptr) bool                              {}
func NSDefaultMallocZone() uintptr                                                {}
func NSDivideRect(a CGRect, b uintptr, c uintptr, d float64, e uint64)            {}
func NSEdgeInsetsEqual(a NSEdgeInsets, b NSEdgeInsets) bool                       {}
func NSEndHashTableEnumeration(a uintptr)                                         {}
func NSEndMapTableEnumeration(a uintptr)                                          {}
func NSEnumerateHashTable(a uintptr) NSHashEnumerator                             {}
func NSEnumerateMapTable(a uintptr) NSMapEnumerator                               {}
func NSEqualPoints(a CGPoint, b CGPoint) bool                                     {}
func NSEqualRects(a CGRect, b CGRect) bool                                        {}
func NSEqualSizes(a CGSize, b CGSize) bool                                        {}
func NSExtraRefCount(a uintptr) uint64                                            {}
func NSFileTypeForHFSTypeCode(a uint32) uintptr                                   {}
func NSFrameAddress(a uint64) uintptr                                             {}
func NSFreeHashTable(a uintptr)                                                   {}
func NSFreeMapTable(a uintptr)                                                    {}
func NSFullUserName() uintptr                                                     {}
func NSGetSizeAndAlignment(a uintptr, b uintptr, c uintptr) uintptr               {}
func NSGetUncaughtExceptionHandler() uintptr                                      {}
func NSHFSTypeCodeFromFileType(a uintptr) uint32                                  {}
func NSHFSTypeOfFile(a uintptr) uintptr                                           {}
func NSHashGet(a uintptr, b uintptr) uintptr                                      {}
func NSHashInsert(a uintptr, b uintptr)                                           {}
func NSHashInsertIfAbsent(a uintptr, b uintptr) uintptr                           {}
func NSHashInsertKnownAbsent(a uintptr, b uintptr)                                {}
func NSHashRemove(a uintptr, b uintptr)                                           {}
func NSHomeDirectory() uintptr                                                    {}
func NSHomeDirectoryForUser(a uintptr) uintptr                                    {}
func NSIncrementExtraRefCount(a uintptr)                                          {}
func NSInsetRect(a CGRect, b float64, c float64) CGRect                           {}
func NSIntegralRect(a CGRect) CGRect                                              {}
func NSIntegralRectWithOptions(a CGRect, b uint64) CGRect                         {}
func NSIntersectionRange(a NSRange, b NSRange) NSRange                            {}
func NSIntersectionRect(a CGRect, b CGRect) CGRect                                {}
func NSIntersectsRect(a CGRect, b CGRect) bool                                    {}
func NSIsEmptyRect(a CGRect) bool                                                 {}
func NSIsFreedObject(a uintptr) bool                                              {}
func NSLog(a uintptr)                                                             {}
func NSLogPageSize() uint64                                                       {}
func NSLogv(a uintptr, b uintptr)                                                 {}
func NSMapGet(a uintptr, b uintptr) uintptr                                       {}
func NSMapInsert(a uintptr, b uintptr, c uintptr)                                 {}
func NSMapInsertIfAbsent(a uintptr, b uintptr, c uintptr) uintptr                 {}
func NSMapInsertKnownAbsent(a uintptr, b uintptr, c uintptr)                      {}
func NSMapMember(a uintptr, b uintptr, c uintptr, d uintptr) bool                 {}
func NSMapRemove(a uintptr, b uintptr)                                            {}
func NSMouseInRect(a CGPoint, b CGRect, c bool) bool                              {}
func NSNextHashEnumeratorItem(a uintptr) uintptr                                  {}
func NSNextMapEnumeratorPair(a uintptr, b uintptr, c uintptr) bool                {}
func NSOffsetRect(a CGRect, b float64, c float64) CGRect                          {}
func NSOpenStepRootDirectory() uintptr                                            {}
func NSPageSize() uint64                                                          {}
func NSPointFromString(a uintptr) CGPoint                                         {}
func NSPointInRect(a CGPoint, b CGRect) bool                                      {}
func NSProtocolFromString(a uintptr) uintptr                                      {}
func NSRangeFromString(a uintptr) NSRange                                         {}
func NSRealMemoryAvailable() uint64                                               {}
func NSReallocateCollectable(a uintptr, b uint64, c uint64) uintptr               {}
func NSRecordAllocationEvent(a int32, b uintptr)                                  {}
func NSRectFromString(a uintptr) CGRect                                           {}
func NSRecycleZone(a uintptr)                                                     {}
func NSResetHashTable(a uintptr)                                                  {}
func NSResetMapTable(a uintptr)                                                   {}
func NSReturnAddress(a uint64) uintptr                                            {}
func NSRoundDownToMultipleOfPageSize(a uint64) uint64                             {}
func NSRoundUpToMultipleOfPageSize(a uint64) uint64                               {}
func NSSearchPathForDirectoriesInDomains(a uint64, b uint64, c bool) uintptr      {}
func NSSelectorFromString(a uintptr) uintptr                                      {}
func NSSetUncaughtExceptionHandler(a uintptr)                                     {}
func NSSetZoneName(a uintptr, b uintptr)                                          {}
func NSShouldRetainWithZone(a uintptr, b uintptr) bool                            {}
func NSSizeFromString(a uintptr) CGSize                                           {}
func NSStringFromClass(a uintptr) uintptr                                         {}
func NSStringFromHashTable(a uintptr) uintptr                                     {}
func NSStringFromMapTable(a uintptr) uintptr                                      {}
func NSStringFromPoint(a CGPoint) uintptr                                         {}
func NSStringFromProtocol(a uintptr) uintptr                                      {}
func NSStringFromRange(a NSRange) uintptr                                         {}
func NSStringFromRect(a CGRect) uintptr                                           {}
func NSStringFromSelector(a uintptr) uintptr                                      {}
func NSStringFromSize(a CGSize) uintptr                                           {}
func NSTemporaryDirectory() uintptr                                               {}
func NSUnionRange(a NSRange, b NSRange) NSRange                                   {}
func NSUnionRect(a CGRect, b CGRect) CGRect                                       {}
func NSUserName() uintptr                                                         {}
func NSZoneCalloc(a uintptr, b uint64, c uint64) uintptr                          {}
func NSZoneFree(a uintptr, b uintptr)                                             {}
func NSZoneFromPointer(a uintptr) uintptr                                         {}
func NSZoneMalloc(a uintptr, b uint64) uintptr                                    {}
func NSZoneName(a uintptr) uintptr                                                {}
func NSZoneRealloc(a uintptr, b uintptr, c uint64) uintptr                        {}
func NXReadNSObjectFromCoder(a uintptr) uintptr                                   {}
