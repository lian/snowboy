package snowboy

//go:generate swig -go -cgo -c++ -intgosize 64 -o snowboy-detect-swig.cc snowboy-detect-swig.i

/*
#cgo CXXFLAGS: -D_GLIBCXX_USE_CXX11_ABI=0
#cgo LDFLAGS: -lcblas -L${SRCDIR} -Wl,-Bstatic -lsnowboy-detect -Wl,-Bdynamic

// if you actually want to use atlas, place static atlas libs into SRCDIR/atlas-lib
//#cgo LDFLAGS: -L${SRCDIR}/atlas-lib -Wl,-Bstatic -lf77blas -lcblas -llapack -latlas -Wl,-Bdynamic -L${SRCDIR} -Wl,-Bstatic -lsnowboy-detect -Wl,-Bdynamic
*/
import "C"
