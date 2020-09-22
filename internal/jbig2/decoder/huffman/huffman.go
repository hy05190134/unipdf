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

package huffman ;import (_e "errors";_d "fmt";_c "github.com/unidoc/unipdf/v3/internal/bitwise";_ad "math";_f "strings";);func _bfc (_gce ,_fbc int32 )int32 {if _gce > _fbc {return _gce ;};return _fbc ;};func _adad (_bcf []*Code ){var _dde int32 ;for _ ,_eae :=range _bcf {_dde =_bfc (_dde ,_eae ._gbg );};_agf :=make ([]int32 ,_dde +1);for _ ,_cfab :=range _bcf {_agf [_cfab ._gbg ]++;};var _agfd int32 ;_gbf :=make ([]int32 ,len (_agf )+1);_agf [0]=0;for _acf :=int32 (1);_acf <=int32 (len (_agf ));_acf ++{_gbf [_acf ]=(_gbf [_acf -1]+(_agf [_acf -1]))<<1;_agfd =_gbf [_acf ];for _ ,_gaf :=range _bcf {if _gaf ._gbg ==_acf {_gaf ._gde =_agfd ;_agfd ++;};};};};func (_ac *EncodedTable )parseTable ()error {var (_ddg []*Code ;_dg ,_ce ,_acg int32 ;_fd uint64 ;_gb error ;);_adb :=_ac .StreamReader ();_gf :=_ac .HtLow ();for _gf < _ac .HtHigh (){_fd ,_gb =_adb .ReadBits (byte (_ac .HtPS ()));if _gb !=nil {return _gb ;};_dg =int32 (_fd );_fd ,_gb =_adb .ReadBits (byte (_ac .HtRS ()));if _gb !=nil {return _gb ;};_ce =int32 (_fd );_ddg =append (_ddg ,NewCode (_dg ,_ce ,_acg ,false ));_gf +=1<<uint (_ce );};_fd ,_gb =_adb .ReadBits (byte (_ac .HtPS ()));if _gb !=nil {return _gb ;};_dg =int32 (_fd );_ce =32;_acg =_ac .HtLow ()-1;_ddg =append (_ddg ,NewCode (_dg ,_ce ,_acg ,true ));_fd ,_gb =_adb .ReadBits (byte (_ac .HtPS ()));if _gb !=nil {return _gb ;};_dg =int32 (_fd );_ce =32;_acg =_ac .HtHigh ();_ddg =append (_ddg ,NewCode (_dg ,_ce ,_acg ,false ));if _ac .HtOOB ()==1{_fd ,_gb =_adb .ReadBits (byte (_ac .HtPS ()));if _gb !=nil {return _gb ;};_dg =int32 (_fd );_ddg =append (_ddg ,NewCode (_dg ,-1,-1,false ));};if _gb =_ac .InitTree (_ddg );_gb !=nil {return _gb ;};return nil ;};func (_aa *StandardTable )InitTree (codeTable []*Code )error {_adad (codeTable );for _ ,_gfg :=range codeTable {if _db :=_aa ._bcd .append (_gfg );_db !=nil {return _db ;};};return nil ;};func _bd (_ga *Code )*OutOfBandNode {return &OutOfBandNode {}};func (_cd *EncodedTable )Decode (r _c .StreamReader )(int64 ,error ){return _cd ._g .Decode (r )};var _ Tabler =&EncodedTable {};func GetStandardTable (number int )(Tabler ,error ){if number <=0||number > len (_cfg ){return nil ,_e .New ("\u0049n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");};_bbcg :=_cfg [number -1];if _bbcg ==nil {var _bbb error ;_bbcg ,_bbb =_fggd (_ee [number -1]);if _bbb !=nil {return nil ,_bbb ;};_cfg [number -1]=_bbcg ;};return _bbcg ,nil ;};func (_cgd *Code )String ()string {var _adbc string ;if _cgd ._gde !=-1{_adbc =_gac (_cgd ._gde ,_cgd ._gbg );}else {_adbc ="\u003f";};return _d .Sprintf ("%\u0073\u002f\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064",_adbc ,_cgd ._gbg ,_cgd ._aea ,_cgd ._fad );};type StandardTable struct{_bcd *InternalNode };func (_cca *OutOfBandNode )Decode (r _c .StreamReader )(int64 ,error ){return int64 (_ad .MaxInt64 ),nil };type FixedSizeTable struct{_feg *InternalNode };func (_eff *ValueNode )String ()string {return _d .Sprintf ("\u0025\u0064\u002f%\u0064",_eff ._ed ,_eff ._dc );};func (_b *FixedSizeTable )Decode (r _c .StreamReader )(int64 ,error ){return _b ._feg .Decode (r )};func _fggd (_eg [][]int32 )(*StandardTable ,error ){var _ccf []*Code ;for _bdaa :=0;_bdaa < len (_eg );_bdaa ++{_cea :=_eg [_bdaa ][0];_dgb :=_eg [_bdaa ][1];_bg :=_eg [_bdaa ][2];var _cfa bool ;if len (_eg [_bdaa ])> 3{_cfa =true ;};_ccf =append (_ccf ,NewCode (_cea ,_dgb ,_bg ,_cfa ));};_cdae :=&StandardTable {_bcd :_cdg (0)};if _dbc :=_cdae .InitTree (_ccf );_dbc !=nil {return nil ,_dbc ;};return _cdae ,nil ;};func (_cce *InternalNode )Decode (r _c .StreamReader )(int64 ,error ){_be ,_acgg :=r .ReadBit ();if _acgg !=nil {return 0,_acgg ;};if _be ==0{return _cce ._cef .Decode (r );};return _cce ._bc .Decode (r );};func (_ea *StandardTable )RootNode ()*InternalNode {return _ea ._bcd };func (_aed *FixedSizeTable )RootNode ()*InternalNode {return _aed ._feg };func (_ba *ValueNode )Decode (r _c .StreamReader )(int64 ,error ){_bda ,_ef :=r .ReadBits (byte (_ba ._ed ));if _ef !=nil {return 0,_ef ;};if _ba ._df {_bda =-_bda ;};return int64 (_ba ._dc )+int64 (_bda ),nil ;};func (_ab *EncodedTable )InitTree (codeTable []*Code )error {_adad (codeTable );for _ ,_gge :=range codeTable {if _fe :=_ab ._g .append (_gge );_fe !=nil {return _fe ;};};return nil ;};func (_dgd *InternalNode )pad (_fc *_f .Builder ){for _gd :=int32 (0);_gd < _dgd ._bbc ;_gd ++{_fc .WriteString ("\u0020\u0020\u0020");};};func (_ae *FixedSizeTable )InitTree (codeTable []*Code )error {_adad (codeTable );for _ ,_fg :=range codeTable {_eb :=_ae ._feg .append (_fg );if _eb !=nil {return _eb ;};};return nil ;};var _ee =[][][]int32 {{{1,4,0},{2,8,16},{3,16,272},{3,32,65808}},{{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{6,32,75},{6,-1,0}},{{8,8,-256},{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{8,32,-257,999},{7,32,75},{6,-1,0}},{{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{5,32,76}},{{7,8,-255},{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{7,32,-256,999},{6,32,76}},{{5,10,-2048},{4,9,-1024},{4,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{2,7,0},{3,7,128},{3,8,256},{4,9,512},{4,10,1024},{6,32,-2049,999},{6,32,2048}},{{4,9,-1024},{3,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{4,5,0},{5,5,32},{5,6,64},{4,7,128},{3,8,256},{3,9,512},{3,10,1024},{5,32,-1025,999},{5,32,2048}},{{8,3,-15},{9,1,-7},{8,1,-5},{9,0,-3},{7,0,-2},{4,0,-1},{2,1,0},{5,0,2},{6,0,3},{3,4,4},{6,1,20},{4,4,22},{4,5,38},{5,6,70},{5,7,134},{6,7,262},{7,8,390},{6,10,646},{9,32,-16,999},{9,32,1670},{2,-1,0}},{{8,4,-31},{9,2,-15},{8,2,-11},{9,1,-7},{7,1,-5},{4,1,-3},{3,1,-1},{3,1,1},{5,1,3},{6,1,5},{3,5,7},{6,2,39},{4,5,43},{4,6,75},{5,7,139},{5,8,267},{6,8,523},{7,9,779},{6,11,1291},{9,32,-32,999},{9,32,3339},{2,-1,0}},{{7,4,-21},{8,0,-5},{7,0,-4},{5,0,-3},{2,2,-2},{5,0,2},{6,0,3},{7,0,4},{8,0,5},{2,6,6},{5,5,70},{6,5,102},{6,6,134},{6,7,198},{6,8,326},{6,9,582},{6,10,1094},{7,11,2118},{8,32,-22,999},{8,32,4166},{2,-1,0}},{{1,0,1},{2,1,2},{4,0,4},{4,1,5},{5,1,7},{5,2,9},{6,2,13},{7,2,17},{7,3,21},{7,4,29},{7,5,45},{7,6,77},{7,32,141}},{{1,0,1},{2,0,2},{3,1,3},{5,0,5},{5,1,6},{6,1,8},{7,0,10},{7,1,11},{7,2,13},{7,3,17},{7,4,25},{8,5,41},{8,32,73}},{{1,0,1},{3,0,2},{4,0,3},{5,0,4},{4,1,5},{3,3,7},{6,1,15},{6,2,17},{6,3,21},{6,4,29},{6,5,45},{7,6,77},{7,32,141}},{{3,0,-2},{3,0,-1},{1,0,0},{3,0,1},{3,0,2}},{{7,4,-24},{6,2,-8},{5,1,-4},{4,0,-2},{3,0,-1},{1,0,0},{3,0,1},{4,0,2},{5,1,3},{6,2,5},{7,4,9},{7,32,-25,999},{7,32,25}}};type InternalNode struct{_bbc int32 ;_cef Node ;_bc Node ;};func _cdg (_cfd int32 )*InternalNode {return &InternalNode {_bbc :_cfd }};func NewEncodedTable (table BasicTabler )(*EncodedTable ,error ){_add :=&EncodedTable {_g :&InternalNode {},BasicTabler :table };if _gg :=_add .parseTable ();_gg !=nil {return nil ,_gg ;};return _add ,nil ;};func _ebe (_ada *Code )*ValueNode {return &ValueNode {_ed :_ada ._aea ,_dc :_ada ._fad ,_df :_ada ._cac }};func NewCode (prefixLength ,rangeLength ,rangeLow int32 ,isLowerRange bool )*Code {return &Code {_gbg :prefixLength ,_aea :rangeLength ,_fad :rangeLow ,_cac :isLowerRange ,_gde :-1};};func (_fcc *StandardTable )Decode (r _c .StreamReader )(int64 ,error ){return _fcc ._bcd .Decode (r )};func NewFixedSizeTable (codeTable []*Code )(*FixedSizeTable ,error ){_aca :=&FixedSizeTable {_feg :&InternalNode {}};if _fdg :=_aca .InitTree (codeTable );_fdg !=nil {return nil ,_fdg ;};return _aca ,nil ;};func (_fb *InternalNode )String ()string {_cfe :=&_f .Builder {};_cfe .WriteString ("\u000a");_fb .pad (_cfe );_cfe .WriteString ("\u0030\u003a\u0020");_cfe .WriteString (_fb ._cef .String ()+"\u000a");_fb .pad (_cfe );_cfe .WriteString ("\u0031\u003a\u0020");_cfe .WriteString (_fb ._bc .String ()+"\u000a");return _cfe .String ();};func (_ddc *EncodedTable )RootNode ()*InternalNode {return _ddc ._g };type Code struct{_gbg int32 ;_aea int32 ;_fad int32 ;_cac bool ;_gde int32 ;};type EncodedTable struct{BasicTabler ;_g *InternalNode ;};func (_fgg *StandardTable )String ()string {return _fgg ._bcd .String ()+"\u000a"};type Node interface{Decode (_cc _c .StreamReader )(int64 ,error );String ()string ;};var _cfg =make ([]Tabler ,len (_ee ));type Tabler interface{Decode (_beb _c .StreamReader )(int64 ,error );InitTree (_bcge []*Code )error ;String ()string ;RootNode ()*InternalNode ;};type BasicTabler interface{HtHigh ()int32 ;HtLow ()int32 ;StreamReader ()_c .StreamReader ;HtPS ()int32 ;HtRS ()int32 ;HtOOB ()int32 ;};func (_cda *InternalNode )append (_bf *Code )(_af error ){if _bf ._gbg ==0{return nil ;};_afe :=_bf ._gbg -1-_cda ._bbc ;if _afe < 0{return _e .New ("\u004e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0073\u0068\u0069\u0066\u0074\u0069n\u0067 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064");};_ge :=(_bf ._gde >>uint (_afe ))&0x1;if _afe ==0{if _bf ._aea ==-1{if _ge ==1{if _cda ._bc !=nil {return _d .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_bf );};_cda ._bc =_bd (_bf );}else {if _cda ._cef !=nil {return _d .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_bf );};_cda ._cef =_bd (_bf );};}else {if _ge ==1{if _cda ._bc !=nil {return _d .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_bf );};_cda ._bc =_ebe (_bf );}else {if _cda ._cef !=nil {return _d .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_bf );};_cda ._cef =_ebe (_bf );};};}else {if _ge ==1{if _cda ._bc ==nil {_cda ._bc =_cdg (_cda ._bbc +1);};if _af =_cda ._bc .(*InternalNode ).append (_bf );_af !=nil {return _af ;};}else {if _cda ._cef ==nil {_cda ._cef =_cdg (_cda ._bbc +1);};if _af =_cda ._cef .(*InternalNode ).append (_bf );_af !=nil {return _af ;};};};return nil ;};func (_cf *FixedSizeTable )String ()string {return _cf ._feg .String ()+"\u000a"};var _ Node =&InternalNode {};type OutOfBandNode struct{};type ValueNode struct{_ed int32 ;_dc int32 ;_df bool ;};var _ Node =&OutOfBandNode {};var _ Node =&ValueNode {};func _gac (_cb ,_gdd int32 )string {var _fcb int32 ;_ag :=make ([]rune ,_gdd );for _ec :=int32 (1);_ec <=_gdd ;_ec ++{_fcb =_cb >>uint (_gdd -_ec )&1;if _fcb !=0{_ag [_ec -1]='1';}else {_ag [_ec -1]='0';};};return string (_ag );};func (_fa *OutOfBandNode )String ()string {return _d .Sprintf ("\u0025\u0030\u00364\u0062",int64 (_ad .MaxInt64 ));};func (_ca *EncodedTable )String ()string {return _ca ._g .String ()+"\u000a"};