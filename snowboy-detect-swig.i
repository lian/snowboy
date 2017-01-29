%module snowboy

#pragma SWIG nowarn=SWIGWARN_PARSE_NESTED_CLASS

%include "std_string.i"

%{
#include "snowboy-detect.h"
%}

%include "snowboy-detect.h"
