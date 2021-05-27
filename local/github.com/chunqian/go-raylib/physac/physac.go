// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 28 Apr 2021 17:24:23 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package physac

/*
#include "../lib/raylib/src/physac.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// InitPhysics function as declared in src/physac.h:206
func InitPhysics() {
	C.InitPhysics()
}

// RunPhysicsStep function as declared in src/physac.h:207
func RunPhysicsStep() {
	C.RunPhysicsStep()
}

// SetPhysicsTimeStep function as declared in src/physac.h:208
func SetPhysicsTimeStep(delta float64) {
	cdelta, _ := (C.double)(delta), cgoAllocsUnknown
	C.SetPhysicsTimeStep(cdelta)
}

// IsPhysicsEnabled function as declared in src/physac.h:209
func IsPhysicsEnabled() bool {
	__ret := C.IsPhysicsEnabled()
	__v := (bool)(__ret)
	return __v
}

// SetPhysicsGravity function as declared in src/physac.h:210
func SetPhysicsGravity(x float32, y float32) {
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	C.SetPhysicsGravity(cx, cy)
}

// CreatePhysicsBodyCircle function as declared in src/physac.h:211
func CreatePhysicsBodyCircle(pos Vector2, radius float32, density float32) *PhysicsBodyData {
	cpos, _ := *(*C.Vector2)(unsafe.Pointer(&pos)), cgoAllocsUnknown
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	cdensity, _ := (C.float)(density), cgoAllocsUnknown
	__ret := C.CreatePhysicsBodyCircle(cpos, cradius, cdensity)
	__v := newPhysicsBodyDataRef(unsafe.Pointer(__ret)).convert()
	return __v
}

// CreatePhysicsBodyRectangle function as declared in src/physac.h:212
func CreatePhysicsBodyRectangle(pos Vector2, width float32, height float32, density float32) *PhysicsBodyData {
	cpos, _ := *(*C.Vector2)(unsafe.Pointer(&pos)), cgoAllocsUnknown
	cwidth, _ := (C.float)(width), cgoAllocsUnknown
	cheight, _ := (C.float)(height), cgoAllocsUnknown
	cdensity, _ := (C.float)(density), cgoAllocsUnknown
	__ret := C.CreatePhysicsBodyRectangle(cpos, cwidth, cheight, cdensity)
	__v := newPhysicsBodyDataRef(unsafe.Pointer(__ret)).convert()
	return __v
}

// CreatePhysicsBodyPolygon function as declared in src/physac.h:213
func CreatePhysicsBodyPolygon(pos Vector2, radius float32, sides int32, density float32) *PhysicsBodyData {
	cpos, _ := *(*C.Vector2)(unsafe.Pointer(&pos)), cgoAllocsUnknown
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	csides, _ := (C.int)(sides), cgoAllocsUnknown
	cdensity, _ := (C.float)(density), cgoAllocsUnknown
	__ret := C.CreatePhysicsBodyPolygon(cpos, cradius, csides, cdensity)
	__v := newPhysicsBodyDataRef(unsafe.Pointer(__ret)).convert()
	return __v
}

// PhysicsAddForce function as declared in src/physac.h:214
func PhysicsAddForce(body *PhysicsBodyData, force Vector2) {
	cbody, _ := (*C.PhysicsBodyData)(unsafe.Pointer(body)), cgoAllocsUnknown
	cforce, _ := *(*C.Vector2)(unsafe.Pointer(&force)), cgoAllocsUnknown
	C.PhysicsAddForce(cbody, cforce)
}

// PhysicsAddTorque function as declared in src/physac.h:215
func PhysicsAddTorque(body *PhysicsBodyData, amount float32) {
	cbody, _ := (*C.PhysicsBodyData)(unsafe.Pointer(body)), cgoAllocsUnknown
	camount, _ := (C.float)(amount), cgoAllocsUnknown
	C.PhysicsAddTorque(cbody, camount)
}

// PhysicsShatter function as declared in src/physac.h:216
func PhysicsShatter(body *PhysicsBodyData, position Vector2, force float32) {
	cbody, _ := (*C.PhysicsBodyData)(unsafe.Pointer(body)), cgoAllocsUnknown
	cposition, _ := *(*C.Vector2)(unsafe.Pointer(&position)), cgoAllocsUnknown
	cforce, _ := (C.float)(force), cgoAllocsUnknown
	C.PhysicsShatter(cbody, cposition, cforce)
}

// GetPhysicsBodiesCount function as declared in src/physac.h:217
func GetPhysicsBodiesCount() int32 {
	__ret := C.GetPhysicsBodiesCount()
	__v := (int32)(__ret)
	return __v
}

// GetPhysicsBody function as declared in src/physac.h:218
func GetPhysicsBody(index int32) *PhysicsBodyData {
	cindex, _ := (C.int)(index), cgoAllocsUnknown
	__ret := C.GetPhysicsBody(cindex)
	__v := newPhysicsBodyDataRef(unsafe.Pointer(__ret)).convert()
	return __v
}

// GetPhysicsShapeType function as declared in src/physac.h:219
func GetPhysicsShapeType(index int32) int32 {
	cindex, _ := (C.int)(index), cgoAllocsUnknown
	__ret := C.GetPhysicsShapeType(cindex)
	__v := (int32)(__ret)
	return __v
}

// GetPhysicsShapeVerticesCount function as declared in src/physac.h:220
func GetPhysicsShapeVerticesCount(index int32) int32 {
	cindex, _ := (C.int)(index), cgoAllocsUnknown
	__ret := C.GetPhysicsShapeVerticesCount(cindex)
	__v := (int32)(__ret)
	return __v
}

// GetPhysicsShapeVertex function as declared in src/physac.h:221
func GetPhysicsShapeVertex(body *PhysicsBodyData, vertex int32) Vector2 {
	cbody, _ := (*C.PhysicsBodyData)(unsafe.Pointer(body)), cgoAllocsUnknown
	cvertex, _ := (C.int)(vertex), cgoAllocsUnknown
	__ret := C.GetPhysicsShapeVertex(cbody, cvertex)
	__v := *newVector2Ref(unsafe.Pointer(&__ret)).convert()
	return __v
}

// SetPhysicsBodyRotation function as declared in src/physac.h:222
func SetPhysicsBodyRotation(body *PhysicsBodyData, radians float32) {
	cbody, _ := (*C.PhysicsBodyData)(unsafe.Pointer(body)), cgoAllocsUnknown
	cradians, _ := (C.float)(radians), cgoAllocsUnknown
	C.SetPhysicsBodyRotation(cbody, cradians)
}

// DestroyPhysicsBody function as declared in src/physac.h:223
func DestroyPhysicsBody(body *PhysicsBodyData) {
	cbody, _ := (*C.PhysicsBodyData)(unsafe.Pointer(body)), cgoAllocsUnknown
	C.DestroyPhysicsBody(cbody)
}

// ResetPhysics function as declared in src/physac.h:224
func ResetPhysics() {
	C.ResetPhysics()
}

// ClosePhysics function as declared in src/physac.h:225
func ClosePhysics() {
	C.ClosePhysics()
}

// FindAvailableBodyIndex function as declared in src/physac.h:377
func FindAvailableBodyIndex() int32 {
	__ret := C.FindAvailableBodyIndex()
	__v := (int32)(__ret)
	return __v
}

// CreateRandomPolygon function as declared in src/physac.h:378
func CreateRandomPolygon(radius float32, sides int32) PolygonData {
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	csides, _ := (C.int)(sides), cgoAllocsUnknown
	__ret := C.CreateRandomPolygon(cradius, csides)
	__v := *newPolygonDataRef(unsafe.Pointer(&__ret)).convert()
	return __v
}

// CreateRectanglePolygon function as declared in src/physac.h:379
func CreateRectanglePolygon(pos Vector2, size Vector2) PolygonData {
	cpos, _ := *(*C.Vector2)(unsafe.Pointer(&pos)), cgoAllocsUnknown
	csize, _ := *(*C.Vector2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.CreateRectanglePolygon(cpos, csize)
	__v := *newPolygonDataRef(unsafe.Pointer(&__ret)).convert()
	return __v
}

// PhysicsLoop function as declared in src/physac.h:380
func PhysicsLoop(arg unsafe.Pointer) unsafe.Pointer {
	carg, _ := arg, cgoAllocsUnknown
	__ret := C.PhysicsLoop(carg)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// PhysicsStep function as declared in src/physac.h:381
func PhysicsStep() {
	C.PhysicsStep()
}

// FindAvailableManifoldIndex function as declared in src/physac.h:382
func FindAvailableManifoldIndex() int32 {
	__ret := C.FindAvailableManifoldIndex()
	__v := (int32)(__ret)
	return __v
}

// CreatePhysicsManifold function as declared in src/physac.h:383
func CreatePhysicsManifold(a *PhysicsBodyData, b *PhysicsBodyData) *PhysicsManifoldData {
	ca, _ := (*C.PhysicsBodyData)(unsafe.Pointer(a)), cgoAllocsUnknown
	cb, _ := (*C.PhysicsBodyData)(unsafe.Pointer(b)), cgoAllocsUnknown
	__ret := C.CreatePhysicsManifold(ca, cb)
	__v := newPhysicsManifoldDataRef(unsafe.Pointer(__ret)).convert()
	return __v
}

// DestroyPhysicsManifold function as declared in src/physac.h:384
func DestroyPhysicsManifold(manifold *PhysicsManifoldData) {
	cmanifold, _ := (*C.PhysicsManifoldData)(unsafe.Pointer(manifold)), cgoAllocsUnknown
	C.DestroyPhysicsManifold(cmanifold)
}

// SolvePhysicsManifold function as declared in src/physac.h:385
func SolvePhysicsManifold(manifold *PhysicsManifoldData) {
	cmanifold, _ := (*C.PhysicsManifoldData)(unsafe.Pointer(manifold)), cgoAllocsUnknown
	C.SolvePhysicsManifold(cmanifold)
}

// SolveCircleToCircle function as declared in src/physac.h:386
func SolveCircleToCircle(manifold *PhysicsManifoldData) {
	cmanifold, _ := (*C.PhysicsManifoldData)(unsafe.Pointer(manifold)), cgoAllocsUnknown
	C.SolveCircleToCircle(cmanifold)
}

// SolveCircleToPolygon function as declared in src/physac.h:387
func SolveCircleToPolygon(manifold *PhysicsManifoldData) {
	cmanifold, _ := (*C.PhysicsManifoldData)(unsafe.Pointer(manifold)), cgoAllocsUnknown
	C.SolveCircleToPolygon(cmanifold)
}

// SolvePolygonToCircle function as declared in src/physac.h:388
func SolvePolygonToCircle(manifold *PhysicsManifoldData) {
	cmanifold, _ := (*C.PhysicsManifoldData)(unsafe.Pointer(manifold)), cgoAllocsUnknown
	C.SolvePolygonToCircle(cmanifold)
}

// SolvePolygonToPolygon function as declared in src/physac.h:389
func SolvePolygonToPolygon(manifold *PhysicsManifoldData) {
	cmanifold, _ := (*C.PhysicsManifoldData)(unsafe.Pointer(manifold)), cgoAllocsUnknown
	C.SolvePolygonToPolygon(cmanifold)
}

// IntegratePhysicsForces function as declared in src/physac.h:390
func IntegratePhysicsForces(body *PhysicsBodyData) {
	cbody, _ := (*C.PhysicsBodyData)(unsafe.Pointer(body)), cgoAllocsUnknown
	C.IntegratePhysicsForces(cbody)
}

// InitializePhysicsManifolds function as declared in src/physac.h:391
func InitializePhysicsManifolds(manifold *PhysicsManifoldData) {
	cmanifold, _ := (*C.PhysicsManifoldData)(unsafe.Pointer(manifold)), cgoAllocsUnknown
	C.InitializePhysicsManifolds(cmanifold)
}

// IntegratePhysicsImpulses function as declared in src/physac.h:392
func IntegratePhysicsImpulses(manifold *PhysicsManifoldData) {
	cmanifold, _ := (*C.PhysicsManifoldData)(unsafe.Pointer(manifold)), cgoAllocsUnknown
	C.IntegratePhysicsImpulses(cmanifold)
}

// IntegratePhysicsVelocity function as declared in src/physac.h:393
func IntegratePhysicsVelocity(body *PhysicsBodyData) {
	cbody, _ := (*C.PhysicsBodyData)(unsafe.Pointer(body)), cgoAllocsUnknown
	C.IntegratePhysicsVelocity(cbody)
}

// CorrectPhysicsPositions function as declared in src/physac.h:394
func CorrectPhysicsPositions(manifold *PhysicsManifoldData) {
	cmanifold, _ := (*C.PhysicsManifoldData)(unsafe.Pointer(manifold)), cgoAllocsUnknown
	C.CorrectPhysicsPositions(cmanifold)
}

// FindAxisLeastPenetration function as declared in src/physac.h:395
func FindAxisLeastPenetration(faceIndex *int32, shapeA PhysicsShape, shapeB PhysicsShape) float32 {
	cfaceIndex, _ := (*C.int)(unsafe.Pointer(faceIndex)), cgoAllocsUnknown
	cshapeA, _ := *(*C.PhysicsShape)(unsafe.Pointer(&shapeA)), cgoAllocsUnknown
	cshapeB, _ := *(*C.PhysicsShape)(unsafe.Pointer(&shapeB)), cgoAllocsUnknown
	__ret := C.FindAxisLeastPenetration(cfaceIndex, cshapeA, cshapeB)
	__v := (float32)(__ret)
	return __v
}

// FindIncidentFace function as declared in src/physac.h:396
func FindIncidentFace(v0 *Vector2, v1 *Vector2, ref PhysicsShape, inc PhysicsShape, index int32) {
	cv0, _ := (*C.Vector2)(unsafe.Pointer(v0)), cgoAllocsUnknown
	cv1, _ := (*C.Vector2)(unsafe.Pointer(v1)), cgoAllocsUnknown
	cref, _ := *(*C.PhysicsShape)(unsafe.Pointer(&ref)), cgoAllocsUnknown
	cinc, _ := *(*C.PhysicsShape)(unsafe.Pointer(&inc)), cgoAllocsUnknown
	cindex, _ := (C.int)(index), cgoAllocsUnknown
	C.FindIncidentFace(cv0, cv1, cref, cinc, cindex)
}

// Clip function as declared in src/physac.h:397
func Clip(normal Vector2, clip float32, faceA *Vector2, faceB *Vector2) int32 {
	cnormal, _ := *(*C.Vector2)(unsafe.Pointer(&normal)), cgoAllocsUnknown
	cclip, _ := (C.float)(clip), cgoAllocsUnknown
	cfaceA, _ := (*C.Vector2)(unsafe.Pointer(faceA)), cgoAllocsUnknown
	cfaceB, _ := (*C.Vector2)(unsafe.Pointer(faceB)), cgoAllocsUnknown
	__ret := C.Clip(cnormal, cclip, cfaceA, cfaceB)
	__v := (int32)(__ret)
	return __v
}

// BiasGreaterThan function as declared in src/physac.h:398
func BiasGreaterThan(valueA float32, valueB float32) bool {
	cvalueA, _ := (C.float)(valueA), cgoAllocsUnknown
	cvalueB, _ := (C.float)(valueB), cgoAllocsUnknown
	__ret := C.BiasGreaterThan(cvalueA, cvalueB)
	__v := (bool)(__ret)
	return __v
}

// TriangleBarycenter function as declared in src/physac.h:399
func TriangleBarycenter(v1 Vector2, v2 Vector2, v3 Vector2) Vector2 {
	cv1, _ := *(*C.Vector2)(unsafe.Pointer(&v1)), cgoAllocsUnknown
	cv2, _ := *(*C.Vector2)(unsafe.Pointer(&v2)), cgoAllocsUnknown
	cv3, _ := *(*C.Vector2)(unsafe.Pointer(&v3)), cgoAllocsUnknown
	__ret := C.TriangleBarycenter(cv1, cv2, cv3)
	__v := *newVector2Ref(unsafe.Pointer(&__ret)).convert()
	return __v
}

// InitTimer function as declared in src/physac.h:401
func InitTimer() {
	C.InitTimer()
}

// GetTimeCount function as declared in src/physac.h:402
func GetTimeCount() uint64 {
	__ret := C.GetTimeCount()
	__v := (uint64)(__ret)
	return __v
}

// GetCurrentTime function as declared in src/physac.h:403
func GetCurrentTime() float64 {
	__ret := C.GetCurrentTime()
	__v := (float64)(__ret)
	return __v
}

// MathCross function as declared in src/physac.h:406
func MathCross(value float32, vector Vector2) Vector2 {
	cvalue, _ := (C.float)(value), cgoAllocsUnknown
	cvector, _ := *(*C.Vector2)(unsafe.Pointer(&vector)), cgoAllocsUnknown
	__ret := C.MathCross(cvalue, cvector)
	__v := *newVector2Ref(unsafe.Pointer(&__ret)).convert()
	return __v
}

// MathCrossVector2 function as declared in src/physac.h:407
func MathCrossVector2(v1 Vector2, v2 Vector2) float32 {
	cv1, _ := *(*C.Vector2)(unsafe.Pointer(&v1)), cgoAllocsUnknown
	cv2, _ := *(*C.Vector2)(unsafe.Pointer(&v2)), cgoAllocsUnknown
	__ret := C.MathCrossVector2(cv1, cv2)
	__v := (float32)(__ret)
	return __v
}

// MathLenSqr function as declared in src/physac.h:408
func MathLenSqr(vector Vector2) float32 {
	cvector, _ := *(*C.Vector2)(unsafe.Pointer(&vector)), cgoAllocsUnknown
	__ret := C.MathLenSqr(cvector)
	__v := (float32)(__ret)
	return __v
}

// MathDot function as declared in src/physac.h:409
func MathDot(v1 Vector2, v2 Vector2) float32 {
	cv1, _ := *(*C.Vector2)(unsafe.Pointer(&v1)), cgoAllocsUnknown
	cv2, _ := *(*C.Vector2)(unsafe.Pointer(&v2)), cgoAllocsUnknown
	__ret := C.MathDot(cv1, cv2)
	__v := (float32)(__ret)
	return __v
}

// DistSqr function as declared in src/physac.h:410
func DistSqr(v1 Vector2, v2 Vector2) float32 {
	cv1, _ := *(*C.Vector2)(unsafe.Pointer(&v1)), cgoAllocsUnknown
	cv2, _ := *(*C.Vector2)(unsafe.Pointer(&v2)), cgoAllocsUnknown
	__ret := C.DistSqr(cv1, cv2)
	__v := (float32)(__ret)
	return __v
}

// MathNormalize function as declared in src/physac.h:411
func MathNormalize(vector *Vector2) {
	cvector, _ := (*C.Vector2)(unsafe.Pointer(vector)), cgoAllocsUnknown
	C.MathNormalize(cvector)
}

// Mat2Radians function as declared in src/physac.h:417
func Mat2Radians(radians float32) Matrix2x2 {
	cradians, _ := (C.float)(radians), cgoAllocsUnknown
	__ret := C.Mat2Radians(cradians)
	__v := *newMatrix2x2Ref(unsafe.Pointer(&__ret)).convert()
	return __v
}

// Mat2Set function as declared in src/physac.h:418
func Mat2Set(matrix *Matrix2x2, radians float32) {
	cmatrix, _ := (*C.Matrix2x2)(unsafe.Pointer(matrix)), cgoAllocsUnknown
	cradians, _ := (C.float)(radians), cgoAllocsUnknown
	C.Mat2Set(cmatrix, cradians)
}

// Mat2Transpose function as declared in src/physac.h:419
func Mat2Transpose(matrix Matrix2x2) Matrix2x2 {
	cmatrix, _ := *(*C.Matrix2x2)(unsafe.Pointer(&matrix)), cgoAllocsUnknown
	__ret := C.Mat2Transpose(cmatrix)
	__v := *newMatrix2x2Ref(unsafe.Pointer(&__ret)).convert()
	return __v
}

// Mat2MultiplyVector2 function as declared in src/physac.h:420
func Mat2MultiplyVector2(matrix Matrix2x2, vector Vector2) Vector2 {
	cmatrix, _ := *(*C.Matrix2x2)(unsafe.Pointer(&matrix)), cgoAllocsUnknown
	cvector, _ := *(*C.Vector2)(unsafe.Pointer(&vector)), cgoAllocsUnknown
	__ret := C.Mat2MultiplyVector2(cmatrix, cvector)
	__v := *newVector2Ref(unsafe.Pointer(&__ret)).convert()
	return __v
}