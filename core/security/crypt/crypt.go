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

package crypt ;import (_fe "crypto/aes";_g "crypto/cipher";_d "crypto/md5";_dc "crypto/rand";_c "crypto/rc4";_e "fmt";_b "github.com/unidoc/unipdf/v3/common";_ef "github.com/unidoc/unipdf/v3/core/security";_f "io";);func init (){_acg ("\u0041\u0045\u0053V\u0032",_eca )};func _ge (_ba ,_ecc uint32 ,_aga []byte ,_ga bool )([]byte ,error ){_bf :=make ([]byte ,len (_aga )+5);for _gf :=0;_gf < len (_aga );_gf ++{_bf [_gf ]=_aga [_gf ];};for _aa :=0;_aa < 3;_aa ++{_bfe :=byte ((_ba >>uint32 (8*_aa ))&0xff);_bf [_aa +len (_aga )]=_bfe ;};for _dag :=0;_dag < 2;_dag ++{_ccf :=byte ((_ecc >>uint32 (8*_dag ))&0xff);_bf [_dag +len (_aga )+3]=_ccf ;};if _ga {_bf =append (_bf ,0x73);_bf =append (_bf ,0x41);_bf =append (_bf ,0x6C);_bf =append (_bf ,0x54);};_gag :=_d .New ();_gag .Write (_bf );_fegb :=_gag .Sum (nil );if len (_aga )+5< 16{return _fegb [0:len (_aga )+5],nil ;};return _fegb ,nil ;};

// Name implements Filter interface.
func (filterAESV2 )Name ()string {return "\u0041\u0045\u0053V\u0032"};func (filterIdentity )EncryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };func _acg (_dec string ,_gfd filterFunc ){if _ ,_dfg :=_dd [_dec ];_dfg {panic ("\u0061l\u0072e\u0061\u0064\u0079\u0020\u0072e\u0067\u0069s\u0074\u0065\u0072\u0065\u0064");};_dd [_dec ]=_gfd ;};func _ag (_ae FilterDict )(Filter ,error ){if _ae .Length ==256{_b .Log .Debug ("\u0041\u0045S\u0056\u0033\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_ae .Length );_ae .Length /=8;};if _ae .Length !=0&&_ae .Length !=32{return nil ,_e .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0033\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_ae .Length );};return filterAESV3 {},nil ;};

// PDFVersion implements Filter interface.
func (_bbef filterV2 )PDFVersion ()[2]int {return [2]int {}};

// KeyLength implements Filter interface.
func (_gff filterV2 )KeyLength ()int {return _gff ._ggf };

// Name implements Filter interface.
func (filterV2 )Name ()string {return "\u0056\u0032"};

// PDFVersion implements Filter interface.
func (filterAESV2 )PDFVersion ()[2]int {return [2]int {1,5}};func init (){_acg ("\u0041\u0045\u0053V\u0033",_ag )};func (filterIdentity )PDFVersion ()[2]int {return [2]int {}};

// KeyLength implements Filter interface.
func (filterAESV2 )KeyLength ()int {return 128/8};

// Name implements Filter interface.
func (filterAESV3 )Name ()string {return "\u0041\u0045\u0053V\u0033"};func (filterAES )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_ca ,_gd :=_fe .NewCipher (okey );if _gd !=nil {return nil ,_gd ;};_b .Log .Trace ("A\u0045\u0053\u0020\u0045nc\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );const _fc =_fe .BlockSize ;_efc :=_fc -len (buf )%_fc ;for _gda :=0;_gda < _efc ;_gda ++{buf =append (buf ,byte (_efc ));};_b .Log .Trace ("\u0050a\u0064d\u0065\u0064\u0020\u0074\u006f \u0025\u0064 \u0062\u0079\u0074\u0065\u0073",len (buf ));_fa :=make ([]byte ,_fc +len (buf ));_ac :=_fa [:_fc ];if _ ,_da :=_f .ReadFull (_dc .Reader ,_ac );_da !=nil {return nil ,_da ;};_eg :=_g .NewCBCEncrypter (_ca ,_ac );_eg .CryptBlocks (_fa [_fc :],buf );buf =_fa ;_b .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );return buf ,nil ;};

// HandlerVersion implements Filter interface.
func (filterAESV3 )HandlerVersion ()(V ,R int ){V ,R =5,6;return ;};func _eca (_eff FilterDict )(Filter ,error ){if _eff .Length ==128{_b .Log .Debug ("\u0041\u0045S\u0056\u0032\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_eff .Length );_eff .Length /=8;};if _eff .Length !=0&&_eff .Length !=16{return nil ,_e .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0032\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_eff .Length );};return filterAESV2 {},nil ;};func (filterAES )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_cee ,_acb :=_fe .NewCipher (okey );if _acb !=nil {return nil ,_acb ;};if len (buf )< 16{_b .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020\u0041\u0045\u0053\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0062\u0075\u0066\u0020\u0025\u0073",buf );return buf ,_e .Errorf ("\u0041\u0045\u0053\u003a B\u0075\u0066\u0020\u006c\u0065\u006e\u0020\u003c\u0020\u0031\u0036\u0020\u0028\u0025d\u0029",len (buf ));};_de :=buf [:16];buf =buf [16:];if len (buf )%16!=0{_b .Log .Debug ("\u0020\u0069\u0076\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (_de ),_de );_b .Log .Debug ("\u0062\u0075\u0066\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );return buf ,_e .Errorf ("\u0041\u0045\u0053\u0020\u0062\u0075\u0066\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069p\u006c\u0065\u0020\u006f\u0066 \u0031\u0036 \u0028\u0025\u0064\u0029",len (buf ));};_ecg :=_g .NewCBCDecrypter (_cee ,_de );_b .Log .Trace ("A\u0045\u0053\u0020\u0044ec\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );_b .Log .Trace ("\u0063\u0068\u006f\u0070\u0020\u0041\u0045\u0053\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u0020\u0028\u0025\u0064\u0029\u003a \u0025\u0020\u0078",len (buf ),buf );_ecg .CryptBlocks (buf ,buf );_b .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );if len (buf )==0{_b .Log .Trace ("\u0045\u006d\u0070\u0074\u0079\u0020b\u0075\u0066\u002c\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0065\u006d\u0070\u0074\u0079\u0020\u0073t\u0072\u0069\u006e\u0067");return buf ,nil ;};_fb :=int (buf [len (buf )-1]);if _fb > len (buf ){_b .Log .Debug ("\u0049\u006c\u006c\u0065g\u0061\u006c\u0020\u0070\u0061\u0064\u0020\u006c\u0065\u006eg\u0074h\u0020\u0028\u0025\u0064\u0020\u003e\u0020%\u0064\u0029",_fb ,len (buf ));return buf ,_e .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0070a\u0064\u0020l\u0065\u006e\u0067\u0074\u0068");};buf =buf [:len (buf )-_fb ];return buf ,nil ;};

// HandlerVersion implements Filter interface.
func (_ea filterV2 )HandlerVersion ()(V ,R int ){V ,R =2,3;return ;};func (filterIdentity )DecryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };type filterFunc func (_cb FilterDict )(Filter ,error );

// Filter is a common interface for crypt filter methods.
type Filter interface{

// Name returns a name of the filter that should be used in CFM field of Encrypt dictionary.
Name ()string ;

// KeyLength returns a length of the encryption key in bytes.
KeyLength ()int ;

// PDFVersion reports the minimal version of PDF document that introduced this filter.
PDFVersion ()[2]int ;

// HandlerVersion reports V and R parameters that should be used for this filter.
HandlerVersion ()(V ,R int );

// MakeKey generates a object encryption key based on file encryption key and object numbers.
// Used only for legacy filters - AESV3 doesn't change the key for each object.
MakeKey (_ace ,_aea uint32 ,_cf []byte )([]byte ,error );

// EncryptBytes encrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and encrypt data in-place.
EncryptBytes (_agac []byte ,_ff []byte )([]byte ,error );

// DecryptBytes decrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and decrypt data in-place.
DecryptBytes (_ee []byte ,_bc []byte )([]byte ,error );};

// PDFVersion implements Filter interface.
func (filterAESV3 )PDFVersion ()[2]int {return [2]int {2,0}};func init (){_acg ("\u0056\u0032",_ad )};func (filterIdentity )Name ()string {return "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"};

// NewFilterAESV3 creates an AES-based filter with a 256 bit key (AESV3).
func NewFilterAESV3 ()Filter {_dce ,_feg :=_ag (FilterDict {});if _feg !=nil {_b .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0033\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_feg );return filterAESV3 {};};return _dce ;};var (_dd =make (map[string ]filterFunc ););type filterAESV2 struct{filterAES };

// FilterDict represents information from a CryptFilter dictionary.
type FilterDict struct{CFM string ;AuthEvent _ef .AuthEvent ;Length int ;};type filterAES struct{};var _ Filter =filterAESV3 {};var _ Filter =filterAESV2 {};

// MakeKey implements Filter interface.
func (filterAESV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _ge (objNum ,genNum ,ekey ,true );};

// MakeKey implements Filter interface.
func (_ed filterV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _ge (objNum ,genNum ,ekey ,false );};func (filterIdentity )KeyLength ()int {return 0};func (filterIdentity )MakeKey (objNum ,genNum uint32 ,fkey []byte )([]byte ,error ){return fkey ,nil };

// NewFilterV2 creates a RC4-based filter with a specified key length (in bytes).
func NewFilterV2 (length int )Filter {_dga ,_bba :=_ad (FilterDict {Length :length });if _bba !=nil {_b .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020R\u0043\u0034\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_bba );return filterV2 {_ggf :length };};return _dga ;};type filterIdentity struct{};

// EncryptBytes implements Filter interface.
func (filterV2 )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_ebd ,_df :=_c .NewCipher (okey );if _df !=nil {return nil ,_df ;};_b .Log .Trace ("\u0052\u00434\u0020\u0045\u006ec\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );_ebd .XORKeyStream (buf ,buf );_b .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};

// NewFilter creates CryptFilter from a corresponding dictionary.
func NewFilter (d FilterDict )(Filter ,error ){_dea ,_bbg :=_acga (d .CFM );if _bbg !=nil {return nil ,_bbg ;};_daa ,_bbg :=_dea (d );if _bbg !=nil {return nil ,_bbg ;};return _daa ,nil ;};func _ad (_bbe FilterDict )(Filter ,error ){if _bbe .Length %8!=0{return nil ,_e .Errorf ("\u0063\u0072\u0079p\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020o\u0066\u0020\u0038\u0020\u0028\u0025\u0064\u0029",_bbe .Length );};if _bbe .Length < 5||_bbe .Length > 16{if _bbe .Length ==40||_bbe .Length ==64||_bbe .Length ==128{_b .Log .Debug ("\u0053\u0054\u0041\u004e\u0044AR\u0044\u0020V\u0049\u004f\u004c\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0043\u0072\u0079\u0070\u0074\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072s\u0020\u0074\u006f \u0062\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u0074\u0073\u0020\u0072\u0061t\u0068\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0020-\u0020\u0061s\u0073u\u006d\u0069\u006e\u0067\u0020\u0062\u0069t\u0073\u0020\u0028\u0025\u0064\u0029",_bbe .Length );_bbe .Length /=8;}else {return nil ,_e .Errorf ("\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074h\u0020\u006e\u006f\u0074\u0020\u0069\u006e \u0072\u0061\u006e\u0067\u0065\u0020\u0034\u0030\u0020\u002d\u00201\u0032\u0038\u0020\u0062\u0069\u0074\u0020\u0028\u0025\u0064\u0029",_bbe .Length );};};return filterV2 {_ggf :_bbe .Length },nil ;};type filterAESV3 struct{filterAES };func (filterIdentity )HandlerVersion ()(V ,R int ){return ;};type filterV2 struct{_ggf int };

// MakeKey implements Filter interface.
func (filterAESV3 )MakeKey (_ ,_ uint32 ,ekey []byte )([]byte ,error ){return ekey ,nil };func _acga (_gdb string )(filterFunc ,error ){_fg :=_dd [string (_gdb )];if _fg ==nil {return nil ,_e .Errorf ("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0072\u0079p\u0074 \u0066\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0071",_gdb );};return _fg ,nil ;};

// KeyLength implements Filter interface.
func (filterAESV3 )KeyLength ()int {return 256/8};

// HandlerVersion implements Filter interface.
func (filterAESV2 )HandlerVersion ()(V ,R int ){V ,R =4,4;return ;};

// NewFilterAESV2 creates an AES-based filter with a 128 bit key (AESV2).
func NewFilterAESV2 ()Filter {_gg ,_bb :=_eca (FilterDict {});if _bb !=nil {_b .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_bb );return filterAESV2 {};};return _gg ;};

// NewIdentity creates an identity filter that bypasses all data without changes.
func NewIdentity ()Filter {return filterIdentity {}};var _ Filter =filterV2 {};

// DecryptBytes implements Filter interface.
func (filterV2 )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_gde ,_ega :=_c .NewCipher (okey );if _ega !=nil {return nil ,_ega ;};_b .Log .Trace ("\u0052\u00434\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );_gde .XORKeyStream (buf ,buf );_b .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};