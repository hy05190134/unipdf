//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package draw has handy features for defining paths which can be used to draw content on a PDF page.  Handles
// defining paths as points, vector calculations and conversion to PDF content stream data which can be used in
// page content streams and XObject forms and thus also in annotation appearance streams.
//
// Also defines utility functions for drawing common shapes such as rectangles, lines and circles (ovals).
package draw ;import (_e "fmt";_gg "github.com/unidoc/unipdf/v3/contentstream";_cb "github.com/unidoc/unipdf/v3/core";_gd "github.com/unidoc/unipdf/v3/internal/transform";_c "github.com/unidoc/unipdf/v3/model";_a "math";);

// Line defines a line shape between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none (regular line),
// or arrows at either end.  The line also has a specified width, color and opacity.
type Line struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor *_c .PdfColorDeviceRGB ;Opacity float64 ;LineWidth float64 ;LineEndingStyle1 LineEndingStyle ;LineEndingStyle2 LineEndingStyle ;LineStyle LineStyle ;};

// Draw draws the circle. Can specify a graphics state (gsName) for setting opacity etc.  Otherwise leave empty ("").
// Returns the content stream as a byte array, the bounding box and an error on failure.
func (_eb Circle )Draw (gsName string )([]byte ,*_c .PdfRectangle ,error ){_gcd :=_eb .Width /2;_ge :=_eb .Height /2;if _eb .BorderEnabled {_gcd -=_eb .BorderWidth /2;_ge -=_eb .BorderWidth /2;};_cg :=0.551784;_caf :=_gcd *_cg ;_gff :=_ge *_cg ;_ceb :=NewCubicBezierPath ();_ceb =_ceb .AppendCurve (NewCubicBezierCurve (-_gcd ,0,-_gcd ,_gff ,-_caf ,_ge ,0,_ge ));_ceb =_ceb .AppendCurve (NewCubicBezierCurve (0,_ge ,_caf ,_ge ,_gcd ,_gff ,_gcd ,0));_ceb =_ceb .AppendCurve (NewCubicBezierCurve (_gcd ,0,_gcd ,-_gff ,_caf ,-_ge ,0,-_ge ));_ceb =_ceb .AppendCurve (NewCubicBezierCurve (0,-_ge ,-_caf ,-_ge ,-_gcd ,-_gff ,-_gcd ,0));_ceb =_ceb .Offset (_gcd ,_ge );if _eb .BorderEnabled {_ceb =_ceb .Offset (_eb .BorderWidth /2,_eb .BorderWidth /2);};if _eb .X !=0||_eb .Y !=0{_ceb =_ceb .Offset (_eb .X ,_eb .Y );};_gaf :=_gg .NewContentCreator ();_gaf .Add_q ();if _eb .FillEnabled {_gaf .Add_rg (_eb .FillColor .R (),_eb .FillColor .G (),_eb .FillColor .B ());};if _eb .BorderEnabled {_gaf .Add_RG (_eb .BorderColor .R (),_eb .BorderColor .G (),_eb .BorderColor .B ());_gaf .Add_w (_eb .BorderWidth );};if len (gsName )> 1{_gaf .Add_gs (_cb .PdfObjectName (gsName ));};DrawBezierPathWithCreator (_ceb ,_gaf );_gaf .Add_h ();if _eb .FillEnabled &&_eb .BorderEnabled {_gaf .Add_B ();}else if _eb .FillEnabled {_gaf .Add_f ();}else if _eb .BorderEnabled {_gaf .Add_S ();};_gaf .Add_Q ();_gdg :=_ceb .GetBoundingBox ();if _eb .BorderEnabled {_gdg .Height +=_eb .BorderWidth ;_gdg .Width +=_eb .BorderWidth ;_gdg .X -=_eb .BorderWidth /2;_gdg .Y -=_eb .BorderWidth /2;};return _gaf .Bytes (),_gdg .ToPdfRectangle (),nil ;};

// Copy returns a clone of the Bezier path.
func (_cac CubicBezierPath )Copy ()CubicBezierPath {_b :=CubicBezierPath {};_b .Curves =[]CubicBezierCurve {};for _ ,_bf :=range _cac .Curves {_b .Curves =append (_b .Curves ,_bf );};return _b ;};

// Vector represents a two-dimensional vector.
type Vector struct{Dx float64 ;Dy float64 ;};

// Scale scales the vector by the specified factor.
func (_cbd Vector )Scale (factor float64 )Vector {_efd :=_cbd .Magnitude ();_bfbd :=_cbd .GetPolarAngle ();_cbd .Dx =factor *_efd *_a .Cos (_bfbd );_cbd .Dy =factor *_efd *_a .Sin (_bfbd );return _cbd ;};

// AppendCurve appends the specified Bezier curve to the path.
func (_ca CubicBezierPath )AppendCurve (curve CubicBezierCurve )CubicBezierPath {_ca .Curves =append (_ca .Curves ,curve );return _ca ;};

// Rotate rotates the vector by the specified angle.
func (_bfd Vector )Rotate (phi float64 )Vector {_ecc :=_bfd .Magnitude ();_aea :=_bfd .GetPolarAngle ();return NewVectorPolar (_ecc ,_aea +phi );};

// Point represents a two-dimensional point.
type Point struct{X float64 ;Y float64 ;};

// Magnitude returns the magnitude of the vector.
func (_bgd Vector )Magnitude ()float64 {return _a .Sqrt (_a .Pow (_bgd .Dx ,2.0)+_a .Pow (_bgd .Dy ,2.0))};

// Draw draws the basic line to PDF. Generates the content stream which can be used in page contents or appearance
// stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error if
// one occurred.
func (_gfd BasicLine )Draw (gsName string )([]byte ,*_c .PdfRectangle ,error ){_dbc :=_gfd .LineWidth ;_bb :=NewPath ();_bb =_bb .AppendPoint (NewPoint (_gfd .X1 ,_gfd .Y1 ));_bb =_bb .AppendPoint (NewPoint (_gfd .X2 ,_gfd .Y2 ));_accc :=_gg .NewContentCreator ();_aad :=_bb .GetBoundingBox ();DrawPathWithCreator (_bb ,_accc );if _gfd .LineStyle ==LineStyleDashed {_accc .Add_d ([]int64 {1,1},0);};_accc .Add_RG (_gfd .LineColor .R (),_gfd .LineColor .G (),_gfd .LineColor .B ()).Add_w (_dbc ).Add_S ().Add_Q ();return _accc .Bytes (),_aad .ToPdfRectangle (),nil ;};

// GetPolarAngle returns the angle the magnitude of the vector forms with the
// positive X-axis going counterclockwise.
func (_egf Vector )GetPolarAngle ()float64 {return _a .Atan2 (_egf .Dy ,_egf .Dx )};

// CubicBezierCurve is defined by:
// R(t) = P0*(1-t)^3 + P1*3*t*(1-t)^2 + P2*3*t^2*(1-t) + P3*t^3
// where P0 is the current point, P1, P2 control points and P3 the final point.
type CubicBezierCurve struct{P0 Point ;P1 Point ;P2 Point ;P3 Point ;};

// GetBoundingBox returns the bounding box of the Bezier path.
func (_ag CubicBezierPath )GetBoundingBox ()Rectangle {_ec :=Rectangle {};_af :=0.0;_ggb :=0.0;_ecb :=0.0;_ce :=0.0;for _fb ,_da :=range _ag .Curves {_gaa :=_da .GetBounds ();if _fb ==0{_af =_gaa .Llx ;_ggb =_gaa .Urx ;_ecb =_gaa .Lly ;_ce =_gaa .Ury ;continue ;};if _gaa .Llx < _af {_af =_gaa .Llx ;};if _gaa .Urx > _ggb {_ggb =_gaa .Urx ;};if _gaa .Lly < _ecb {_ecb =_gaa .Lly ;};if _gaa .Ury > _ce {_ce =_gaa .Ury ;};};_ec .X =_af ;_ec .Y =_ecb ;_ec .Width =_ggb -_af ;_ec .Height =_ce -_ecb ;return _ec ;};

// NewPath returns a new empty path.
func NewPath ()Path {return Path {}};

// Draw draws the polygon. A graphics state name can be specified for
// setting the polygon properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polygon
// bounding box.
func (_bae Polygon )Draw (gsName string )([]byte ,*_c .PdfRectangle ,error ){_cff :=_gg .NewContentCreator ();_cff .Add_q ();_bae .FillEnabled =_bae .FillEnabled &&_bae .FillColor !=nil ;if _bae .FillEnabled {_cff .Add_rg (_bae .FillColor .R (),_bae .FillColor .G (),_bae .FillColor .B ());};_bae .BorderEnabled =_bae .BorderEnabled &&_bae .BorderColor !=nil ;if _bae .BorderEnabled {_cff .Add_RG (_bae .BorderColor .R (),_bae .BorderColor .G (),_bae .BorderColor .B ());_cff .Add_w (_bae .BorderWidth );};if len (gsName )> 1{_cff .Add_gs (_cb .PdfObjectName (gsName ));};_fcf :=NewPath ();for _ ,_cgg :=range _bae .Points {for _bab ,_dd :=range _cgg {_fcf =_fcf .AppendPoint (_dd );if _bab ==0{_cff .Add_m (_dd .X ,_dd .Y );}else {_cff .Add_l (_dd .X ,_dd .Y );};};_cff .Add_h ();};if _bae .FillEnabled &&_bae .BorderEnabled {_cff .Add_B ();}else if _bae .FillEnabled {_cff .Add_f ();}else if _bae .BorderEnabled {_cff .Add_S ();};_cff .Add_Q ();return _cff .Bytes (),_fcf .GetBoundingBox ().ToPdfRectangle (),nil ;};

// BasicLine defines a line between point 1 (X1,Y1) and point 2 (X2,Y2). The line has a specified width, color and opacity.
type BasicLine struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor *_c .PdfColorDeviceRGB ;Opacity float64 ;LineWidth float64 ;LineStyle LineStyle ;};

// Draw draws the rectangle. Can specify a graphics state (gsName) for setting opacity etc.
// Otherwise leave empty (""). Returns the content stream as a byte array, bounding box and an error on failure.
func (_ccb Rectangle )Draw (gsName string )([]byte ,*_c .PdfRectangle ,error ){_ee :=NewPath ();_ee =_ee .AppendPoint (NewPoint (0,0));_ee =_ee .AppendPoint (NewPoint (0,_ccb .Height ));_ee =_ee .AppendPoint (NewPoint (_ccb .Width ,_ccb .Height ));_ee =_ee .AppendPoint (NewPoint (_ccb .Width ,0));_ee =_ee .AppendPoint (NewPoint (0,0));if _ccb .X !=0||_ccb .Y !=0{_ee =_ee .Offset (_ccb .X ,_ccb .Y );};_dca :=_gg .NewContentCreator ();_dca .Add_q ();if _ccb .FillEnabled {_dca .Add_rg (_ccb .FillColor .R (),_ccb .FillColor .G (),_ccb .FillColor .B ());};if _ccb .BorderEnabled {_dca .Add_RG (_ccb .BorderColor .R (),_ccb .BorderColor .G (),_ccb .BorderColor .B ());_dca .Add_w (_ccb .BorderWidth );};if len (gsName )> 1{_dca .Add_gs (_cb .PdfObjectName (gsName ));};DrawPathWithCreator (_ee ,_dca );_dca .Add_h ();if _ccb .FillEnabled &&_ccb .BorderEnabled {_dca .Add_B ();}else if _ccb .FillEnabled {_dca .Add_f ();}else if _ccb .BorderEnabled {_dca .Add_S ();};_dca .Add_Q ();return _dca .Bytes (),_ee .GetBoundingBox ().ToPdfRectangle (),nil ;};

// GetPointNumber returns the path point at the index specified by number.
// The index is 1-based.
func (_bc Path )GetPointNumber (number int )Point {if number < 1||number > len (_bc .Points ){return Point {};};return _bc .Points [number -1];};

// Polygon is a multi-point shape that can be drawn to a PDF content stream.
type Polygon struct{Points [][]Point ;FillEnabled bool ;FillColor *_c .PdfColorDeviceRGB ;BorderEnabled bool ;BorderColor *_c .PdfColorDeviceRGB ;BorderWidth float64 ;};

// NewVectorPolar returns a new vector calculated from the specified
// magnitude and angle.
func NewVectorPolar (length float64 ,theta float64 )Vector {_fcbe :=Vector {};_fcbe .Dx =length *_a .Cos (theta );_fcbe .Dy =length *_a .Sin (theta );return _fcbe ;};

// NewPoint returns a new point with the coordinates x, y.
func NewPoint (x ,y float64 )Point {return Point {X :x ,Y :y }};

// Rotate returns a new Point at `p` rotated by `theta` degrees.
func (_eae Point )Rotate (theta float64 )Point {_ecg :=_gd .NewPoint (_eae .X ,_eae .Y ).Rotate (theta );return NewPoint (_ecg .X ,_ecg .Y );};

// ToPdfRectangle returns the rectangle as a PDF rectangle.
func (_ccfa Rectangle )ToPdfRectangle ()*_c .PdfRectangle {return &_c .PdfRectangle {Llx :_ccfa .X ,Lly :_ccfa .Y ,Urx :_ccfa .X +_ccfa .Width ,Ury :_ccfa .Y +_ccfa .Height };};

// DrawBezierPathWithCreator makes the bezier path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawBezierPathWithCreator (bpath CubicBezierPath ,creator *_gg .ContentCreator ){for _cffg ,_ece :=range bpath .Curves {if _cffg ==0{creator .Add_m (_ece .P0 .X ,_ece .P0 .Y );};creator .Add_c (_ece .P1 .X ,_ece .P1 .Y ,_ece .P2 .X ,_ece .P2 .Y ,_ece .P3 .X ,_ece .P3 .Y );};};

// CubicBezierPath represents a collection of cubic Bezier curves.
type CubicBezierPath struct{Curves []CubicBezierCurve ;};

// GetBounds returns the bounding box of the Bezier curve.
func (_dcg CubicBezierCurve )GetBounds ()_c .PdfRectangle {_ad :=_dcg .P0 .X ;_aa :=_dcg .P0 .X ;_f :=_dcg .P0 .Y ;_fa :=_dcg .P0 .Y ;for _ga :=0.0;_ga <=1.0;_ga +=0.001{Rx :=_dcg .P0 .X *_a .Pow (1-_ga ,3)+_dcg .P1 .X *3*_ga *_a .Pow (1-_ga ,2)+_dcg .P2 .X *3*_a .Pow (_ga ,2)*(1-_ga )+_dcg .P3 .X *_a .Pow (_ga ,3);Ry :=_dcg .P0 .Y *_a .Pow (1-_ga ,3)+_dcg .P1 .Y *3*_ga *_a .Pow (1-_ga ,2)+_dcg .P2 .Y *3*_a .Pow (_ga ,2)*(1-_ga )+_dcg .P3 .Y *_a .Pow (_ga ,3);if Rx < _ad {_ad =Rx ;};if Rx > _aa {_aa =Rx ;};if Ry < _f {_f =Ry ;};if Ry > _fa {_fa =Ry ;};};_gdd :=_c .PdfRectangle {};_gdd .Llx =_ad ;_gdd .Lly =_f ;_gdd .Urx =_aa ;_gdd .Ury =_fa ;return _gdd ;};

// Circle represents a circle shape with fill and border properties that can be drawn to a PDF content stream.
type Circle struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;FillEnabled bool ;FillColor *_c .PdfColorDeviceRGB ;BorderEnabled bool ;BorderWidth float64 ;BorderColor *_c .PdfColorDeviceRGB ;Opacity float64 ;};

// Flip changes the sign of the vector: -vector.
func (_bbc Vector )Flip ()Vector {_gddf :=_bbc .Magnitude ();_ecge :=_bbc .GetPolarAngle ();_bbc .Dx =_gddf *_a .Cos (_ecge +_a .Pi );_bbc .Dy =_gddf *_a .Sin (_ecge +_a .Pi );return _bbc ;};

// Polyline defines a slice of points that are connected as straight lines.
type Polyline struct{Points []Point ;LineColor *_c .PdfColorDeviceRGB ;LineWidth float64 ;};

// Add adds the specified vector to the current one and returns the result.
func (_dcc Vector )Add (other Vector )Vector {_dcc .Dx +=other .Dx ;_dcc .Dy +=other .Dy ;return _dcc };

// GetBoundingBox returns the bounding box of the path.
func (_dfb Path )GetBoundingBox ()BoundingBox {_gc :=BoundingBox {};_aee :=0.0;_cf :=0.0;_cad :=0.0;_afg :=0.0;for _bd ,_ba :=range _dfb .Points {if _bd ==0{_aee =_ba .X ;_cf =_ba .X ;_cad =_ba .Y ;_afg =_ba .Y ;continue ;};if _ba .X < _aee {_aee =_ba .X ;};if _ba .X > _cf {_cf =_ba .X ;};if _ba .Y < _cad {_cad =_ba .Y ;};if _ba .Y > _afg {_afg =_ba .Y ;};};_gc .X =_aee ;_gc .Y =_cad ;_gc .Width =_cf -_aee ;_gc .Height =_afg -_cad ;return _gc ;};

// DrawPathWithCreator makes the path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawPathWithCreator (path Path ,creator *_gg .ContentCreator ){for _ffa ,_de :=range path .Points {if _ffa ==0{creator .Add_m (_de .X ,_de .Y );}else {creator .Add_l (_de .X ,_de .Y );};};};

// PolyBezierCurve represents a composite curve that is the result of
// joining multiple cubic Bezier curves.
type PolyBezierCurve struct{Curves []CubicBezierCurve ;BorderWidth float64 ;BorderColor *_c .PdfColorDeviceRGB ;FillEnabled bool ;FillColor *_c .PdfColorDeviceRGB ;};

// RemovePoint removes the point at the index specified by number from the
// path. The index is 1-based.
func (_ggbb Path )RemovePoint (number int )Path {if number < 1||number > len (_ggbb .Points ){return _ggbb ;};_ac :=number -1;_ggbb .Points =append (_ggbb .Points [:_ac ],_ggbb .Points [_ac +1:]...);return _ggbb ;};

// ToPdfRectangle returns the bounding box as a PDF rectangle.
func (_fad BoundingBox )ToPdfRectangle ()*_c .PdfRectangle {return &_c .PdfRectangle {Llx :_fad .X ,Lly :_fad .Y ,Urx :_fad .X +_fad .Width ,Ury :_fad .Y +_fad .Height };};

// Path consists of straight line connections between each point defined in an array of points.
type Path struct{Points []Point ;};const (LineEndingStyleNone LineEndingStyle =0;LineEndingStyleArrow LineEndingStyle =1;LineEndingStyleButt LineEndingStyle =2;);

// Length returns the number of points in the path.
func (_cec Path )Length ()int {return len (_cec .Points )};

// FlipX flips the sign of the Dx component of the vector.
func (_acg Vector )FlipX ()Vector {_acg .Dx =-_acg .Dx ;return _acg };

// LineStyle refers to how the line will be created.
type LineStyle int ;

// Draw draws the line to PDF contentstream. Generates the content stream which can be used in page contents or
// appearance stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error
// if one occurred.
func (_gde Line )Draw (gsName string )([]byte ,*_c .PdfRectangle ,error ){_aaf ,_fcb :=_gde .X1 ,_gde .X2 ;_adgg ,_acca :=_gde .Y1 ,_gde .Y2 ;_cba :=_acca -_adgg ;_ef :=_fcb -_aaf ;_dab :=_a .Atan2 (_cba ,_ef );L :=_a .Sqrt (_a .Pow (_ef ,2.0)+_a .Pow (_cba ,2.0));_efe :=_gde .LineWidth ;_ccbg :=_a .Pi ;_fdd :=1.0;if _ef < 0{_fdd *=-1.0;};if _cba < 0{_fdd *=-1.0;};VsX :=_fdd *(-_efe /2*_a .Cos (_dab +_ccbg /2));VsY :=_fdd *(-_efe /2*_a .Sin (_dab +_ccbg /2)+_efe *_a .Sin (_dab +_ccbg /2));V1X :=VsX +_efe /2*_a .Cos (_dab +_ccbg /2);V1Y :=VsY +_efe /2*_a .Sin (_dab +_ccbg /2);V2X :=VsX +_efe /2*_a .Cos (_dab +_ccbg /2)+L *_a .Cos (_dab );V2Y :=VsY +_efe /2*_a .Sin (_dab +_ccbg /2)+L *_a .Sin (_dab );V3X :=VsX +_efe /2*_a .Cos (_dab +_ccbg /2)+L *_a .Cos (_dab )+_efe *_a .Cos (_dab -_ccbg /2);V3Y :=VsY +_efe /2*_a .Sin (_dab +_ccbg /2)+L *_a .Sin (_dab )+_efe *_a .Sin (_dab -_ccbg /2);V4X :=VsX +_efe /2*_a .Cos (_dab -_ccbg /2);V4Y :=VsY +_efe /2*_a .Sin (_dab -_ccbg /2);_dda :=NewPath ();_dda =_dda .AppendPoint (NewPoint (V1X ,V1Y ));_dda =_dda .AppendPoint (NewPoint (V2X ,V2Y ));_dda =_dda .AppendPoint (NewPoint (V3X ,V3Y ));_dda =_dda .AppendPoint (NewPoint (V4X ,V4Y ));_ddf :=_gde .LineEndingStyle1 ;_aeg :=_gde .LineEndingStyle2 ;_cadc :=3*_efe ;_bcg :=3*_efe ;_badf :=(_bcg -_efe )/2;if _aeg ==LineEndingStyleArrow {_adf :=_dda .GetPointNumber (2);_bfe :=NewVectorPolar (_cadc ,_dab +_ccbg );_eag :=_adf .AddVector (_bfe );_cga :=NewVectorPolar (_bcg /2,_dab +_ccbg /2);_ccc :=NewVectorPolar (_cadc ,_dab );_fbc :=NewVectorPolar (_badf ,_dab +_ccbg /2);_bac :=_eag .AddVector (_fbc );_cgf :=_ccc .Add (_cga .Flip ());_eaa :=_bac .AddVector (_cgf );_cbe :=_cga .Scale (2).Flip ().Add (_cgf .Flip ());_fe :=_eaa .AddVector (_cbe );_gcdb :=_eag .AddVector (NewVectorPolar (_efe ,_dab -_ccbg /2));_aeb :=NewPath ();_aeb =_aeb .AppendPoint (_dda .GetPointNumber (1));_aeb =_aeb .AppendPoint (_eag );_aeb =_aeb .AppendPoint (_bac );_aeb =_aeb .AppendPoint (_eaa );_aeb =_aeb .AppendPoint (_fe );_aeb =_aeb .AppendPoint (_gcdb );_aeb =_aeb .AppendPoint (_dda .GetPointNumber (4));_dda =_aeb ;};if _ddf ==LineEndingStyleArrow {_gafe :=_dda .GetPointNumber (1);_cfg :=_dda .GetPointNumber (_dda .Length ());_bea :=NewVectorPolar (_efe /2,_dab +_ccbg +_ccbg /2);_bg :=_gafe .AddVector (_bea );_dad :=NewVectorPolar (_cadc ,_dab ).Add (NewVectorPolar (_bcg /2,_dab +_ccbg /2));_aeed :=_bg .AddVector (_dad );_add :=NewVectorPolar (_badf ,_dab -_ccbg /2);_fgd :=_aeed .AddVector (_add );_efc :=NewVectorPolar (_cadc ,_dab );_adfb :=_cfg .AddVector (_efc );_bfb :=NewVectorPolar (_badf ,_dab +_ccbg +_ccbg /2);_bce :=_adfb .AddVector (_bfb );_baed :=_bg ;_gda :=NewPath ();_gda =_gda .AppendPoint (_bg );_gda =_gda .AppendPoint (_aeed );_gda =_gda .AppendPoint (_fgd );for _ ,_baedg :=range _dda .Points [1:len (_dda .Points )-1]{_gda =_gda .AppendPoint (_baedg );};_gda =_gda .AppendPoint (_adfb );_gda =_gda .AppendPoint (_bce );_gda =_gda .AppendPoint (_baed );_dda =_gda ;};_faa :=_gg .NewContentCreator ();_faa .Add_q ().Add_rg (_gde .LineColor .R (),_gde .LineColor .G (),_gde .LineColor .B ());if len (gsName )> 1{_faa .Add_gs (_cb .PdfObjectName (gsName ));};_dda =_dda .Offset (_gde .X1 ,_gde .Y1 );_bed :=_dda .GetBoundingBox ();DrawPathWithCreator (_dda ,_faa );if _gde .LineStyle ==LineStyleDashed {_faa .Add_d ([]int64 {1,1},0).Add_S ().Add_f ().Add_Q ();}else {_faa .Add_f ().Add_Q ();};return _faa .Bytes (),_bed .ToPdfRectangle (),nil ;};

// BoundingBox represents the smallest rectangular area that encapsulates an object.
type BoundingBox struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;};

// Draw draws the composite Bezier curve. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array and
// the curve bounding box.
func (_ab PolyBezierCurve )Draw (gsName string )([]byte ,*_c .PdfRectangle ,error ){if _ab .BorderColor ==nil {_ab .BorderColor =_c .NewPdfColorDeviceRGB (0,0,0);};_bad :=NewCubicBezierPath ();for _ ,_eg :=range _ab .Curves {_bad =_bad .AppendCurve (_eg );};_ccf :=_gg .NewContentCreator ();_ccf .Add_q ();_ab .FillEnabled =_ab .FillEnabled &&_ab .FillColor !=nil ;if _ab .FillEnabled {_ccf .Add_rg (_ab .FillColor .R (),_ab .FillColor .G (),_ab .FillColor .B ());};_ccf .Add_RG (_ab .BorderColor .R (),_ab .BorderColor .G (),_ab .BorderColor .B ());_ccf .Add_w (_ab .BorderWidth );if len (gsName )> 1{_ccf .Add_gs (_cb .PdfObjectName (gsName ));};for _ ,_be :=range _bad .Curves {_ccf .Add_m (_be .P0 .X ,_be .P0 .Y );_ccf .Add_c (_be .P1 .X ,_be .P1 .Y ,_be .P2 .X ,_be .P2 .Y ,_be .P3 .X ,_be .P3 .Y );};if _ab .FillEnabled {_ccf .Add_h ();_ccf .Add_B ();}else {_ccf .Add_S ();};_ccf .Add_Q ();return _ccf .Bytes (),_bad .GetBoundingBox ().ToPdfRectangle (),nil ;};

// AddVector adds vector to a point.
func (_acc Point )AddVector (v Vector )Point {_acc .X +=v .Dx ;_acc .Y +=v .Dy ;return _acc };

// NewVectorBetween returns a new vector with the direction specified by
// the subtraction of point a from point b (b-a).
func NewVectorBetween (a Point ,b Point )Vector {_feb :=Vector {};_feb .Dx =b .X -a .X ;_feb .Dy =b .Y -a .Y ;return _feb ;};

// NewCubicBezierCurve returns a new cubic Bezier curve.
func NewCubicBezierCurve (x0 ,y0 ,x1 ,y1 ,x2 ,y2 ,x3 ,y3 float64 )CubicBezierCurve {_d :=CubicBezierCurve {};_d .P0 =NewPoint (x0 ,y0 );_d .P1 =NewPoint (x1 ,y1 );_d .P2 =NewPoint (x2 ,y2 );_d .P3 =NewPoint (x3 ,y3 );return _d ;};

// Offset shifts the Bezier path with the specified offsets.
func (_gf CubicBezierPath )Offset (offX ,offY float64 )CubicBezierPath {for _df ,_ff :=range _gf .Curves {_gf .Curves [_df ]=_ff .AddOffsetXY (offX ,offY );};return _gf ;};

// Rectangle is a shape with a specified Width and Height and a lower left corner at (X,Y) that can be
// drawn to a PDF content stream.  The rectangle can optionally have a border and a filling color.
// The Width/Height includes the border (if any specified), i.e. is positioned inside.
type Rectangle struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;FillEnabled bool ;FillColor *_c .PdfColorDeviceRGB ;BorderEnabled bool ;BorderWidth float64 ;BorderColor *_c .PdfColorDeviceRGB ;Opacity float64 ;};

// Copy returns a clone of the path.
func (_fc Path )Copy ()Path {_dcgg :=Path {};_dcgg .Points =[]Point {};for _ ,_fd :=range _fc .Points {_dcgg .Points =append (_dcgg .Points ,_fd );};return _dcgg ;};

// LineEndingStyle defines the line ending style for lines.
// The currently supported line ending styles are None, Arrow (ClosedArrow) and Butt.
type LineEndingStyle int ;

// Add shifts the coordinates of the point with dx, dy and returns the result.
func (_db Point )Add (dx ,dy float64 )Point {_db .X +=dx ;_db .Y +=dy ;return _db };

// AddOffsetXY adds X,Y offset to all points on a curve.
func (_dc CubicBezierCurve )AddOffsetXY (offX ,offY float64 )CubicBezierCurve {_dc .P0 .X +=offX ;_dc .P1 .X +=offX ;_dc .P2 .X +=offX ;_dc .P3 .X +=offX ;_dc .P0 .Y +=offY ;_dc .P1 .Y +=offY ;_dc .P2 .Y +=offY ;_dc .P3 .Y +=offY ;return _dc ;};func (_gb Point )String ()string {return _e .Sprintf ("(\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u0029",_gb .X ,_gb .Y );};

// FlipY flips the sign of the Dy component of the vector.
func (_fab Vector )FlipY ()Vector {_fab .Dy =-_fab .Dy ;return _fab };const (LineStyleSolid LineStyle =0;LineStyleDashed LineStyle =1;);

// Draw draws the polyline. A graphics state name can be specified for
// setting the polyline properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polyline
// bounding box.
func (_fcc Polyline )Draw (gsName string )([]byte ,*_c .PdfRectangle ,error ){if _fcc .LineColor ==nil {_fcc .LineColor =_c .NewPdfColorDeviceRGB (0,0,0);};_eff :=NewPath ();for _ ,_addg :=range _fcc .Points {_eff =_eff .AppendPoint (_addg );};_fca :=_gg .NewContentCreator ();_fca .Add_q ();_fca .Add_RG (_fcc .LineColor .R (),_fcc .LineColor .G (),_fcc .LineColor .B ());_fca .Add_w (_fcc .LineWidth );if len (gsName )> 1{_fca .Add_gs (_cb .PdfObjectName (gsName ));};DrawPathWithCreator (_eff ,_fca );_fca .Add_S ();_fca .Add_Q ();return _fca .Bytes (),_eff .GetBoundingBox ().ToPdfRectangle (),nil ;};

// AppendPoint adds the specified point to the path.
func (_fg Path )AppendPoint (point Point )Path {_fg .Points =append (_fg .Points ,point );return _fg };

// NewVector returns a new vector with the direction specified by dx and dy.
func NewVector (dx ,dy float64 )Vector {_fdb :=Vector {};_fdb .Dx =dx ;_fdb .Dy =dy ;return _fdb };

// NewCubicBezierPath returns a new empty cubic Bezier path.
func NewCubicBezierPath ()CubicBezierPath {_cc :=CubicBezierPath {};_cc .Curves =[]CubicBezierCurve {};return _cc ;};

// Offset shifts the path with the specified offsets.
func (_afe Path )Offset (offX ,offY float64 )Path {for _fdg ,_fgc :=range _afe .Points {_afe .Points [_fdg ]=_fgc .Add (offX ,offY );};return _afe ;};