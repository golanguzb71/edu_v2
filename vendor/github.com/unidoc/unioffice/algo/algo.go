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

package algo ;import _c "strconv";

// NaturalLess compares two strings in a human manner so rId2 sorts less than rId10
func NaturalLess (lhs ,rhs string )bool {_d ,_cc :=0,0;for _d < len (lhs )&&_cc < len (rhs ){_a :=lhs [_d ];_ff :=rhs [_cc ];_cg :=_b (_a );_db :=_b (_ff );switch {case _cg &&!_db :return true ;case !_cg &&_db :return false ;case !_cg &&!_db :if _a !=_ff {return _a < _ff ;
};_d ++;_cc ++;default:_fg :=_d +1;_be :=_cc +1;for _fg < len (lhs )&&_b (lhs [_fg ]){_fg ++;};for _be < len (rhs )&&_b (rhs [_be ]){_be ++;};_de ,_ :=_c .ParseUint (lhs [_d :_fg ],10,64);_bec ,_ :=_c .ParseUint (rhs [_d :_be ],10,64);if _de !=_bec {return _de < _bec ;
};_d =_fg ;_cc =_be ;};};return len (lhs )< len (rhs );};func _b (_f byte )bool {return _f >='0'&&_f <='9'};func RepeatString (s string ,cnt int )string {if cnt <=0{return "";};_ag :=make ([]byte ,len (s )*cnt );_fb :=[]byte (s );for _agb :=0;_agb < cnt ;
_agb ++{copy (_ag [_agb :],_fb );};return string (_ag );};