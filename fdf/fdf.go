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

// Package fdf provides support for loading form field data from Form Field Data (FDF) files.
package fdf ;import (_ca "bufio";_fc "bytes";_cf "encoding/hex";_eg "errors";_d "fmt";_b "github.com/unidoc/unipdf/v3/common";_db "github.com/unidoc/unipdf/v3/core";_cg "io";_da "os";_c "regexp";_egg "sort";_g "strconv";_f "strings";);func (_dfc *fdfParser )parseString ()(*_db .PdfObjectString ,error ){_dfc ._eca .ReadByte ();var _bgba _fc .Buffer ;_ecac :=1;for {_cgd ,_gac :=_dfc ._eca .Peek (1);if _gac !=nil {return _db .MakeString (_bgba .String ()),_gac ;};if _cgd [0]=='\\'{_dfc ._eca .ReadByte ();_bef ,_dae :=_dfc ._eca .ReadByte ();if _dae !=nil {return _db .MakeString (_bgba .String ()),_dae ;};if _db .IsOctalDigit (_bef ){_eab ,_eae :=_dfc ._eca .Peek (2);if _eae !=nil {return _db .MakeString (_bgba .String ()),_eae ;};var _gbg []byte ;_gbg =append (_gbg ,_bef );for _ ,_ddg :=range _eab {if _db .IsOctalDigit (_ddg ){_gbg =append (_gbg ,_ddg );}else {break ;};};_dfc ._eca .Discard (len (_gbg )-1);_b .Log .Trace ("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022",_gbg );_fed ,_eae :=_g .ParseUint (string (_gbg ),8,32);if _eae !=nil {return _db .MakeString (_bgba .String ()),_eae ;};_bgba .WriteByte (byte (_fed ));continue ;};switch _bef {case 'n':_bgba .WriteRune ('\n');case 'r':_bgba .WriteRune ('\r');case 't':_bgba .WriteRune ('\t');case 'b':_bgba .WriteRune ('\b');case 'f':_bgba .WriteRune ('\f');case '(':_bgba .WriteRune ('(');case ')':_bgba .WriteRune (')');case '\\':_bgba .WriteRune ('\\');};continue ;}else if _cgd [0]=='('{_ecac ++;}else if _cgd [0]==')'{_ecac --;if _ecac ==0{_dfc ._eca .ReadByte ();break ;};};_bbd ,_ :=_dfc ._eca .ReadByte ();_bgba .WriteByte (_bbd );};return _db .MakeString (_bgba .String ()),nil ;};type fdfParser struct{_ce int ;_ad int ;_bdb map[int64 ]_db .PdfObject ;_dcc _cg .ReadSeeker ;_eca *_ca .Reader ;_cca int64 ;_deg *_db .PdfObjectDictionary ;};func (_dfb *fdfParser )trace (_efc _db .PdfObject )_db .PdfObject {switch _eef :=_efc .(type ){case *_db .PdfObjectReference :_bbe ,_gdg :=_dfb ._bdb [_eef .ObjectNumber ].(*_db .PdfIndirectObject );if _gdg {return _bbe .PdfObject ;};_b .Log .Debug ("\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");return nil ;case *_db .PdfIndirectObject :return _eef .PdfObject ;};return _efc ;};func (_gfbc *fdfParser )parse ()error {_gfbc ._dcc .Seek (0,_cg .SeekStart );_gfbc ._eca =_ca .NewReader (_gfbc ._dcc );for {_gfbc .skipComments ();_afe ,_bfd :=_gfbc ._eca .Peek (20);if _bfd !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");return _bfd ;};if _f .HasPrefix (string (_afe ),"\u0074r\u0061\u0069\u006c\u0065\u0072"){_gfbc ._eca .Discard (7);_gfbc .skipSpaces ();_gfbc .skipComments ();_bdbf ,_ :=_gfbc .parseDict ();_gfbc ._deg =_bdbf ;break ;};_eagf :=_fca .FindStringSubmatchIndex (string (_afe ));if len (_eagf )< 6{_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_afe ));return _eg .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");};_cee ,_bfd :=_gfbc .parseIndirectObject ();if _bfd !=nil {return _bfd ;};switch _agb :=_cee .(type ){case *_db .PdfIndirectObject :_gfbc ._bdb [_agb .ObjectNumber ]=_agb ;case *_db .PdfObjectStream :_gfbc ._bdb [_agb .ObjectNumber ]=_agb ;default:return _eg .New ("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");};};return nil ;};func (_acdf *fdfParser )seekFdfVersionTopDown ()(int ,int ,error ){_acdf ._dcc .Seek (0,_cg .SeekStart );_acdf ._eca =_ca .NewReader (_acdf ._dcc );_ceeg :=20;_adag :=make ([]byte ,_ceeg );for {_aga ,_bbg :=_acdf ._eca .ReadByte ();if _bbg !=nil {if _bbg ==_cg .EOF {break ;}else {return 0,0,_bbg ;};};if _db .IsDecimalDigit (_aga )&&_adag [_ceeg -1]=='.'&&_db .IsDecimalDigit (_adag [_ceeg -2])&&_adag [_ceeg -3]=='-'&&_adag [_ceeg -4]=='F'&&_adag [_ceeg -5]=='D'&&_adag [_ceeg -6]=='P'{_cfc :=int (_adag [_ceeg -2]-'0');_dagf :=int (_aga -'0');return _cfc ,_dagf ,nil ;};_adag =append (_adag [1:_ceeg ],_aga );};return 0,0,_eg .New ("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064");};

// Root returns the Root of the FDF document.
func (_ddgd *fdfParser )Root ()(*_db .PdfObjectDictionary ,error ){if _ddgd ._deg !=nil {if _dfg ,_dcd :=_ddgd .trace (_ddgd ._deg .Get ("\u0052\u006f\u006f\u0074")).(*_db .PdfObjectDictionary );_dcd {if _ecaf ,_bdae :=_ddgd .trace (_dfg .Get ("\u0046\u0044\u0046")).(*_db .PdfObjectDictionary );_bdae {return _ecaf ,nil ;};};};var _aeb []int64 ;for _dda :=range _ddgd ._bdb {_aeb =append (_aeb ,_dda );};_egg .Slice (_aeb ,func (_acc ,_ge int )bool {return _aeb [_acc ]< _aeb [_ge ]});for _ ,_cbc :=range _aeb {_efg :=_ddgd ._bdb [_cbc ];if _dfd ,_cac :=_ddgd .trace (_efg ).(*_db .PdfObjectDictionary );_cac {if _bagg ,_adga :=_ddgd .trace (_dfd .Get ("\u0046\u0044\u0046")).(*_db .PdfObjectDictionary );_adga {return _bagg ,nil ;};};};return nil ,_eg .New ("\u0046\u0044\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};var _bfa =_c .MustCompile ("\u0025F\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029");var _abb =_c .MustCompile ("\u0025\u0025\u0045O\u0046");var _bgb =_c .MustCompile ("^\u005c\u0073\u002a\u0028\\d\u002b)\u005c\u0073\u002b\u0028\u005cd\u002b\u0029\u005c\u0073\u002b\u0052");func (_deff *fdfParser )parseFdfVersion ()(int ,int ,error ){_deff ._dcc .Seek (0,_cg .SeekStart );_dce :=20;_gceb :=make ([]byte ,_dce );_deff ._dcc .Read (_gceb );_bad :=_bfa .FindStringSubmatch (string (_gceb ));if len (_bad )< 3{_gbffg ,_bce ,_ebd :=_deff .seekFdfVersionTopDown ();if _ebd !=nil {_b .Log .Debug ("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e");return 0,0,_ebd ;};return _gbffg ,_bce ,nil ;};_eea ,_dbec :=_g .Atoi (_bad [1]);if _dbec !=nil {return 0,0,_dbec ;};_adg ,_dbec :=_g .Atoi (_bad [2]);if _dbec !=nil {return 0,0,_dbec ;};_b .Log .Debug ("\u0046\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064",_eea ,_adg );return int (_eea ),int (_adg ),nil ;};func (_ff *fdfParser )skipComments ()error {if _ ,_fb :=_ff .skipSpaces ();_fb !=nil {return _fb ;};_cd :=true ;for {_cec ,_ae :=_ff ._eca .Peek (1);if _ae !=nil {_b .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_ae .Error ());return _ae ;};if _cd &&_cec [0]!='%'{return nil ;};_cd =false ;if (_cec [0]!='\r')&&(_cec [0]!='\n'){_ff ._eca .ReadByte ();}else {break ;};};return _ff .skipComments ();};func _gfb (_beag _cg .ReadSeeker )(*fdfParser ,error ){_bga :=&fdfParser {};_bga ._dcc =_beag ;_bga ._bdb =map[int64 ]_db .PdfObject {};_egc ,_ecc ,_dffa :=_bga .parseFdfVersion ();if _dffa !=nil {_b .Log .Error ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076",_dffa );return nil ,_dffa ;};_bga ._ce =_egc ;_bga ._ad =_ecc ;_dffa =_bga .parse ();return _bga ,_dffa ;};

// FieldDictionaries returns a map of field names to field dictionaries.
func (fdf *Data )FieldDictionaries ()(map[string ]*_db .PdfObjectDictionary ,error ){_ed :=map[string ]*_db .PdfObjectDictionary {};for _bd :=0;_bd < fdf ._gb .Len ();_bd ++{_dag ,_dd :=_db .GetDict (fdf ._gb .Get (_bd ));if _dd {_egd ,_ :=_db .GetString (_dag .Get ("\u0054"));if _egd !=nil {_ed [_egd .Str ()]=_dag ;};};};return _ed ,nil ;};func (_fe *fdfParser )getFileOffset ()int64 {_af ,_ :=_fe ._dcc .Seek (0,_cg .SeekCurrent );_af -=int64 (_fe ._eca .Buffered ());return _af ;};

// LoadFromPath loads FDF form data from file path `fdfPath`.
func LoadFromPath (fdfPath string )(*Data ,error ){_fd ,_ee :=_da .Open (fdfPath );if _ee !=nil {return nil ,_ee ;};defer _fd .Close ();return Load (_fd );};func (_gdc *fdfParser )parseObject ()(_db .PdfObject ,error ){_b .Log .Trace ("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074");_gdc .skipSpaces ();for {_gbff ,_feaa :=_gdc ._eca .Peek (2);if _feaa !=nil {return nil ,_feaa ;};_b .Log .Trace ("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073",string (_gbff ));if _gbff [0]=='/'{_ded ,_ega :=_gdc .parseName ();_b .Log .Trace ("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027",_ded );return &_ded ,_ega ;}else if _gbff [0]=='('{_b .Log .Trace ("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!");return _gdc .parseString ();}else if _gbff [0]=='['{_b .Log .Trace ("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021");return _gdc .parseArray ();}else if (_gbff [0]=='<')&&(_gbff [1]=='<'){_b .Log .Trace ("\u002d>\u0044\u0069\u0063\u0074\u0021");return _gdc .parseDict ();}else if _gbff [0]=='<'{_b .Log .Trace ("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021");return _gdc .parseHexString ();}else if _gbff [0]=='%'{_gdc .readComment ();_gdc .skipSpaces ();}else {_b .Log .Trace ("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f");_gbff ,_ =_gdc ._eca .Peek (15);_gaeb :=string (_gbff );_b .Log .Trace ("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073",_gaeb );if (len (_gaeb )> 3)&&(_gaeb [:4]=="\u006e\u0075\u006c\u006c"){_bdd ,_dg :=_gdc .parseNull ();return &_bdd ,_dg ;}else if (len (_gaeb )> 4)&&(_gaeb [:5]=="\u0066\u0061\u006cs\u0065"){_daf ,_eedc :=_gdc .parseBool ();return &_daf ,_eedc ;}else if (len (_gaeb )> 3)&&(_gaeb [:4]=="\u0074\u0072\u0075\u0065"){_cbg ,_dege :=_gdc .parseBool ();return &_cbg ,_dege ;};_bff :=_bgb .FindStringSubmatch (string (_gaeb ));if len (_bff )> 1{_gbff ,_ =_gdc ._eca .ReadBytes ('R');_b .Log .Trace ("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027",string (_gbff [:]));_ffe ,_dec :=_eag (string (_gbff ));return &_ffe ,_dec ;};_dab :=_ea .FindStringSubmatch (string (_gaeb ));if len (_dab )> 1{_b .Log .Trace ("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021");return _gdc .parseNumber ();};_dab =_fcf .FindStringSubmatch (string (_gaeb ));if len (_dab )> 1{_b .Log .Trace ("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021");_b .Log .Trace ("\u0025\u0020\u0073",_dab );return _gdc .parseNumber ();};_b .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)",_gaeb );return nil ,_eg .New ("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e");};};};var _fca =_c .MustCompile ("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a");func (_gaef *fdfParser )parseHexString ()(*_db .PdfObjectString ,error ){_gaef ._eca .ReadByte ();var _fbe _fc .Buffer ;for {_eeb ,_fcg :=_gaef ._eca .Peek (1);if _fcg !=nil {return _db .MakeHexString (""),_fcg ;};if _eeb [0]=='>'{_gaef ._eca .ReadByte ();break ;};_bag ,_ :=_gaef ._eca .ReadByte ();if !_db .IsWhiteSpace (_bag ){_fbe .WriteByte (_bag );};};if _fbe .Len ()%2==1{_fbe .WriteRune ('0');};_gfd ,_abbd :=_cf .DecodeString (_fbe .String ());if _abbd !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0078\u0020\u0073\u0074r\u0069\u006e\u0067\u003a\u0020\u0027\u0025\u0073\u0027 \u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0061n\u0020\u0065\u006d\u0070\u0074\u0079 \u0073\u0074\u0072i\u006e\u0067",_fbe .String ());return _db .MakeHexString (""),nil ;};return _db .MakeHexString (string (_gfd )),nil ;};func (_bgc *fdfParser )parseArray ()(*_db .PdfObjectArray ,error ){_ddd :=_db .MakeArray ();_bgc ._eca .ReadByte ();for {_bgc .skipSpaces ();_cgb ,_gce :=_bgc ._eca .Peek (1);if _gce !=nil {return _ddd ,_gce ;};if _cgb [0]==']'{_bgc ._eca .ReadByte ();break ;};_fcfe ,_gce :=_bgc .parseObject ();if _gce !=nil {return _ddd ,_gce ;};_ddd .Append (_fcfe );};return _ddd ,nil ;};func (_dbe *fdfParser )parseNumber ()(_db .PdfObject ,error ){return _db .ParseNumber (_dbe ._eca )};func (_egb *fdfParser )readAtLeast (_faa []byte ,_bg int )(int ,error ){_ef :=_bg ;_faag :=0;_daa :=0;for _ef > 0{_ec ,_df :=_egb ._eca .Read (_faa [_faag :]);if _df !=nil {_b .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073",_ec ,_daa ,_df .Error ());return _faag ,_eg .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067");};_daa ++;_faag +=_ec ;_ef -=_ec ;};return _faag ,nil ;};func _eag (_gaee string )(_db .PdfObjectReference ,error ){_acg :=_db .PdfObjectReference {};_ceb :=_bgb .FindStringSubmatch (string (_gaee ));if len (_ceb )< 3{_b .Log .Debug ("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065");return _acg ,_eg .New ("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e");};_gbc ,_dbg :=_g .Atoi (_ceb [1]);if _dbg !=nil {_b .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069n\u0067\u0020\u006fb\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072 '\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0020\u003d\u0020\u0030",_ceb [1]);return _acg ,nil ;};_acg .ObjectNumber =int64 (_gbc );_aag ,_dbg :=_g .Atoi (_ceb [2]);if _dbg !=nil {_b .Log .Debug ("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020g\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0027\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0020\u003d\u0020\u0030",_ceb [2]);return _acg ,nil ;};_acg .GenerationNumber =int64 (_aag );return _acg ,nil ;};func (_eee *fdfParser )readTextLine ()(string ,error ){var _cga _fc .Buffer ;for {_cad ,_bee :=_eee ._eca .Peek (1);if _bee !=nil {_b .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_bee .Error ());return _cga .String (),_bee ;};if (_cad [0]!='\r')&&(_cad [0]!='\n'){_ecb ,_ :=_eee ._eca .ReadByte ();_cga .WriteByte (_ecb );}else {break ;};};return _cga .String (),nil ;};func (_cab *fdfParser )parseIndirectObject ()(_db .PdfObject ,error ){_cdb :=_db .PdfIndirectObject {};_b .Log .Trace ("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a");_faae ,_cdf :=_cab ._eca .Peek (20);if _cdf !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");return &_cdb ,_cdf ;};_b .Log .Trace ("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022",string (_faae ));_befe :=_fca .FindStringSubmatchIndex (string (_faae ));if len (_befe )< 6{_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_faae ));return &_cdb ,_eg .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");};_cab ._eca .Discard (_befe [0]);_b .Log .Trace ("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064",_befe );_add :=_befe [1]-_befe [0];_gdb :=make ([]byte ,_add );_ ,_cdf =_cab .readAtLeast (_gdb ,_add );if _cdf !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073",_cdf );return nil ,_cdf ;};_b .Log .Trace ("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073",_gdb );_beg :=_fca .FindStringSubmatch (string (_gdb ));if len (_beg )< 3{_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_gdb ));return &_cdb ,_eg .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");};_cfg ,_ :=_g .Atoi (_beg [1]);_gbd ,_ :=_g .Atoi (_beg [2]);_cdb .ObjectNumber =int64 (_cfg );_cdb .GenerationNumber =int64 (_gbd );for {_bba ,_gcce :=_cab ._eca .Peek (2);if _gcce !=nil {return &_cdb ,_gcce ;};_b .Log .Trace ("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_bba ),string (_bba ));if _db .IsWhiteSpace (_bba [0]){_cab .skipSpaces ();}else if _bba [0]=='%'{_cab .skipComments ();}else if (_bba [0]=='<')&&(_bba [1]=='<'){_b .Log .Trace ("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074");_cdb .PdfObject ,_gcce =_cab .parseDict ();_b .Log .Trace ("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076",_gcce );if _gcce !=nil {return &_cdb ,_gcce ;};_b .Log .Trace ("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e");}else if (_bba [0]=='/')||(_bba [0]=='(')||(_bba [0]=='[')||(_bba [0]=='<'){_cdb .PdfObject ,_gcce =_cab .parseObject ();if _gcce !=nil {return &_cdb ,_gcce ;};_b .Log .Trace ("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e");}else {if _bba [0]=='e'{_ecg ,_gff :=_cab .readTextLine ();if _gff !=nil {return nil ,_gff ;};if len (_ecg )>=6&&_ecg [0:6]=="\u0065\u006e\u0064\u006f\u0062\u006a"{break ;};}else if _bba [0]=='s'{_bba ,_ =_cab ._eca .Peek (10);if string (_bba [:6])=="\u0073\u0074\u0072\u0065\u0061\u006d"{_abc :=6;if len (_bba )> 6{if _db .IsWhiteSpace (_bba [_abc ])&&_bba [_abc ]!='\r'&&_bba [_abc ]!='\n'{_b .Log .Debug ("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0046\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072");_abc ++;};if _bba [_abc ]=='\r'{_abc ++;if _bba [_abc ]=='\n'{_abc ++;};}else if _bba [_abc ]=='\n'{_abc ++;};};_cab ._eca .Discard (_abc );_gga ,_eafa :=_cdb .PdfObject .(*_db .PdfObjectDictionary );if !_eafa {return nil ,_eg .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079");};_b .Log .Trace ("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073",_gga );_aba ,_gbb :=_gga .Get ("\u004c\u0065\u006e\u0067\u0074\u0068").(*_db .PdfObjectInteger );if !_gbb {return nil ,_eg .New ("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072");};_acgb :=*_aba ;if _acgb < 0{return nil ,_eg .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030");};if int64 (_acgb )> _cab ._cca {_b .Log .Debug ("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065");return nil ,_eg .New ("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065");};_cgag :=make ([]byte ,_acgb );_ ,_gcce =_cab .readAtLeast (_cgag ,int (_acgb ));if _gcce !=nil {_b .Log .Debug ("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058",len (_cgag ),_cgag );_b .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_gcce );return nil ,_gcce ;};_badf :=_db .PdfObjectStream {};_badf .Stream =_cgag ;_badf .PdfObjectDictionary =_cdb .PdfObject .(*_db .PdfObjectDictionary );_badf .ObjectNumber =_cdb .ObjectNumber ;_badf .GenerationNumber =_cdb .GenerationNumber ;_cab .skipSpaces ();_cab ._eca .Discard (9);_cab .skipSpaces ();return &_badf ,nil ;};};_cdb .PdfObject ,_gcce =_cab .parseObject ();return &_cdb ,_gcce ;};};_b .Log .Trace ("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021");return &_cdb ,nil ;};var _ea =_c .MustCompile ("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029");func (_ab *fdfParser )setFileOffset (_dde int64 ){_ab ._dcc .Seek (_dde ,_cg .SeekStart );_ab ._eca =_ca .NewReader (_ab ._dcc );};func _aff (_feaag string )(*fdfParser ,error ){_beeg :=fdfParser {};_fdf :=[]byte (_feaag );_aagg :=_fc .NewReader (_fdf );_beeg ._dcc =_aagg ;_beeg ._bdb =map[int64 ]_db .PdfObject {};_aad :=_ca .NewReader (_aagg );_beeg ._eca =_aad ;_beeg ._cca =int64 (len (_feaag ));return &_beeg ,_beeg .parse ();};func (_ada *fdfParser )skipSpaces ()(int ,error ){_gab :=0;for {_gg ,_gf :=_ada ._eca .ReadByte ();if _gf !=nil {return 0,_gf ;};if _db .IsWhiteSpace (_gg ){_gab ++;}else {_ada ._eca .UnreadByte ();break ;};};return _gab ,nil ;};

// Load loads FDF form data from `r`.
func Load (r _cg .ReadSeeker )(*Data ,error ){_gc ,_gd :=_gfb (r );if _gd !=nil {return nil ,_gd ;};_be ,_gd :=_gc .Root ();if _gd !=nil {return nil ,_gd ;};_cc ,_gcf :=_db .GetArray (_be .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));if !_gcf {return nil ,_eg .New ("\u0066\u0069\u0065\u006c\u0064\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067");};return &Data {_gb :_cc ,_dc :_be },nil ;};func (_gae *fdfParser )parseName ()(_db .PdfObjectName ,error ){var _baf _fc .Buffer ;_cfe :=false ;for {_bb ,_fea :=_gae ._eca .Peek (1);if _fea ==_cg .EOF {break ;};if _fea !=nil {return _db .PdfObjectName (_baf .String ()),_fea ;};if !_cfe {if _bb [0]=='/'{_cfe =true ;_gae ._eca .ReadByte ();}else if _bb [0]=='%'{_gae .readComment ();_gae .skipSpaces ();}else {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029",_bb ,_bb );return _db .PdfObjectName (_baf .String ()),_d .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029",_bb [0]);};}else {if _db .IsWhiteSpace (_bb [0]){break ;}else if (_bb [0]=='/')||(_bb [0]=='[')||(_bb [0]=='(')||(_bb [0]==']')||(_bb [0]=='<')||(_bb [0]=='>'){break ;}else if _bb [0]=='#'{_acd ,_dff :=_gae ._eca .Peek (3);if _dff !=nil {return _db .PdfObjectName (_baf .String ()),_dff ;};_gae ._eca .Discard (3);_ddb ,_dff :=_cf .DecodeString (string (_acd [1:3]));if _dff !=nil {return _db .PdfObjectName (_baf .String ()),_dff ;};_baf .Write (_ddb );}else {_cbf ,_ :=_gae ._eca .ReadByte ();_baf .WriteByte (_cbf );};};};return _db .PdfObjectName (_baf .String ()),nil ;};

// Data represents forms data format (FDF) file data.
type Data struct{_dc *_db .PdfObjectDictionary ;_gb *_db .PdfObjectArray ;};func (_cde *fdfParser )parseDict ()(*_db .PdfObjectDictionary ,error ){_b .Log .Trace ("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020D\u0069\u0063\u0074\u0021");_ece :=_db .MakeDict ();_gcb ,_ :=_cde ._eca .ReadByte ();if _gcb !='<'{return nil ,_eg .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");};_gcb ,_ =_cde ._eca .ReadByte ();if _gcb !='<'{return nil ,_eg .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");};for {_cde .skipSpaces ();_cde .skipComments ();_fcae ,_cff :=_cde ._eca .Peek (2);if _cff !=nil {return nil ,_cff ;};_b .Log .Trace ("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_fcae ),string (_fcae ));if (_fcae [0]=='>')&&(_fcae [1]=='>'){_b .Log .Trace ("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079");_cde ._eca .ReadByte ();_cde ._eca .ReadByte ();break ;};_b .Log .Trace ("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021");_bcb ,_cff :=_cde .parseName ();_b .Log .Trace ("\u004be\u0079\u003a\u0020\u0025\u0073",_bcb );if _cff !=nil {_b .Log .Debug ("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073",_cff );return nil ,_cff ;};if len (_bcb )> 4&&_bcb [len (_bcb )-4:]=="\u006e\u0075\u006c\u006c"{_eaf :=_bcb [0:len (_bcb )-4];_b .Log .Debug ("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029",_bcb );_b .Log .Debug ("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c",_eaf );_cde .skipSpaces ();_fec ,_ :=_cde ._eca .Peek (1);if _fec [0]=='/'{_ece .Set (_eaf ,_db .MakeNull ());continue ;};};_cde .skipSpaces ();_caa ,_cff :=_cde .parseObject ();if _cff !=nil {return nil ,_cff ;};_ece .Set (_bcb ,_caa );_b .Log .Trace ("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073",_bcb ,_caa .String ());};_b .Log .Trace ("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021");return _ece ,nil ;};var _fcf =_c .MustCompile ("\u005e\u005b\u005c+-\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e]\u002b)\u0065[\u005c+\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e\u005d\u002b\u0029");func (_bacc *fdfParser )parseNull ()(_db .PdfObjectNull ,error ){_ ,_cce :=_bacc ._eca .Discard (4);return _db .PdfObjectNull {},_cce ;};func (_gcff *fdfParser )readComment ()(string ,error ){var _dfa _fc .Buffer ;_ ,_aa :=_gcff .skipSpaces ();if _aa !=nil {return _dfa .String (),_aa ;};_ac :=true ;for {_gbf ,_bac :=_gcff ._eca .Peek (1);if _bac !=nil {_b .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_bac .Error ());return _dfa .String (),_bac ;};if _ac &&_gbf [0]!='%'{return _dfa .String (),_eg .New ("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025");};_ac =false ;if (_gbf [0]!='\r')&&(_gbf [0]!='\n'){_gcd ,_ :=_gcff ._eca .ReadByte ();_dfa .WriteByte (_gcd );}else {break ;};};return _dfa .String (),nil ;};func (_fad *fdfParser )parseBool ()(_db .PdfObjectBool ,error ){_caf ,_cag :=_fad ._eca .Peek (4);if _cag !=nil {return _db .PdfObjectBool (false ),_cag ;};if (len (_caf )>=4)&&(string (_caf [:4])=="\u0074\u0072\u0075\u0065"){_fad ._eca .Discard (4);return _db .PdfObjectBool (true ),nil ;};_caf ,_cag =_fad ._eca .Peek (5);if _cag !=nil {return _db .PdfObjectBool (false ),_cag ;};if (len (_caf )>=5)&&(string (_caf [:5])=="\u0066\u0061\u006cs\u0065"){_fad ._eca .Discard (5);return _db .PdfObjectBool (false ),nil ;};return _db .PdfObjectBool (false ),_eg .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg");};func (_eaec *fdfParser )seekToEOFMarker (_bafb int64 )error {_ecbf :=int64 (0);_gbgc :=int64 (1000);for _ecbf < _bafb {if _bafb <=(_gbgc +_ecbf ){_gbgc =_bafb -_ecbf ;};_ ,_agc :=_eaec ._dcc .Seek (-_ecbf -_gbgc ,_cg .SeekEnd );if _agc !=nil {return _agc ;};_gcdc :=make ([]byte ,_gbgc );_eaec ._dcc .Read (_gcdc );_b .Log .Trace ("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022",string (_gcdc ));_eaa :=_abb .FindAllStringIndex (string (_gcdc ),-1);if _eaa !=nil {_edd :=_eaa [len (_eaa )-1];_b .Log .Trace ("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064",_eaa );_eaec ._dcc .Seek (-_ecbf -_gbgc +int64 (_edd [0]),_cg .SeekEnd );return nil ;};_b .Log .Debug ("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067");_ecbf +=_gbgc ;};_b .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e");return _eg .New ("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};

// FieldValues implements interface model.FieldValueProvider.
// Returns a map of field names to values (PdfObjects).
func (fdf *Data )FieldValues ()(map[string ]_db .PdfObject ,error ){_fa ,_a :=fdf .FieldDictionaries ();if _a !=nil {return nil ,_a ;};var _de []string ;for _ddc :=range _fa {_de =append (_de ,_ddc );};_egg .Strings (_de );_eb :=map[string ]_db .PdfObject {};for _ ,_ebb :=range _de {_ba :=_fa [_ebb ];_gcc :=_db .TraceToDirectObject (_ba .Get ("\u0056"));_eb [_ebb ]=_gcc ;};return _eb ,nil ;};